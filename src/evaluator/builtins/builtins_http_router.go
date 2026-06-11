package builtins

import (
	"BanglaCode/src/object"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"
)

// Route holds a compiled route pattern with its handler.
type Route struct {
	pattern string
	params  []string       // param names in order: /users/:id → ["id"]
	re      *regexp.Regexp // precompiled once at AddRoute time
	handler object.Object
}

// CORSOptions holds CORS configuration.
type CORSOptions struct {
	Origin  string
	Methods string
	Headers string
	MaxAge  string
}

// FileRoute maps a URL prefix to a static file handler.
type FileRoute struct {
	prefix  string
	handler http.Handler
}

// RateLimiter implements a per-IP sliding-window rate limiter.
type RateLimiter struct {
	max    int
	window time.Duration
	mu     sync.Mutex
	hits   map[string][]time.Time
}

// Allow reports whether the given IP may proceed; trims expired hits.
func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	now := time.Now()
	cutoff := now.Add(-rl.window)
	prev := rl.hits[ip]
	valid := prev[:0]
	for _, t := range prev {
		if t.After(cutoff) {
			valid = append(valid, t)
		}
	}
	if len(valid) >= rl.max {
		rl.hits[ip] = valid
		return false
	}
	rl.hits[ip] = append(valid, now)
	return true
}

// Router is a modular HTTP router with Express.js-level features.
type Router struct {
	basePath     string
	routes       map[string][]Route // HTTP method → ordered route slice
	middlewares  []object.Object    // run before every route handler
	fileRoutes   []FileRoute
	errorHandler object.Object // bhul_sambhalo handler
	corsEnabled  bool
	corsOptions  CORSOptions
	gzipEnabled  bool
	logEnabled   bool
	timeout      time.Duration
	maxBodyBytes int64
	rateLimiter  *RateLimiter
	mu           sync.RWMutex
}

// NewRouter creates a Router with all HTTP methods pre-initialized.
func NewRouter(basePath string) *Router {
	return &Router{
		basePath: basePath,
		routes: map[string][]Route{
			"GET": {}, "POST": {}, "PUT": {}, "DELETE": {},
			"PATCH": {}, "HEAD": {}, "OPTIONS": {},
		},
	}
}

// compilePattern converts a path pattern into a precompiled regex and param list.
// Example: /users/:id/posts/:pid → ^/users/([^/]+)/posts/([^/]+)$, ["id","pid"]
func compilePattern(pattern string) (*regexp.Regexp, []string) {
	parts := strings.Split(pattern, "/")
	var paramNames []string
	var regexParts []string
	for _, part := range parts {
		if strings.HasPrefix(part, ":") {
			paramNames = append(paramNames, part[1:])
			regexParts = append(regexParts, "([^/]+)")
		} else {
			regexParts = append(regexParts, regexp.QuoteMeta(part))
		}
	}
	re := regexp.MustCompile("^" + strings.Join(regexParts, "/") + "$")
	return re, paramNames
}

// AddRoute registers a route; compiles the pattern once.
func (r *Router) AddRoute(method, pattern string, handler object.Object) {
	if !strings.HasPrefix(pattern, "/") {
		pattern = "/" + pattern
	}
	re, params := compilePattern(pattern)
	r.mu.Lock()
	defer r.mu.Unlock()
	r.routes[method] = append(r.routes[method], Route{
		pattern: pattern, params: params, re: re, handler: handler,
	})
}

// FindRoute returns the first matching route and extracted path params.
func (r *Router) FindRoute(method, path string) (*Route, map[string]string, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for i := range r.routes[method] {
		route := &r.routes[method][i]
		matches := route.re.FindStringSubmatch(path)
		if matches == nil {
			continue
		}
		params := make(map[string]string, len(route.params))
		for j, name := range route.params {
			params[name] = matches[j+1]
		}
		return route, params, true
	}
	return nil, nil, false
}

// AddMiddleware appends a middleware to the router's chain.
func (r *Router) AddMiddleware(handler object.Object) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.middlewares = append(r.middlewares, handler)
}

// MountSubRouter copies all routes from subRouter prefixed with mountPath.
func (r *Router) MountSubRouter(mountPath string, subRouter *Router) {
	if !strings.HasPrefix(mountPath, "/") {
		mountPath = "/" + mountPath
	}
	mountPath = strings.TrimSuffix(mountPath, "/")

	subRouter.mu.RLock()
	defer subRouter.mu.RUnlock()
	r.mu.Lock()
	defer r.mu.Unlock()

	for method, routes := range subRouter.routes {
		for _, route := range routes {
			fullPattern := mountPath + route.pattern
			re, params := compilePattern(fullPattern)
			r.routes[method] = append(r.routes[method], Route{
				pattern: fullPattern, params: params, re: re, handler: route.handler,
			})
		}
	}
}

// callHandler invokes either a *object.Function (via EvalFunc) or *object.Builtin directly.
func callHandler(handler object.Object, args []object.Object) object.Object {
	switch h := handler.(type) {
	case *object.Function:
		if EvalFunc != nil {
			return EvalFunc(h, args)
		}
	case *object.Builtin:
		return h.Fn(args...)
	}
	return object.NULL
}

// ServeHTTP implements http.Handler with the full Express.js-style pipeline.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	start := time.Now()

	// 1. Read body with optional size limit
	var bodyReader io.Reader = req.Body
	if r.maxBodyBytes > 0 {
		bodyReader = io.LimitReader(req.Body, r.maxBodyBytes)
	}
	body, _ := io.ReadAll(bodyReader)

	// 2. CORS headers (before any WriteHeader)
	if r.corsEnabled {
		setCORSHeaders(w, r.corsOptions)
	}

	// 3. OPTIONS preflight
	if req.Method == "OPTIONS" && r.corsEnabled {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// 4. Static file routes
	for _, fr := range r.fileRoutes {
		if strings.HasPrefix(req.URL.Path, fr.prefix) {
			fr.handler.ServeHTTP(w, req)
			return
		}
	}

	// 5. Rate limiting
	if r.rateLimiter != nil {
		ip := getClientIP(req)
		if !r.rateLimiter.Allow(ip) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprint(w, `{"error":"গতি সীমা অতিক্রান্ত / Rate limit exceeded"}`)
			return
		}
	}

	// 6. Route matching
	route, params, ok := r.FindRoute(req.Method, req.URL.Path)
	if !ok {
		http.NotFound(w, req)
		return
	}

	// 7. Build BanglaCode request / response maps
	reqMap := buildRequestMap(req, body, params)
	resMap := buildResponseMap()

	// 8. Middleware chain + route handler with optional timeout
	middlewares := r.middlewares
	var execute func(idx int)
	execute = func(idx int) {
		if idx < len(middlewares) {
			mw := middlewares[idx]
			nextFn := &object.Builtin{
				Fn: func(_ ...object.Object) object.Object {
					execute(idx + 1)
					return object.NULL
				},
			}
			callHandler(mw, []object.Object{reqMap, resMap, nextFn})
		} else {
			result := callHandler(route.handler, []object.Object{reqMap, resMap})
			if result != nil && result.Type() == object.ERROR_OBJ && r.errorHandler != nil {
				callHandler(r.errorHandler, []object.Object{result, reqMap, resMap})
			}
		}
	}

	if r.timeout > 0 {
		ctx, cancel := context.WithTimeout(req.Context(), r.timeout)
		defer cancel()
		done := make(chan struct{})
		go func() { execute(0); close(done) }()
		select {
		case <-done:
		case <-ctx.Done():
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusGatewayTimeout)
			fmt.Fprint(w, `{"error":"Request timeout"}`)
			return
		}
	} else {
		execute(0)
	}

	// 9. Logging
	if r.logEnabled {
		status := 200
		if s, ok2 := resMap.Pairs["status"].(*object.Number); ok2 {
			status = int(s.Value)
		}
		fmt.Printf("🔵 [BanglaCode] %s %s → %d (%v)\n", req.Method, req.URL.Path, status, time.Since(start))
	}

	// 10. Write HTTP response (gzip if requested and enabled)
	useGzip := r.gzipEnabled && strings.Contains(req.Header.Get("Accept-Encoding"), "gzip")
	writeHTTPResponse(w, resMap, useGzip)
}

// setCORSHeaders writes the CORS headers to the response.
func setCORSHeaders(w http.ResponseWriter, opts CORSOptions) {
	w.Header().Set("Access-Control-Allow-Origin", opts.Origin)
	w.Header().Set("Access-Control-Allow-Methods", opts.Methods)
	w.Header().Set("Access-Control-Allow-Headers", opts.Headers)
	if opts.MaxAge != "" {
		w.Header().Set("Access-Control-Max-Age", opts.MaxAge)
	}
}

// getClientIP extracts the real client IP, respecting proxy headers.
func getClientIP(req *http.Request) string {
	if xff := req.Header.Get("X-Forwarded-For"); xff != "" {
		return strings.TrimSpace(strings.SplitN(xff, ",", 2)[0])
	}
	if xri := req.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return req.RemoteAddr
	}
	return ip
}

// buildRequestMap constructs the BanglaCode req object with all parsed fields.
func buildRequestMap(req *http.Request, body []byte, params map[string]string) *object.Map {
	m := &object.Map{Pairs: make(map[string]object.Object, 10)}
	m.Pairs["method"] = &object.String{Value: req.Method}
	m.Pairs["path"] = &object.String{Value: req.URL.Path}
	m.Pairs["ip"] = &object.String{Value: getClientIP(req)}

	// Headers
	headersMap := &object.Map{Pairs: make(map[string]object.Object, len(req.Header))}
	for k, v := range req.Header {
		if len(v) > 0 {
			headersMap.Pairs[k] = &object.String{Value: v[0]}
		}
	}
	m.Pairs["headers"] = headersMap

	// Raw body
	m.Pairs["body"] = &object.String{Value: string(body)}

	// Auto JSON parse
	ct := req.Header.Get("Content-Type")
	if strings.Contains(ct, "application/json") && len(body) > 0 {
		m.Pairs["json"] = parseJSON(string(body))
	} else {
		m.Pairs["json"] = object.NULL
	}

	// URL-encoded form data
	if strings.Contains(ct, "application/x-www-form-urlencoded") {
		if formVals, err := url.ParseQuery(string(body)); err == nil {
			formMap := &object.Map{Pairs: make(map[string]object.Object, len(formVals))}
			for k, v := range formVals {
				if len(v) > 0 {
					formMap.Pairs[k] = &object.String{Value: v[0]}
				}
			}
			m.Pairs["form"] = formMap
		} else {
			m.Pairs["form"] = object.NULL
		}
	} else {
		m.Pairs["form"] = object.NULL
	}

	// Path params
	paramsMap := &object.Map{Pairs: make(map[string]object.Object, len(params))}
	for k, v := range params {
		paramsMap.Pairs[k] = &object.String{Value: v}
	}
	m.Pairs["params"] = paramsMap

	// Query string (parsed MAP + raw)
	queryMap := &object.Map{Pairs: make(map[string]object.Object)}
	for k, v := range req.URL.Query() {
		if len(v) > 0 {
			queryMap.Pairs[k] = &object.String{Value: v[0]}
		}
	}
	m.Pairs["query"] = queryMap
	m.Pairs["query_raw"] = &object.String{Value: req.URL.RawQuery}

	// Cookies
	kukisMap := &object.Map{Pairs: make(map[string]object.Object)}
	for _, c := range req.Cookies() {
		kukisMap.Pairs[c.Name] = &object.String{Value: c.Value}
	}
	m.Pairs["kukis"] = kukisMap

	return m
}

// buildResponseMap creates the initial BanglaCode res object.
func buildResponseMap() *object.Map {
	m := &object.Map{Pairs: make(map[string]object.Object, 3)}
	m.Pairs["status"] = &object.Number{Value: 200}
	m.Pairs["body"] = &object.String{Value: ""}
	m.Pairs["headers"] = &object.Map{Pairs: make(map[string]object.Object)}
	return m
}

// writeHTTPResponse writes the BanglaCode res map to the HTTP response.
// Headers are always set before WriteHeader to comply with HTTP/1.1.
func writeHTTPResponse(w http.ResponseWriter, resMap *object.Map, useGzip bool) {
	// Set custom headers first
	if h, ok := resMap.Pairs["headers"].(*object.Map); ok {
		for k, v := range h.Pairs {
			w.Header().Set(k, v.Inspect())
		}
	}
	status := 200
	if s, ok := resMap.Pairs["status"].(*object.Number); ok {
		status = int(s.Value)
	}
	body := ""
	if b, ok := resMap.Pairs["body"]; ok {
		body = b.Inspect()
	}

	if useGzip {
		w.Header().Set("Content-Encoding", "gzip")
		w.WriteHeader(status)
		gz := gzip.NewWriter(w)
		defer gz.Close()
		fmt.Fprint(gz, body)
	} else {
		w.WriteHeader(status)
		fmt.Fprint(w, body)
	}
}

// Global router registry — maps pointer string → *Router.
var (
	routerRegistry   = make(map[string]*Router)
	routerRegistryMu sync.RWMutex
)

func registerRouter(r *Router) {
	routerRegistryMu.Lock()
	defer routerRegistryMu.Unlock()
	routerRegistry[fmt.Sprintf("%p", r)] = r
}

func getRouter(id string) (*Router, bool) {
	routerRegistryMu.RLock()
	defer routerRegistryMu.RUnlock()
	r, ok := routerRegistry[id]
	return r, ok
}

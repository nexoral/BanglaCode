package url

import (
	"BanglaCode/src/object"
	"net/url"
	"strings"
)

// Builtins contains all URL-related built-in functions
var Builtins = map[string]*object.Builtin{
	"url_parse": {
		Fn: urlParse,
	},
	"url_query_params": {
		Fn: urlQueryParams,
	},
	"url_query_get": {
		Fn: urlQueryGet,
	},
	"url_query_set": {
		Fn: urlQuerySet,
	},
	"url_query_append": {
		Fn: urlQueryAppend,
	},
	"url_query_delete": {
		Fn: urlQueryDelete,
	},
	"url_query_has": {
		Fn: urlQueryHas,
	},
	"url_query_keys": {
		Fn: urlQueryKeys,
	},
	"url_query_values": {
		Fn: urlQueryValues,
	},
	"url_query_toString": {
		Fn: urlQueryToString,
	},
}

// urlParse parses a URL string into a URL object
// Usage: url_parse("https://example.com:8080/path?query=value#fragment")
func urlParse(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "url_parse() takes exactly 1 argument"}
	}

	urlStr, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "url_parse() argument must be a string"}
	}

	// Parse URL using Go's net/url package
	parsedURL, err := url.Parse(urlStr.Value)
	if err != nil {
		return &object.Error{Message: "Invalid URL: " + err.Error()}
	}

	// Extract components
	protocol := parsedURL.Scheme
	if protocol != "" {
		protocol += ":"
	}

	hostname := parsedURL.Hostname()
	port := parsedURL.Port()
	host := parsedURL.Host

	username := parsedURL.User.Username()
	password, _ := parsedURL.User.Password()

	pathname := parsedURL.Path
	if pathname == "" {
		pathname = "/"
	}

	search := parsedURL.RawQuery
	if search != "" {
		search = "?" + search
	}

	hash := parsedURL.Fragment
	if hash != "" {
		hash = "#" + hash
	}

	// Construct origin (protocol + hostname + port)
	origin := protocol + "//" + hostname
	if port != "" && port != "80" && port != "443" {
		origin += ":" + port
	}

	return &object.URL{
		Href:     urlStr.Value,
		Protocol: protocol,
		Username: username,
		Password: password,
		Hostname: hostname,
		Port:     port,
		Host:     host,
		Pathname: pathname,
		Search:   search,
		Hash:     hash,
		Origin:   origin,
	}
}

// urlQueryParams creates a URLSearchParams object from a query string or URL
// Usage: url_query_params("key1=value1&key2=value2")
// Usage: url_query_params(urlObject)
func urlQueryParams(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "url_query_params() takes exactly 1 argument"}
	}

	var queryString string

	switch arg := args[0].(type) {
	case *object.String:
		queryString = arg.Value
		// Remove leading "?" if present
		queryString = strings.TrimPrefix(queryString, "?")

	case *object.URL:
		queryString = strings.TrimPrefix(arg.Search, "?")

	default:
		return &object.Error{Message: "url_query_params() argument must be a string or URL object"}
	}

	// Parse query string
	values, err := url.ParseQuery(queryString)
	if err != nil {
		return &object.Error{Message: "Invalid query string: " + err.Error()}
	}

	// Convert to our URLSearchParams format
	params := make(map[string][]string)
	for key, vals := range values {
		params[key] = vals
	}

	return &object.URLSearchParams{
		Params: params,
	}
}

// urlQueryGet gets the first value for a given key
// Usage: url_query_get(params, "key")
func urlQueryGet(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "url_query_get() takes exactly 2 arguments"}
	}

	params, ok := args[0].(*object.URLSearchParams)
	if !ok {
		return &object.Error{Message: "url_query_get() first argument must be URLSearchParams"}
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "url_query_get() second argument must be a string"}
	}

	values, exists := params.Params[key.Value]
	if !exists || len(values) == 0 {
		return object.NULL
	}

	return &object.String{Value: values[0]}
}

// urlQuerySet sets a key to a single value, replacing existing values
// Usage: url_query_set(params, "key", "value")
func urlQuerySet(args ...object.Object) object.Object {
	if len(args) != 3 {
		return &object.Error{Message: "url_query_set() takes exactly 3 arguments"}
	}

	params, ok := args[0].(*object.URLSearchParams)
	if !ok {
		return &object.Error{Message: "url_query_set() first argument must be URLSearchParams"}
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "url_query_set() second argument must be a string"}
	}

	value, ok := args[2].(*object.String)
	if !ok {
		return &object.Error{Message: "url_query_set() third argument must be a string"}
	}

	// Replace all values with a single value
	params.Params[key.Value] = []string{value.Value}

	return params
}

// urlQueryAppend appends a value to a key (allowing multiple values)
// Usage: url_query_append(params, "key", "value")
func urlQueryAppend(args ...object.Object) object.Object {
	if len(args) != 3 {
		return &object.Error{Message: "url_query_append() takes exactly 3 arguments"}
	}

	params, ok := args[0].(*object.URLSearchParams)
	if !ok {
		return &object.Error{Message: "url_query_append() first argument must be URLSearchParams"}
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "url_query_append() second argument must be a string"}
	}

	value, ok := args[2].(*object.String)
	if !ok {
		return &object.Error{Message: "url_query_append() third argument must be a string"}
	}

	// Append value to existing values
	params.Params[key.Value] = append(params.Params[key.Value], value.Value)

	return params
}

// urlQueryDelete removes all values for a given key
// Usage: url_query_delete(params, "key")
func urlQueryDelete(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "url_query_delete() takes exactly 2 arguments"}
	}

	params, ok := args[0].(*object.URLSearchParams)
	if !ok {
		return &object.Error{Message: "url_query_delete() first argument must be URLSearchParams"}
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "url_query_delete() second argument must be a string"}
	}

	delete(params.Params, key.Value)

	return params
}

// urlQueryHas checks if a key exists
// Usage: url_query_has(params, "key")
func urlQueryHas(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "url_query_has() takes exactly 2 arguments"}
	}

	params, ok := args[0].(*object.URLSearchParams)
	if !ok {
		return &object.Error{Message: "url_query_has() first argument must be URLSearchParams"}
	}

	key, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "url_query_has() second argument must be a string"}
	}

	_, exists := params.Params[key.Value]
	return object.NativeBoolToBooleanObject(exists)
}

// urlQueryKeys returns all keys as an array
// Usage: url_query_keys(params)
func urlQueryKeys(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "url_query_keys() takes exactly 1 argument"}
	}

	params, ok := args[0].(*object.URLSearchParams)
	if !ok {
		return &object.Error{Message: "url_query_keys() argument must be URLSearchParams"}
	}

	keys := []object.Object{}
	for key := range params.Params {
		keys = append(keys, &object.String{Value: key})
	}

	return &object.Array{Elements: keys}
}

// urlQueryValues returns all values as an array (flattened)
// Usage: url_query_values(params)
func urlQueryValues(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "url_query_values() takes exactly 1 argument"}
	}

	params, ok := args[0].(*object.URLSearchParams)
	if !ok {
		return &object.Error{Message: "url_query_values() argument must be URLSearchParams"}
	}

	values := []object.Object{}
	for _, vals := range params.Params {
		for _, val := range vals {
			values = append(values, &object.String{Value: val})
		}
	}

	return &object.Array{Elements: values}
}

// urlQueryToString converts URLSearchParams to a query string
// Usage: url_query_toString(params)
func urlQueryToString(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "url_query_toString() takes exactly 1 argument"}
	}

	params, ok := args[0].(*object.URLSearchParams)
	if !ok {
		return &object.Error{Message: "url_query_toString() argument must be URLSearchParams"}
	}

	// Build query string
	values := url.Values{}
	for key, vals := range params.Params {
		for _, val := range vals {
			values.Add(key, val)
		}
	}

	return &object.String{Value: values.Encode()}
}

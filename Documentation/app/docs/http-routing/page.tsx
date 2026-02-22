import { Metadata } from "next";

export const metadata: Metadata = {
  title: "HTTP Routing - BanglaCode Documentation",
  description:
    "Learn Express.js-style modular routing in BanglaCode. Create routers, define routes, mount sub-routers, and build modular REST APIs with pure Banglish keywords.",
};

export default function HTTPRoutingPage() {
  return (
    <div className="prose prose-slate dark:prose-invert max-w-none">
      <h1>HTTP Routing</h1>
      <p className="lead">
        Build modular, Express.js-style REST APIs with BanglaCode's powerful routing system using pure Banglish keywords.
      </p>

      <h2>Overview</h2>
      <p>
        BanglaCode provides an Express.js-inspired router system that allows you to organize your HTTP endpoints into modular, reusable components. This makes it easy to build complex web applications with clean, maintainable code structure.
      </p>

      <h3>Key Features</h3>
      <ul>
        <li>✅ <strong>Modular Routing</strong> - Organize routes in separate files</li>
        <li>✅ <strong>All HTTP Methods</strong> - GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS support</li>
        <li>✅ <strong>Router Mounting</strong> - Mount sub-routers on paths</li>
        <li>✅ <strong>Method Chaining</strong> - Define multiple routes fluently</li>
        <li>✅ <strong>Pure Banglish</strong> - Bengali keywords throughout</li>
      </ul>

      <hr />

      <h2>Basic Router</h2>
      <p>Create a basic HTTP router with multiple endpoints:</p>

      <h3>router_banao() - Create Router</h3>
      <pre><code className="language-banglacode">{`// Create new router
dhoro app = router_banao();

// Define GET route (আনা - fetch)
app.ana("/", kaj(req, res) {
    uttor(res, "Welcome to BanglaCode!");
});

// Define POST route (পাঠানো - send)
app.pathano("/submit", kaj(req, res) {
    json_uttor(res, {"message": "Data received"}, 201);
});

// Start server with router
server_chalu(3000, app);`}</code></pre>

      <hr />

      <h2>HTTP Methods</h2>
      <p>All 7 standard HTTP methods are supported with pure Banglish keywords:</p>

      <h3>ana() - GET (আনা - fetch)</h3>
      <pre><code className="language-banglacode">{`app.ana("/users", kaj(req, res) {
    dhoro users = [
        {"id": 1, "name": "Ankan"},
        {"id": 2, "name": "Rahim"}
    ];
    json_uttor(res, {"users": users});
});`}</code></pre>

      <h3>pathano() - POST (পাঠানো - send)</h3>
      <pre><code className="language-banglacode">{`app.pathano("/users", kaj(req, res) {
    // In real app, would parse req["body"]
    dhoro newUser = {"id": 3, "name": "New User"};
    json_uttor(res, {"user": newUser}, 201);
});`}</code></pre>

      <h3>bodlano() - PUT (বদলানো - update/change)</h3>
      <pre><code className="language-banglacode">{`app.bodlano("/users/update", kaj(req, res) {
    json_uttor(res, {
        "message": "User updated successfully"
    });
});`}</code></pre>

      <h3>mujhe_felo() - DELETE (মুছে ফেলো - remove)</h3>
      <pre><code className="language-banglacode">{`app.mujhe_felo("/users/delete", kaj(req, res) {
    uttor(res, "User deleted", 204);
});`}</code></pre>

      <h3>songshodhon() - PATCH (সংশোধন - modify)</h3>
      <pre><code className="language-banglacode">{`app.songshodhon("/users/modify", kaj(req, res) {
    json_uttor(res, {"message": "User modified"});
});`}</code></pre>

      <h3>matha() - HEAD (মাথা - retrieve headers)</h3>
      <pre><code className="language-banglacode">{`app.matha("/users", kaj(req, res) {
    res["headers"] = {"Content-Type": "application/json"};
    res["status"] = 200;
});`}</code></pre>

      <h3>nirdharon() - OPTIONS (নির্ধারণ - determine options)</h3>
      <pre><code className="language-banglacode">{`app.nirdharon("/users", kaj(req, res) {
    res["headers"] = {"Allow": "GET, POST, PUT, DELETE, PATCH"};
    res["status"] = 200;
});`}</code></pre>

      <hr />

      <h2>Modular Routing</h2>
      <p>
        The real power of BanglaCode's routing system is the ability to organize routes into separate modules, just like Express.js.
      </p>

      <h3>Step 1: Create Route Modules</h3>

      <h4>File: routes_auth.bang</h4>
      <pre><code className="language-banglacode">{`// Authentication Routes Module
dhoro authRouter = router_banao();

authRouter.pathano("/login", kaj(req, res) {
    json_uttor(res, {
        "token": "jwt_token_here",
        "user": {"name": "Ankan"}
    });
});

authRouter.pathano("/register", kaj(req, res) {
    json_uttor(res, {
        "message": "User registered successfully"
    }, 201);
});

authRouter.pathano("/logout", kaj(req, res) {
    json_uttor(res, {"message": "Logged out"});
});

// Export router
pathao authRouter;`}</code></pre>

      <h4>File: routes_users.bang</h4>
      <pre><code className="language-banglacode">{`// User Management Routes Module
dhoro usersRouter = router_banao();

dhoro users = [
    {"id": 1, "name": "Ankan"},
    {"id": 2, "name": "Rahim"}
];

usersRouter.ana("/", kaj(req, res) {
    json_uttor(res, {"users": users});
});

usersRouter.pathano("/", kaj(req, res) {
    json_uttor(res, {
        "message": "User created"
    }, 201);
});

usersRouter.mujhe_felo("/delete", kaj(req, res) {
    json_uttor(res, {"message": "User deleted"});
});

// Export router
pathao usersRouter;`}</code></pre>

      <h3>Step 2: Mount Routers in Main App</h3>

      <h4>File: main.bang</h4>
      <pre><code className="language-banglacode">{`// Import route modules
dhoro authRouter = ano("routes_auth.bang");
dhoro usersRouter = ano("routes_users.bang");

// Create main app
dhoro app = router_banao();

// Main routes
app.ana("/", kaj(req, res) {
    json_uttor(res, {
        "message": "Welcome to API",
        "endpoints": {
            "auth": "/api/auth",
            "users": "/api/users"
        }
    });
});

// Mount sub-routers using bebohar (ব্যবহার - use/mount)
app.bebohar("/api/auth", authRouter);
app.bebohar("/api/users", usersRouter);

// Start server
dekho("Server starting on port 3000...");
server_chalu(3000, app);`}</code></pre>

      <h3>Available Routes</h3>
      <p>After mounting, these routes are available:</p>
      <ul>
        <li><code>GET /</code> - Main welcome message</li>
        <li><code>POST /api/auth/login</code> - User login</li>
        <li><code>POST /api/auth/register</code> - User registration</li>
        <li><code>POST /api/auth/logout</code> - User logout</li>
        <li><code>GET /api/users</code> - Get all users</li>
        <li><code>POST /api/users</code> - Create new user</li>
        <li><code>DELETE /api/users/delete</code> - Delete user</li>
      </ul>

      <hr />

      <h2>Method Chaining</h2>
      <p>Define multiple routes fluently with method chaining:</p>

      <pre><code className="language-banglacode">{`dhoro app = router_banao();

app
    .ana("/", kaj(req, res) {
        uttor(res, "Home");
    })
    .ana("/about", kaj(req, res) {
        uttor(res, "About");
    })
    .pathano("/submit", kaj(req, res) {
        uttor(res, "Submitted");
    })
    .bodlano("/update", kaj(req, res) {
        uttor(res, "Updated");
    });`}</code></pre>

      <hr />

      <h2>Request & Response Objects</h2>

      <h3>Request Object (req)</h3>
      <p>The request object contains information about the incoming HTTP request:</p>

      <pre><code className="language-banglacode">{`app.ana("/inspect", kaj(req, res) {
    dhoro method = req["method"];      // "GET"
    dhoro path = req["path"];          // "/inspect"
    dhoro query = req["query"];        // Query string
    dhoro headers = req["headers"];    // Request headers
    dhoro body = req["body"];          // Request body

    json_uttor(res, {
        "method": method,
        "path": path,
        "query": query
    });
});`}</code></pre>

      <h3>Response Object (res)</h3>
      <p>Use helper functions to send responses:</p>

      <h4>uttor() - Simple Response</h4>
      <pre><code className="language-banglacode">{`// Basic response
uttor(res, "Hello, World!");

// With status code
uttor(res, "Not Found", 404);

// With content type
uttor(res, "Hello", 200, "text/plain");`}</code></pre>

      <h4>json_uttor() - JSON Response</h4>
      <pre><code className="language-banglacode">{`// JSON response (auto sets content-type)
json_uttor(res, {"message": "Success"});

// With status code
json_uttor(res, {"error": "Not Found"}, 404);`}</code></pre>

      <hr />

      <h2>Complete Example</h2>
      <p>A full modular REST API with authentication and users:</p>

      <pre><code className="language-banglacode">{`// File: main.bang - Complete Modular API

// Import modules
dhoro authRouter = ano("routes_auth.bang");
dhoro usersRouter = ano("routes_users.bang");

// Create app
dhoro app = router_banao();

// Health check
app.ana("/health", kaj(req, res) {
    json_uttor(res, {
        "status": "healthy",
        "uptime": somoy()
    });
});

// API info
app.ana("/api", kaj(req, res) {
    json_uttor(res, {
        "version": "1.0.0",
        "endpoints": {
            "auth": "/api/auth (login, register)",
            "users": "/api/users (CRUD)"
        }
    });
});

// Mount routers
app.bebohar("/api/auth", authRouter);
app.bebohar("/api/users", usersRouter);

// 404 handler
app.ana("/*", kaj(req, res) {
    json_uttor(res, {
        "error": "Route not found"
    }, 404);
});

// Start server
dekho("🚀 Server starting...");
dekho("📍 http://localhost:3000");
server_chalu(3000, app);`}</code></pre>

      <hr />

      <h2>Best Practices</h2>

      <h3>1. Organize by Feature</h3>
      <p>Group related routes in separate modules:</p>
      <ul>
        <li><code>routes_auth.bang</code> - Authentication</li>
        <li><code>routes_users.bang</code> - User management</li>
        <li><code>routes_products.bang</code> - Product catalog</li>
        <li><code>routes_orders.bang</code> - Order processing</li>
      </ul>

      <h3>2. Use Consistent Naming</h3>
      <pre><code className="language-banglacode">{`// Good - descriptive router names
dhoro authRouter = router_banao();
dhoro usersRouter = router_banao();

// Bad - unclear names
dhoro r1 = router_banao();
dhoro temp = router_banao();`}</code></pre>

      <h3>3. Mount on Logical Paths</h3>
      <pre><code className="language-banglacode">{`// Good - grouped by API version and resource
app.bebohar("/api/v1/auth", authRouter);
app.bebohar("/api/v1/users", usersRouter);

// Bad - inconsistent paths
app.bebohar("/auth", authRouter);
app.bebohar("/api/users", usersRouter);`}</code></pre>

      <h3>4. Export from Modules</h3>
      <pre><code className="language-banglacode">{`// At end of route module file
pathao authRouter;  // Always export your router`}</code></pre>

      <hr />

      <h2>API Reference</h2>

      <h3>router_banao()</h3>
      <p><strong>Returns:</strong> Router object</p>
      <p><strong>Description:</strong> Creates a new Express-style router instance</p>

      <h3>router.ana(path, handler)</h3>
      <p><strong>Method:</strong> GET (আনা - fetch)</p>
      <ul>
        <li><code>path</code> (String) - Route path</li>
        <li><code>handler</code> (Function) - Request handler</li>
        <li><strong>Returns:</strong> Router (for chaining)</li>
      </ul>

      <h3>router.pathano(path, handler)</h3>
      <p><strong>Method:</strong> POST (পাঠানো - send)</p>
      <ul>
        <li><code>path</code> (String) - Route path</li>
        <li><code>handler</code> (Function) - Request handler</li>
        <li><strong>Returns:</strong> Router (for chaining)</li>
      </ul>

      <h3>router.bodlano(path, handler)</h3>
      <p><strong>Method:</strong> PUT (বদলানো - update/change)</p>
      <ul>
        <li><code>path</code> (String) - Route path</li>
        <li><code>handler</code> (Function) - Request handler</li>
        <li><strong>Returns:</strong> Router (for chaining)</li>
      </ul>

      <h3>router.mujhe_felo(path, handler)</h3>
      <p><strong>Method:</strong> DELETE (মুছে ফেলো - remove)</p>
      <ul>
        <li><code>path</code> (String) - Route path</li>
        <li><code>handler</code> (Function) - Request handler</li>
        <li><strong>Returns:</strong> Router (for chaining)</li>
      </ul>

      <h3>router.songshodhon(path, handler)</h3>
      <p><strong>Method:</strong> PATCH (সংশোধন - modify)</p>
      <ul>
        <li><code>path</code> (String) - Route path</li>
        <li><code>handler</code> (Function) - Request handler</li>
        <li><strong>Returns:</strong> Router (for chaining)</li>
      </ul>

      <h3>router.matha(path, handler)</h3>
      <p><strong>Method:</strong> HEAD (মাথা - retrieve headers)</p>
      <ul>
        <li><code>path</code> (String) - Route path</li>
        <li><code>handler</code> (Function) - Request handler</li>
        <li><strong>Returns:</strong> Router (for chaining)</li>
      </ul>

      <h3>router.nirdharon(path, handler)</h3>
      <p><strong>Method:</strong> OPTIONS (নির্ধারণ - determine options)</p>
      <ul>
        <li><code>path</code> (String) - Route path</li>
        <li><code>handler</code> (Function) - Request handler</li>
        <li><strong>Returns:</strong> Router (for chaining)</li>
      </ul>

      <h3>router.bebohar(mountPath, subRouter)</h3>
      <p><strong>Method:</strong> Mount sub-router (ব্যবহার - use/mount)</p>
      <ul>
        <li><code>mountPath</code> (String) - Base path for sub-router</li>
        <li><code>subRouter</code> (Router) - Router to mount</li>
        <li><strong>Returns:</strong> Router (for chaining)</li>
      </ul>

      <h3>server_chalu(port, handler)</h3>
      <ul>
        <li><code>port</code> (Number) - Port to listen on</li>
        <li><code>handler</code> (Function OR Router) - Request handler or router</li>
      </ul>

      <hr />

      <h2>Related Topics</h2>
      <ul>
        <li><a href="/docs/http-server">HTTP Server Basics</a></li>
        <li><a href="/docs/modules">Modules & Imports</a></li>
        <li><a href="/docs/functions">Functions</a></li>
        <li><a href="/docs/builtins">Built-in Functions</a></li>
      </ul>
    </div>
  );
}

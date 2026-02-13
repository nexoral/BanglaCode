import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function HttpServer() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Advanced
        </span>
      </div>

      <h1>HTTP Server</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        BanglaCode can create HTTP servers using the <code>server_chalu</code> (start server)
        function. This allows you to build web applications and APIs.
      </p>

      <h2>Starting a Server</h2>

      <p>
        Use <code>server_chalu</code> to start an HTTP server on a specified port:
      </p>

      <CodeBlock
        filename="server.bang"
        code={`// Basic HTTP server
server_chalu(8080, kaj(req, res) {
    uttor(res, "Namaskar from BanglaCode!");
});

dekho("Server running on port 8080");`}
      />

      <p>Run the server and visit <code>http://localhost:8080</code> in your browser.</p>

      <h2>Request Object</h2>

      <p>The request object contains information about the incoming HTTP request:</p>

      <div className="overflow-x-auto my-6">
        <table>
          <thead>
            <tr>
              <th>Property</th>
              <th>Description</th>
              <th>Example</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>method</code></td><td>HTTP method</td><td>&quot;GET&quot;, &quot;POST&quot;</td></tr>
            <tr><td><code>path</code></td><td>URL path</td><td>&quot;/users&quot;</td></tr>
            <tr><td><code>query</code></td><td>Query string</td><td>&quot;id=123&quot;</td></tr>
            <tr><td><code>headers</code></td><td>Request headers</td><td>Map of headers</td></tr>
            <tr><td><code>body</code></td><td>Request body</td><td>String or parsed JSON</td></tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`server_chalu(8080, kaj(req, res) {
    dekho("Method:", req.method);
    dekho("Path:", req.path);
    dekho("Query:", req.query);
    dekho("Headers:", req.headers);

    uttor(res, "Request received!");
});`}
      />

      <h2>Sending Responses</h2>

      <h3>Text Response (uttor)</h3>

      <CodeBlock
        code={`// Basic text response
server_chalu(8080, kaj(req, res) {
    uttor(res, "Hello World");
});

// With status code
server_chalu(8080, kaj(req, res) {
    uttor(res, "Not Found", 404);
});

// With content type
server_chalu(8080, kaj(req, res) {
    uttor(res, "<h1>Hello</h1>", 200, "text/html");
});`}
      />

      <h3>JSON Response (json_uttor)</h3>

      <CodeBlock
        code={`server_chalu(8080, kaj(req, res) {
    dhoro data = {
        status: "success",
        message: "User created",
        user: {
            id: 1,
            naam: "Rahim"
        }
    };

    json_uttor(res, data);
});

// With custom status code
server_chalu(8080, kaj(req, res) {
    json_uttor(res, {error: "Not found"}, 404);
});`}
      />

      <h2>Routing</h2>

      <CodeBlock
        code={`server_chalu(8080, kaj(req, res) {
    // Simple routing based on path
    jodi (req.path == "/") {
        uttor(res, "Home Page", 200, "text/html");
    } nahole jodi (req.path == "/about") {
        uttor(res, "About Page", 200, "text/html");
    } nahole jodi (req.path == "/api/users") {
        json_uttor(res, {users: ["Rahim", "Karim"]});
    } nahole {
        uttor(res, "404 Not Found", 404);
    }
});`}
      />

      <h3>Method-based Routing</h3>

      <CodeBlock
        code={`server_chalu(8080, kaj(req, res) {
    jodi (req.path == "/api/users") {
        jodi (req.method == "GET") {
            // List users
            json_uttor(res, {users: getAllUsers()});

        } nahole jodi (req.method == "POST") {
            // Create user
            dhoro body = json_poro(req.body);
            dhoro user = createUser(body);
            json_uttor(res, user, 201);

        } nahole {
            json_uttor(res, {error: "Method not allowed"}, 405);
        }
    } nahole {
        json_uttor(res, {error: "Not found"}, 404);
    }
});`}
      />

      <h2>Practical Examples</h2>

      <h3>Simple API Server</h3>

      <CodeBlock
        filename="api_server.bang"
        code={`// In-memory data store
dhoro users = [];
dhoro nextId = 1;

kaj findUser(id) {
    ghuriye (dhoro i = 0; i < dorghyo(users); i = i + 1) {
        jodi (users[i].id == id) {
            ferao users[i];
        }
    }
    ferao khali;
}

server_chalu(8080, kaj(req, res) {
    // CORS headers for API
    res.headers["Access-Control-Allow-Origin"] = "*";

    // Routes
    jodi (req.path == "/api/users" ebong req.method == "GET") {
        // List all users
        json_uttor(res, {
            success: sotti,
            data: users,
            count: dorghyo(users)
        });

    } nahole jodi (req.path == "/api/users" ebong req.method == "POST") {
        // Create user
        chesta {
            dhoro body = json_poro(req.body);

            dhoro user = {
                id: nextId,
                naam: body.naam,
                email: body.email,
                createdAt: somoy()
            };

            dhokao(users, user);
            nextId = nextId + 1;

            json_uttor(res, {success: sotti, data: user}, 201);
        } dhoro_bhul (e) {
            json_uttor(res, {success: mittha, error: "Invalid data"}, 400);
        }

    } nahole {
        json_uttor(res, {success: mittha, error: "Not found"}, 404);
    }
});

dekho("API Server running on http://localhost:8080");`}
      />

      <h3>HTML Web Server</h3>

      <CodeBlock
        filename="web_server.bang"
        code={`kaj renderHTML(title, content) {
    ferao "<!DOCTYPE html>
<html>
<head>
    <title>" + title + "</title>
    <style>
        body { font-family: sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
        h1 { color: #7c3aed; }
    </style>
</head>
<body>
    " + content + "
</body>
</html>";
}

server_chalu(8080, kaj(req, res) {
    jodi (req.path == "/") {
        dhoro html = renderHTML("Home", "
            <h1>Namaskar!</h1>
            <p>Welcome to BanglaCode Web Server</p>
            <ul>
                <li><a href='/about'>About</a></li>
                <li><a href='/contact'>Contact</a></li>
            </ul>
        ");
        uttor(res, html, 200, "text/html");

    } nahole jodi (req.path == "/about") {
        dhoro html = renderHTML("About", "
            <h1>About Us</h1>
            <p>BanglaCode is a Bengali programming language.</p>
            <a href='/'>Back to Home</a>
        ");
        uttor(res, html, 200, "text/html");

    } nahole {
        dhoro html = renderHTML("404", "<h1>Page Not Found</h1>");
        uttor(res, html, 404, "text/html");
    }
});

dekho("Web server running on http://localhost:8080");`}
      />

      <h3>Form Handler</h3>

      <CodeBlock
        code={`server_chalu(8080, kaj(req, res) {
    jodi (req.path == "/" ebong req.method == "GET") {
        dhoro form = "<!DOCTYPE html>
<html>
<body>
    <h1>Contact Form</h1>
    <form method='POST' action='/submit'>
        <p>Name: <input name='name'></p>
        <p>Email: <input name='email' type='email'></p>
        <p>Message: <textarea name='message'></textarea></p>
        <button type='submit'>Send</button>
    </form>
</body>
</html>";
        uttor(res, form, 200, "text/html");

    } nahole jodi (req.path == "/submit" ebong req.method == "POST") {
        // Process form data
        dekho("Form submitted:", req.body);

        dhoro response = "<!DOCTYPE html>
<html>
<body>
    <h1>Thank You!</h1>
    <p>Your message has been received.</p>
    <a href='/'>Back</a>
</body>
</html>";
        uttor(res, response, 200, "text/html");
    }
});`}
      />

      <h2>HTTP Client (anun)</h2>

      <p>
        Use <code>anun</code> (meaning &quot;fetch&quot; or &quot;bring&quot;) to make HTTP requests:
      </p>

      <CodeBlock
        code={`// Simple GET request
dhoro response = anun("https://api.example.com/data");
dekho(response);

// Parse JSON response
dhoro data = json_poro(response.body);
dekho(data);`}
      />

      <h2>Best Practices</h2>

      <ul>
        <li><strong>Always validate input</strong> - Never trust client data</li>
        <li><strong>Use appropriate status codes</strong> - 200 for success, 404 for not found, etc.</li>
        <li><strong>Set correct content types</strong> - Especially for JSON and HTML</li>
        <li><strong>Handle errors gracefully</strong> - Wrap handlers in try-catch</li>
        <li><strong>Use JSON for APIs</strong> - It&apos;s the standard for data exchange</li>
      </ul>

      <DocNavigation currentPath="/docs/http-server" />
    </div>
  );
}

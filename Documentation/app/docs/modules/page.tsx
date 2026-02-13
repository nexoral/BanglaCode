import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function Modules() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Advanced
        </span>
      </div>

      <h1>Modules</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        BanglaCode supports a module system for organizing code into separate files.
        Use <code>pathao</code> (export) and <code>ano</code> (import) to share code between files.
      </p>

      <h2>Exporting (pathao)</h2>

      <p>
        Use the <code>pathao</code> keyword (meaning &quot;send&quot;) to export functions, variables,
        or classes:
      </p>

      <CodeBlock
        filename="math.bang"
        code={`// Export a function
pathao kaj add(a, b) {
    ferao a + b;
}

pathao kaj subtract(a, b) {
    ferao a - b;
}

pathao kaj multiply(a, b) {
    ferao a * b;
}

pathao kaj divide(a, b) {
    jodi (b == 0) {
        felo "Cannot divide by zero";
    }
    ferao a / b;
}

// Export a constant
pathao dhoro PI = 3.14159;
pathao dhoro E = 2.71828;`}
      />

      <h3>Exporting Classes</h3>

      <CodeBlock
        filename="shapes.bang"
        code={`pathao sreni Rectangle {
    shuru(width, height) {
        ei.width = width;
        ei.height = height;
    }

    kaj area() {
        ferao ei.width * ei.height;
    }

    kaj perimeter() {
        ferao 2 * (ei.width + ei.height);
    }
}

pathao sreni Circle {
    shuru(radius) {
        ei.radius = radius;
    }

    kaj area() {
        ferao 3.14159 * ei.radius * ei.radius;
    }

    kaj circumference() {
        ferao 2 * 3.14159 * ei.radius;
    }
}`}
      />

      <h2>Importing (ano)</h2>

      <p>
        Use the <code>ano</code> keyword (meaning &quot;bring&quot;) with <code>hisabe</code>
        (meaning &quot;as&quot;) to import modules with an alias:
      </p>

      <CodeBlock
        filename="main.bang"
        code={`// Import with alias
ano "math.bang" hisabe math;

// Use imported functions
dekho(math.add(5, 3));       // 8
dekho(math.multiply(4, 7));  // 28
dekho(math.PI);              // 3.14159

// Import shapes module
ano "shapes.bang" hisabe shapes;

dhoro rect = notun shapes.Rectangle(10, 5);
dekho("Rectangle area:", rect.area());  // 50

dhoro circle = notun shapes.Circle(7);
dekho("Circle area:", circle.area());  // ~153.94`}
      />

      <h2>Import Paths</h2>

      <CodeBlock
        code={`// Relative imports
ano "utils.bang" hisabe utils;           // Same directory
ano "./helpers/string.bang" hisabe str;  // Subdirectory
ano "../common/config.bang" hisabe cfg;  // Parent directory

// Use imported modules
dhoro result = utils.processData(data);
dhoro formatted = str.capitalize(text);
dekho(cfg.APP_NAME);`}
      />

      <h2>Module Organization</h2>

      <p>
        A well-organized project structure:
      </p>

      <CodeBlock
        language="text"
        showLineNumbers={false}
        code={`project/
├── main.bang
├── config.bang
├── utils/
│   ├── string.bang
│   ├── array.bang
│   └── math.bang
├── models/
│   ├── user.bang
│   └── product.bang
└── services/
    ├── auth.bang
    └── database.bang`}
      />

      <h3>Example: Utils Module</h3>

      <CodeBlock
        filename="utils/string.bang"
        code={`pathao kaj capitalize(str) {
    jodi (dorghyo(str) == 0) {
        ferao str;
    }
    ferao boroHater(str[0]) + chotoHater(angsho(str, 1));
}

pathao kaj repeat(str, times) {
    dhoro result = "";
    ghuriye (dhoro i = 0; i < times; i = i + 1) {
        result = result + str;
    }
    ferao result;
}

pathao kaj truncate(str, maxLen) {
    jodi (dorghyo(str) <= maxLen) {
        ferao str;
    }
    ferao angsho(str, 0, maxLen - 3) + "...";
}`}
      />

      <CodeBlock
        filename="main.bang"
        code={`ano "utils/string.bang" hisabe strUtils;

dekho(strUtils.capitalize("hello"));     // "Hello"
dekho(strUtils.repeat("ab", 3));         // "ababab"
dekho(strUtils.truncate("Hello World", 8));  // "Hello..."`}
      />

      <h2>JSON Imports</h2>

      <p>
        You can import JSON files as data:
      </p>

      <CodeBlock
        filename="config.json"
        language="json"
        code={`{
    "app_name": "MyApp",
    "version": "1.0.0",
    "database": {
        "host": "localhost",
        "port": 5432
    },
    "features": {
        "dark_mode": true,
        "notifications": true
    }
}`}
      />

      <CodeBlock
        filename="main.bang"
        code={`// Import JSON as object
ano "config.json" hisabe config;

dekho("App:", config.app_name);
dekho("Version:", config.version);
dekho("DB Host:", config.database.host);

jodi (config.features.dark_mode) {
    dekho("Dark mode is enabled");
}`}
      />

      <h2>Module Patterns</h2>

      <h3>Service Module</h3>

      <CodeBlock
        filename="services/auth.bang"
        code={`dhoro users = {};
dhoro sessions = {};

pathao kaj register(username, password) {
    jodi (users[username] != khali) {
        felo "User already exists";
    }
    users[username] = {
        password: password,
        createdAt: somoy()
    };
    dekho("User registered:", username);
    ferao sotti;
}

pathao kaj login(username, password) {
    dhoro user = users[username];
    jodi (user == khali) {
        felo "User not found";
    }
    jodi (user.password != password) {
        felo "Invalid password";
    }

    dhoro sessionId = lipi(somoy()) + "_" + username;
    sessions[sessionId] = {
        username: username,
        loginTime: somoy()
    };
    ferao sessionId;
}

pathao kaj isLoggedIn(sessionId) {
    ferao sessions[sessionId] != khali;
}

pathao kaj logout(sessionId) {
    sessions[sessionId] = khali;
    dekho("Logged out");
}`}
      />

      <CodeBlock
        filename="main.bang"
        code={`ano "services/auth.bang" hisabe auth;

// Register and login
auth.register("rahim", "password123");
dhoro session = auth.login("rahim", "password123");

dekho("Logged in:", auth.isLoggedIn(session));  // sotti

auth.logout(session);
dekho("Logged in:", auth.isLoggedIn(session));  // mittha`}
      />

      <h3>Model Module</h3>

      <CodeBlock
        filename="models/user.bang"
        code={`pathao sreni User {
    shuru(id, naam, email) {
        ei.id = id;
        ei.naam = naam;
        ei.email = email;
        ei.createdAt = somoy();
    }

    kaj toJSON() {
        ferao {
            id: ei.id,
            naam: ei.naam,
            email: ei.email,
            createdAt: ei.createdAt
        };
    }

    kaj validate() {
        jodi (dorghyo(ei.naam) < 2) {
            ferao {valid: mittha, error: "Name too short"};
        }
        jodi (khojo(ei.email, "@") < 0) {
            ferao {valid: mittha, error: "Invalid email"};
        }
        ferao {valid: sotti};
    }
}`}
      />

      <h2>Circular Import Protection</h2>

      <p>
        BanglaCode has built-in protection against circular imports. If module A imports B
        and B imports A, the interpreter handles it gracefully:
      </p>

      <CodeBlock
        code={`// The interpreter caches loaded modules to prevent
// infinite loops from circular imports.
// Each module is only executed once, and subsequent
// imports return the cached exports.`}
      />

      <h2>Best Practices</h2>

      <ul>
        <li><strong>One responsibility per module</strong> - Keep modules focused on a single task</li>
        <li><strong>Use meaningful names</strong> - Module names should describe their purpose</li>
        <li><strong>Export only what&apos;s needed</strong> - Keep internal helpers private</li>
        <li><strong>Document exports</strong> - Add comments explaining what each export does</li>
        <li><strong>Avoid circular dependencies</strong> - Structure code to prevent circular imports</li>
      </ul>

      <DocNavigation currentPath="/docs/modules" />
    </div>
  );
}

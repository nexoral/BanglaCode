import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function Maps() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Data Structures
        </span>
      </div>

      <h1>Maps (Objects)</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        Maps in BanglaCode are collections of key-value pairs, similar to objects in JavaScript
        or dictionaries in Python. Keys are strings and values can be any type.
      </p>

      <h2>Creating Maps</h2>

      <CodeBlock
        code={`// Empty map
dhoro empty = {};

// Map with properties
dhoro person = {
    naam: "Rahim Khan",
    boyosh: 25,
    city: "Dhaka"
};

// Properties can be any value type
dhoro data = {
    string: "hello",
    number: 42,
    boolean: sotti,
    array: [1, 2, 3],
    nested: {
        x: 10,
        y: 20
    }
};`}
      />

      <h2>Accessing Properties</h2>

      <CodeBlock
        code={`dhoro person = {
    naam: "Rahim",
    boyosh: 25,
    city: "Dhaka"
};

// Dot notation
dekho(person.naam);   // "Rahim"
dekho(person.boyosh); // 25

// Bracket notation
dekho(person["city"]);  // "Dhaka"

// Dynamic property access
dhoro key = "naam";
dekho(person[key]);  // "Rahim"

// Nested access
dhoro company = {
    name: "TechCorp",
    address: {
        street: "123 Main St",
        city: "Kolkata",
        zip: "700001"
    }
};

dekho(company.address.city);      // "Kolkata"
dekho(company["address"]["zip"]); // "700001"`}
      />

      <h2>Modifying Maps</h2>

      <CodeBlock
        code={`dhoro user = {
    naam: "Rahim",
    email: "rahim@example.com"
};

// Modify existing property
user.naam = "Rahim Khan";
dekho(user.naam);  // "Rahim Khan"

// Add new property
user.phone = "01234567890";
user["address"] = "Dhaka, Bangladesh";

dekho(user);
// {naam: "Rahim Khan", email: "rahim@example.com", phone: "01234567890", address: "Dhaka, Bangladesh"}`}
      />

      <h2>Map Built-in Functions</h2>

      <h3>chabi - Get All Keys</h3>

      <CodeBlock
        code={`dhoro person = {
    naam: "Karim",
    boyosh: 30,
    job: "Engineer",
    city: "Chittagong"
};

dhoro keys = chabi(person);
dekho(keys);  // ["naam", "boyosh", "job", "city"]

// Iterate over keys
ghuriye (dhoro i = 0; i < dorghyo(keys); i = i + 1) {
    dhoro key = keys[i];
    dekho(key + ":", person[key]);
}`}
      />

      <h2>Checking for Properties</h2>

      <CodeBlock
        code={`dhoro config = {
    host: "localhost",
    port: 8080
};

// Check if property exists
dhoro keys = chabi(config);

jodi (ache(keys, "host")) {
    dekho("Host is configured:", config.host);
}

jodi (na ache(keys, "timeout")) {
    dekho("Timeout not configured, using default");
    config.timeout = 30;
}

// Check for null/undefined value
jodi (config.database == khali) {
    dekho("Database not configured");
}`}
      />

      <h2>Nested Maps</h2>

      <CodeBlock
        code={`dhoro organization = {
    name: "Tech Company",
    departments: {
        engineering: {
            head: "Alice",
            employees: 50
        },
        marketing: {
            head: "Bob",
            employees: 20
        },
        hr: {
            head: "Carol",
            employees: 10
        }
    },
    locations: ["Dhaka", "Kolkata", "Mumbai"]
};

// Deep access
dekho(organization.departments.engineering.head);  // "Alice"
dekho(organization.departments.marketing.employees);  // 20

// Modify nested
organization.departments.engineering.employees = 55;

// Add new nested property
organization.departments.sales = {
    head: "David",
    employees: 15
};`}
      />

      <h2>Maps with Methods</h2>

      <CodeBlock
        code={`dhoro calculator = {
    value: 0,

    add: kaj(n) {
        ei.value = ei.value + n;
        ferao ei;
    },

    subtract: kaj(n) {
        ei.value = ei.value - n;
        ferao ei;
    },

    multiply: kaj(n) {
        ei.value = ei.value * n;
        ferao ei;
    },

    getResult: kaj() {
        ferao ei.value;
    },

    reset: kaj() {
        ei.value = 0;
        ferao ei;
    }
};

// Note: 'ei' in map methods refers to the map itself
calculator.add(10).multiply(2).subtract(5);
dekho(calculator.getResult());  // 15`}
      />

      <h2>Iterating Over Maps</h2>

      <CodeBlock
        code={`dhoro scores = {
    Rahim: 85,
    Karim: 92,
    Jamil: 78,
    Nasir: 88
};

dhoro keys = chabi(scores);

// Print all entries
ghuriye (dhoro i = 0; i < dorghyo(keys); i = i + 1) {
    dhoro name = keys[i];
    dekho(name + ": " + lipi(scores[name]));
}

// Calculate average
dhoro total = 0;
ghuriye (dhoro i = 0; i < dorghyo(keys); i = i + 1) {
    total = total + scores[keys[i]];
}
dekho("Average:", total / dorghyo(keys));

// Find highest score
dhoro highest = 0;
dhoro topStudent = "";
ghuriye (dhoro i = 0; i < dorghyo(keys); i = i + 1) {
    jodi (scores[keys[i]] > highest) {
        highest = scores[keys[i]];
        topStudent = keys[i];
    }
}
dekho("Top student:", topStudent, "with", highest);`}
      />

      <h2>Practical Examples</h2>

      <h3>Configuration Object</h3>

      <CodeBlock
        code={`dhoro config = {
    app: {
        name: "MyApp",
        version: "1.0.0",
        debug: mittha
    },
    database: {
        host: "localhost",
        port: 5432,
        name: "mydb"
    },
    features: {
        darkMode: sotti,
        notifications: sotti,
        analytics: mittha
    }
};

// Access configuration
jodi (config.features.darkMode) {
    dekho("Dark mode enabled");
}

// Build connection string
dhoro db = config.database;
dhoro connStr = db.host + ":" + lipi(db.port) + "/" + db.name;
dekho("Connection:", connStr);`}
      />

      <h3>Word Counter</h3>

      <CodeBlock
        code={`dhoro text = "apple banana apple cherry banana apple";
dhoro words = bhag(text, " ");
dhoro counts = {};

ghuriye (dhoro i = 0; i < dorghyo(words); i = i + 1) {
    dhoro word = words[i];
    jodi (counts[word] == khali) {
        counts[word] = 0;
    }
    counts[word] = counts[word] + 1;
}

dekho(counts);
// {apple: 3, banana: 2, cherry: 1}

// Find most common word
dhoro keys = chabi(counts);
dhoro maxCount = 0;
dhoro maxWord = "";

ghuriye (dhoro i = 0; i < dorghyo(keys); i = i + 1) {
    jodi (counts[keys[i]] > maxCount) {
        maxCount = counts[keys[i]];
        maxWord = keys[i];
    }
}

dekho("Most common:", maxWord, "(", maxCount, "times)");`}
      />

      <h3>Data Transformation</h3>

      <CodeBlock
        code={`// Array of maps
dhoro users = [
    {id: 1, naam: "Rahim", role: "admin"},
    {id: 2, naam: "Karim", role: "user"},
    {id: 3, naam: "Jamil", role: "user"},
    {id: 4, naam: "Nasir", role: "moderator"}
];

// Group by role
dhoro byRole = {};

ghuriye (dhoro i = 0; i < dorghyo(users); i = i + 1) {
    dhoro user = users[i];
    dhoro role = user.role;

    jodi (byRole[role] == khali) {
        byRole[role] = [];
    }
    dhokao(byRole[role], user.naam);
}

dekho(byRole);
// {admin: ["Rahim"], user: ["Karim", "Jamil"], moderator: ["Nasir"]}

// Create lookup by ID
dhoro byId = {};
ghuriye (dhoro i = 0; i < dorghyo(users); i = i + 1) {
    dhoro user = users[i];
    byId[lipi(user.id)] = user;
}

// Quick lookup
dekho(byId["2"].naam);  // "Karim"`}
      />

      <h3>JSON-like Data</h3>

      <CodeBlock
        code={`dhoro apiResponse = {
    status: "success",
    code: 200,
    data: {
        users: [
            {id: 1, naam: "User 1", active: sotti},
            {id: 2, naam: "User 2", active: mittha}
        ],
        total: 2,
        page: 1
    },
    meta: {
        timestamp: somoy(),
        server: "api-1"
    }
};

// Process response
jodi (apiResponse.status == "success") {
    dhoro users = apiResponse.data.users;

    ghuriye (dhoro i = 0; i < dorghyo(users); i = i + 1) {
        dhoro user = users[i];
        jodi (user.active) {
            dekho("Active user:", user.naam);
        }
    }
}

// Convert to JSON string
dhoro jsonStr = json_banao(apiResponse);
dekho(jsonStr);`}
      />

      <DocNavigation currentPath="/docs/maps" />
    </div>
  );
}

import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function FileIO() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Advanced
        </span>
      </div>

      <h1>File I/O</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        BanglaCode provides built-in functions for reading and writing files:
        <code>poro</code> (read) and <code>lekho</code> (write).
      </p>

      <h2>Reading Files (poro)</h2>

      <p>
        The <code>poro</code> function (meaning &quot;read&quot;) reads the entire content of a file:
      </p>

      <CodeBlock
        code={`// Read a text file
dhoro content = poro("data.txt");
dekho(content);

// Read and process line by line
dhoro lines = bhag(content, "\\n");
ghuriye (dhoro i = 0; i < dorghyo(lines); i = i + 1) {
    dekho("Line", i + 1, ":", lines[i]);
}`}
      />

      <h3>Reading JSON Files</h3>

      <CodeBlock
        code={`// Read JSON file
dhoro jsonStr = poro("config.json");
dhoro config = json_poro(jsonStr);

dekho("App name:", config.name);
dekho("Version:", config.version);`}
      />

      <h3>Error Handling for File Reading</h3>

      <CodeBlock
        code={`chesta {
    dhoro content = poro("myfile.txt");
    dekho("File content:", content);
} dhoro_bhul (error) {
    dekho("Could not read file:", error);
}

// Check if file exists by attempting to read
kaj fileExists(path) {
    chesta {
        poro(path);
        ferao sotti;
    } dhoro_bhul (e) {
        ferao mittha;
    }
}

jodi (fileExists("data.txt")) {
    dekho("File exists!");
} nahole {
    dekho("File not found");
}`}
      />

      <h2>Writing Files (lekho)</h2>

      <p>
        The <code>lekho</code> function (meaning &quot;write&quot;) writes content to a file:
      </p>

      <CodeBlock
        code={`// Write text to file
lekho("output.txt", "Hello, World!");

// Write multiple lines
dhoro content = "Line 1\\nLine 2\\nLine 3";
lekho("multiline.txt", content);

// Write from array
dhoro lines = ["First line", "Second line", "Third line"];
lekho("fromarray.txt", joro(lines, "\\n"));`}
      />

      <h3>Writing JSON Data</h3>

      <CodeBlock
        code={`dhoro data = {
    users: [
        {naam: "Rahim", boyosh: 25},
        {naam: "Karim", boyosh: 30}
    ],
    count: 2,
    lastUpdated: somoy()
};

// Convert to JSON and write
dhoro jsonStr = json_banao(data);
lekho("users.json", jsonStr);

dekho("Data saved to users.json");`}
      />

      <h3>Appending to Files</h3>

      <CodeBlock
        code={`// BanglaCode overwrites by default
// To append, read first then write combined content

kaj appendToFile(path, newContent) {
    dhoro existing = "";

    chesta {
        existing = poro(path);
    } dhoro_bhul (e) {
        // File doesn't exist, start fresh
        existing = "";
    }

    lekho(path, existing + newContent);
}

// Usage
appendToFile("log.txt", "New log entry\\n");
appendToFile("log.txt", "Another entry\\n");`}
      />

      <h2>Practical Examples</h2>

      <h3>CSV Processing</h3>

      <CodeBlock
        code={`// Read CSV file
dhoro csvContent = poro("data.csv");
dhoro lines = bhag(csvContent, "\\n");

// Parse header
dhoro header = bhag(lines[0], ",");
dekho("Columns:", header);

// Parse data rows
dhoro data = [];
ghuriye (dhoro i = 1; i < dorghyo(lines); i = i + 1) {
    jodi (dorghyo(lines[i]) > 0) {
        dhoro values = bhag(lines[i], ",");
        dhoro row = {};

        ghuriye (dhoro j = 0; j < dorghyo(header); j = j + 1) {
            row[header[j]] = values[j];
        }

        dhokao(data, row);
    }
}

// Use parsed data
dekho("Total rows:", dorghyo(data));
ghuriye (dhoro i = 0; i < dorghyo(data); i = i + 1) {
    dekho(data[i]);
}`}
      />

      <h3>Writing CSV</h3>

      <CodeBlock
        code={`dhoro users = [
    {naam: "Rahim", boyosh: 25, city: "Dhaka"},
    {naam: "Karim", boyosh: 30, city: "Chittagong"},
    {naam: "Jamil", boyosh: 28, city: "Sylhet"}
];

// Build CSV
dhoro header = "naam,boyosh,city";
dhoro rows = [header];

ghuriye (dhoro i = 0; i < dorghyo(users); i = i + 1) {
    dhoro u = users[i];
    dhokao(rows, u.naam + "," + lipi(u.boyosh) + "," + u.city);
}

dhoro csv = joro(rows, "\\n");
lekho("users.csv", csv);

dekho("Saved to users.csv");`}
      />

      <h3>Log File</h3>

      <CodeBlock
        code={`kaj log(message, level) {
    jodi (level == khali) {
        level = "INFO";
    }

    // Format: [TIMESTAMP] [LEVEL] message
    dhoro timestamp = somoy();
    dhoro entry = "[" + lipi(timestamp) + "] [" + level + "] " + message + "\\n";

    // Append to log file
    dhoro existing = "";
    chesta {
        existing = poro("app.log");
    } dhoro_bhul (e) {
        existing = "";
    }

    lekho("app.log", existing + entry);
}

// Usage
log("Application started");
log("Processing data...");
log("Something looks wrong", "WARN");
log("Critical failure!", "ERROR");`}
      />

      <h3>Configuration File</h3>

      <CodeBlock
        code={`dhoro CONFIG_FILE = "config.json";

kaj loadConfig() {
    chesta {
        dhoro content = poro(CONFIG_FILE);
        ferao json_poro(content);
    } dhoro_bhul (e) {
        dekho("Config not found, using defaults");
        ferao {
            theme: "dark",
            language: "bn",
            fontSize: 14
        };
    }
}

kaj saveConfig(config) {
    dhoro json = json_banao(config);
    lekho(CONFIG_FILE, json);
    dekho("Config saved");
}

// Usage
dhoro config = loadConfig();
dekho("Current theme:", config.theme);

// Update config
config.theme = "light";
saveConfig(config);`}
      />

      <h3>Data Backup</h3>

      <CodeBlock
        code={`kaj backup(sourcePath, backupDir) {
    chesta {
        // Read source file
        dhoro content = poro(sourcePath);

        // Create backup filename with timestamp
        dhoro timestamp = somoy();
        dhoro backupPath = backupDir + "/backup_" + lipi(timestamp) + ".txt";

        // Write backup
        lekho(backupPath, content);

        dekho("Backup created:", backupPath);
        ferao sotti;
    } dhoro_bhul (error) {
        dekho("Backup failed:", error);
        ferao mittha;
    }
}

// Create backup
backup("important_data.txt", "backups");`}
      />

      <h3>Simple Database</h3>

      <CodeBlock
        code={`dhoro DB_FILE = "database.json";

kaj loadDB() {
    chesta {
        dhoro content = poro(DB_FILE);
        ferao json_poro(content);
    } dhoro_bhul (e) {
        ferao {users: [], nextId: 1};
    }
}

kaj saveDB(db) {
    lekho(DB_FILE, json_banao(db));
}

kaj addUser(naam, email) {
    dhoro db = loadDB();

    dhoro user = {
        id: db.nextId,
        naam: naam,
        email: email,
        createdAt: somoy()
    };

    dhokao(db.users, user);
    db.nextId = db.nextId + 1;

    saveDB(db);
    ferao user;
}

kaj findUser(id) {
    dhoro db = loadDB();

    ghuriye (dhoro i = 0; i < dorghyo(db.users); i = i + 1) {
        jodi (db.users[i].id == id) {
            ferao db.users[i];
        }
    }

    ferao khali;
}

// Usage
dhoro newUser = addUser("Rahim", "rahim@example.com");
dekho("Created user:", newUser);

dhoro found = findUser(1);
dekho("Found user:", found);`}
      />

      <h2>Best Practices</h2>

      <ul>
        <li><strong>Always handle errors</strong> - Files might not exist or be inaccessible</li>
        <li><strong>Use relative paths carefully</strong> - Consider the working directory</li>
        <li><strong>Close resources in finally</strong> - Though BanglaCode handles this automatically</li>
        <li><strong>Validate before writing</strong> - Ensure data is valid before persisting</li>
        <li><strong>Create backups</strong> - Before overwriting important files</li>
      </ul>

      <DocNavigation currentPath="/docs/file-io" />
    </div>
  );
}

import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function ErrorHandling() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Advanced
        </span>
      </div>

      <h1>Error Handling</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        BanglaCode provides a try-catch-finally mechanism using Bengali keywords:
        <code>chesta</code> (try), <code>dhoro_bhul</code> (catch), <code>shesh</code> (finally),
        and <code>felo</code> (throw).
      </p>

      <h2>Try-Catch (chesta-dhoro_bhul)</h2>

      <p>
        Use <code>chesta</code> (meaning &quot;try&quot; or &quot;attempt&quot;) to wrap code that might throw errors,
        and <code>dhoro_bhul</code> (meaning &quot;catch error&quot;) to handle them:
      </p>

      <CodeBlock
        code={`chesta {
    // Code that might fail
    dhoro result = 10 / 0;
    dekho(result);
} dhoro_bhul (error) {
    // Handle the error
    dekho("An error occurred:", error);
}

// Output: An error occurred: division by zero`}
      />

      <h2>Custom Error Types (v7.0.16)</h2>

      <p>
        BanglaCode provides JavaScript-compatible error types for more specific error handling:
      </p>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Error Type</th>
              <th>Use Case</th>
              <th>Example</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>Error</code></td>
              <td>Generic errors</td>
              <td>General failures</td>
            </tr>
            <tr>
              <td><code>TypeError</code></td>
              <td>Type mismatches</td>
              <td>Expected number, got string</td>
            </tr>
            <tr>
              <td><code>ReferenceError</code></td>
              <td>Undefined variables</td>
              <td>Variable not defined</td>
            </tr>
            <tr>
              <td><code>RangeError</code></td>
              <td>Out of range values</td>
              <td>Index out of bounds</td>
            </tr>
            <tr>
              <td><code>SyntaxError</code></td>
              <td>Syntax issues</td>
              <td>Invalid JSON format</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h3>Creating Custom Errors</h3>

      <CodeBlock
        code={`// Create typed errors
dhoro err1 = Error("Something went wrong");
dhoro err2 = TypeError("Expected number, got string");
dhoro err3 = ReferenceError("Variable 'x' is not defined");
dhoro err4 = RangeError("Index out of bounds");
dhoro err5 = SyntaxError("Invalid JSON");

// Throw custom errors
kaj validateInput(value) {
    jodi (dhoron(value) != "NUMBER") {
        felo TypeError("Input must be a number");
    }
    jodi (value < 0 ba value > 100) {
        felo RangeError("Value must be between 0 and 100");
    }
    ferao sotti;
}`}
      />

      <h3>Error Utility Functions</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Description</th>
              <th>Returns</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>bhul_message(error)</code></td>
              <td>Get error message</td>
              <td>String</td>
            </tr>
            <tr>
              <td><code>bhul_naam(error)</code></td>
              <td>Get error type name</td>
              <td>String</td>
            </tr>
            <tr>
              <td><code>bhul_stack(error)</code></td>
              <td>Get stack trace</td>
              <td>String</td>
            </tr>
            <tr>
              <td><code>is_error(value)</code></td>
              <td>Check if value is error</td>
              <td>Boolean</td>
            </tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`// Using error utilities
chesta {
    felo TypeError("Invalid type");
} dhoro_bhul(e) {
    dekho("Error name:", bhul_naam(e));       // "TypeError"
    dekho("Error message:", bhul_message(e)); // "Invalid type"
    dekho("Stack trace:", bhul_stack(e));     // Stack information
}

// Check if value is an error
dhoro value = TypeError("test");
jodi (is_error(value)) {
    dekho("This is an error object");
}`}
      />

      <h3>Real-World Validation Example</h3>

      <CodeBlock
        code={`// User input validation with proper error types
kaj validateUser(user) {
    // Type checking
    jodi (dhoron(user) != "MAP") {
        felo TypeError("User must be an object");
    }
    
    // Required fields
    dhoro keys = chabi(user);
    dhoro hasName = mittha;
    dhoro hasAge = mittha;
    
    ghuriye (dhoro i = 0; i < kato(keys); i = i + 1) {
        jodi (keys[i] == "name") { hasName = sotti; }
        jodi (keys[i] == "age") { hasAge = sotti; }
    }
    
    jodi (!hasName) {
        felo ReferenceError("User must have 'name' property");
    }
    jodi (!hasAge) {
        felo ReferenceError("User must have 'age' property");
    }
    
    // Value validation
    dhoro age = user["age"];
    jodi (dhoron(age) != "NUMBER") {
        felo TypeError("Age must be a number");
    }
    jodi (age < 0 ba age > 150) {
        felo RangeError("Age must be between 0 and 150");
    }
    
    ferao sotti;
}

// Usage with detailed error handling
dhoro user = {"name": "Rahim", "age": 25};

chesta {
    validateUser(user);
    dekho("User is valid!");
} dhoro_bhul(e) {
    dhoro errorType = bhul_naam(e);
    dhoro errorMsg = bhul_message(e);
    
    dekho("Validation failed:");
    dekho("  Type:", errorType);
    dekho("  Message:", errorMsg);
    
    // Handle different error types
    jodi (errorType == "TypeError") {
        dekho("  → Fix data types");
    } nahole jodi (errorType == "ReferenceError") {
        dekho("  → Add missing fields");
    } nahole jodi (errorType == "RangeError") {
        dekho("  → Check value ranges");
    }
}`}
      />

      <h2>Finally Block (shesh)</h2>

      <p>
        The <code>shesh</code> block (meaning &quot;end&quot; or &quot;finally&quot;) always executes, whether
        an error occurred or not. It&apos;s useful for cleanup:
      </p>

      <CodeBlock
        code={`chesta {
    dekho("Opening file...");
    // Risky operation
    dhoro data = poro("nonexistent.txt");
} dhoro_bhul (error) {
    dekho("Error reading file:", error);
} shesh {
    dekho("Cleanup: closing resources");
}

// Output:
// Opening file...
// Error reading file: file not found
// Cleanup: closing resources`}
      />

      <h3>Finally Without Catch</h3>

      <CodeBlock
        code={`// You can use try-finally without catch
chesta {
    dekho("Processing...");
    // Do something
} shesh {
    dekho("Always runs");
}`}
      />

      <h2>Throwing Errors (felo)</h2>

      <p>
        Use <code>felo</code> (meaning &quot;throw&quot;) to throw custom errors:
      </p>

      <CodeBlock
        code={`kaj divide(a, b) {
    jodi (b == 0) {
        felo "Cannot divide by zero!";
    }
    ferao a / b;
}

chesta {
    dhoro result = divide(10, 0);
    dekho(result);
} dhoro_bhul (error) {
    dekho("Error:", error);
}

// Output: Error: Cannot divide by zero!`}
      />

      <h3>Throwing Different Types</h3>

      <CodeBlock
        code={`// Throw string
felo "Something went wrong";

// Throw with details
felo "Invalid input: expected number";

// In practice, keep error messages descriptive
kaj validateAge(age) {
    jodi (dhoron(age) != "int" ebong dhoron(age) != "float") {
        felo "Age must be a number";
    }
    jodi (age < 0) {
        felo "Age cannot be negative";
    }
    jodi (age > 150) {
        felo "Age seems unrealistic";
    }
    ferao sotti;
}`}
      />

      <h2>Error Propagation</h2>

      <p>
        Errors propagate up the call stack until caught:
      </p>

      <CodeBlock
        code={`kaj level3() {
    dekho("Level 3: throwing error");
    felo "Error from level 3";
}

kaj level2() {
    dekho("Level 2: calling level3");
    level3();
    dekho("Level 2: this won't execute");
}

kaj level1() {
    dekho("Level 1: calling level2");
    level2();
    dekho("Level 1: this won't execute");
}

chesta {
    level1();
} dhoro_bhul (error) {
    dekho("Caught error:", error);
}

// Output:
// Level 1: calling level2
// Level 2: calling level3
// Level 3: throwing error
// Caught error: Error from level 3`}
      />

      <h2>Re-throwing Errors</h2>

      <CodeBlock
        code={`kaj processData(data) {
    chesta {
        // Process data
        jodi (data == khali) {
            felo "Data is null";
        }
        // ... processing logic
    } dhoro_bhul (error) {
        dekho("Logging error:", error);
        // Re-throw after logging
        felo error;
    }
}

chesta {
    processData(khali);
} dhoro_bhul (error) {
    dekho("Main handler caught:", error);
}`}
      />

      <h2>Practical Patterns</h2>

      <h3>Input Validation</h3>

      <CodeBlock
        code={`kaj validateUser(user) {
    jodi (user == khali) {
        felo "User object is required";
    }
    jodi (user.naam == khali ba dorghyo(user.naam) < 2) {
        felo "Name must be at least 2 characters";
    }
    jodi (user.email == khali ba khojo(user.email, "@") < 0) {
        felo "Valid email is required";
    }
    jodi (user.boyosh == khali ba user.boyosh < 0) {
        felo "Age must be a positive number";
    }
    ferao sotti;
}

chesta {
    dhoro newUser = {
        naam: "R",
        email: "invalid-email",
        boyosh: -5
    };
    validateUser(newUser);
    dekho("User is valid");
} dhoro_bhul (error) {
    dekho("Validation failed:", error);
}`}
      />

      <h3>Safe Division</h3>

      <CodeBlock
        code={`kaj safeDivide(a, b) {
    chesta {
        jodi (b == 0) {
            felo "Division by zero";
        }
        ferao {
            success: sotti,
            value: a / b,
            error: khali
        };
    } dhoro_bhul (error) {
        ferao {
            success: mittha,
            value: khali,
            error: error
        };
    }
}

dhoro result = safeDivide(10, 0);

jodi (result.success) {
    dekho("Result:", result.value);
} nahole {
    dekho("Error:", result.error);
}`}
      />

      <h3>API Response Handler</h3>

      <CodeBlock
        code={`kaj fetchData(url) {
    chesta {
        dhoro response = anun(url);

        jodi (response.status != 200) {
            felo "HTTP Error: " + lipi(response.status);
        }

        ferao response.body;
    } dhoro_bhul (error) {
        dekho("Failed to fetch data:", error);
        ferao khali;
    }
}

kaj processApiData() {
    dhoro data = fetchData("https://api.example.com/data");

    jodi (data == khali) {
        dekho("Using fallback data");
        data = {default: sotti};
    }

    // Process data...
    ferao data;
}`}
      />

      <h3>Retry Logic</h3>

      <CodeBlock
        code={`kaj retryOperation(operation, maxRetries) {
    dhoro attempt = 0;

    jotokkhon (attempt < maxRetries) {
        chesta {
            ferao operation();
        } dhoro_bhul (error) {
            attempt = attempt + 1;
            dekho("Attempt", attempt, "failed:", error);

            jodi (attempt >= maxRetries) {
                felo "Max retries exceeded. Last error: " + error;
            }

            // Wait before retry
            ghum(1000);  // Wait 1 second
        }
    }
}

// Usage
chesta {
    dhoro result = retryOperation(kaj() {
        // Simulate unreliable operation
        jodi (lotto() < 0.7) {
            felo "Random failure";
        }
        ferao "Success!";
    }, 5);

    dekho("Final result:", result);
} dhoro_bhul (error) {
    dekho("Operation failed completely:", error);
}`}
      />

      <h3>Resource Cleanup</h3>

      <CodeBlock
        code={`kaj processFile(filename) {
    dhoro file = khali;

    chesta {
        dekho("Opening file:", filename);
        file = poro(filename);

        // Process file content
        dekho("Processing", dorghyo(file), "characters");

        // Simulate error
        jodi (khojo(file, "ERROR") >= 0) {
            felo "File contains error marker";
        }

        dekho("Processing complete");
    } dhoro_bhul (error) {
        dekho("Error processing file:", error);
    } shesh {
        // Cleanup always runs
        jodi (file != khali) {
            dekho("Closing file resources");
            file = khali;
        }
    }
}`}
      />

      <h2>Best Practices</h2>

      <ul>
        <li><strong>Be specific with error messages</strong> - Include relevant details about what went wrong</li>
        <li><strong>Don&apos;t catch errors you can&apos;t handle</strong> - Let them propagate to where they can be properly handled</li>
        <li><strong>Use finally for cleanup</strong> - Always release resources in the shesh block</li>
        <li><strong>Log before re-throwing</strong> - Capture error context before propagating</li>
        <li><strong>Validate early</strong> - Check inputs at function entry to fail fast</li>
        <li><strong>Return result objects</strong> - For expected failures, consider returning success/error objects instead of throwing</li>
      </ul>

      <DocNavigation currentPath="/docs/error-handling" />
    </div>
  );
}

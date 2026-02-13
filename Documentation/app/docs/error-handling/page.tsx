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

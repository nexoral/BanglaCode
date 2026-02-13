import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function BreakContinue() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Control Flow
        </span>
      </div>

      <h1>Break & Continue</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        BanglaCode provides <code>thamo</code> (break) to exit loops early and
        <code>chharo</code> (continue) to skip to the next iteration.
      </p>

      <h2>Break Statement (thamo)</h2>

      <p>
        The <code>thamo</code> keyword (meaning &quot;stop&quot;) immediately exits the current loop:
      </p>

      <CodeBlock
        code={`// Exit loop when condition is met
ghuriye (dhoro i = 1; i <= 10; i = i + 1) {
    jodi (i == 6) {
        dekho("Stopping at", i);
        thamo;
    }
    dekho(i);
}

dekho("Loop ended");

// Output:
// 1
// 2
// 3
// 4
// 5
// Stopping at 6
// Loop ended`}
      />

      <h3>Break in While Loop</h3>

      <CodeBlock
        code={`dhoro found = mittha;
dhoro numbers = [4, 8, 15, 16, 23, 42];
dhoro target = 16;
dhoro index = 0;

jotokkhon (index < dorghyo(numbers)) {
    jodi (numbers[index] == target) {
        found = sotti;
        dekho("Found", target, "at index", index);
        thamo;
    }
    index = index + 1;
}

jodi (na found) {
    dekho(target, "not found");
}`}
      />

      <h3>Break with Infinite Loop</h3>

      <CodeBlock
        code={`// Common pattern: infinite loop with break condition
jotokkhon (sotti) {
    dekho("Menu:");
    dekho("1. Play");
    dekho("2. Settings");
    dekho("3. Exit");

    dhoro choice = nao("Enter choice: ");

    jodi (choice == "3") {
        dekho("Goodbye!");
        thamo;
    }

    jodi (choice == "1") {
        dekho("Starting game...");
    } nahole jodi (choice == "2") {
        dekho("Opening settings...");
    } nahole {
        dekho("Invalid choice");
    }
}`}
      />

      <h2>Continue Statement (chharo)</h2>

      <p>
        The <code>chharo</code> keyword (meaning &quot;skip&quot; or &quot;leave&quot;) skips the rest of the
        current iteration and continues with the next:
      </p>

      <CodeBlock
        code={`// Skip even numbers
ghuriye (dhoro i = 1; i <= 10; i = i + 1) {
    jodi (i % 2 == 0) {
        chharo;  // Skip even numbers
    }
    dekho(i);  // Only prints odd numbers
}

// Output: 1 3 5 7 9`}
      />

      <h3>Continue in While Loop</h3>

      <CodeBlock
        code={`dhoro i = 0;

jotokkhon (i < 10) {
    i = i + 1;

    // Skip multiples of 3
    jodi (i % 3 == 0) {
        chharo;
    }

    dekho(i);
}

// Output: 1 2 4 5 7 8 10`}
      />

      <h3>Skipping Invalid Data</h3>

      <CodeBlock
        code={`dhoro data = [10, -5, 20, khali, 30, 0, 15];

dhoro sum = 0;
dhoro count = 0;

ghuriye (dhoro i = 0; i < dorghyo(data); i = i + 1) {
    dhoro value = data[i];

    // Skip invalid values
    jodi (value == khali ba value <= 0) {
        chharo;
    }

    sum = sum + value;
    count = count + 1;
}

dekho("Sum of positive numbers:", sum);      // 75
dekho("Count of valid numbers:", count);     // 4
dekho("Average:", sum / count);              // 18.75`}
      />

      <h2>Nested Loops</h2>

      <p>
        <code>thamo</code> and <code>chharo</code> only affect the innermost loop:
      </p>

      <CodeBlock
        code={`// Break only exits inner loop
ghuriye (dhoro i = 1; i <= 3; i = i + 1) {
    dekho("Outer loop i =", i);

    ghuriye (dhoro j = 1; j <= 5; j = j + 1) {
        jodi (j == 3) {
            thamo;  // Only breaks inner loop
        }
        dekho("  Inner loop j =", j);
    }
}

// Output:
// Outer loop i = 1
//   Inner loop j = 1
//   Inner loop j = 2
// Outer loop i = 2
//   Inner loop j = 1
//   Inner loop j = 2
// Outer loop i = 3
//   Inner loop j = 1
//   Inner loop j = 2`}
      />

      <h3>Breaking Outer Loop with Flag</h3>

      <CodeBlock
        code={`// Use a flag to break outer loop
dhoro shouldBreak = mittha;

ghuriye (dhoro i = 1; i <= 5; i = i + 1) {
    ghuriye (dhoro j = 1; j <= 5; j = j + 1) {
        dekho("i =", i, "j =", j);

        jodi (i == 2 ebong j == 3) {
            shouldBreak = sotti;
            thamo;
        }
    }

    jodi (shouldBreak) {
        thamo;
    }
}

dekho("Exited both loops");`}
      />

      <h2>Practical Examples</h2>

      <h3>Finding First Match</h3>

      <CodeBlock
        code={`kaj findFirst(arr, predicate) {
    ghuriye (dhoro i = 0; i < dorghyo(arr); i = i + 1) {
        jodi (predicate(arr[i])) {
            ferao arr[i];
        }
    }
    ferao khali;
}

dhoro numbers = [1, 4, 9, 16, 25, 36];

// Find first number > 10
dhoro result = findFirst(numbers, kaj(n) { ferao n > 10; });
dekho("First > 10:", result);  // 16`}
      />

      <h3>Processing Until Error</h3>

      <CodeBlock
        code={`dhoro items = ["apple", "banana", "ERROR", "cherry", "date"];

ghuriye (dhoro i = 0; i < dorghyo(items); i = i + 1) {
    dhoro item = items[i];

    jodi (item == "ERROR") {
        dekho("Error encountered, stopping processing");
        thamo;
    }

    dekho("Processing:", item);
}

// Output:
// Processing: apple
// Processing: banana
// Error encountered, stopping processing`}
      />

      <h3>Filtering with Continue</h3>

      <CodeBlock
        code={`dhoro scores = [85, 42, 91, 67, 73, 38, 95, 88];

// Print only passing scores (>= 60)
dekho("Passing scores:");

ghuriye (dhoro i = 0; i < dorghyo(scores); i = i + 1) {
    jodi (scores[i] < 60) {
        chharo;  // Skip failing scores
    }
    dekho(scores[i]);
}

// Output: 85 91 67 73 95 88`}
      />

      <h3>Game Loop with Multiple Exit Conditions</h3>

      <CodeBlock
        code={`dhoro health = 100;
dhoro rounds = 0;
dhoro maxRounds = 10;

jotokkhon (sotti) {
    rounds = rounds + 1;

    // Simulate damage
    dhoro damage = kache(lotto() * 30);
    health = health - damage;

    dekho("Round", rounds, "- Damage:", damage, "- Health:", health);

    // Check win condition
    jodi (rounds >= maxRounds) {
        dekho("You survived all rounds!");
        thamo;
    }

    // Check lose condition
    jodi (health <= 0) {
        dekho("Game Over! You lasted", rounds, "rounds");
        thamo;
    }
}`}
      />

      <h2>Best Practices</h2>

      <ul>
        <li>Use <code>thamo</code> for early exit when a goal is achieved (e.g., finding an element)</li>
        <li>Use <code>chharo</code> to skip invalid or unwanted iterations</li>
        <li>Avoid excessive use of <code>thamo</code>/<code>chharo</code> as it can make code harder to follow</li>
        <li>Consider extracting complex loop logic into functions with <code>ferao</code> instead</li>
        <li>For nested loops, use flags or functions to manage breaking outer loops</li>
      </ul>

      <DocNavigation currentPath="/docs/break-continue" />
    </div>
  );
}

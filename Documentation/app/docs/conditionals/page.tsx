import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function Conditionals() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Control Flow
        </span>
      </div>

      <h1>Conditionals</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        BanglaCode uses <code>jodi</code> (if) and <code>nahole</code> (else) for conditional
        execution. These keywords allow your program to make decisions based on conditions.
      </p>

      <h2>Basic If Statement</h2>

      <p>
        The <code>jodi</code> keyword executes a block of code only if the condition is true:
      </p>

      <CodeBlock
        code={`dhoro age = 20;

jodi (age >= 18) {
    dekho("You are an adult");
}

// Output: You are an adult`}
      />

      <h2>If-Else Statement</h2>

      <p>
        Use <code>nahole</code> (else) to specify code to run when the condition is false:
      </p>

      <CodeBlock
        code={`dhoro age = 15;

jodi (age >= 18) {
    dekho("You are an adult");
} nahole {
    dekho("You are a minor");
}

// Output: You are a minor`}
      />

      <h2>Else-If Chain</h2>

      <p>
        Chain multiple conditions using <code>nahole jodi</code>:
      </p>

      <CodeBlock
        code={`dhoro score = 85;

jodi (score >= 90) {
    dekho("Grade: A - Excellent!");
} nahole jodi (score >= 80) {
    dekho("Grade: B - Very Good!");
} nahole jodi (score >= 70) {
    dekho("Grade: C - Good");
} nahole jodi (score >= 60) {
    dekho("Grade: D - Satisfactory");
} nahole {
    dekho("Grade: F - Needs Improvement");
}

// Output: Grade: B - Very Good!`}
      />

      <h2>Nested Conditionals</h2>

      <p>
        You can nest conditionals inside each other:
      </p>

      <CodeBlock
        code={`dhoro age = 25;
dhoro hasLicense = sotti;
dhoro hasInsurance = sotti;

jodi (age >= 18) {
    jodi (hasLicense) {
        jodi (hasInsurance) {
            dekho("You can drive legally!");
        } nahole {
            dekho("You need insurance to drive.");
        }
    } nahole {
        dekho("You need a license to drive.");
    }
} nahole {
    dekho("You must be 18 or older to drive.");
}`}
      />

      <h2>Using Logical Operators</h2>

      <p>
        Combine conditions using <code>ebong</code> (AND), <code>ba</code> (OR), and <code>na</code> (NOT):
      </p>

      <CodeBlock
        code={`dhoro age = 25;
dhoro hasLicense = sotti;
dhoro isSuspended = mittha;

// AND - all conditions must be true
jodi (age >= 18 ebong hasLicense ebong na isSuspended) {
    dekho("You can drive!");
}

// OR - at least one condition must be true
dhoro isWeekend = sotti;
dhoro isHoliday = mittha;

jodi (isWeekend ba isHoliday) {
    dekho("No work today!");
}

// NOT - inverts the condition
jodi (na isSuspended) {
    dekho("License is valid");
}

// Complex conditions with parentheses
dhoro role = "admin";
dhoro isActive = sotti;

jodi ((role == "admin" ba role == "moderator") ebong isActive) {
    dekho("Access granted");
}`}
      />

      <h2>Conditional Expressions</h2>

      <p>
        Conditions can include any expression that evaluates to a boolean:
      </p>

      <CodeBlock
        code={`// Comparison operators
dhoro x = 10;

jodi (x == 10) { dekho("Equal to 10"); }
jodi (x != 5) { dekho("Not equal to 5"); }
jodi (x > 5) { dekho("Greater than 5"); }
jodi (x < 20) { dekho("Less than 20"); }
jodi (x >= 10) { dekho("Greater than or equal to 10"); }
jodi (x <= 10) { dekho("Less than or equal to 10"); }

// String comparison
dhoro name = "Rahim";

jodi (name == "Rahim") {
    dekho("Hello Rahim!");
}

// Array membership check
dhoro fruits = ["apple", "banana", "mango"];

jodi (ache(fruits, "mango")) {
    dekho("Mango is in the list!");
}`}
      />

      <h2>Truthiness in Conditions</h2>

      <p>
        Any value can be used in a condition. BanglaCode evaluates truthiness as follows:
      </p>

      <CodeBlock
        code={`// Falsy values (treated as false)
dhoro falsyValues = [mittha, khali, 0];

jodi (mittha) { dekho("Won't print"); }
jodi (khali) { dekho("Won't print"); }
jodi (0) { dekho("Won't print"); }

// Truthy values (treated as true)
jodi (sotti) { dekho("Prints: boolean true"); }
jodi (1) { dekho("Prints: non-zero number"); }
jodi ("hello") { dekho("Prints: non-empty string"); }
jodi ([]) { dekho("Prints: empty array is truthy"); }
jodi ({}) { dekho("Prints: empty map is truthy"); }

// Practical example: checking for value
dhoro user = getUserFromDatabase();

jodi (user) {
    dekho("User found:", user.naam);
} nahole {
    dekho("User not found");
}`}
      />

      <h2>Single Statement Conditionals</h2>

      <p>
        Always use braces for conditional blocks. BanglaCode requires braces even for single statements:
      </p>

      <CodeBlock
        code={`// Correct way
jodi (x > 0) {
    dekho("Positive");
}

// Also correct - single line
jodi (x > 0) { dekho("Positive"); }

// Multiple statements need braces
jodi (x > 0) {
    dekho("Positive");
    dekho("Greater than zero");
}`}
      />

      <h2>Practical Examples</h2>

      <h3>Input Validation</h3>

      <CodeBlock
        code={`kaj validateAge(age) {
    jodi (age < 0) {
        dekho("Error: Age cannot be negative");
        ferao mittha;
    } nahole jodi (age > 150) {
        dekho("Error: Age seems unrealistic");
        ferao mittha;
    }
    ferao sotti;
}

jodi (validateAge(25)) {
    dekho("Age is valid");
}`}
      />

      <h3>Menu Selection</h3>

      <CodeBlock
        code={`dhoro choice = 2;

jodi (choice == 1) {
    dekho("Starting new game...");
} nahole jodi (choice == 2) {
    dekho("Loading saved game...");
} nahole jodi (choice == 3) {
    dekho("Opening settings...");
} nahole jodi (choice == 4) {
    dekho("Goodbye!");
} nahole {
    dekho("Invalid choice. Please try again.");
}`}
      />

      <h3>Range Checking</h3>

      <CodeBlock
        code={`kaj getTemperatureMessage(temp) {
    jodi (temp < 0) {
        ferao "Freezing cold!";
    } nahole jodi (temp < 15) {
        ferao "Cold";
    } nahole jodi (temp < 25) {
        ferao "Pleasant";
    } nahole jodi (temp < 35) {
        ferao "Warm";
    } nahole {
        ferao "Hot!";
    }
}

dekho(getTemperatureMessage(22));  // "Pleasant"`}
      />

      <DocNavigation currentPath="/docs/conditionals" />
    </div>
  );
}

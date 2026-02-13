import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function Variables() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Language Basics
        </span>
      </div>

      <h1>Variables</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        Variables in BanglaCode are declared using the <code>dhoro</code> keyword, meaning
        &quot;hold&quot; in Bengali. BanglaCode is dynamically typed, so you don&apos;t need to specify types.
      </p>

      <h2>Declaring Variables</h2>

      <p>
        Use the <code>dhoro</code> keyword followed by the variable name and an initial value:
      </p>

      <CodeBlock
        code={`// Basic variable declaration
dhoro naam = "Rahim";
dhoro boyosh = 25;
dhoro active = sotti;

// Variables can hold any type
dhoro message = "Hello";
message = 42;        // Now holds a number
message = [1, 2, 3]; // Now holds an array`}
      />

      <h2>Naming Rules</h2>

      <p>Variable names must follow these rules:</p>

      <ul>
        <li>Must start with a letter (a-z, A-Z) or underscore (_)</li>
        <li>Can contain letters, digits (0-9), and underscores</li>
        <li>Cannot be a reserved keyword</li>
        <li>Are case-sensitive (<code>naam</code> and <code>Naam</code> are different)</li>
      </ul>

      <CodeBlock
        code={`// Valid variable names
dhoro firstName = "Rahim";
dhoro _private = "secret";
dhoro count2 = 100;
dhoro userAge = 25;

// Invalid variable names (will cause errors)
// dhoro 2count = 5;      // Cannot start with digit
// dhoro first-name = ""; // Hyphens not allowed
// dhoro dhoro = 5;       // Cannot use reserved keywords`}
      />

      <h2>Assignment Operators</h2>

      <p>
        BanglaCode supports compound assignment operators for convenience:
      </p>

      <CodeBlock
        code={`dhoro x = 10;

// Simple assignment
x = 20;

// Compound assignment operators
x += 5;   // x = x + 5  (x is now 25)
x -= 3;   // x = x - 3  (x is now 22)
x *= 2;   // x = x * 2  (x is now 44)
x /= 4;   // x = x / 4  (x is now 11)

dekho(x);  // Output: 11`}
      />

      <h2>Variable Scope</h2>

      <p>
        Variables have block scope. A variable declared inside a block (between <code>{`{}`}</code>)
        is only accessible within that block:
      </p>

      <CodeBlock
        code={`dhoro global = "I am global";

jodi (sotti) {
    dhoro local = "I am local";
    dekho(global);  // Works - can access outer variable
    dekho(local);   // Works - can access local variable
}

dekho(global);  // Works
// dekho(local);  // Error! local is not defined here`}
      />

      <h3>Closures</h3>

      <p>
        Functions capture variables from their surrounding scope:
      </p>

      <CodeBlock
        code={`kaj makeCounter() {
    dhoro count = 0;

    ferao kaj() {
        count = count + 1;
        ferao count;
    };
}

dhoro counter = makeCounter();
dekho(counter());  // 1
dekho(counter());  // 2
dekho(counter());  // 3`}
      />

      <h2>Type Checking</h2>

      <p>
        Use the <code>dhoron</code> (type) built-in function to check a variable&apos;s type:
      </p>

      <CodeBlock
        code={`dhoro num = 42;
dhoro str = "hello";
dhoro bool = sotti;
dhoro arr = [1, 2, 3];
dhoro obj = {x: 1, y: 2};
dhoro nothing = khali;

dekho(dhoron(num));      // "int" or "float"
dekho(dhoron(str));      // "string"
dekho(dhoron(bool));     // "boolean"
dekho(dhoron(arr));      // "array"
dekho(dhoron(obj));      // "map"
dekho(dhoron(nothing));  // "null"`}
      />

      <h2>Type Conversion</h2>

      <p>
        BanglaCode provides built-in functions for type conversion:
      </p>

      <CodeBlock
        code={`// Convert to string
dhoro num = 42;
dhoro str = lipi(num);  // "42"

// Convert to number
dhoro text = "123";
dhoro number = sonkha(text);  // 123

// Check results
dekho(dhoron(str));     // "string"
dekho(dhoron(number));  // "int"`}
      />

      <h2>Constants</h2>

      <p>
        BanglaCode doesn&apos;t have a separate constant keyword. By convention, use UPPERCASE
        names for values that shouldn&apos;t change:
      </p>

      <CodeBlock
        code={`// Convention: UPPERCASE for constants
dhoro PI = 3.14159;
dhoro MAX_SIZE = 100;
dhoro APP_NAME = "MyApp";

// These CAN be changed (no enforcement)
// but by convention, they shouldn't be`}
      />

      <h2>Multiple Variables</h2>

      <p>
        Each variable must be declared separately:
      </p>

      <CodeBlock
        code={`// Declare multiple variables
dhoro a = 1;
dhoro b = 2;
dhoro c = 3;

// Variables can reference each other
dhoro x = 10;
dhoro y = x * 2;
dhoro z = x + y;

dekho(z);  // 30`}
      />

      <h2>Truthiness</h2>

      <p>
        When used in boolean contexts (like <code>jodi</code>), values are evaluated for truthiness:
      </p>

      <CodeBlock
        code={`// Falsy values
dhoro f1 = mittha;  // false
dhoro f2 = khali;   // null
dhoro f3 = 0;       // zero

// Truthy values (everything else)
dhoro t1 = sotti;   // true
dhoro t2 = 1;       // non-zero numbers
dhoro t3 = "hello"; // non-empty strings
dhoro t4 = [];      // arrays (even empty)
dhoro t5 = {};      // maps (even empty)

// Example usage
jodi (f1) {
    dekho("This won't print");
}

jodi (t2) {
    dekho("This will print");
}`}
      />

      <DocNavigation currentPath="/docs/variables" />
    </div>
  );
}

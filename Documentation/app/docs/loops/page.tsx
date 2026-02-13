import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function Loops() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Control Flow
        </span>
      </div>

      <h1>Loops</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        BanglaCode provides two loop constructs: <code>jotokkhon</code> (while) for
        condition-based loops and <code>ghuriye</code> (for) for counter-based iteration.
      </p>

      <h2>While Loop (jotokkhon)</h2>

      <p>
        The <code>jotokkhon</code> loop repeats a block of code while a condition is true.
        The Bengali word means &quot;as long as&quot;:
      </p>

      <CodeBlock
        code={`dhoro i = 0;

jotokkhon (i < 5) {
    dekho("Count:", i);
    i = i + 1;
}

// Output:
// Count: 0
// Count: 1
// Count: 2
// Count: 3
// Count: 4`}
      />

      <h3>While Loop Syntax</h3>

      <CodeBlock
        code={`jotokkhon (condition) {
    // Code to repeat
    // Must update condition to avoid infinite loop
}`}
      />

      <h3>Example: Sum of Numbers</h3>

      <CodeBlock
        code={`dhoro sum = 0;
dhoro n = 1;

jotokkhon (n <= 10) {
    sum = sum + n;
    n = n + 1;
}

dekho("Sum of 1 to 10:", sum);  // 55`}
      />

      <h3>Example: User Input Loop</h3>

      <CodeBlock
        code={`dhoro input = "";

jotokkhon (input != "quit") {
    dekho("Enter a command (type 'quit' to exit):");
    input = nao();

    jodi (input != "quit") {
        dekho("You entered:", input);
    }
}

dekho("Goodbye!");`}
      />

      <h2>For Loop (ghuriye)</h2>

      <p>
        The <code>ghuriye</code> loop is used for counter-based iteration.
        The Bengali word means &quot;rotating&quot; or &quot;looping&quot;:
      </p>

      <CodeBlock
        code={`ghuriye (dhoro i = 0; i < 5; i = i + 1) {
    dekho("Number:", i);
}

// Output:
// Number: 0
// Number: 1
// Number: 2
// Number: 3
// Number: 4`}
      />

      <h3>For Loop Syntax</h3>

      <CodeBlock
        code={`ghuriye (initialization; condition; update) {
    // Code to repeat
}

// Parts:
// - initialization: Run once before loop starts
// - condition: Checked before each iteration
// - update: Run after each iteration`}
      />

      <h3>Counting Up</h3>

      <CodeBlock
        code={`// Count from 1 to 10
ghuriye (dhoro i = 1; i <= 10; i = i + 1) {
    dekho(i);
}

// Count by 2s
ghuriye (dhoro i = 0; i <= 10; i = i + 2) {
    dekho(i);  // 0, 2, 4, 6, 8, 10
}`}
      />

      <h3>Counting Down</h3>

      <CodeBlock
        code={`// Countdown from 10 to 1
ghuriye (dhoro i = 10; i >= 1; i = i - 1) {
    dekho(i);
}

dekho("Blast off!");`}
      />

      <h2>Iterating Over Arrays</h2>

      <CodeBlock
        code={`dhoro fruits = ["Apple", "Banana", "Mango", "Orange"];

// Using index-based for loop
ghuriye (dhoro i = 0; i < dorghyo(fruits); i = i + 1) {
    dekho(i + 1, "-", fruits[i]);
}

// Output:
// 1 - Apple
// 2 - Banana
// 3 - Mango
// 4 - Orange`}
      />

      <h2>Nested Loops</h2>

      <CodeBlock
        code={`// Multiplication table
ghuriye (dhoro i = 1; i <= 5; i = i + 1) {
    dhoro row = "";
    ghuriye (dhoro j = 1; j <= 5; j = j + 1) {
        row = row + lipi(i * j) + "\\t";
    }
    dekho(row);
}

// Output:
// 1  2  3  4  5
// 2  4  6  8  10
// 3  6  9  12 15
// 4  8  12 16 20
// 5  10 15 20 25`}
      />

      <h3>Pattern Printing</h3>

      <CodeBlock
        code={`// Print a triangle pattern
dhoro n = 5;

ghuriye (dhoro i = 1; i <= n; i = i + 1) {
    dhoro stars = "";
    ghuriye (dhoro j = 1; j <= i; j = j + 1) {
        stars = stars + "* ";
    }
    dekho(stars);
}

// Output:
// *
// * *
// * * *
// * * * *
// * * * * *`}
      />

      <h2>Loop Control</h2>

      <p>
        Use <code>thamo</code> (break) to exit a loop early and <code>chharo</code> (continue)
        to skip to the next iteration. See the <a href="/docs/break-continue" className="text-primary hover:underline">Break & Continue</a> page for details.
      </p>

      <CodeBlock
        code={`// Using thamo (break)
ghuriye (dhoro i = 1; i <= 10; i = i + 1) {
    jodi (i == 6) {
        thamo;  // Exit loop when i is 6
    }
    dekho(i);
}
// Output: 1 2 3 4 5

// Using chharo (continue)
ghuriye (dhoro i = 1; i <= 5; i = i + 1) {
    jodi (i == 3) {
        chharo;  // Skip when i is 3
    }
    dekho(i);
}
// Output: 1 2 4 5`}
      />

      <h2>Practical Examples</h2>

      <h3>Finding Prime Numbers</h3>

      <CodeBlock
        code={`kaj isPrime(n) {
    jodi (n < 2) {
        ferao mittha;
    }

    ghuriye (dhoro i = 2; i * i <= n; i = i + 1) {
        jodi (n % i == 0) {
            ferao mittha;
        }
    }

    ferao sotti;
}

// Print primes up to 50
dekho("Prime numbers up to 50:");
ghuriye (dhoro n = 2; n <= 50; n = n + 1) {
    jodi (isPrime(n)) {
        dekho(n);
    }
}`}
      />

      <h3>Fibonacci Sequence</h3>

      <CodeBlock
        code={`dhoro n = 10;
dhoro a = 0;
dhoro b = 1;

dekho("First", n, "Fibonacci numbers:");

ghuriye (dhoro i = 0; i < n; i = i + 1) {
    dekho(a);
    dhoro temp = a + b;
    a = b;
    b = temp;
}`}
      />

      <h3>Array Processing</h3>

      <CodeBlock
        code={`dhoro numbers = [4, 2, 9, 1, 7, 5, 3, 8, 6];

// Find maximum
dhoro max = numbers[0];
ghuriye (dhoro i = 1; i < dorghyo(numbers); i = i + 1) {
    jodi (numbers[i] > max) {
        max = numbers[i];
    }
}
dekho("Maximum:", max);  // 9

// Calculate sum
dhoro total = 0;
ghuriye (dhoro i = 0; i < dorghyo(numbers); i = i + 1) {
    total = total + numbers[i];
}
dekho("Sum:", total);  // 45

// Calculate average
dekho("Average:", total / dorghyo(numbers));  // 5`}
      />

      <h3>String Manipulation</h3>

      <CodeBlock
        code={`dhoro text = "Hello";

// Reverse a string
dhoro reversed = "";
ghuriye (dhoro i = dorghyo(text) - 1; i >= 0; i = i - 1) {
    reversed = reversed + text[i];
}
dekho("Reversed:", reversed);  // "olleH"

// Count vowels
dhoro vowels = "aeiouAEIOU";
dhoro count = 0;
ghuriye (dhoro i = 0; i < dorghyo(text); i = i + 1) {
    jodi (khojo(vowels, text[i]) >= 0) {
        count = count + 1;
    }
}
dekho("Vowels:", count);  // 2`}
      />

      <h2>Infinite Loops</h2>

      <p>
        Be careful to ensure your loop condition eventually becomes false, or use <code>thamo</code>
        to exit:
      </p>

      <CodeBlock
        code={`// Infinite loop with break
jotokkhon (sotti) {
    dhoro input = nao("Enter command: ");

    jodi (input == "exit") {
        thamo;  // Exit the infinite loop
    }

    dekho("Processing:", input);
}

// Warning: This would run forever!
// dhoro i = 0;
// jotokkhon (i < 10) {
//     dekho(i);
//     // Forgot to increment i!
// }`}
      />

      <DocNavigation currentPath="/docs/loops" />
    </div>
  );
}

import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function Precedence() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Reference
        </span>
      </div>

      <h1>Operator Precedence</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        Operator precedence determines the order in which operations are evaluated.
        Higher precedence operators are evaluated first.
      </p>

      <h2>Precedence Table</h2>

      <p>Operators are listed from highest precedence (1) to lowest (11):</p>

      <div className="overflow-x-auto my-6">
        <table>
          <thead>
            <tr>
              <th>Level</th>
              <th>Operators</th>
              <th>Description</th>
              <th>Associativity</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><strong>1</strong></td>
              <td><code>[]</code> <code>.</code></td>
              <td>Member access, index</td>
              <td>Left to right</td>
            </tr>
            <tr>
              <td><strong>2</strong></td>
              <td><code>()</code></td>
              <td>Function call</td>
              <td>Left to right</td>
            </tr>
            <tr>
              <td><strong>3</strong></td>
              <td><code>-</code> <code>!</code> <code>na</code></td>
              <td>Unary minus, logical NOT</td>
              <td>Right to left</td>
            </tr>
            <tr>
              <td><strong>4</strong></td>
              <td><code>**</code></td>
              <td>Exponentiation</td>
              <td>Right to left</td>
            </tr>
            <tr>
              <td><strong>5</strong></td>
              <td><code>*</code> <code>/</code> <code>%</code></td>
              <td>Multiplication, division, modulo</td>
              <td>Left to right</td>
            </tr>
            <tr>
              <td><strong>6</strong></td>
              <td><code>+</code> <code>-</code></td>
              <td>Addition, subtraction</td>
              <td>Left to right</td>
            </tr>
            <tr>
              <td><strong>7</strong></td>
              <td><code>&lt;</code> <code>&gt;</code> <code>&lt;=</code> <code>&gt;=</code></td>
              <td>Comparison</td>
              <td>Left to right</td>
            </tr>
            <tr>
              <td><strong>8</strong></td>
              <td><code>==</code> <code>!=</code></td>
              <td>Equality</td>
              <td>Left to right</td>
            </tr>
            <tr>
              <td><strong>9</strong></td>
              <td><code>ebong</code></td>
              <td>Logical AND</td>
              <td>Left to right</td>
            </tr>
            <tr>
              <td><strong>10</strong></td>
              <td><code>ba</code></td>
              <td>Logical OR</td>
              <td>Left to right</td>
            </tr>
            <tr>
              <td><strong>11</strong></td>
              <td><code>=</code> <code>+=</code> <code>-=</code> <code>*=</code> <code>/=</code></td>
              <td>Assignment</td>
              <td>Right to left</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h2>Arithmetic Precedence</h2>

      <CodeBlock
        code={`// Multiplication before addition
dekho(2 + 3 * 4);       // 14 (not 20)
// Equivalent to: 2 + (3 * 4)

// Division before subtraction
dekho(10 - 6 / 2);      // 7 (not 2)
// Equivalent to: 10 - (6 / 2)

// Left to right for same precedence
dekho(20 / 4 * 2);      // 10
// Evaluated as: (20 / 4) * 2

// Exponentiation is right-associative
dekho(2 ** 3 ** 2);     // 512
// Evaluated as: 2 ** (3 ** 2) = 2 ** 9`}
      />

      <h2>Comparison vs Arithmetic</h2>

      <CodeBlock
        code={`// Arithmetic evaluated before comparison
dekho(5 + 3 > 6);       // sotti
// Evaluated as: (5 + 3) > 6 = 8 > 6

dekho(10 - 2 == 4 * 2); // sotti
// Evaluated as: (10 - 2) == (4 * 2) = 8 == 8

dekho(2 * 3 <= 1 + 5);  // sotti
// Evaluated as: (2 * 3) <= (1 + 5) = 6 <= 6`}
      />

      <h2>Logical Operators</h2>

      <CodeBlock
        code={`// Comparison before logical operators
dekho(5 > 3 ebong 2 < 4);   // sotti
// Evaluated as: (5 > 3) ebong (2 < 4)

// AND before OR
dekho(sotti ba mittha ebong mittha);  // sotti
// Evaluated as: sotti ba (mittha ebong mittha) = sotti ba mittha = sotti

// Use parentheses for clarity
dekho((sotti ba mittha) ebong mittha);  // mittha

// NOT has highest precedence among logical
dekho(na sotti ebong sotti);  // mittha
// Evaluated as: (na sotti) ebong sotti = mittha ebong sotti`}
      />

      <h2>Member Access</h2>

      <CodeBlock
        code={`dhoro obj = {
    arr: [1, 2, 3],
    fn: kaj(x) { ferao x * 2; }
};

// Member access has highest precedence
dekho(obj.arr[0]);      // 1
dekho(obj.fn(5));       // 10

// Chained access
dhoro nested = {a: {b: {c: 42}}};
dekho(nested.a.b.c);    // 42`}
      />

      <h2>Assignment</h2>

      <CodeBlock
        code={`// Assignment has lowest precedence
dhoro x = 5 + 3;        // x = 8
dhoro y = 2 * 4 > 6;    // y = sotti

// Assignment is right-associative
dhoro a = 1;
dhoro b = 2;
dhoro c = 3;

a = b = c = 10;
dekho(a, b, c);  // 10 10 10
// Evaluated as: a = (b = (c = 10))

// Compound assignment
dhoro n = 10;
n += 5 * 2;     // n = n + (5 * 2) = 10 + 10 = 20`}
      />

      <h2>Using Parentheses</h2>

      <p>
        Parentheses can override default precedence:
      </p>

      <CodeBlock
        code={`// Without parentheses
dekho(2 + 3 * 4);       // 14

// With parentheses
dekho((2 + 3) * 4);     // 20

// Complex expression
dhoro result = (10 + 5) * (20 - 10) / (2 + 3);
dekho(result);          // 30

// Logical expressions
dhoro a = sotti;
dhoro b = mittha;
dhoro c = sotti;

dekho(a ba b ebong c);      // sotti (a ba (b ebong c))
dekho((a ba b) ebong c);    // sotti ((a ba b) ebong c)`}
      />

      <h2>Common Mistakes</h2>

      <CodeBlock
        code={`// Mistake 1: Assuming left-to-right for everything
dhoro x = 2 + 3 * 4;    // x = 14, not 20

// Mistake 2: Forgetting AND has higher precedence than OR
dhoro result = sotti ba mittha ebong mittha;
// This is: sotti ba (mittha ebong mittha) = sotti
// Not: (sotti ba mittha) ebong mittha = mittha

// Mistake 3: Comparison chaining doesn't work as expected
dhoro age = 25;
// This doesn't do what you think:
// 18 < age < 65

// Correct way:
jodi (age > 18 ebong age < 65) {
    dekho("Working age");
}

// Mistake 4: Assignment in condition
// This assigns, doesn't compare:
// jodi (x = 5) { ... }

// Correct:
jodi (x == 5) { dekho("Equal"); }`}
      />

      <h2>Best Practices</h2>

      <ul>
        <li><strong>Use parentheses for clarity</strong> - Even when not strictly needed, they make intent clear</li>
        <li><strong>Break complex expressions</strong> - Use intermediate variables for readability</li>
        <li><strong>Be careful with logical operators</strong> - AND has higher precedence than OR</li>
        <li><strong>Test edge cases</strong> - Verify your expressions work as expected</li>
      </ul>

      <CodeBlock
        code={`// Good: Clear with parentheses
dhoro isValid = (age >= 18) ebong (hasPermission ba isAdmin);

// Good: Broken into steps
dhoro basePrice = 100;
dhoro discount = 20;
dhoro tax = 10;
dhoro priceAfterDiscount = basePrice - discount;
dhoro finalPrice = priceAfterDiscount + (priceAfterDiscount * tax / 100);

// Instead of:
// dhoro finalPrice = basePrice - discount + (basePrice - discount) * tax / 100;`}
      />

      <DocNavigation currentPath="/docs/precedence" />
    </div>
  );
}

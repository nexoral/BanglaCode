import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function Operators() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Language Basics
        </span>
      </div>

      <h1>Operators</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        BanglaCode provides a comprehensive set of operators for arithmetic, comparison,
        logical operations, and more. Some logical operators use Bengali keywords.
      </p>

      <h2>Arithmetic Operators</h2>

      <div className="overflow-x-auto my-6">
        <table>
          <thead>
            <tr>
              <th>Operator</th>
              <th>Name</th>
              <th>Example</th>
              <th>Result</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>+</code></td><td>Addition</td><td><code>5 + 3</code></td><td>8</td></tr>
            <tr><td><code>-</code></td><td>Subtraction</td><td><code>5 - 3</code></td><td>2</td></tr>
            <tr><td><code>*</code></td><td>Multiplication</td><td><code>5 * 3</code></td><td>15</td></tr>
            <tr><td><code>/</code></td><td>Division</td><td><code>10 / 4</code></td><td>2.5</td></tr>
            <tr><td><code>%</code></td><td>Modulo</td><td><code>10 % 3</code></td><td>1</td></tr>
            <tr><td><code>**</code></td><td>Exponent</td><td><code>2 ** 3</code></td><td>8</td></tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`dhoro a = 10;
dhoro b = 3;

dekho(a + b);   // 13 (Addition)
dekho(a - b);   // 7  (Subtraction)
dekho(a * b);   // 30 (Multiplication)
dekho(a / b);   // 3.333... (Division)
dekho(a % b);   // 1  (Modulo - remainder)
dekho(a ** b);  // 1000 (Power: 10^3)

// String concatenation with +
dhoro greeting = "Hello" + " " + "World";
dekho(greeting);  // "Hello World"`}
      />

      <h2>Comparison Operators</h2>

      <div className="overflow-x-auto my-6">
        <table>
          <thead>
            <tr>
              <th>Operator</th>
              <th>Name</th>
              <th>Example</th>
              <th>Result</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>==</code></td><td>Equal</td><td><code>5 == 5</code></td><td><code>sotti</code></td></tr>
            <tr><td><code>!=</code></td><td>Not Equal</td><td><code>5 != 3</code></td><td><code>sotti</code></td></tr>
            <tr><td><code>&lt;</code></td><td>Less Than</td><td><code>3 &lt; 5</code></td><td><code>sotti</code></td></tr>
            <tr><td><code>&gt;</code></td><td>Greater Than</td><td><code>5 &gt; 3</code></td><td><code>sotti</code></td></tr>
            <tr><td><code>&lt;=</code></td><td>Less or Equal</td><td><code>5 &lt;= 5</code></td><td><code>sotti</code></td></tr>
            <tr><td><code>&gt;=</code></td><td>Greater or Equal</td><td><code>5 &gt;= 3</code></td><td><code>sotti</code></td></tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`dhoro x = 10;
dhoro y = 5;

dekho(x == y);   // mittha (false)
dekho(x != y);   // sotti (true)
dekho(x < y);    // mittha
dekho(x > y);    // sotti
dekho(x <= 10);  // sotti
dekho(x >= 10);  // sotti

// Comparing strings
dekho("apple" == "apple");  // sotti
dekho("abc" < "abd");       // sotti (lexicographic)`}
      />

      <h2>Logical Operators</h2>

      <p>
        BanglaCode uses Bengali keywords for logical operators:
      </p>

      <div className="overflow-x-auto my-6">
        <table>
          <thead>
            <tr>
              <th>Operator</th>
              <th>Bengali</th>
              <th>Meaning</th>
              <th>Example</th>
              <th>Result</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>ebong</code></td><td>এবং</td><td>AND</td><td><code>sotti ebong mittha</code></td><td><code>mittha</code></td></tr>
            <tr><td><code>ba</code></td><td>বা</td><td>OR</td><td><code>sotti ba mittha</code></td><td><code>sotti</code></td></tr>
            <tr><td><code>na</code> or <code>!</code></td><td>না</td><td>NOT</td><td><code>na sotti</code></td><td><code>mittha</code></td></tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`// AND - both must be true
dekho(sotti ebong sotti);   // sotti
dekho(sotti ebong mittha);  // mittha
dekho(mittha ebong mittha); // mittha

// OR - at least one must be true
dekho(sotti ba sotti);   // sotti
dekho(sotti ba mittha);  // sotti
dekho(mittha ba mittha); // mittha

// NOT - inverts the value
dekho(na sotti);   // mittha
dekho(na mittha);  // sotti
dekho(!sotti);     // mittha (alternative syntax)

// Complex expressions
dhoro age = 25;
dhoro hasLicense = sotti;

jodi (age >= 18 ebong hasLicense) {
    dekho("Can drive!");
}

jodi (age < 13 ba age > 65) {
    dekho("Special discount!");
}`}
      />

      <h2>Assignment Operators</h2>

      <div className="overflow-x-auto my-6">
        <table>
          <thead>
            <tr>
              <th>Operator</th>
              <th>Example</th>
              <th>Equivalent To</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>=</code></td><td><code>x = 5</code></td><td>Assign 5 to x</td></tr>
            <tr><td><code>+=</code></td><td><code>x += 3</code></td><td><code>x = x + 3</code></td></tr>
            <tr><td><code>-=</code></td><td><code>x -= 3</code></td><td><code>x = x - 3</code></td></tr>
            <tr><td><code>*=</code></td><td><code>x *= 3</code></td><td><code>x = x * 3</code></td></tr>
            <tr><td><code>/=</code></td><td><code>x /= 3</code></td><td><code>x = x / 3</code></td></tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`dhoro x = 10;

x += 5;   // x is now 15
dekho(x);

x -= 3;   // x is now 12
dekho(x);

x *= 2;   // x is now 24
dekho(x);

x /= 4;   // x is now 6
dekho(x);`}
      />

      <h2>Unary Operators</h2>

      <div className="overflow-x-auto my-6">
        <table>
          <thead>
            <tr>
              <th>Operator</th>
              <th>Name</th>
              <th>Example</th>
              <th>Result</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>-</code></td><td>Negation</td><td><code>-5</code></td><td>-5</td></tr>
            <tr><td><code>!</code> or <code>na</code></td><td>Logical NOT</td><td><code>!sotti</code></td><td><code>mittha</code></td></tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`dhoro num = 5;
dekho(-num);     // -5
dekho(-(-num));  // 5

dhoro flag = sotti;
dekho(!flag);    // mittha
dekho(na flag);  // mittha
dekho(!!flag);   // sotti (double negation)`}
      />

      <h2>Member Access Operators</h2>

      <div className="overflow-x-auto my-6">
        <table>
          <thead>
            <tr>
              <th>Operator</th>
              <th>Name</th>
              <th>Example</th>
              <th>Use Case</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>.</code></td><td>Dot notation</td><td><code>obj.prop</code></td><td>Access known property</td></tr>
            <tr><td><code>[]</code></td><td>Bracket notation</td><td><code>obj[&quot;prop&quot;]</code></td><td>Dynamic property access</td></tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`// Object property access
dhoro person = {naam: "Rahim", boyosh: 25};

dekho(person.naam);        // "Rahim" (dot notation)
dekho(person["boyosh"]);   // 25 (bracket notation)

// Dynamic property name
dhoro key = "naam";
dekho(person[key]);        // "Rahim"

// Array index access
dhoro arr = [10, 20, 30, 40];
dekho(arr[0]);   // 10
dekho(arr[2]);   // 30

// String character access
dhoro str = "Hello";
dekho(str[0]);   // "H"
dekho(str[4]);   // "o"

// Nested access
dhoro data = {
    users: [
        {naam: "A"},
        {naam: "B"}
    ]
};
dekho(data.users[0].naam);  // "A"`}
      />

      <h2>Operator Precedence</h2>

      <p>
        Operators are evaluated in the following order (highest to lowest):
      </p>

      <div className="overflow-x-auto my-6">
        <table>
          <thead>
            <tr>
              <th>Priority</th>
              <th>Operators</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr><td>1 (highest)</td><td><code>[]</code>, <code>.</code></td><td>Member access</td></tr>
            <tr><td>2</td><td><code>()</code></td><td>Function call</td></tr>
            <tr><td>3</td><td><code>-</code>, <code>!</code>, <code>na</code></td><td>Unary operators</td></tr>
            <tr><td>4</td><td><code>**</code></td><td>Exponentiation</td></tr>
            <tr><td>5</td><td><code>*</code>, <code>/</code>, <code>%</code></td><td>Multiplication, division</td></tr>
            <tr><td>6</td><td><code>+</code>, <code>-</code></td><td>Addition, subtraction</td></tr>
            <tr><td>7</td><td><code>&lt;</code>, <code>&gt;</code>, <code>&lt;=</code>, <code>&gt;=</code></td><td>Comparison</td></tr>
            <tr><td>8</td><td><code>==</code>, <code>!=</code></td><td>Equality</td></tr>
            <tr><td>9</td><td><code>ebong</code></td><td>Logical AND</td></tr>
            <tr><td>10</td><td><code>ba</code></td><td>Logical OR</td></tr>
            <tr><td>11 (lowest)</td><td><code>=</code>, <code>+=</code>, etc.</td><td>Assignment</td></tr>
          </tbody>
        </table>
      </div>

      <CodeBlock
        code={`// Precedence examples
dekho(2 + 3 * 4);       // 14 (not 20, * before +)
dekho((2 + 3) * 4);     // 20 (parentheses override)

dekho(2 ** 3 ** 2);     // 512 (right-to-left: 2^(3^2) = 2^9)

dekho(5 > 3 ebong 2 < 4);  // sotti (comparison before AND)

// Use parentheses for clarity
dhoro result = (a + b) * (c - d);`}
      />

      <DocNavigation currentPath="/docs/operators" />
    </div>
  );
}

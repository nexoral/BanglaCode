import Link from "next/link";

export default function OOP() {
  return (
    <div className="space-y-6">
      <h1>Object Oriented Programming</h1>
      <p className="lead text-xl text-muted-foreground">
        Classes, Objects, Getters, Setters, Static Properties, and Private Fields in BanglaCode.
      </p>

      <h2>Classes (`sreni`)</h2>
      <p>Define a class using <code>sreni</code> and constructor with <code>shuru</code> (constructor).</p>

      <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
        <pre><code><span className="text-purple-400">sreni</span> Manush &#123;
          <span className="text-yellow-200">shuru</span>(naam, boyosh) &#123;
          <span className="text-blue-400">ei</span>.naam = naam;
          <span className="text-blue-400">ei</span>.boyosh = boyosh;
          &#125;

          <span className="text-purple-400">kaj</span> <span className="text-yellow-200">porichoy</span>() &#123;
          <span className="text-yellow-200">dekho</span>(<span className="text-green-400">"Amar naam "</span>, <span className="text-blue-400">ei</span>.naam);
          &#125;
          &#125;</code></pre>
      </div>

      <h2>Creating Objects</h2>
      <p>Use <code>notun</code> to create an instance.</p>
      <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
        <pre><code><span className="text-blue-400">dhoro</span> person = <span className="text-purple-400">notun</span> Manush(<span className="text-green-400">"Ankan"</span>, 25);
          person.porichoy();</code></pre>
      </div>

      <h2>Getters (`pao`)</h2>
      <p>Getters allow you to define computed properties that are accessed like regular properties.</p>
      <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
        <pre><code><span className="text-purple-400">sreni</span> Person &#123;
          <span className="text-yellow-200">shuru</span>(naam, boichhor) &#123;
          <span className="text-blue-400">ei</span>.naam = naam;
          <span className="text-blue-400">ei</span>.boichhor = boichhor;
          &#125;

          <span className="text-purple-400">pao</span> <span className="text-yellow-200">boshi</span>() &#123;
          <span className="text-blue-400">dhoro</span> current = 2026;
          <span className="text-pink-400">ferao</span> current - <span className="text-blue-400">ei</span>.boichhor;
          &#125;
          &#125;

          <span className="text-blue-400">dhoro</span> p = <span className="text-purple-400">notun</span> Person(<span className="text-green-400">"Ankan"</span>, 1995);
          <span className="text-yellow-200">dekho</span>(p.boshi); <span className="text-gray-400">// 31 (computed automatically)</span></code></pre>
      </div>

      <h2>Setters (`set`)</h2>
      <p>Setters allow you to define how properties are assigned, with validation or transformation.</p>
      <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
        <pre><code><span className="text-purple-400">sreni</span> Rectangle &#123;
          <span className="text-yellow-200">shuru</span>() &#123;
          <span className="text-blue-400">ei</span>._width = 0;
          <span className="text-blue-400">ei</span>._height = 0;
          &#125;

          <span className="text-purple-400">pao</span> <span className="text-yellow-200">area</span>() &#123;
          <span className="text-pink-400">ferao</span> <span className="text-blue-400">ei</span>._width * <span className="text-blue-400">ei</span>._height;
          &#125;

          <span className="text-purple-400">set</span> <span className="text-yellow-200">width</span>(w) &#123;
          <span className="text-blue-400">ei</span>._width = w;
          &#125;

          <span className="text-purple-400">set</span> <span className="text-yellow-200">height</span>(h) &#123;
          <span className="text-blue-400">ei</span>._height = h;
          &#125;
          &#125;

          <span className="text-blue-400">dhoro</span> rect = <span className="text-purple-400">notun</span> Rectangle();
          rect.width = 10;  <span className="text-gray-400">// Calls setter</span>
          rect.height = 5;  <span className="text-gray-400">// Calls setter</span>
          <span className="text-yellow-200">dekho</span>(rect.area); <span className="text-gray-400">// 50</span></code></pre>
      </div>

      <h2>Static Properties (`sthir`)</h2>
      <p>Static properties belong to the class itself, not instances. Access them via <code>ClassName.property</code>.</p>
      <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
        <pre><code><span className="text-purple-400">sreni</span> Circle &#123;
          <span className="text-pink-400">sthir</span> PI = 3.14159;

          <span className="text-yellow-200">shuru</span>(radius) &#123;
          <span className="text-blue-400">ei</span>.radius = radius;
          &#125;

          <span className="text-purple-400">kaj</span> <span className="text-yellow-200">area</span>() &#123;
          <span className="text-pink-400">ferao</span> Circle.PI * <span className="text-blue-400">ei</span>.radius * <span className="text-blue-400">ei</span>.radius;
          &#125;
          &#125;

          <span className="text-yellow-200">dekho</span>(Circle.PI); <span className="text-gray-400">// 3.14159</span>
          <span className="text-blue-400">dhoro</span> c = <span className="text-purple-400">notun</span> Circle(10);
          <span className="text-yellow-200">dekho</span>(c.area()); <span className="text-gray-400">// 314.159</span></code></pre>
      </div>

      <h2>Private Fields (Convention: `_prefix`)</h2>
      <p>Use underscore prefix for private fields. While accessible, it signals internal use only.</p>
      <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
        <pre><code><span className="text-purple-400">sreni</span> BankAccount &#123;
          <span className="text-yellow-200">shuru</span>(balance) &#123;
          <span className="text-blue-400">ei</span>._balance = balance;  <span className="text-gray-400">// Private field</span>
          &#125;

          <span className="text-purple-400">kaj</span> <span className="text-yellow-200">deposit</span>(amount) &#123;
          <span className="text-blue-400">ei</span>._balance = <span className="text-blue-400">ei</span>._balance + amount;
          &#125;

          <span className="text-purple-400">kaj</span> <span className="text-yellow-200">withdraw</span>(amount) &#123;
          <span className="text-pink-400">jodi</span> (amount &lt;= <span className="text-blue-400">ei</span>._balance) &#123;
          <span className="text-blue-400">ei</span>._balance = <span className="text-blue-400">ei</span>._balance - amount;
          <span className="text-pink-400">ferao</span> sotti;
          &#125;
          <span className="text-pink-400">ferao</span> mittha;
          &#125;

          <span className="text-purple-400">pao</span> <span className="text-yellow-200">balance</span>() &#123;
          <span className="text-pink-400">ferao</span> <span className="text-blue-400">ei</span>._balance;  <span className="text-gray-400">// Controlled access</span>
          &#125;
          &#125;

          <span className="text-blue-400">dhoro</span> account = <span className="text-purple-400">notun</span> BankAccount(1000);
          account.deposit(500);
          <span className="text-yellow-200">dekho</span>(account.balance); <span className="text-gray-400">// 1500 (via getter)</span></code></pre>
      </div>

      <h2>Real-World Example: Temperature Converter</h2>
      <p>Combining getters and setters for bidirectional conversion:</p>
      <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
        <pre><code><span className="text-purple-400">sreni</span> Temperature &#123;
          <span className="text-yellow-200">shuru</span>() &#123;
          <span className="text-blue-400">ei</span>._celsius = 0;
          &#125;

          <span className="text-purple-400">pao</span> <span className="text-yellow-200">celsius</span>() &#123;
          <span className="text-pink-400">ferao</span> <span className="text-blue-400">ei</span>._celsius;
          &#125;

          <span className="text-purple-400">set</span> <span className="text-yellow-200">celsius</span>(c) &#123;
          <span className="text-blue-400">ei</span>._celsius = c;
          &#125;

          <span className="text-purple-400">pao</span> <span className="text-yellow-200">fahrenheit</span>() &#123;
          <span className="text-pink-400">ferao</span> (<span className="text-blue-400">ei</span>._celsius * 9 / 5) + 32;
          &#125;

          <span className="text-purple-400">set</span> <span className="text-yellow-200">fahrenheit</span>(f) &#123;
          <span className="text-blue-400">ei</span>._celsius = (f - 32) * 5 / 9;
          &#125;
          &#125;

          <span className="text-blue-400">dhoro</span> temp = <span className="text-purple-400">notun</span> Temperature();
          temp.celsius = 100;
          <span className="text-yellow-200">dekho</span>(temp.fahrenheit); <span className="text-gray-400">// 212</span>

          temp.fahrenheit = 32;
          <span className="text-yellow-200">dekho</span>(temp.celsius); <span className="text-gray-400">// 0</span></code></pre>
      </div>

      <h2>Real-World Example: Product with Tax</h2>
      <p>Using static properties for shared constants:</p>
      <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
        <pre><code><span className="text-purple-400">sreni</span> Product &#123;
          <span className="text-pink-400">sthir</span> TAX_RATE = 0.15;  <span className="text-gray-400">// 15% tax</span>

          <span className="text-yellow-200">shuru</span>(naam, dam, quantity) &#123;
          <span className="text-blue-400">ei</span>._naam = naam;
          <span className="text-blue-400">ei</span>._dam = dam;
          <span className="text-blue-400">ei</span>._quantity = quantity;
          &#125;

          <span className="text-purple-400">pao</span> <span className="text-yellow-200">subtotal</span>() &#123;
          <span className="text-pink-400">ferao</span> <span className="text-blue-400">ei</span>._dam * <span className="text-blue-400">ei</span>._quantity;
          &#125;

          <span className="text-purple-400">pao</span> <span className="text-yellow-200">tax</span>() &#123;
          <span className="text-pink-400">ferao</span> <span className="text-blue-400">ei</span>.subtotal * Product.TAX_RATE;
          &#125;

          <span className="text-purple-400">pao</span> <span className="text-yellow-200">total</span>() &#123;
          <span className="text-pink-400">ferao</span> <span className="text-blue-400">ei</span>.subtotal + <span className="text-blue-400">ei</span>.tax;
          &#125;
          &#125;

          <span className="text-blue-400">dhoro</span> product = <span className="text-purple-400">notun</span> Product(<span className="text-green-400">"Laptop"</span>, 1000, 2);
          <span className="text-yellow-200">dekho</span>(product.subtotal); <span className="text-gray-400">// 2000</span>
          <span className="text-yellow-200">dekho</span>(product.tax);      <span className="text-gray-400">// 300</span>
          <span className="text-yellow-200">dekho</span>(product.total);    <span className="text-gray-400">// 2300</span></code></pre>
      </div>

      <h2>OOP Features Summary</h2>
      <div className="space-y-4">
        <div className="border border-border rounded-lg p-4">
          <h3 className="text-lg font-semibold mb-2">Getters (`pao`)</h3>
          <p className="text-sm text-muted-foreground">
            • Computed properties accessed like regular properties<br/>
            • No parameters, must return a value<br/>
            • Perfect for derived data and calculations
          </p>
        </div>

        <div className="border border-border rounded-lg p-4">
          <h3 className="text-lg font-semibold mb-2">Setters (`set`)</h3>
          <p className="text-sm text-muted-foreground">
            • Control how properties are assigned<br/>
            • Exactly one parameter required<br/>
            • Enable validation and transformation
          </p>
        </div>

        <div className="border border-border rounded-lg p-4">
          <h3 className="text-lg font-semibold mb-2">Static Properties (`sthir`)</h3>
          <p className="text-sm text-muted-foreground">
            • Belong to the class, not instances<br/>
            • Access via <code>ClassName.property</code><br/>
            • Shared across all instances<br/>
            • Can be modified at runtime
          </p>
        </div>

        <div className="border border-border rounded-lg p-4">
          <h3 className="text-lg font-semibold mb-2">Private Fields (`_prefix`)</h3>
          <p className="text-sm text-muted-foreground">
            • Use underscore prefix for internal fields<br/>
            • Signals "internal use only" by convention<br/>
            • Access via getters/setters for encapsulation<br/>
            • Separate storage from public properties
          </p>
        </div>
      </div>

      <div className="mt-8 p-4 bg-blue-500/10 border border-blue-500/20 rounded-lg">
        <h3 className="text-lg font-semibold mb-2">💡 Best Practices</h3>
        <ul className="list-disc list-inside space-y-1 text-sm">
          <li>Use getters for computed properties that don't require parameters</li>
          <li>Use setters for validation before assignment</li>
          <li>Use static properties for constants shared across instances</li>
          <li>Use `_prefix` for private fields and expose via getters/setters</li>
          <li>Combine getters and setters for bidirectional conversions</li>
          <li>Access static properties via <code>ClassName.property</code>, not <code>instance.property</code></li>
        </ul>
      </div>
    </div>
  );
}

import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function Functions() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Functions & OOP
        </span>
      </div>

      <h1>Functions</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        Functions in BanglaCode are defined using the <code>kaj</code> keyword (meaning &quot;work&quot;
        or &quot;task&quot;). Functions are first-class values and support closures.
      </p>

      <h2>New Utility Methods</h2>

      <p>
        BanglaCode now includes additional JavaScript-style helpers for arrays, strings, number parsing,
        and URI encoding.
      </p>

      <CodeBlock
        code={`// Array search helpers
dhoro first = khojo_prothom([1, 3, 8, 10], kaj(x) { ferao x > 5; });      // 8
dhoro idx = khojo_index([1, 3, 8, 10], kaj(x) { ferao x > 5; });           // 2
dhoro last = khojo_shesh([1, 3, 8, 10], kaj(x) { ferao x % 2 == 0; });     // 10

// String helpers
ache_text("banglacode", "code");        // sotti
shuru_diye("banglacode", "bang");       // sotti
shesh_diye("banglacode", "code");       // sotti
baro("ha", 3);                          // "hahaha"
text_at("bangla", -1);                  // "a"

// Number + URI helpers
purno_sonkhya("42");                    // 42
doshomik_sonkhya("3.14abc");            // 3.14
sonkhya_na("abc");                      // sotti
uri_ongsho_encode("hello world");       // "hello%20world"

// Date + Regex helpers
dhoro ts = tarikh_ekhon();
tarikh_format(ts, "2006-01-02");
regex_test("[a-z]+", "bangla");         // sotti
regex_search("la", "bangla");           // 4

// Object helpers
nijer_ache({a: 1}, "a");                // sotti
jora_theke([["x", 10], ["y", 20]]);     // {x: 10, y: 20}
ekoi_ki(1, 1);                           // sotti

// Timers
dhoro id = setInterval(kaj() { dekho("tick"); }, 1000);
clearInterval(id);`}
      />

      <h2>Defining Functions</h2>

      <p>
        Use the <code>kaj</code> keyword followed by the function name and parameters:
      </p>

      <CodeBlock
        code={`// Basic function definition
kaj greet() {
    dekho("Namaskar!");
}

// Call the function
greet();  // Output: Namaskar!`}
      />

      <h3>Functions with Parameters</h3>

      <CodeBlock
        code={`kaj greetPerson(naam) {
    dekho("Namaskar,", naam);
}

greetPerson("Rahim");   // Namaskar, Rahim
greetPerson("Karim");   // Namaskar, Karim

// Multiple parameters
kaj add(a, b) {
    dekho(a, "+", b, "=", a + b);
}

add(5, 3);  // 5 + 3 = 8`}
      />

      <h2>Return Values</h2>

      <p>
        Use the <code>ferao</code> keyword (meaning &quot;return&quot;) to return a value:
      </p>

      <CodeBlock
        code={`kaj add(a, b) {
    ferao a + b;
}

dhoro sum = add(10, 20);
dekho(sum);  // 30

// Use return value directly
dekho(add(5, 3) * 2);  // 16

// Function without return returns null
kaj sayHello() {
    dekho("Hello!");
}

dhoro result = sayHello();
dekho(result == khali);  // sotti`}
      />

      <h3>Early Return</h3>

      <CodeBlock
        code={`kaj getGrade(score) {
    jodi (score >= 90) {
        ferao "A";
    }
    jodi (score >= 80) {
        ferao "B";
    }
    jodi (score >= 70) {
        ferao "C";
    }
    jodi (score >= 60) {
        ferao "D";
    }
    ferao "F";
}

dekho(getGrade(85));  // B
dekho(getGrade(55));  // F`}
      />

      <h2>Function Expressions</h2>

      <p>
        Functions can be assigned to variables (anonymous functions):
      </p>

      <CodeBlock
        code={`// Anonymous function assigned to variable
dhoro multiply = kaj(x, y) {
    ferao x * y;
};

dekho(multiply(4, 5));  // 20

// Functions as array elements
dhoro operations = [
    kaj(a, b) { ferao a + b; },
    kaj(a, b) { ferao a - b; },
    kaj(a, b) { ferao a * b; },
    kaj(a, b) { ferao a / b; }
];

dekho(operations[0](10, 5));  // 15 (addition)
dekho(operations[2](10, 5));  // 50 (multiplication)`}
      />

      <h2>Higher-Order Functions</h2>

      <p>
        Functions can take other functions as arguments or return functions:
      </p>

      <CodeBlock
        code={`// Function that takes a function as argument
kaj applyOperation(a, b, operation) {
    ferao operation(a, b);
}

kaj add(x, y) { ferao x + y; }
kaj multiply(x, y) { ferao x * y; }

dekho(applyOperation(5, 3, add));       // 8
dekho(applyOperation(5, 3, multiply));  // 15

// Function that returns a function
kaj makeMultiplier(factor) {
    ferao kaj(n) {
        ferao n * factor;
    };
}

dhoro double = makeMultiplier(2);
dhoro triple = makeMultiplier(3);

dekho(double(5));   // 10
dekho(triple(5));   // 15`}
      />

      <h2>Closures</h2>

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

dhoro counter1 = makeCounter();
dhoro counter2 = makeCounter();

dekho(counter1());  // 1
dekho(counter1());  // 2
dekho(counter1());  // 3

dekho(counter2());  // 1 (separate counter)
dekho(counter2());  // 2`}
      />

      <h3>Closure with State</h3>

      <CodeBlock
        code={`kaj createBankAccount(initialBalance) {
    dhoro balance = initialBalance;

    ferao {
        deposit: kaj(amount) {
            balance = balance + amount;
            dekho("Deposited:", amount, "- Balance:", balance);
        },
        withdraw: kaj(amount) {
            jodi (amount > balance) {
                dekho("Insufficient funds!");
                ferao mittha;
            }
            balance = balance - amount;
            dekho("Withdrew:", amount, "- Balance:", balance);
            ferao sotti;
        },
        getBalance: kaj() {
            ferao balance;
        }
    };
}

dhoro account = createBankAccount(1000);
account.deposit(500);     // Deposited: 500 - Balance: 1500
account.withdraw(200);    // Withdrew: 200 - Balance: 1300
dekho(account.getBalance());  // 1300`}
      />

      <h2>Recursion</h2>

      <p>
        Functions can call themselves:
      </p>

      <CodeBlock
        code={`// Factorial
kaj factorial(n) {
    jodi (n <= 1) {
        ferao 1;
    }
    ferao n * factorial(n - 1);
}

dekho(factorial(5));  // 120

// Fibonacci
kaj fibonacci(n) {
    jodi (n <= 1) {
        ferao n;
    }
    ferao fibonacci(n - 1) + fibonacci(n - 2);
}

dekho(fibonacci(10));  // 55

// Sum of array using recursion
kaj sumArray(arr, index) {
    jodi (index >= dorghyo(arr)) {
        ferao 0;
    }
    ferao arr[index] + sumArray(arr, index + 1);
}

dekho(sumArray([1, 2, 3, 4, 5], 0));  // 15`}
      />

      <h2>Default-Like Behavior</h2>

      <p>
        BanglaCode doesn&apos;t have default parameters, but you can simulate them:
      </p>

      <CodeBlock
        code={`kaj greet(naam, greeting) {
    // Check if parameter is provided
    jodi (greeting == khali) {
        greeting = "Namaskar";
    }
    dekho(greeting + ",", naam);
}

greet("Rahim", "Hello");     // Hello, Rahim
greet("Karim", khali);       // Namaskar, Karim`}
      />

      <h2>Rest Parameters</h2>

      <p>
        Use <code>...</code> (spread/rest operator) to collect remaining arguments into an array:
      </p>

      <CodeBlock
        code={`// Variadic function with rest parameter
kaj sum(...numbers) {
    dhoro total = 0;
    ghuriye (dhoro i = 0; i < dorghyo(numbers); i = i + 1) {
        total = total + numbers[i];
    }
    ferao total;
}

dekho(sum(1, 2, 3));           // 6
dekho(sum(1, 2, 3, 4, 5));     // 15
dekho(sum());                   // 0

// Mixed regular parameters with rest
kaj greetAll(greeting, ...names) {
    ghuriye (dhoro i = 0; i < dorghyo(names); i = i + 1) {
        dekho(greeting, names[i]);
    }
}

greetAll("Hello", "Alice", "Bob", "Charlie");
// Hello Alice
// Hello Bob
// Hello Charlie`}
      />

      <h2>Spread Operator</h2>

      <p>
        Use <code>...</code> to expand arrays in function calls or array literals:
      </p>

      <CodeBlock
        code={`// Spread in function calls
kaj sum(...numbers) {
    dhoro total = 0;
    ghuriye (dhoro i = 0; i < dorghyo(numbers); i = i + 1) {
        total = total + numbers[i];
    }
    ferao total;
}

dhoro nums = [1, 2, 3, 4, 5];
dekho(sum(...nums));  // 15

// Combine with regular arguments
dekho(sum(10, ...nums, 20));  // 10 + 1 + 2 + 3 + 4 + 5 + 20 = 45

// Spread in array literals
dhoro arr1 = [1, 2];
dhoro arr2 = [3, 4];
dhoro combined = [...arr1, ...arr2];  // [1, 2, 3, 4]
dhoro withExtra = [0, ...arr1, 99];   // [0, 1, 2, 99]

// Clone an array
dhoro original = [1, 2, 3];
dhoro copy = [...original];

// Spread with dekho
dhoro items = ["apple", "banana", "cherry"];
dekho(...items);  // apple banana cherry`}
      />

      <h2>Callback Pattern</h2>

      <CodeBlock
        code={`// Process array with callback
kaj forEach(arr, callback) {
    ghuriye (dhoro i = 0; i < dorghyo(arr); i = i + 1) {
        callback(arr[i], i);
    }
}

dhoro numbers = [1, 2, 3, 4, 5];

forEach(numbers, kaj(value, index) {
    dekho("Index", index, ":", value);
});

// Custom map function
kaj map(arr, transform) {
    dhoro result = [];
    ghuriye (dhoro i = 0; i < dorghyo(arr); i = i + 1) {
        dhokao(result, transform(arr[i]));
    }
    ferao result;
}

dhoro doubled = map([1, 2, 3], kaj(n) { ferao n * 2; });
dekho(doubled);  // [2, 4, 6]

// Custom filter function
kaj filter(arr, predicate) {
    dhoro result = [];
    ghuriye (dhoro i = 0; i < dorghyo(arr); i = i + 1) {
        jodi (predicate(arr[i])) {
            dhokao(result, arr[i]);
        }
    }
    ferao result;
}

dhoro evens = filter([1, 2, 3, 4, 5, 6], kaj(n) { ferao n % 2 == 0; });
dekho(evens);  // [2, 4, 6]`}
      />

      <h2>Practical Examples</h2>

      <h3>Memoization</h3>

      <CodeBlock
        code={`// Memoized Fibonacci
kaj createMemoizedFib() {
    dhoro cache = {};

    ferao kaj(n) {
        jodi (n <= 1) {
            ferao n;
        }

        dhoro key = lipi(n);
        jodi (cache[key] != khali) {
            ferao cache[key];
        }

        dhoro result = ei(n - 1) + ei(n - 2);
        cache[key] = result;
        ferao result;
    };
}

dhoro fib = createMemoizedFib();
dekho(fib(40));  // Much faster with memoization!`}
      />

      <h3>Compose Functions</h3>

      <CodeBlock
        code={`kaj compose(f, g) {
    ferao kaj(x) {
        ferao f(g(x));
    };
}

kaj addOne(n) { ferao n + 1; }
kaj double(n) { ferao n * 2; }

dhoro addOneThenDouble = compose(double, addOne);
dekho(addOneThenDouble(5));  // 12 ((5+1)*2)

dhoro doubleThenAddOne = compose(addOne, double);
dekho(doubleThenAddOne(5));  // 11 ((5*2)+1)`}
      />

      <DocNavigation currentPath="/docs/functions" />
    </div>
  );
}

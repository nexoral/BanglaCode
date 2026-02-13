import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function Classes() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Functions & OOP
        </span>
      </div>

      <h1>Classes</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        BanglaCode supports object-oriented programming with classes using the <code>sreni</code>
        keyword (meaning &quot;class&quot; or &quot;category&quot;).
      </p>

      <h2>Defining a Class</h2>

      <p>
        Use the <code>sreni</code> keyword to define a class:
      </p>

      <CodeBlock
        code={`// Basic class definition
sreni Person {
    // Constructor
    shuru(naam, boyosh) {
        ei.naam = naam;
        ei.boyosh = boyosh;
    }

    // Method
    kaj introduce() {
        dekho("Hi, I am", ei.naam, "and I am", ei.boyosh, "years old");
    }
}`}
      />

      <h2>Creating Instances</h2>

      <p>
        Use the <code>notun</code> keyword (meaning &quot;new&quot;) to create an instance:
      </p>

      <CodeBlock
        code={`// Create instance
dhoro person1 = notun Person("Rahim", 25);
dhoro person2 = notun Person("Karim", 30);

// Call methods
person1.introduce();  // Hi, I am Rahim and I am 25 years old
person2.introduce();  // Hi, I am Karim and I am 30 years old

// Access properties
dekho(person1.naam);   // Rahim
dekho(person2.boyosh); // 30`}
      />

      <h2>Constructor (shuru)</h2>

      <p>
        The constructor is defined using the <code>shuru</code> keyword (meaning &quot;start&quot; or &quot;begin&quot;).
        It&apos;s called automatically when creating a new instance:
      </p>

      <CodeBlock
        code={`sreni Rectangle {
    shuru(width, height) {
        ei.width = width;
        ei.height = height;
        dekho("Created rectangle:", width, "x", height);
    }

    kaj area() {
        ferao ei.width * ei.height;
    }

    kaj perimeter() {
        ferao 2 * (ei.width + ei.height);
    }
}

dhoro rect = notun Rectangle(10, 5);
// Output: Created rectangle: 10 x 5

dekho("Area:", rect.area());           // Area: 50
dekho("Perimeter:", rect.perimeter()); // Perimeter: 30`}
      />

      <h2>The &apos;ei&apos; Keyword (this)</h2>

      <p>
        The <code>ei</code> keyword (meaning &quot;this&quot;) refers to the current instance:
      </p>

      <CodeBlock
        code={`sreni Counter {
    shuru(start) {
        ei.count = start;
    }

    kaj increment() {
        ei.count = ei.count + 1;
        ferao ei;  // Return this for chaining
    }

    kaj decrement() {
        ei.count = ei.count - 1;
        ferao ei;
    }

    kaj getValue() {
        ferao ei.count;
    }
}

dhoro counter = notun Counter(0);
counter.increment().increment().increment();
dekho(counter.getValue());  // 3`}
      />

      <h2>Methods</h2>

      <p>
        Methods are defined using the <code>kaj</code> keyword inside the class:
      </p>

      <CodeBlock
        code={`sreni BankAccount {
    shuru(owner, balance) {
        ei.owner = owner;
        ei.balance = balance;
    }

    kaj deposit(amount) {
        jodi (amount > 0) {
            ei.balance = ei.balance + amount;
            dekho("Deposited:", amount);
        } nahole {
            dekho("Invalid deposit amount");
        }
    }

    kaj withdraw(amount) {
        jodi (amount > ei.balance) {
            dekho("Insufficient funds!");
            ferao mittha;
        }
        ei.balance = ei.balance - amount;
        dekho("Withdrew:", amount);
        ferao sotti;
    }

    kaj getBalance() {
        ferao ei.balance;
    }

    kaj transfer(toAccount, amount) {
        jodi (ei.withdraw(amount)) {
            toAccount.deposit(amount);
            dekho("Transferred", amount, "to", toAccount.owner);
        }
    }
}

dhoro account1 = notun BankAccount("Rahim", 1000);
dhoro account2 = notun BankAccount("Karim", 500);

account1.deposit(500);
account1.transfer(account2, 300);

dekho(account1.getBalance());  // 1200
dekho(account2.getBalance());  // 800`}
      />

      <h2>Instance Properties</h2>

      <CodeBlock
        code={`sreni Product {
    shuru(name, price) {
        ei.name = name;
        ei.price = price;
        ei.quantity = 0;  // Default property
    }

    kaj addStock(amount) {
        ei.quantity = ei.quantity + amount;
    }

    kaj sell(amount) {
        jodi (amount > ei.quantity) {
            dekho("Not enough stock!");
            ferao mittha;
        }
        ei.quantity = ei.quantity - amount;
        ferao sotti;
    }

    kaj getValue() {
        ferao ei.price * ei.quantity;
    }
}

dhoro laptop = notun Product("Laptop", 50000);
laptop.addStock(10);

// Access and modify properties directly
dekho(laptop.name);     // Laptop
dekho(laptop.quantity); // 10

laptop.price = 55000;   // Modify property
dekho(laptop.getValue()); // 550000`}
      />

      <h2>Practical Examples</h2>

      <h3>Todo List</h3>

      <CodeBlock
        code={`sreni TodoList {
    shuru() {
        ei.tasks = [];
    }

    kaj add(task) {
        dhokao(ei.tasks, {
            text: task,
            completed: mittha
        });
        dekho("Added:", task);
    }

    kaj complete(index) {
        jodi (index >= 0 ebong index < dorghyo(ei.tasks)) {
            ei.tasks[index].completed = sotti;
            dekho("Completed:", ei.tasks[index].text);
        }
    }

    kaj list() {
        ghuriye (dhoro i = 0; i < dorghyo(ei.tasks); i = i + 1) {
            dhoro task = ei.tasks[i];
            dhoro status = task.completed ? "[X]" : "[ ]";
            dekho(status, task.text);
        }
    }
}

dhoro todos = notun TodoList();
todos.add("Learn BanglaCode");
todos.add("Build a project");
todos.add("Share with friends");
todos.complete(0);
todos.list();`}
      />

      <h3>Game Character</h3>

      <CodeBlock
        code={`sreni Character {
    shuru(name, health, attack) {
        ei.name = name;
        ei.maxHealth = health;
        ei.health = health;
        ei.attack = attack;
        ei.isAlive = sotti;
    }

    kaj takeDamage(damage) {
        ei.health = ei.health - damage;
        dekho(ei.name, "took", damage, "damage!");

        jodi (ei.health <= 0) {
            ei.health = 0;
            ei.isAlive = mittha;
            dekho(ei.name, "has been defeated!");
        } nahole {
            dekho(ei.name, "has", ei.health, "HP remaining");
        }
    }

    kaj attackEnemy(enemy) {
        jodi (ei.isAlive ebong enemy.isAlive) {
            dekho(ei.name, "attacks", enemy.name);
            enemy.takeDamage(ei.attack);
        }
    }

    kaj heal(amount) {
        jodi (ei.isAlive) {
            ei.health = ei.health + amount;
            jodi (ei.health > ei.maxHealth) {
                ei.health = ei.maxHealth;
            }
            dekho(ei.name, "healed to", ei.health, "HP");
        }
    }
}

dhoro hero = notun Character("Hero", 100, 25);
dhoro monster = notun Character("Goblin", 60, 15);

hero.attackEnemy(monster);
monster.attackEnemy(hero);
hero.heal(10);
hero.attackEnemy(monster);
hero.attackEnemy(monster);`}
      />

      <h3>Stack Data Structure</h3>

      <CodeBlock
        code={`sreni Stack {
    shuru() {
        ei.items = [];
    }

    kaj push(item) {
        dhokao(ei.items, item);
    }

    kaj pop() {
        jodi (ei.isEmpty()) {
            dekho("Stack is empty!");
            ferao khali;
        }
        ferao berKoro(ei.items);
    }

    kaj peek() {
        jodi (ei.isEmpty()) {
            ferao khali;
        }
        ferao ei.items[dorghyo(ei.items) - 1];
    }

    kaj isEmpty() {
        ferao dorghyo(ei.items) == 0;
    }

    kaj size() {
        ferao dorghyo(ei.items);
    }
}

dhoro stack = notun Stack();
stack.push(1);
stack.push(2);
stack.push(3);

dekho(stack.peek());  // 3
dekho(stack.pop());   // 3
dekho(stack.pop());   // 2
dekho(stack.size());  // 1`}
      />

      <DocNavigation currentPath="/docs/classes" />
    </div>
  );
}

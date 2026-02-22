package test

import (
	"testing"

	"BanglaCode/src/evaluator"
	"BanglaCode/src/evaluator/builtins"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
)

func evalOOPInput(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	builtins.InitializeEnvironmentWithConstants(env)
	return evaluator.Eval(program, env)
}

// Test getter methods
func TestClassGetter(t *testing.T) {
	input := `
		sreni Person {
			shuru(naam, boichhor) {
				ei.naam = naam;
				ei.boichhor = boichhor;
			}

			pao boshi() {
				dhoro current = 2026;
				ferao current - ei.boichhor;
			}
		}

		dhoro p = notun Person("Ankan", 1995);
		p.boshi
	`

	result := evalOOPInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if num.Value != 31 {
		t.Errorf("Expected boshi = 31, got %f", num.Value)
	}
}

// Test setter methods
func TestClassSetter(t *testing.T) {
	input := `
		sreni Rectangle {
			shuru() {
				ei._width = 0;
				ei._height = 0;
			}

			pao area() {
				ferao ei._width * ei._height;
			}

			set width(w) {
				ei._width = w;
			}

			set height(h) {
				ei._height = h;
			}
		}

		dhoro rect = notun Rectangle();
		rect.width = 10;
		rect.height = 5;
		rect.area
	`

	result := evalOOPInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if num.Value != 50 {
		t.Errorf("Expected area = 50, got %f", num.Value)
	}
}

// Test getter and setter together
func TestGetterSetterTogether(t *testing.T) {
	input := `
		sreni Temperature {
			shuru() {
				ei._celsius = 0;
			}

			pao celsius() {
				ferao ei._celsius;
			}

			set celsius(c) {
				ei._celsius = c;
			}

			pao fahrenheit() {
				ferao (ei._celsius * 9 / 5) + 32;
			}

			set fahrenheit(f) {
				ei._celsius = (f - 32) * 5 / 9;
			}
		}

		dhoro temp = notun Temperature();
		temp.celsius = 100;
		temp.fahrenheit
	`

	result := evalOOPInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if num.Value != 212 {
		t.Errorf("Expected 212°F, got %f", num.Value)
	}
}

// Test fahrenheit setter
func TestFahrenheitSetter(t *testing.T) {
	input := `
		sreni Temperature {
			shuru() {
				ei._celsius = 0;
			}

			pao celsius() {
				ferao ei._celsius;
			}

			set fahrenheit(f) {
				ei._celsius = (f - 32) * 5 / 9;
			}
		}

		dhoro temp = notun Temperature();
		temp.fahrenheit = 32;
		temp.celsius
	`

	result := evalOOPInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if num.Value != 0 {
		t.Errorf("Expected 0°C, got %f", num.Value)
	}
}

// Test static properties
func TestStaticProperties(t *testing.T) {
	input := `
		sreni Circle {
			sthir PI = 3.14159;

			shuru(radius) {
				ei.radius = radius;
			}

			kaj area() {
				ferao Circle.PI * ei.radius * ei.radius;
			}
		}

		Circle.PI
	`

	result := evalOOPInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if num.Value != 3.14159 {
		t.Errorf("Expected PI = 3.14159, got %f", num.Value)
	}
}

// Test static property usage in instance method
func TestStaticPropertyInMethod(t *testing.T) {
	input := `
		sreni Circle {
			sthir PI = 3.14159;

			shuru(radius) {
				ei.radius = radius;
			}

			kaj area() {
				ferao Circle.PI * ei.radius * ei.radius;
			}
		}

		dhoro c = notun Circle(10);
		c.area()
	`

	result := evalOOPInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	expected := 3.14159 * 10 * 10
	if num.Value != expected {
		t.Errorf("Expected area = %f, got %f", expected, num.Value)
	}
}

// Test multiple static properties
func TestMultipleStaticProperties(t *testing.T) {
	input := `
		sreni Math {
			sthir PI = 3.14159;
			sthir E = 2.71828;
			sthir GOLDEN_RATIO = 1.618;
		}

		Math.E
	`

	result := evalOOPInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if num.Value != 2.71828 {
		t.Errorf("Expected E = 2.71828, got %f", num.Value)
	}
}

// Test private fields with underscore prefix
func TestPrivateFields(t *testing.T) {
	input := `
		sreni BankAccount {
			shuru(balance) {
				ei._balance = balance;
			}

			kaj deposit(amount) {
				ei._balance = ei._balance + amount;
			}

			kaj withdraw(amount) {
				jodi (amount <= ei._balance) {
					ei._balance = ei._balance - amount;
					ferao sotti;
				}
				ferao mittha;
			}

			pao balance() {
				ferao ei._balance;
			}
		}

		dhoro account = notun BankAccount(1000);
		account.deposit(500);
		account.withdraw(300);
		account.balance
	`

	result := evalOOPInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if num.Value != 1200 {
		t.Errorf("Expected balance = 1200, got %f", num.Value)
	}
}

// Test accessing private field directly
func TestPrivateFieldDirectAccess(t *testing.T) {
	input := `
		sreni Secret {
			shuru() {
				ei._secret = "hidden";
			}
		}

		dhoro s = notun Secret();
		s._secret
	`

	result := evalOOPInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	if str.Value != "hidden" {
		t.Errorf("Expected 'hidden', got '%s'", str.Value)
	}
}

// Test complex getter with logic
func TestComplexGetter(t *testing.T) {
	input := `
		sreni User {
			shuru(naam, admin) {
				ei.naam = naam;
				ei.admin = admin;
			}

			pao display() {
				jodi (ei.admin) {
					ferao ei.naam + " (Admin)";
				} nahole {
					ferao ei.naam + " (User)";
				}
			}
		}

		dhoro user = notun User("Ankan", sotti);
		user.display
	`

	result := evalOOPInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	if str.Value != "Ankan (Admin)" {
		t.Errorf("Expected 'Ankan (Admin)', got '%s'", str.Value)
	}
}

// Test setter with validation
func TestSetterWithValidation(t *testing.T) {
	input := `
		sreni Product {
			shuru() {
				ei._price = 0;
			}

			pao price() {
				ferao ei._price;
			}

			set price(p) {
				jodi (p >= 0) {
					ei._price = p;
				}
			}
		}

		dhoro product = notun Product();
		product.price = 100;
		product.price = -50;  // This should be ignored
		product.price
	`

	result := evalOOPInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if num.Value != 100 {
		t.Errorf("Expected price = 100 (negative ignored), got %f", num.Value)
	}
}

// Test combining getters, setters, private fields and static properties
func TestCombinedOOPFeatures(t *testing.T) {
	input := `
		sreni Counter {
			sthir instanceCount = 0;
			
			shuru() {
				ei._value = 0;
				Counter.instanceCount = Counter.instanceCount + 1;
			}

			pao value() {
				ferao ei._value;
			}

			set value(v) {
				ei._value = v;
			}

			kaj increment() {
				ei._value = ei._value + 1;
			}
		}

		dhoro c1 = notun Counter();
		dhoro c2 = notun Counter();
		c1.value = 10;
		c1.increment();
		
		// Return array with results
		[c1.value, Counter.instanceCount]
	`

	result := evalOOPInput(input)
	arr, ok := result.(*object.Array)
	if !ok {
		t.Fatalf("Expected Array, got %T", result)
	}

	if len(arr.Elements) != 2 {
		t.Fatalf("Expected 2 elements, got %d", len(arr.Elements))
	}

	val := arr.Elements[0].(*object.Number).Value
	if val != 11 {
		t.Errorf("Expected counter value = 11, got %f", val)
	}

	count := arr.Elements[1].(*object.Number).Value
	if count != 2 {
		t.Errorf("Expected instance count = 2, got %f", count)
	}
}

// Test real-world example: Product with computed properties
func TestProductWithComputedProperties(t *testing.T) {
	input := `
		sreni Product {
			sthir TAX_RATE = 0.15;
			
			shuru(naam, dam, quantity) {
				ei._naam = naam;
				ei._dam = dam;
				ei._quantity = quantity;
			}

			pao naam() {
				ferao ei._naam;
			}

			pao dam() {
				ferao ei._dam;
			}

			set dam(d) {
				jodi (d >= 0) {
					ei._dam = d;
				}
			}

			pao quantity() {
				ferao ei._quantity;
			}

			set quantity(q) {
				jodi (q >= 0) {
					ei._quantity = q;
				}
			}

			pao subtotal() {
				ferao ei._dam * ei._quantity;
			}

			pao tax() {
				ferao ei.subtotal * Product.TAX_RATE;
			}

			pao total() {
				ferao ei.subtotal + ei.tax;
			}
		}

		dhoro product = notun Product("Laptop", 1000, 2);
		product.total
	`

	result := evalOOPInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	// subtotal = 1000 * 2 = 2000
	// tax = 2000 * 0.15 = 300
	// total = 2000 + 300 = 2300
	expected := 2300.0
	if num.Value != expected {
		t.Errorf("Expected total = %f, got %f", expected, num.Value)
	}
}

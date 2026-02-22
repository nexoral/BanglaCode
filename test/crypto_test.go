package test

import (
	"testing"

	"BanglaCode/src/evaluator"
	"BanglaCode/src/evaluator/builtins"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
)

func evalCryptoInput(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	builtins.InitializeEnvironmentWithConstants(env)
	return evaluator.Eval(program, env)
}

func TestAESEncryptionDecryption(t *testing.T) {
	input := `
dhoro plaintext = "Hello, World!";
dhoro key = "my-secret-key";
dhoro encrypted = crypto_encrypt_aes(plaintext, key);
dhoro decrypted = crypto_decrypt_aes(encrypted, key);
decrypted
`

	result := evalCryptoInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	if str.Value != "Hello, World!" {
		t.Errorf("Expected 'Hello, World!', got %s", str.Value)
	}
}

func TestRSAKeyPairGeneration(t *testing.T) {
	input := `crypto_generate_keypair()`
	result := evalCryptoInput(input)

	mapObj, ok := result.(*object.Map)
	if !ok {
		t.Fatalf("Expected Map, got %T", result)
	}

	if _, ok := mapObj.Pairs["privateKey"]; !ok {
		t.Fatal("Missing privateKey in result")
	}
	if _, ok := mapObj.Pairs["publicKey"]; !ok {
		t.Fatal("Missing publicKey in result")
	}
}

func TestRSAEncryptionDecryption(t *testing.T) {
	input := `
dhoro plaintext = "Secret Message";
dhoro keys = crypto_generate_keypair();
dhoro encrypted = crypto_encrypt_rsa(plaintext, keys["publicKey"]);
dhoro decrypted = crypto_decrypt_rsa(encrypted, keys["privateKey"]);
decrypted
`

	result := evalCryptoInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	if str.Value != "Secret Message" {
		t.Errorf("Expected 'Secret Message', got %s", str.Value)
	}
}

func TestDigitalSignature(t *testing.T) {
	input := `
dhoro message = "Important document";
dhoro keys = crypto_generate_keypair();
dhoro signature = crypto_sign(message, keys["privateKey"]);
dhoro verified = crypto_verify(message, signature, keys["publicKey"]);
verified
`

	result := evalCryptoInput(input)
	boolean, ok := result.(*object.Boolean)
	if !ok {
		t.Fatalf("Expected Boolean, got %T", result)
	}

	if boolean.Value != true {
		t.Error("Expected signature verification to succeed")
	}
}

func TestSignatureVerificationFail(t *testing.T) {
	input := `
dhoro message = "Important document";
dhoro keys1 = crypto_generate_keypair();
dhoro keys2 = crypto_generate_keypair();
dhoro signature = crypto_sign(message, keys1["privateKey"]);
dhoro verified = crypto_verify(message, signature, keys2["publicKey"]);
verified
`

	result := evalCryptoInput(input)
	boolean, ok := result.(*object.Boolean)
	if !ok {
		t.Fatalf("Expected Boolean, got %T", result)
	}

	if boolean.Value != false {
		t.Error("Expected signature verification to fail with wrong key")
	}
}

func TestBcryptHashVerify(t *testing.T) {
	input := `
dhoro password = "mypassword";
dhoro hash = crypto_bcrypt(password);
dhoro verified = crypto_bcrypt_verify(password, hash);
verified
`

	result := evalCryptoInput(input)
	boolean, ok := result.(*object.Boolean)
	if !ok {
		t.Fatalf("Expected Boolean, got %T", result)
	}

	if boolean.Value != true {
		t.Error("Expected bcrypt verification to succeed")
	}
}

func TestBcryptWrongPassword(t *testing.T) {
	input := `
dhoro hash = crypto_bcrypt("correct");
dhoro verified = crypto_bcrypt_verify("wrong", hash);
verified
`

	result := evalCryptoInput(input)
	boolean, ok := result.(*object.Boolean)
	if !ok {
		t.Fatalf("Expected Boolean, got %T", result)
	}

	if boolean.Value != false {
		t.Error("Expected bcrypt verification to fail with wrong password")
	}
}

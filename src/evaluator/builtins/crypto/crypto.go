package crypto

import (
"BanglaCode/src/object"
"crypto/aes"
"crypto/cipher"
"crypto/rand"
"crypto/rsa"
"crypto/sha256"
"crypto/x509"
"crypto"
"encoding/base64"
"encoding/pem"
"fmt"
"io"

"golang.org/x/crypto/bcrypt"
)

// Builtins contains all Crypto-related built-in functions
var Builtins = map[string]*object.Builtin{
// AES Encryption (crypto_encrypt_aes)
"crypto_encrypt_aes": {
Fn: func(args ...object.Object) object.Object {
if len(args) != 2 {
return newError("crypto_encrypt_aes requires 2 arguments (plaintext, key)")
}

plaintext, ok := args[0].(*object.String)
if !ok {
return newError("plaintext must be STRING, got %s", args[0].Type())
}

key, ok := args[1].(*object.String)
if !ok {
return newError("key must be STRING, got %s", args[1].Type())
}

// Generate 32-byte key from password (SHA-256)
keyHash := sha256.Sum256([]byte(key.Value))

// Create AES cipher
block, err := aes.NewCipher(keyHash[:])
if err != nil {
return newError("failed to create cipher: %s", err.Error())
}

// Use GCM mode (Galois/Counter Mode) - authenticated encryption
gcm, err := cipher.NewGCM(block)
if err != nil {
return newError("failed to create GCM: %s", err.Error())
}

// Generate nonce (number used once)
nonce := make([]byte, gcm.NonceSize())
if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
return newError("failed to generate nonce: %s", err.Error())
}

// Encrypt and append authentication tag
ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext.Value), nil)

// Encode to base64 for safe string storage
encoded := base64.StdEncoding.EncodeToString(ciphertext)

return &object.String{Value: encoded}
},
},

// AES Decryption (crypto_decrypt_aes)
"crypto_decrypt_aes": {
Fn: func(args ...object.Object) object.Object {
if len(args) != 2 {
return newError("crypto_decrypt_aes requires 2 arguments (ciphertext, key)")
}

ciphertextStr, ok := args[0].(*object.String)
if !ok {
return newError("ciphertext must be STRING, got %s", args[0].Type())
}

key, ok := args[1].(*object.String)
if !ok {
return newError("key must be STRING, got %s", args[1].Type())
}

// Decode from base64
ciphertext, err := base64.StdEncoding.DecodeString(ciphertextStr.Value)
if err != nil {
return newError("invalid ciphertext encoding: %s", err.Error())
}

// Generate 32-byte key from password (SHA-256)
keyHash := sha256.Sum256([]byte(key.Value))

// Create AES cipher
block, err := aes.NewCipher(keyHash[:])
if err != nil {
return newError("failed to create cipher: %s", err.Error())
}

// Use GCM mode
gcm, err := cipher.NewGCM(block)
if err != nil {
return newError("failed to create GCM: %s", err.Error())
}

// Extract nonce
nonceSize := gcm.NonceSize()
if len(ciphertext) < nonceSize {
return newError("ciphertext too short")
}

nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

// Decrypt and verify authentication tag
plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
if err != nil {
return newError("decryption failed: %s", err.Error())
}

return &object.String{Value: string(plaintext)}
},
},

// Generate RSA Key Pair (crypto_generate_keypair)
"crypto_generate_keypair": {
Fn: func(args ...object.Object) object.Object {
bits := 2048 // Default RSA key size
if len(args) > 0 {
if num, ok := args[0].(*object.Number); ok {
bits = int(num.Value)
if bits < 1024 || bits > 8192 {
return newError("key size must be between 1024 and 8192 bits")
}
}
}

// Generate RSA key pair
privateKey, err := rsa.GenerateKey(rand.Reader, bits)
if err != nil {
return newError("failed to generate key pair: %s", err.Error())
}

// Encode private key to PEM
privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
privateKeyPEM := pem.EncodeToMemory(&pem.Block{
Type:  "RSA PRIVATE KEY",
Bytes: privateKeyBytes,
})

// Encode public key to PEM
publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
if err != nil {
return newError("failed to marshal public key: %s", err.Error())
}
publicKeyPEM := pem.EncodeToMemory(&pem.Block{
Type:  "RSA PUBLIC KEY",
Bytes: publicKeyBytes,
})

// Return as map with privateKey and publicKey
result := make(map[string]object.Object)
result["privateKey"] = &object.String{Value: string(privateKeyPEM)}
result["publicKey"] = &object.String{Value: string(publicKeyPEM)}

return &object.Map{Pairs: result}
},
},

// RSA Encrypt (crypto_encrypt_rsa)
"crypto_encrypt_rsa": {
Fn: func(args ...object.Object) object.Object {
if len(args) != 2 {
return newError("crypto_encrypt_rsa requires 2 arguments (plaintext, publicKey)")
}

plaintext, ok := args[0].(*object.String)
if !ok {
return newError("plaintext must be STRING, got %s", args[0].Type())
}

publicKeyPEM, ok := args[1].(*object.String)
if !ok {
return newError("publicKey must be STRING, got %s", args[1].Type())
}

// Decode PEM public key
block, _ := pem.Decode([]byte(publicKeyPEM.Value))
if block == nil {
return newError("failed to decode public key PEM")
}

// Parse public key
publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
if err != nil {
return newError("failed to parse public key: %s", err.Error())
}

rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
if !ok {
return newError("not an RSA public key")
}

// Encrypt with RSA-OAEP
ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, rsaPublicKey, []byte(plaintext.Value), nil)
if err != nil {
return newError("encryption failed: %s", err.Error())
}

// Encode to base64
encoded := base64.StdEncoding.EncodeToString(ciphertext)
return &object.String{Value: encoded}
},
},

// RSA Decrypt (crypto_decrypt_rsa)
"crypto_decrypt_rsa": {
Fn: func(args ...object.Object) object.Object {
if len(args) != 2 {
return newError("crypto_decrypt_rsa requires 2 arguments (ciphertext, privateKey)")
}

ciphertextStr, ok := args[0].(*object.String)
if !ok {
return newError("ciphertext must be STRING, got %s", args[0].Type())
}

privateKeyPEM, ok := args[1].(*object.String)
if !ok {
return newError("privateKey must be STRING, got %s", args[1].Type())
}

// Decode from base64
ciphertext, err := base64.StdEncoding.DecodeString(ciphertextStr.Value)
if err != nil {
return newError("invalid ciphertext encoding: %s", err.Error())
}

// Decode PEM private key
block, _ := pem.Decode([]byte(privateKeyPEM.Value))
if block == nil {
return newError("failed to decode private key PEM")
}

// Parse private key
privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
if err != nil {
return newError("failed to parse private key: %s", err.Error())
}

// Decrypt with RSA-OAEP
plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
if err != nil {
return newError("decryption failed: %s", err.Error())
}

return &object.String{Value: string(plaintext)}
},
},

// Sign Data (crypto_sign)
"crypto_sign": {
Fn: func(args ...object.Object) object.Object {
if len(args) != 2 {
return newError("crypto_sign requires 2 arguments (message, privateKey)")
}

message, ok := args[0].(*object.String)
if !ok {
return newError("message must be STRING, got %s", args[0].Type())
}

privateKeyPEM, ok := args[1].(*object.String)
if !ok {
return newError("privateKey must be STRING, got %s", args[1].Type())
}

// Decode PEM private key
block, _ := pem.Decode([]byte(privateKeyPEM.Value))
if block == nil {
return newError("failed to decode private key PEM")
}

// Parse private key
privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
if err != nil {
return newError("failed to parse private key: %s", err.Error())
}

// Hash the message
msgHash := sha256.Sum256([]byte(message.Value))

// Sign with RSA-PSS
signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHash[:], nil)
if err != nil {
return newError("signing failed: %s", err.Error())
}

// Encode to base64
encoded := base64.StdEncoding.EncodeToString(signature)
return &object.String{Value: encoded}
},
},

// Verify Signature (crypto_verify)
"crypto_verify": {
Fn: func(args ...object.Object) object.Object {
if len(args) != 3 {
return newError("crypto_verify requires 3 arguments (message, signature, publicKey)")
}

message, ok := args[0].(*object.String)
if !ok {
return newError("message must be STRING, got %s", args[0].Type())
}

signatureStr, ok := args[1].(*object.String)
if !ok {
return newError("signature must be STRING, got %s", args[1].Type())
}

publicKeyPEM, ok := args[2].(*object.String)
if !ok {
return newError("publicKey must be STRING, got %s", args[2].Type())
}

// Decode from base64
signature, err := base64.StdEncoding.DecodeString(signatureStr.Value)
if err != nil {
return newError("invalid signature encoding: %s", err.Error())
}

// Decode PEM public key
block, _ := pem.Decode([]byte(publicKeyPEM.Value))
if block == nil {
return newError("failed to decode public key PEM")
}

// Parse public key
publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
if err != nil {
return newError("failed to parse public key: %s", err.Error())
}

rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
if !ok {
return newError("not an RSA public key")
}

// Hash the message
msgHash := sha256.Sum256([]byte(message.Value))

// Verify with RSA-PSS
err = rsa.VerifyPSS(rsaPublicKey, crypto.SHA256, msgHash[:], signature, nil)
if err != nil {
return object.FALSE
}

return object.TRUE
},
},

// Bcrypt Hash Password (crypto_bcrypt)
"crypto_bcrypt": {
Fn: func(args ...object.Object) object.Object {
if len(args) < 1 || len(args) > 2 {
return newError("crypto_bcrypt requires 1-2 arguments (password, [cost])")
}

password, ok := args[0].(*object.String)
if !ok {
return newError("password must be STRING, got %s", args[0].Type())
}

cost := bcrypt.DefaultCost // 10
if len(args) == 2 {
if num, ok := args[1].(*object.Number); ok {
cost = int(num.Value)
if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
return newError("cost must be between %d and %d", bcrypt.MinCost, bcrypt.MaxCost)
}
}
}

// Hash password with bcrypt
hash, err := bcrypt.GenerateFromPassword([]byte(password.Value), cost)
if err != nil {
return newError("hashing failed: %s", err.Error())
}

return &object.String{Value: string(hash)}
},
},

// Bcrypt Verify Password (crypto_bcrypt_verify)
"crypto_bcrypt_verify": {
Fn: func(args ...object.Object) object.Object {
if len(args) != 2 {
return newError("crypto_bcrypt_verify requires 2 arguments (password, hash)")
}

password, ok := args[0].(*object.String)
if !ok {
return newError("password must be STRING, got %s", args[0].Type())
}

hash, ok := args[1].(*object.String)
if !ok {
return newError("hash must be STRING, got %s", args[1].Type())
}

// Compare password with hash
err := bcrypt.CompareHashAndPassword([]byte(hash.Value), []byte(password.Value))
if err != nil {
return object.FALSE
}

return object.TRUE
},
},
}

// Helper function to create error objects
func newError(format string, args ...interface{}) *object.Error {
return &object.Error{Message: fmt.Sprintf(format, args...)}
}

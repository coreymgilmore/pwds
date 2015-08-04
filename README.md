# pwds.go
Simple Golang Library for Creating and Verifying Bcrypt Passwords

This library just simplifies usage of the `golang.org/x/crypto/bcrypt` library.

---
###Functions

#### Create(input string)
Creates a hash of a given string input (password) using bcrypt.  Returns the hash as a string.

#### Verify(rawPassword string, hashedPassword string)
Compares a plain text password against a hash.  Returns True (boolean) if the given inputs match.

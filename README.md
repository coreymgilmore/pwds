# pwds
Simple Golang Library for Creating and Verifying Bcrypt Passwords

The intended use of this library is to make using golang.org/x/crypto/bcrypt easier.  This library just cleans up usage of the bcrypt library.

---

### Create(input string)
Creates a hash of a given string input (password) using bcrypt.  Returns the hash as a string.

### Verify(rawPassword string, hashedPassword string)
Compares a plain text password against a hash.  Returns True (boolean) if the given inputs match.

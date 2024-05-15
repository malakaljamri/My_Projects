import hashlib
import os

def hash_password(password, salt=None):
    if salt is None:
        salt = os.urandom(16)  # Generate a random 16-byte salt
    
    # Combine the password and salt, then hash the result
    hashed_password = hashlib.sha256(salt + password.encode()).hexdigest()
    
    return hashed_password, salt

# Example usage:
password = "securepassword"
hashed_password, salt = hash_password(password)

print(f"Password: {password}")
print(f"Hashed Password: {hashed_password}")
print(f"Salt: {salt}")

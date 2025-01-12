def is_prime_number(n):
    if n <= 1:
        return False
    if n == 2 or n == 3:
        return True
    if n % 2 == 0 or n % 3 == 0:
        return False
    for d in range(5, int(n ** 0.5) + 1, 2):
        if n % d == 0:
            return False
    return True


prime_numbers = []
for n in range(1, 1000):
    if is_prime_number(n):
        prime_numbers.append(n)

print(f"There are {len(prime_numbers)} prime numbers smaller than 1000, they are: {prime_numbers}")

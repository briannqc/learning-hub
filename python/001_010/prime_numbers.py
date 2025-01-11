def is_prime(n):
    for i in range(2, int(n ** 0.5) + 1):
        if n % i == 0:
            return False
    return True


def print_prime_numbers():
    for i in range(1, 101):
        if is_prime(i):
            print(i, "is a prime number")


print_prime_numbers()

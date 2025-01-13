def factorial(n):
    if n < 0:
        return None
    if n <= 2:
        return n
    return n * factorial(n - 1)


def factorial_iterative(n):
    if n < 0:
        return None
    if n == 0:
        return 1

    result = 1
    for i in range(1, n + 1):
        result = result * i

    return result


for n in range(0, 21):
    print(f"{n}! = {factorial(n)} (second method result: {factorial_iterative(n)})")

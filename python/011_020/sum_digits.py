def sum_digits(n):
    total = 0
    while n > 0:
        total = total + n % 10
        n = n // 10
    return total


tests = [123456, 0, 1, 10, 100, 22]
for t in tests:
    print("Sum of all digits of", t, "is", sum_digits(t))

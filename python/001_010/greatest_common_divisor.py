def gcd(a, b):
    if b == 0:
        return a
    return gcd(b, a % b)


tests = [
    [10, 0],
    [10, 1],
    [10, 5],
    [10, 8],
    [12, 8],
    [100, 250],
    [4, 5]
]

for pair in tests:
    print("Greatest common divisor of", pair, "is", gcd(pair[0], pair[1]))

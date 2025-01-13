def power(base, exponent):
    if exponent == 0:
        return 1
    result = 1
    for n in range(1, exponent + 1):
        result = result * base
    return result


tests = [
    (1, 1, 1),
    (2, 0, 1),
    (2, 1, 2),
    (2, 2, 4),
    (3, 3, 27)
]

for base, exp, expected in tests:
    actual = power(base, exp)
    if actual == expected:
        print(f"PASS: power({base}, {exp}) = {actual} as expected")
    else:
        print(f"FAILED: power({base}, {exp}) should be {expected} but was {actual}")

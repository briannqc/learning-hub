def max3(a, b, c):
    if a >= b and a >= c:
        return a
    if b >= a and b >= c:
        return b
    return c


tests = [
    [1, 2, 3],
    [4, 3, 2],
    [5, 7, 6]
]

for t in tests:
    print("The biggest number in", t, "is", max3(t[0], t[1], t[2]))

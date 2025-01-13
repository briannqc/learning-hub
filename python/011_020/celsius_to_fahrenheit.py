def celsius_to_fahrenheit(c):
    f = (c * 9 / 5) + 32
    return round(f, 2)


tests = [
    (-69, -92.2),
    (-5, 23.0),
    (68, 154.4),
    (-23, -9.4),
    (-77, -106.6),
    (-14, 6.8),
    (89, 192.2),
    (-88, -126.4),
    (-67, -88.6),
    (24, 75.2),
    (67, 152.6),
    (86, 186.8),
    (-50, -58.0),
    (-100, -148.0),
    (82, 179.6),
    (16, 60.8),
    (9, 48.2),
    (-11, 12.2),
    (-24, -11.2),
    (3, 37.4)
]

for c, f in tests:
    actual = celsius_to_fahrenheit(c)
    if actual == f:
        print(f"PASSED: {c} Celsius = {f} Fahrenheit")
    else:
        print(f"FAILED {c} Celsius should equal to {f} Fahrenheit, but it was {actual}")

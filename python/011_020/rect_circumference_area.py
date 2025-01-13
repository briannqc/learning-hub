def perimeter(length, width):
    return (length + width) * 2


def area(length, width):
    return length * width


tests = [
    (62, 22, 168, 1364),
    (25, 70, 190, 1750),
    (87, 9, 192, 783),
    (50, 7, 114, 350),
    (60, 67, 254, 4020),
    (17, 56, 146, 952),
    (7, 94, 202, 658),
    (73, 41, 228, 2993),
    (100, 62, 324, 6200),
    (67, 77, 288, 5159),
    (8, 48, 112, 384),
    (28, 79, 214, 2212),
    (1, 27, 56, 27),
    (72, 73, 290, 5256),
    (49, 23, 144, 1127),
    (65, 28, 186, 1820),
    (64, 96, 320, 6144),
    (66, 32, 196, 2112),
    (36, 12, 96, 432),
    (90, 1, 182, 90)
]

for l, w, expected_perimeter, expected_area in tests:
    actual_perimeter = perimeter(l, w)
    actual_area = area(l, w)
    if actual_area == expected_area and actual_perimeter == expected_perimeter:
        print(f"PASSED: perimeter({l}, {w}) = {actual_perimeter} and area({l}, {w}) = {actual_area} as expected")
        continue

    if actual_perimeter != expected_perimeter:
        print(f"FAILED: perimeter({l}, {w}) should be {expected_perimeter} but it was {actual_perimeter}")
    if actual_area != expected_area:
        print(f"FAILED: area({l}, {w}) should be {expected_area} but it was {actual_area}")

import math


def circumference(r):
    return round(2 * math.pi * r, 2)


def are(r):
    return round(math.pi * r * r, 2)


tests = [
    (88, 552.92, 24328.49),
    (60, 376.99, 11309.73),
    (29, 182.21, 2642.08),
    (19, 119.38, 1134.11),
    (1, 6.28, 3.14),
    (75, 471.24, 17671.46),
    (70, 439.82, 15393.80),
    (45, 282.74, 6361.73),
    (99, 622.04, 30790.75),
    (82, 515.22, 21124.07),
    (24, 150.80, 1809.56),
    (56, 351.86, 9852.03),
    (15, 94.25, 706.86),
    (31, 194.78, 3019.07),
    (42, 263.89, 5541.77),
    (48, 301.59, 7238.23),
    (98, 615.75, 30171.86),
    (3, 18.85, 28.27),
    (92, 578.05, 26590.44),
    (34, 213.63, 3631.68)
]

for radius, expected_circumference, expected_area in tests:
    actual_circumference = circumference(radius)
    actual_area = area(radius)
    if actual_area == expected_area and actual_circumference == expected_circumference:
        print(f"PASSED: circumference({radius} = {actual_circumference} and area({radius}) = {actual_area} as expected")
        continue

    if actual_circumference != expected_circumference:
        print(f"FAILED: circumference({radius} should be {expected_circumference} but it was {actual_circumference}")
    if actual_area != expected_area:
        print(f"FAILED: area({radius} should be {expected_area} but it was {actual_area}")

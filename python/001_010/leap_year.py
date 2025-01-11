def is_leap_year(y):
    return (y % 4 == 0 and y % 100 != 0) or (y % 400 == 0)


tests = [2000, 4000, 4040, 4400, 1500, 2025, 2028]
for y in tests:
    if is_leap_year(y):
        print(y, "is a leap year")
    else:
        print(y, "is NOT a leap year")

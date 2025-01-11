def check_negative_positive_number(n):
    if n < 0:
        return "Negative"
    if n == 0:
        return "Zero"
    return "Positive"


for n in [10, 0, -10, 2, 4]:
    print(n, "is", check_negative_positive_number(n))

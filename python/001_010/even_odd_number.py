def check_even_odd(n):
    if n % 2 == 0:
        return "Even"
    return "Odd"


for n in [123, 2, 0, 3, 5, -1, -10]:
    print(n, "is", check_even_odd(n))

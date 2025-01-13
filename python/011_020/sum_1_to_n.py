def sum_1_to_n(n):
    return n * (n + 1) // 2


for n in range(1, 101):
    print(f"sum[1..{n}] = {sum_1_to_n(n)}")

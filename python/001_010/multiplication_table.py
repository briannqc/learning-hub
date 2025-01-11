def print_multiplication_table():
    for i in range(1, 11):
        for j in range(1, 11):
            print(f"{i:2d} * {j:2d} = {i * j:3d}")
        print()


print_multiplication_table()

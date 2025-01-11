def count_even_odd(numbers):
    even_count = 0
    odd_count = 0
    for n in numbers:
        if n % 2 == 0:
            even_count = even_count + 1
        else:
            odd_count = odd_count + 1
    return even_count, odd_count


numbers = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 16, 33, 12, 10]
even_count, odd_count = count_even_odd(numbers)
print("There are", even_count, "even numbers and", odd_count, "odd numbers in the list")

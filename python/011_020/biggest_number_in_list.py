def find_biggest_number(nums):
    biggest = nums[0]
    for n in nums:
        if biggest < n:
            biggest = n
    return biggest


tests = [
    [1, 2, 3, 4, 5, 6, 6, 7],
    [7, 8, 9, 20, 11, 2],
    [234, 2, 3, 655, 2]
]
for t in tests:
    biggest = find_biggest_number(t)
    print("Biggest number in", t, "is", biggest)

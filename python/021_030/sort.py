nums = [1, 2, 4, 1, 2, 5, 2, 0]
nums.sort()
assert nums == [0, 1, 1, 2, 2, 2, 4, 5]

nums = [1, 2, 4, 1, 2, 5, 2, 0]
nums.sort(reverse=True)
assert nums == [5, 4, 2, 2, 2, 1, 1, 0]

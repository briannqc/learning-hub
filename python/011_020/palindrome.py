def is_palindrome(s):
    if len(s) <= 1:
        return True
    for i in range(int(len(s) / 2)):
        if s[i] != s[-(i + 1)]:
            return False
    return True


tests = [
    "abba",
    "aba",
    "abc",
    "a",
    "mama",
    "daddy",
    "dadda",
    "madam",
    "racecar"
]
for s in tests:
    print(f"Is '{s}' a palindrome?", is_palindrome(s))

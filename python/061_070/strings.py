import unittest


class StringTestCase(unittest.TestCase):
    def test_define_strings(self):
        str1 = "The simplest way to define a string is using double quotes, hence we can include single quote ' character inside it conveniently"
        self.assertIn("'", str1)

        str2 = 'We can also using single quotes to define string, which allow us include double quote " character inside it conveniently'
        self.assertIn('"', str2)

        str3 = """We can use triple double-quotes
        to define
        multiple lines
        string
        """
        self.assertIn("\n", str3)

        str4 = '''Same for triple single-quotes
        We can also use them
        for multiple lines string
        '''
        self.assertIn("\n", str4)

        str5 = str(3)
        self.assertEqual("3", str5, "str() function converts objects to strings")

        str6 = "We can use '+' operator " + "to concat strings"
        self.assertIn("+", str6)

        str7 = f"{'f-string'} is a convenient way to format string, it is available from {'v3.6'}"
        self.assertIn("v3.6", str7)

        str8 = "We can use format() function to {} too!".format("format string")
        self.assertIn("format string", str8)

    def test_concatenate_strings(self):
        s1 = "Hello, "
        s2 = "World!"
        expected = "Hello, World!"
        self.assertEqual(expected, s1 + s2, "'+' concatenates 2 strings")
        self.assertEqual(expected, "".join([s1, s2]), "join() function joins string iterables")
        self.assertEqual(expected, f"{s1}{s2}", "f-string formats string can be used to concatenate strings too")
        self.assertEqual(expected, "{}{}".format(s1, s2), "same for format() function")

    def test_len_of_string(self):
        s = "Hello, World!"
        self.assertEqual(13, len(s), "len() function returns len of a string")

    def test_upper_string(self):
        s = "Hello"
        self.assertEqual("HELLO", s.upper(), "s.upper() function creates a new string which is upper of s")

    def test_lower_string(self):
        s = "Hello"
        self.assertEqual("hello", s.lower(), "s.lower() function creates a new string which is lower of s")

    def test_count_char_in_string(self):
        s = "Hello"
        self.assertEqual(1, s.count('H'), "count() function counts a char in string")
        self.assertEqual(2, s.count('l'), "count() function counts a char in string")

    def test_substring(self):
        s = "Hello, World!"
        self.assertTrue("Hello" in s, "Use 'a in s' to check if a is a substring of s")
        self.assertFalse("world" in s, "Use 'a in s' to check if a is a substring of s")

    def test_replace_word_in_string(self):
        s = "HellO, WOrld!"
        s = s.replace("O", "o")
        self.assertEqual("Hello, World!", s,
                         "replace() returns a copy with all occurrences of substring old replaced by new.")

    def test_split_string(self):
        s = "Hello world, welcome to the world of Python."
        self.assertListEqual(["Hello", "world,", "welcome", "to", "the", "world", "of", "Python."], s.split(),
                             "split() returns a list of the substrings in the string, using sep as the separator string.")
        self.assertEqual(["Hello world", "welcome to the world of Python."], s.split(", "),
                         "We can pass to split() the separator too.")

    def test_join_strings(self):
        l = ["Hello", "world"]
        self.assertEqual("Hello, world", ", ".join(l), "join() joins list of strings into one string")


if __name__ == '__main__':
    unittest.main()

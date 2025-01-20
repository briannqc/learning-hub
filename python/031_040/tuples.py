import unittest


class TupleTestCase(unittest.TestCase):
    def test_declare_tuple(self):
        l = [1, 2, 3]
        t = (1, 2, 3)
        self.assertSequenceEqual(l, t)

    def test_access_element_of_tuple(self):
        t = (12, 23, 34, 45, 56)
        self.assertEqual(45, t[3], "Access forward")
        self.assertEqual(45, t[-2], "Access backward")

    def test_concat_tuples(self):
        t1 = (1, 2, 3)
        t2 = ("a", "b", "c")
        self.assertTupleEqual(t1 + t2, (1, 2, 3, "a", "b", "c"))
        self.assertTupleEqual(t2 + t1, ("a", "b", "c", 1, 2, 3))

    def test_len_of_tuple(self):
        t = (12, 23, 34, 45, 56)
        self.assertEqual(5, len(t))

    def test_tuple_to_list(self):
        t = (1, 2, 3)
        l = list(t)
        self.assertListEqual([1, 2, 3], l)

    def test_list_to_tuple(self):
        l = [1, 2, 3]
        t = tuple(l)
        self.assertTupleEqual((1, 2, 3), t)

    def test_element_in_tuple(self):
        t = (1, 2, 3)
        self.assertTrue(1 in t)
        self.assertFalse(4 in t)

    def test_count_in_tuple(self):
        t = (1, 2, 3, 4, 1, 1, 2, 3)
        self.assertEqual(3, t.count(1))
        self.assertEqual(1, t.count(4))
        self.assertEqual(0, t.count(5))

    def test_max_in_tuple(self):
        self.assertEqual("c", max(("a", "b", "c")))
        self.assertEqual(3, max((1, 3, 2)))

    def test_min_in_tuple(self):
        self.assertEqual("a", min(("a", "b", "c")))
        self.assertEqual(1, min((1, 3, 2)))


if __name__ == '__main__':
    unittest.main()

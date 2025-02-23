import unittest


class SetTestCase(unittest.TestCase):
    def test_define_set(self):
        s1 = {"head", "shoulder", "knee", "toes"}
        self.assertTrue(isinstance(s1, set))
        self.assertEqual(4, len(s1))

        s2 = set({})
        self.assertTrue(isinstance(s2, set))
        self.assertEqual(0, len(s2))

        s3 = {1, 2, 3, 1, 1, 3}
        self.assertTrue(isinstance(s3, set))
        self.assertEqual(3, len(s3))

    def test_add_elements_to_set(self):
        s = {1, 2, 3}
        self.assertEqual(3, len(s))

        # add() adds a single element to set
        s.add(4)
        self.assertEqual(4, len(s))
        self.assertIn(4, s)

        # update() adds elements of an iterable to set
        s.update([1, 4, 5])
        self.assertEqual(5, len(s))
        self.assertIn(5, s)

    def test_remove_elements_from_set(self):
        s = {1, 2, 3}
        self.assertEqual(3, len(s))

        s.remove(1)
        self.assertEqual(2, len(s))
        self.assertNotIn(1, s)

        self.assertRaises(KeyError, lambda _: s.remove(1), "remove() raises KeyError when element not exist")

        s.discard(2)
        self.assertEqual(1, len(s))
        self.assertNotIn(2, s)

        got_key_err = False
        try:
            s.discard(2)
        except KeyError:
            got_key_err = True
        self.assertFalse(got_key_err, "discard() does NOT raise Error when element not exist")

    def test_element_in_set(self):
        self.assertTrue(1 in {1, 2, 3}, "Use 'in' operator to check an element exists in set")

    def test_len_of_set(self):
        self.assertEqual(4, len({1, 2, 3, 4, 4, 4}), "Use len() function to check len of a set")

    def test_union_of_sets(self):
        s1 = {1, 2, 3}
        s2 = {2, 3, 4}
        s3 = s1.union(s2)
        s4 = s1 | s2

        self.assertSetEqual({1, 2, 3, 4}, s3, "Use union() function or | operator to find union of sets")
        self.assertSetEqual({1, 2, 3, 4}, s4, "Use union() function or | operator to find union of sets")

    def test_intersection_of_sets(self):
        s1 = {1, 2, 3}
        s2 = {2, 3, 4}
        s3 = s1.intersection(s2)
        s4 = s1 & s2

        self.assertSetEqual({2, 3}, s3, "Use intersection() function or & operator to find intersection of sets")
        self.assertSetEqual({2, 3}, s4, "Use intersection() function or & operator to find intersection of sets")

    def test_difference_of_sets(self):
        s1 = {1, 2, 3}
        s2 = {2, 3, 4}
        s3 = s1.difference(s2)
        s4 = s1 - s2

        self.assertSetEqual({1}, s3, "Use difference() function or - operator to find intersection of sets")
        self.assertSetEqual({1}, s4, "Use difference() function or - operator to find intersection of sets")

    def test_clear_set(self):
        s = {1, 2, 3}
        s.clear()
        self.assertEqual(0, len(s), "clear() removes all elements of set")

    def test_list_to_set(self):
        l = [1, 2, 3, 1, 1, 2, 3, 3, 4, 2, 2, 4]
        s = set(l)
        self.assertSetEqual({1, 2, 3, 4}, s, "set([...]) converts a list to a set")


if __name__ == '__main__':
    unittest.main()

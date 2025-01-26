import unittest


class DictionaryTestCase(unittest.TestCase):
    def test_041_declare_dictionary(self):
        my_dict = {}
        self.assertEqual(len(my_dict), 0, "The dict should be empty")

        my_dict['name'] = 'NQC'
        my_dict['city'] = 'HCMC'
        self.assertEqual(len(my_dict), 2, "The dict should contain 2 elements")

    def test_042_access_dict(self):
        my_dict = {
            'name': 'NQC'
        }
        self.assertEqual(my_dict['name'], "NQC")

    def test_043_add_pair_to_dict(self):
        my_dict = {
            'name': 'NQC'
        }
        self.assertEqual(len(my_dict), 1)

        my_dict['city'] = 'HCMC'
        self.assertEqual(len(my_dict), 2)
        self.assertEqual(my_dict['city'], 'HCMC')

    def test_044_delete_pair_from_dict(self):
        my_dict = {
            'name': 'NQC',
            'city': 'HCMC'
        }
        self.assertEqual(my_dict['city'], 'HCMC')

        del my_dict['city']
        self.assertEqual(my_dict.get('city'), None)

        name = my_dict.pop('name')
        self.assertEqual(name, 'NQC')
        self.assertEqual(my_dict.get('name'), None)

    def test_045_check_key_exists_in_dict(self):
        my_dict = {
            'name': 'NQC',
            'city': 'HCMC'
        }
        self.assertTrue('name' in my_dict)
        self.assertTrue('city' in my_dict)
        self.assertFalse('age' in my_dict)

    def test_046_keys_of_dict(self):
        my_dict = {
            'name': 'NQC',
            'city': 'HCMC'
        }
        keys = my_dict.keys()
        self.assertEqual(sorted(keys), ['city', 'name'])

    def test_047_values_of_dict(self):
        my_dict = {
            'name': 'NQC',
            'city': 'HCMC'
        }
        values = my_dict.values()
        self.assertEqual(sorted(values), ['HCMC', 'NQC'])

    def test_048_items_of_dict(self):
        my_dict = {
            'name': 'NQC',
            'city': 'HCMC'
        }
        items = my_dict.items()  # dict_items([('name', 'NQC'), ('city', 'HCMC')]
        self.assertIn(('name', 'NQC'), items)
        self.assertIn(('city', 'HCMC'), items)

    def test_049_count_items_in_dict(self):
        my_dict = {
            'name': 'NQC',
            'city': 'HCMC'
        }
        self.assertEqual(2, len(my_dict))
        self.assertEqual(2, len(my_dict.items()))

    def test_050_join_dicts(self):
        dict1 = {
            'name': 'Alice',
            'age': 25
        }

        dict2 = {
            'city': 'New York',
            'country': 'USA'
        }

        dict3 = dict1.copy()
        dict3.update(dict2)

        joined_dict = {
            'name': 'Alice',
            'age': 25,
            'city': 'New York',
            'country': 'USA'
        }
        self.assertDictEqual(joined_dict, dict3)

        dict4 = {**dict1, **dict2}
        self.assertDictEqual(joined_dict, dict4)


if __name__ == '__main__':
    unittest.main()

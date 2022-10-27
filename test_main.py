import unittest

import unittest

from main import *

class TestMain(unittest.TestCase):
  def test_maximum(self):
    val = maximum([1, 2, 3, 4, 5])
    self.assertEqual(val, 5)

  def test_minimum(self):
    val = minimum([-1, 6, 7, 0, -2])
    self.assertEqual(val, -2)
  
  def test_average(self):
    val = average([1, 6, 4, 3, 7, 3, 2])
    # 26 / 7
    self.assertEqual(val, 26 / 7)

  def test_failure(self):
    self.assertEqual(1, 2)

if __name__ == "__main__":
  unittest.main()
  
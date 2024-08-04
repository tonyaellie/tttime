import unittest
from datetime import datetime
from tttime import TTTime

class TestTTTime(unittest.TestCase):
  
  def test_constructor_with_datetime(self):
    time = TTTime(datetime(2020, 10, 21, 18, 0, 0))
    self.assertEqual(time.get_year(), 0)
    self.assertEqual(time.get_month(), 0)
    self.assertEqual(time.get_day(), 0)
    self.assertEqual(time.get_hour(), 0)
    self.assertEqual(time.get_minute(), 0)
    self.assertEqual(time.get_second(), 0)
    self.assertEqual(time.get_millisecond(), 0)

  def test_constructor_with_dict(self):
    time = TTTime({
      'years': 1202,
      'months': 3,
      'days': 5,
      'hours': 12,
      'minutes': 23,
      'seconds': 24,
    })
    self.assertEqual(time.get_year(), 1202)
    self.assertEqual(time.get_month(), 3)
    self.assertEqual(time.get_day(), 5)
    self.assertEqual(time.get_hour(), 12)
    self.assertEqual(time.get_minute(), 23)
    self.assertEqual(time.get_second(), 24)

  def test_constructor_with_iso_string(self):
    time = TTTime('2020-10-21T18:00:00.000Z')
    self.assertEqual(time.get_year(), 0)
    self.assertEqual(time.get_month(), 0)
    self.assertEqual(time.get_day(), 0)
    self.assertEqual(time.get_hour(), 0)
    self.assertEqual(time.get_minute(), 0)
    self.assertEqual(time.get_second(), 0)
    self.assertEqual(time.get_millisecond(), 0)

  def test_to_standard(self):
    time = TTTime({})
    self.assertEqual(time.to_standard().isoformat(), '2020-10-21T18:00:00+00:00')

  def test_get_year(self):
    time = TTTime({})
    self.assertEqual(time.get_year(), 0)

  def test_get_month(self):
    time = TTTime({})
    self.assertEqual(time.get_month(), 0)

  def test_get_day(self):
    time = TTTime({})
    self.assertEqual(time.get_day(), 0)

  def test_get_hour(self):
    time = TTTime({})
    self.assertEqual(time.get_hour(), 0)

  def test_get_minute(self):
    time = TTTime({})
    self.assertEqual(time.get_minute(), 0)

  def test_get_second(self):
    time = TTTime({})
    self.assertEqual(time.get_second(), 0)

  def test_get_millisecond(self):
    time = TTTime({})
    self.assertEqual(time.get_millisecond(), 0)

  def test_set_year(self):
    time = TTTime()
    time.set_year(2020)
    self.assertEqual(time.get_year(), 2020)

  def test_set_month(self):
    time = TTTime()
    time.set_month(9)
    self.assertEqual(time.get_month(), 9)

  def test_set_day(self):
    time = TTTime()
    time.set_day(2)
    self.assertEqual(time.get_day(), 2)

  def test_set_hour(self):
    time = TTTime()
    time.set_hour(18)
    self.assertEqual(time.get_hour(), 18)

  def test_set_minute(self):
    time = TTTime()
    time.set_minute(0)
    self.assertEqual(time.get_minute(), 0)

  def test_set_second(self):
    time = TTTime()
    time.set_second(0)
    self.assertEqual(time.get_second(), 0)

  def test_set_millisecond(self):
    time = TTTime()
    time.set_millisecond(0)
    self.assertEqual(time.get_millisecond(), 0)

  def test_to_string(self):
    time = TTTime({})
    self.assertEqual(time.to_string('%Y-%m-%d %H:%M:%S'), '0000-0-0 00:00:00')

  def test_to_standard_milliseconds(self):
    time = TTTime({})
    self.assertEqual(time.to_standard_milliseconds(), int(datetime(2020, 10, 21, 18, 0, 0).timestamp() * 1000))

  def test_add(self):
    time = TTTime({})
    time.add(TTTime({'hours': 1, 'milliseconds': 23}))
    self.assertEqual(time.get_year(), 0)
    self.assertEqual(time.get_month(), 0)
    self.assertEqual(time.get_day(), 0)
    self.assertEqual(time.get_hour(), 1)
    self.assertEqual(time.get_minute(), 0)
    self.assertEqual(time.get_second(), 0)
    self.assertEqual(time.get_millisecond(), 23)

  def test_subtract(self):
    time = TTTime({'years': 2, 'hours': 20, 'milliseconds': 50})
    time.subtract({'months': 1, 'hours': 1, 'milliseconds': 23})
    self.assertEqual(time.get_year(), 1)
    self.assertEqual(time.get_month(), 9)
    self.assertEqual(time.get_day(), 0)
    self.assertEqual(time.get_hour(), 19)
    self.assertEqual(time.get_minute(), 0)
    self.assertEqual(time.get_second(), 0)
    self.assertEqual(time.get_millisecond(), 27)

  def test_is_before(self):
    time1 = TTTime({'years': 2, 'hours': 20, 'milliseconds': 50})
    time2 = TTTime({'years': 2, 'days': 2, 'hours': 20, 'milliseconds': 50})
    self.assertTrue(time1.is_before(time2))
    self.assertFalse(time2.is_before(time1))

  def test_is_after(self):
    time1 = TTTime({'years': 2, 'hours': 20, 'milliseconds': 50})
    time2 = TTTime({'years': 2, 'days': 2, 'hours': 20, 'milliseconds': 50})
    self.assertFalse(time1.is_after(time2))
    self.assertTrue(time2.is_after(time1))

  def test_is_equal(self):
    time1 = TTTime({'years': 2, 'hours': 20, 'milliseconds': 50})
    time2 = TTTime({'years': 2, 'days': 2, 'hours': 20, 'milliseconds': 50})
    self.assertFalse(time1.is_equal(time2))
    self.assertTrue(time1.is_equal(time1))

if __name__ == '__main__':
  unittest.main()

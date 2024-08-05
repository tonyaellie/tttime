# TTTime Python Library

TTTime is a base 10 time representation. It is used to represent time in a way that is easy to understand and work with. Time is measured from 18:00:00 - 21/10/2020 UTC Standard Time (Day Zero).

To find out more or to contribute, please visit the [github page](https://github.com/tonyaellie/tttime).

```
Day Zero: 18:00:00 - 21/10/2020 UTC
1 TTT Hour = 1 Hour Standard Time
100 TTT Seconds in a TTT Minute
100 TTT Minutes in a TTT Hour
100 TTT Hours in a TTT Day
10 TTT Days in a TTT Month
10 TTT Months in a TTT Year
```

## Benefits

- Easy to understand and work with
- No time zones or daylight savings time
- No leap anything

## Installing and Using

```bash
pip install tttime
```

```python
from tttime import TTTime
time = TTTime()
time.set_year(2020)
time.set_month(9)
time.set_day(2)
time.set_hour(18)
time.set_minute(0)
time.set_second(0)
time.set_millisecond(0)
print(time.to_string('%Y-%m-%d %H:%M:%S'))
# 2020-09-02 18:00:00
```

## API

```python
class TTTime:
  def __init__(self, input: Union[None, Dict[str, int], str, datetime] = None) -> None:
  def to_standard(self) -> datetime:

  def get_year(self) -> int:
  def get_month(self) -> int:
  def get_day(self) -> int:
  def get_hour(self) -> int:
  def get_minute(self) -> int:
  def get_second(self) -> int:
  def get_millisecond(self) -> int:

  def set_year(self, year: int) -> None:
  def set_month(self, month: int) -> None:
  def set_day(self, day: int) -> None:
  def set_hour(self, hour: int) -> None:
  def set_minute(self, minute: int) -> None:
  def set_second(self, second: int) -> None:
  def set_millisecond(self, millisecond: int) -> None:

  def add(self, time: Union['TTTime', Dict[str, int], str, datetime]) -> None:
  def subtract(self, time: Union['TTTime', Dict[str, int], str, datetime]) -> None:

  def is_before(self, time: Union['TTTime', Dict[str, int], str, datetime]) -> bool:
  def is_after(self, time: Union['TTTime', Dict[str, int], str, datetime]) -> bool:
  def is_equal(self, time: Union['TTTime', Dict[str, int], str, datetime]) -> bool:

  def to_string(self, format: str = '%Y-%m-%d %H:%M:%S') -> str:
  def to_standard_milliseconds(self) -> int:
```

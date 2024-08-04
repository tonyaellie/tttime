from datetime import datetime, timedelta, timezone
from typing import Dict, Union

MILLISECOND_MULTIPLIER = 1
SECOND_MULTIPLIER = MILLISECOND_MULTIPLIER * 1000
MINUTE_MULTIPLIER = SECOND_MULTIPLIER * 100
HOUR_MULTIPLIER = MINUTE_MULTIPLIER * 100
DAY_MULTIPLIER = HOUR_MULTIPLIER * 100
MONTH_MULTIPLIER = DAY_MULTIPLIER * 10
YEAR_MULTIPLIER = MONTH_MULTIPLIER * 10

ONE_DIGIT_MASK = 10
TWO_DIGIT_MASK = 100
THREE_DIGIT_MASK = 1000

TTTimePartial = Dict[str, int]

class TTTime:
  # Epoch is 18:00:00 - 21/10/2020 UTC
  epoch: datetime = datetime(2020, 10, 21, 18, 0, 0, tzinfo=timezone.utc)

  def __init__(self, input: Union[None, Dict[str, int], str, datetime] = None) -> None:
    self.__time: int = 0
    if isinstance(input, dict):
      if 'years' in input:
        self.set_year(input['years'])
      if 'months' in input:
        self.set_month(input['months'])
      if 'days' in input:
        self.set_day(input['days'])
      if 'hours' in input:
        self.set_hour(input['hours'])
      if 'minutes' in input:
        self.set_minute(input['minutes'])
      if 'seconds' in input:
        self.set_second(input['seconds'])
      if 'milliseconds' in input:
        self.set_millisecond(input['milliseconds'])
    elif isinstance(input, str):
      self.__time = self.__from_standard(self.__parse_isoformat(input))
    elif isinstance(input, datetime):
      self.__time = self.__from_standard(input)
    else:
      self.__time = self.__from_standard(datetime.now(timezone.utc))

  def __parse_isoformat(self, date_string: str) -> datetime:
    return datetime.fromisoformat(date_string.replace("Z", "+00:00"))

  def __from_standard(self, standard__time: datetime) -> int:
    if standard__time.tzinfo is None:
      standard__time = standard__time.replace(tzinfo=timezone.utc)
    hours_since_epoch = (standard__time - TTTime.epoch) / timedelta(hours=1)
    return int(hours_since_epoch * HOUR_MULTIPLIER)

  def to_standard(self) -> datetime:
    return TTTime.epoch + timedelta(hours=(self.__time / HOUR_MULTIPLIER))

  def get_year(self) -> int:
    return self.__time // YEAR_MULTIPLIER

  def get_month(self) -> int:
    return (self.__time // MONTH_MULTIPLIER) % ONE_DIGIT_MASK

  def get_day(self) -> int:
    return (self.__time // DAY_MULTIPLIER) % ONE_DIGIT_MASK

  def get_hour(self) -> int:
    return (self.__time // HOUR_MULTIPLIER) % TWO_DIGIT_MASK

  def get_minute(self) -> int:
    return (self.__time // MINUTE_MULTIPLIER) % TWO_DIGIT_MASK

  def get_second(self) -> int:
    return (self.__time // SECOND_MULTIPLIER) % TWO_DIGIT_MASK

  def get_millisecond(self) -> int:
    return (self.__time // MILLISECOND_MULTIPLIER) % THREE_DIGIT_MASK

  def set_year(self, year: int) -> None:
    self.__time -= self.get_year() * YEAR_MULTIPLIER
    self.__time += year * YEAR_MULTIPLIER

  def set_month(self, month: int) -> None:
    if month > 9:
      raise ValueError('TTTime: Month can only be between 0-9')
    self.__time -= self.get_month() * MONTH_MULTIPLIER
    self.__time += month * MONTH_MULTIPLIER

  def set_day(self, day: int) -> None:
    if day > 9:
      raise ValueError('TTTime: Day can only be between 0-9')
    self.__time -= self.get_day() * DAY_MULTIPLIER
    self.__time += day * DAY_MULTIPLIER

  def set_hour(self, hour: int) -> None:
    if hour > 99:
      raise ValueError('TTTime: Hour can only be between 0-99')
    self.__time -= self.get_hour() * HOUR_MULTIPLIER
    self.__time += hour * HOUR_MULTIPLIER

  def set_minute(self, minute: int) -> None:
    if minute > 99:
      raise ValueError('TTTime: Minute can only be between 0-99')
    self.__time -= self.get_minute() * MINUTE_MULTIPLIER
    self.__time += minute * MINUTE_MULTIPLIER

  def set_second(self, second: int) -> None:
    if second > 99:
      raise ValueError('TTTime: Second can only be between 0-99')
    self.__time -= self.get_second() * SECOND_MULTIPLIER
    self.__time += second * SECOND_MULTIPLIER

  def set_millisecond(self, millisecond: int) -> None:
    if millisecond > 999:
      raise ValueError('TTTime: Millisecond can only be between 0-999')
    self.__time -= self.get_millisecond() * MILLISECOND_MULTIPLIER
    self.__time += millisecond * MILLISECOND_MULTIPLIER

  def add(self, time: Union['TTTime', Dict[str, int], str, datetime]) -> None:
    if isinstance(time, TTTime):
      self.__time += time.__time
    else:
      self.__time += TTTime(time).__time

  def subtract(self, time: Union['TTTime', Dict[str, int], str, datetime]) -> None:
    if isinstance(time, TTTime):
      self.__time -= time.__time
    else:
      self.__time -= TTTime(time).__time

  def is_before(self, time: Union['TTTime', Dict[str, int], str, datetime]) -> bool:
    if isinstance(time, TTTime):
      return self.__time < time.__time
    else:
      return self.__time < TTTime(time).__time

  def is_after(self, time: Union['TTTime', Dict[str, int], str, datetime]) -> bool:
    if isinstance(time, TTTime):
      return self.__time > time.__time
    else:
      return self.__time > TTTime(time).__time

  def is_equal(self, time: Union['TTTime', Dict[str, int], str, datetime]) -> bool:
    if isinstance(time, TTTime):
      return self.__time == time.__time
    else:
      return self.__time == TTTime(time).__time

  def to_string(self, format: str = '%Y-%m-%d %H:%M:%S') -> str:
    format_map = {
      '%Y': str(self.get_year()).zfill(4),
      '%m': str(self.get_month()),
      '%d': str(self.get_day()),
      '%H': str(self.get_hour()).zfill(2),
      '%M': str(self.get_minute()).zfill(2),
      '%S': str(self.get_second()).zfill(2),
      '%f': str(self.get_millisecond()).zfill(3)
    }
    for key, value in format_map.items():
      format = format.replace(key, value)
    return format

  def to_standard_milliseconds(self) -> int:
    return int(self.to_standard().timestamp() * 1000)

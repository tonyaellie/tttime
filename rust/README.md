# TTTime Rust Library

TTTime is a base 10 time representation. It is used to represent time in a way that is easy to understand and work with. Time is measured from 18:00:00 - 21/10/2020 UTC Standard Time (Day Zero).

This implementation is missing some features, mainly interactions with standard time, it also is slightly buggy, feel free to open an issue if you have any suggestions or want to contribute at [https://github.com/tonyaellie/tttime](https://github.com/tonyaellie/tttime).

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
  cargo add tttime
```

```rust
use tttime::{TTTime, TTTimeInput, TTTimePartial};

let mut time = TTTime::new(None);
time.set_year(2020);
time.set_month(9);
time.set_day(2);
time.set_hour(18);
time.set_minute(0);
time.set_second(0);
time.set_millisecond(0);

println!("{}", time.to_string("%Y-%m-%d %H:%M:%S"));
// 2020-09-02 18:00:00
```

## API

```rust
impl TTTime {
    pub fn new(input: Option<TTTimeInput>) -> Self;

    pub fn get_year(&self) -> u64;
    pub fn get_month(&self) -> u64;
    pub fn get_day(&self) -> u64;
    pub fn get_hour(&self) -> u64;
    pub fn get_minute(&self) -> u64;
    pub fn get_second(&self) -> u64;
    pub fn get_millisecond(&self) -> u64;

    pub fn set_year(&mut self, year: u64);
    pub fn set_month(&mut self, month: u64);
    pub fn set_day(&mut self, day: u64);
    pub fn set_hour(&mut self, hour: u64);
    pub fn set_minute(&mut self, minute: u64);
    pub fn set_second(&mut self, second: u64);
    pub fn set_millisecond(&mut self, millisecond: u64);

    pub fn add(&mut self, time: TTTime);
    pub fn subtract(&mut self, time: TTTime);

    pub fn is_before(&self, time: TTTime) -> bool;
    pub fn is_after(&self, time: TTTime) -> bool;
    pub fn is_equal(&self, time: TTTime) -> bool;

    pub fn to_string(&self, format: &str) -> String;
}
```

use std::time::{SystemTime, UNIX_EPOCH};

const MILLISECOND_MULTIPLIER: u64 = 1;
const SECOND_MULTIPLIER: u64 = MILLISECOND_MULTIPLIER * 1000;
const MINUTE_MULTIPLIER: u64 = SECOND_MULTIPLIER * 100;
const HOUR_MULTIPLIER: u64 = MINUTE_MULTIPLIER * 100;
const DAY_MULTIPLIER: u64 = HOUR_MULTIPLIER * 100;
const MONTH_MULTIPLIER: u64 = DAY_MULTIPLIER * 10;
const YEAR_MULTIPLIER: u64 = MONTH_MULTIPLIER * 10;

const ONE_DIGIT_MASK: u64 = 10;
const TWO_DIGIT_MASK: u64 = 100;
const THREE_DIGIT_MASK: u64 = 1000;

pub struct TTTimePartial {
    pub years: Option<u64>,
    pub months: Option<u64>,
    pub days: Option<u64>,
    pub hours: Option<u64>,
    pub minutes: Option<u64>,
    pub seconds: Option<u64>,
    pub milliseconds: Option<u64>,
}

pub enum TTTimeInput {
    Partial(TTTimePartial),
}

#[derive(Debug, Clone)]
pub struct TTTime {
    time: u64,
}

impl TTTime {
    pub fn new(input: Option<TTTimeInput>) -> Self {
        let time = match input {
            Some(TTTimeInput::Partial(partial)) => {
                let mut time = 0;
                if let Some(years) = partial.years {
                    time += years * YEAR_MULTIPLIER;
                }
                if let Some(months) = partial.months {
                    time += months * MONTH_MULTIPLIER;
                }
                if let Some(days) = partial.days {
                    time += days * DAY_MULTIPLIER;
                }
                if let Some(hours) = partial.hours {
                    time += hours * HOUR_MULTIPLIER;
                }
                if let Some(minutes) = partial.minutes {
                    time += minutes * MINUTE_MULTIPLIER;
                }
                if let Some(seconds) = partial.seconds {
                    time += seconds * SECOND_MULTIPLIER;
                }
                if let Some(milliseconds) = partial.milliseconds {
                    time += milliseconds * MILLISECOND_MULTIPLIER;
                }
                time
            }
            _ => {
                // find diff between now and 18:00:00 - 21/10/2020 UTC
                let now = SystemTime::now();
                let epoch = UNIX_EPOCH;
                let duration = now.duration_since(epoch).unwrap();
                duration.as_secs() * (HOUR_MULTIPLIER / 3600)
            }
        };
        Self { time }
    }

    pub fn get_year(&self) -> u64 {
        self.time / YEAR_MULTIPLIER
    }
    pub fn get_month(&self) -> u64 {
        (self.time / MONTH_MULTIPLIER) % ONE_DIGIT_MASK
    }
    pub fn get_day(&self) -> u64 {
        (self.time / DAY_MULTIPLIER) % ONE_DIGIT_MASK
    }
    pub fn get_hour(&self) -> u64 {
        (self.time / HOUR_MULTIPLIER) % TWO_DIGIT_MASK
    }
    pub fn get_minute(&self) -> u64 {
        (self.time / MINUTE_MULTIPLIER) % TWO_DIGIT_MASK
    }
    pub fn get_second(&self) -> u64 {
        (self.time / SECOND_MULTIPLIER) % TWO_DIGIT_MASK
    }
    pub fn get_millisecond(&self) -> u64 {
        (self.time / MILLISECOND_MULTIPLIER) % THREE_DIGIT_MASK
    }

    pub fn set_year(&mut self, year: u64) {
        self.time -= self.get_year() * YEAR_MULTIPLIER;
        self.time += year * YEAR_MULTIPLIER;
    }
    pub fn set_month(&mut self, month: u64) {
        self.time -= self.get_month() * MONTH_MULTIPLIER;
        self.time += month * MONTH_MULTIPLIER;
    }
    pub fn set_day(&mut self, day: u64) {
        self.time -= self.get_day() * DAY_MULTIPLIER;
        self.time += day * DAY_MULTIPLIER;
    }
    pub fn set_hour(&mut self, hour: u64) {
        self.time -= self.get_hour() * HOUR_MULTIPLIER;
        self.time += hour * HOUR_MULTIPLIER;
    }
    pub fn set_minute(&mut self, minute: u64) {
        self.time -= self.get_minute() * MINUTE_MULTIPLIER;
        self.time += minute * MINUTE_MULTIPLIER;
    }
    pub fn set_second(&mut self, second: u64) {
        self.time -= self.get_second() * SECOND_MULTIPLIER;
        self.time += second * SECOND_MULTIPLIER;
    }
    pub fn set_millisecond(&mut self, millisecond: u64) {
        self.time -= self.get_millisecond() * MILLISECOND_MULTIPLIER;
        self.time += millisecond * MILLISECOND_MULTIPLIER;
    }

    pub fn add(&mut self, time: TTTime) {
        self.time += time.time;
    }
    pub fn subtract(&mut self, time: TTTime) {
        self.time -= time.time;
    }

    pub fn is_before(&self, time: TTTime) -> bool {
        self.time < time.time
    }
    pub fn is_after(&self, time: TTTime) -> bool {
        self.time > time.time
    }
    pub fn is_equal(&self, time: TTTime) -> bool {
        self.time == time.time
    }

    fn pad_start(&self, value: u64, length: usize) -> String {
        let mut value = value.to_string();
        while value.len() < length {
            value = format!("0{}", value);
        }
        value
    }

    pub fn to_string(&self, format: &str) -> String {
        let mut format = format.to_string();
        format = format.replace("%Y", self.pad_start(self.get_year(), 4).as_str());
        format = format.replace("%m", self.pad_start(self.get_month(), 1).as_str());
        format = format.replace("%d", self.pad_start(self.get_day(), 1).as_str());
        format = format.replace("%H", self.pad_start(self.get_hour(), 2).as_str());
        format = format.replace("%M", self.pad_start(self.get_minute(), 2).as_str());
        format = format.replace("%S", self.pad_start(self.get_second(), 2).as_str());
        format = format.replace("%f", self.pad_start(self.get_millisecond(), 3).as_str());
        format
    }

    pub fn clone(&self) -> Self {
        Self {
            time: self.time,
        }
    }
}

#[cfg(test)]
mod test;

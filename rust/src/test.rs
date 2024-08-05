use crate::{TTTime, TTTimeInput, TTTimePartial};

#[test]
fn tttime_constructor_with_now() {
    let time = TTTime::new(None);
    assert_ne!(time.time, 0);
}

#[test]
fn tttime_constructor_with_partial() {
    let time = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: Some(1202),
        months: Some(3),
        days: Some(5),
        hours: Some(12),
        minutes: Some(23),
        seconds: Some(24),
        milliseconds: None,
    })));
    assert_eq!(time.get_year(), 1202);
    assert_eq!(time.get_month(), 3);
    assert_eq!(time.get_day(), 5);
    assert_eq!(time.get_hour(), 12);
    assert_eq!(time.get_minute(), 23);
    assert_eq!(time.get_second(), 24);
}

#[test]
fn tttime_get_year() {
    let time = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: Some(1202),
        months: None,
        days: None,
        hours: None,
        minutes: None,
        seconds: None,
        milliseconds: None,
    })));
    assert_eq!(time.get_year(), 1202);
}

#[test]
fn tttime_get_month() {
    let time = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: Some(3),
        days: None,
        hours: None,
        minutes: None,
        seconds: None,
        milliseconds: None,
    })));
    assert_eq!(time.get_month(), 3);
}

#[test]
fn tttime_get_day() {
    let time = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: Some(5),
        hours: None,
        minutes: None,
        seconds: None,
        milliseconds: None,
    })));
    assert_eq!(time.get_day(), 5);
}

#[test]
fn tttime_get_hour() {
    let time = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: Some(12),
        minutes: None,
        seconds: None,
        milliseconds: None,
    })));
    assert_eq!(time.get_hour(), 12);
}

#[test]
fn tttime_get_minute() {
    let time = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: None,
        minutes: Some(23),
        seconds: None,
        milliseconds: None,
    })));
    assert_eq!(time.get_minute(), 23);
}

#[test]
fn tttime_get_second() {
    let time = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: None,
        minutes: None,
        seconds: Some(24),
        milliseconds: None,
    })));
    assert_eq!(time.get_second(), 24);
}

#[test]
fn tttime_get_millisecond() {
    let time = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: None,
        minutes: None,
        seconds: None,
        milliseconds: Some(23),
    })));
    assert_eq!(time.get_millisecond(), 23);
}

#[test]
fn tttime_set_year() {
    let mut time = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: None,
        minutes: None,
        seconds: None,
        milliseconds: None,
    })));
    time.set_year(2020);
    assert_eq!(time.get_year(), 2020);
}

#[test]
fn tttime_set_month() {
    let mut time = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: None,
        minutes: None,
        seconds: None,
        milliseconds: None,
    })));
    time.set_month(9);
    assert_eq!(time.get_month(), 9);
}

#[test]
fn tttime_set_day() {
    let mut time = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: None,
        minutes: None,
        seconds: None,
        milliseconds: None,
    })));
    time.set_day(2);
    assert_eq!(time.get_day(), 2);
}

#[test]
fn tttime_set_hour() {
    let mut time = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: None,
        minutes: None,
        seconds: None,
        milliseconds: None,
    })));
    time.set_hour(18);
    assert_eq!(time.get_hour(), 18);
}

#[test]
fn tttime_set_minute() {
    let mut time = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: None,
        minutes: None,
        seconds: None,
        milliseconds: None,
    })));
    time.set_minute(0);
    assert_eq!(time.get_minute(), 0);
}

#[test]
fn tttime_set_second() {
    let mut time = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: None,
        minutes: None,
        seconds: None,
        milliseconds: None,
    })));
    time.set_second(0);
    assert_eq!(time.get_second(), 0);
}

#[test]
fn tttime_set_millisecond() {
    let mut time = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: None,
        minutes: None,
        seconds: None,
        milliseconds: None,
    })));
    time.set_millisecond(0);
    assert_eq!(time.get_millisecond(), 0);
}

#[test]
fn tttime_add() {
    let mut time = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: None,
        minutes: None,
        seconds: None,
        milliseconds: None,
    })));
    time.add(TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: Some(1),
        minutes: None,
        seconds: None,
        milliseconds: None,
    }))));
    assert_eq!(time.get_year(), 0);
    assert_eq!(time.get_month(), 0);
    assert_eq!(time.get_day(), 0);
    assert_eq!(time.get_hour(), 1);
    assert_eq!(time.get_minute(), 0);
    assert_eq!(time.get_second(), 0);
    assert_eq!(time.get_millisecond(), 0);
}

#[test]
fn tttime_subtract() {
    let mut time = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: Some(2),
        minutes: None,
        seconds: None,
        milliseconds: None,
    })));
    time.subtract(TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: Some(1),
        minutes: None,
        seconds: None,
        milliseconds: None,
    }))));
    assert_eq!(time.get_year(), 0);
    assert_eq!(time.get_month(), 0);
    assert_eq!(time.get_day(), 0);
    assert_eq!(time.get_hour(), 1);
    assert_eq!(time.get_minute(), 0);
    assert_eq!(time.get_second(), 0);
    assert_eq!(time.get_millisecond(), 0);
}

#[test]
fn tttime_is_before() {
    let time1 = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: None,
        minutes: None,
        seconds: None,
        milliseconds: None,
    })));
    let time2 = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: Some(1),
        minutes: None,
        seconds: None,
        milliseconds: None,
    })));
    assert!(time1.is_before(time2.clone()));
    assert!(!time2.is_before(time1));
}

#[test]
fn tttime_is_after() {
    let time1 = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: Some(1),
        minutes: None,
        seconds: None,
        milliseconds: None,
    })));
    let time2 = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: Some(2),
        minutes: None,
        seconds: None,
        milliseconds: None,
    })));
    assert!(!time1.is_after(time2.clone()));
    assert!(time2.is_after(time1));
}

#[test]
fn tttime_is_equal() {
    let time1 = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: Some(1),
        minutes: None,
        seconds: None,
        milliseconds: None,
    })));
    let time2 = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: None,
        months: None,
        days: None,
        hours: Some(1),
        minutes: None,
        seconds: None,
        milliseconds: None,
    })));
    assert!(time1.is_equal(time2));
    assert!(time1.is_equal(time1.clone()));
}

#[test]
fn tttime_to_string() {
    let time = TTTime::new(Some(TTTimeInput::Partial(TTTimePartial {
        years: Some(1202),
        months: Some(3),
        days: Some(5),
        hours: Some(12),
        minutes: Some(23),
        seconds: Some(24),
        milliseconds: None,
    })));
    assert_eq!(time.to_string("%Y-%m-%d %H:%M:%S"), "1202-3-5 12:23:24");
}
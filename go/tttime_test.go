package tttime

import (
	"testing"
	"time"
)

func TestTTTimeConstructorWithDate(t *testing.T) {
	time := NewWithDate(time.Date(2020, 10, 21, 18, 0, 0, 0, time.UTC))
	if time.GetYear() != 0 {
		t.Errorf("Expected year to be 0, got %d", time.GetYear())
	}
	if time.GetMonth() != 0 {
		t.Errorf("Expected month to be 0, got %d", time.GetMonth())
	}
	if time.GetDay() != 0 {
		t.Errorf("Expected day to be 0, got %d", time.GetDay())
	}
	if time.GetHour() != 0 {
		t.Errorf("Expected hour to be 0, got %d", time.GetHour())
	}
	if time.GetMinute() != 0 {
		t.Errorf("Expected minute to be 0, got %d", time.GetMinute())
	}
	if time.GetSecond() != 0 {
		t.Errorf("Expected second to be 0, got %d", time.GetSecond())
	}
	if time.GetMillisecond() != 0 {
		t.Errorf("Expected millisecond to be 0, got %d", time.GetMillisecond())
	}
}

func TestTTTimeConstructorWithISO(t *testing.T) {
	time := NewWithISO("2020-10-21T18:00:00.000Z")
	if time.GetYear() != 0 {
		t.Errorf("Expected year to be 0, got %d", time.GetYear())
	}
	if time.GetMonth() != 0 {
		t.Errorf("Expected month to be 0, got %d", time.GetMonth())
	}
	if time.GetDay() != 0 {
		t.Errorf("Expected day to be 0, got %d", time.GetDay())
	}
	if time.GetHour() != 0 {
		t.Errorf("Expected hour to be 0, got %d", time.GetHour())
	}
	if time.GetMinute() != 0 {
		t.Errorf("Expected minute to be 0, got %d", time.GetMinute())
	}
	if time.GetSecond() != 0 {
		t.Errorf("Expected second to be 0, got %d", time.GetSecond())
	}
	if time.GetMillisecond() != 0 {
		t.Errorf("Expected millisecond to be 0, got %d", time.GetMillisecond())
	}
}

func TestTTTimeConstructorWithPartial(t *testing.T) {
	time := NewWithPartial(TTTimePartial{
		years:   1202,
		months:  3,
		days:    5,
		hours:   12,
		minutes: 23,
		seconds: 24,
	})
	if time.GetYear() != 1202 {
		t.Errorf("Expected year to be 1202, got %d", time.GetYear())
	}
	if time.GetMonth() != 3 {
		t.Errorf("Expected month to be 3, got %d", time.GetMonth())
	}
	if time.GetDay() != 5 {
		t.Errorf("Expected day to be 5, got %d", time.GetDay())
	}
	if time.GetHour() != 12 {
		t.Errorf("Expected hour to be 12, got %d", time.GetHour())
	}
	if time.GetMinute() != 23 {
		t.Errorf("Expected minute to be 23, got %d", time.GetMinute())
	}
	if time.GetSecond() != 24 {
		t.Errorf("Expected second to be 24, got %d", time.GetSecond())
	}
}

func TestTTTimeGetYear(t *testing.T) {
	time := NewWithPartial(TTTimePartial{
		years:        1202,
		months:       0,
		days:         0,
		hours:        0,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	})
	if time.GetYear() != 1202 {
		t.Errorf("Expected year to be 1202, got %d", time.GetYear())
	}
}

func TestTTTimeGetMonth(t *testing.T) {
	time := NewWithPartial(TTTimePartial{
		years:        0,
		months:       3,
		days:         0,
		hours:        0,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	})
	if time.GetMonth() != 3 {
		t.Errorf("Expected month to be 3, got %d", time.GetMonth())
	}
}

func TestTTTimeGetDay(t *testing.T) {
	time := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         5,
		hours:        0,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	})
	if time.GetDay() != 5 {
		t.Errorf("Expected day to be 5, got %d", time.GetDay())
	}
}

func TestTTTimeGetHour(t *testing.T) {
	time := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        12,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	})
	if time.GetHour() != 12 {
		t.Errorf("Expected hour to be 12, got %d", time.GetHour())
	}
}

func TestTTTimeGetMinute(t *testing.T) {
	time := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        0,
		minutes:      23,
		seconds:      0,
		milliseconds: 0,
	})
	if time.GetMinute() != 23 {
		t.Errorf("Expected minute to be 23, got %d", time.GetMinute())
	}
}

func TestTTTimeGetSecond(t *testing.T) {
	time := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        0,
		minutes:      0,
		seconds:      24,
		milliseconds: 0,
	})
	if time.GetSecond() != 24 {
		t.Errorf("Expected second to be 24, got %d", time.GetSecond())
	}
}

func TestTTTimeGetMillisecond(t *testing.T) {
	time := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        0,
		minutes:      0,
		seconds:      0,
		milliseconds: 23,
	})
	if time.GetMillisecond() != 23 {
		t.Errorf("Expected millisecond to be 23, got %d", time.GetMillisecond())
	}
}

func TestTTTimeSetYear(t *testing.T) {
	time := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        0,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	})
	time.SetYear(2020)
	if time.GetYear() != 2020 {
		t.Errorf("Expected year to be 2020, got %d", time.GetYear())
	}
}

func TestTTTimeSetMonth(t *testing.T) {
	time := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        0,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	})
	time.SetMonth(9)
	if time.GetMonth() != 9 {
		t.Errorf("Expected month to be 9, got %d", time.GetMonth())
	}
}

func TestTTTimeSetDay(t *testing.T) {
	time := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        0,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	})
	time.SetDay(2)
	if time.GetDay() != 2 {
		t.Errorf("Expected day to be 2, got %d", time.GetDay())
	}
}

func TestTTTimeSetHour(t *testing.T) {
	time := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        0,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	})
	time.SetHour(18)
	if time.GetHour() != 18 {
		t.Errorf("Expected hour to be 18, got %d", time.GetHour())
	}
}

func TestTTTimeSetMinute(t *testing.T) {
	time := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        0,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	})
	time.SetMinute(10)
	if time.GetMinute() != 10 {
		t.Errorf("Expected minute to be 0, got %d", time.GetMinute())
	}
}

func TestTTTimeSetSecond(t *testing.T) {
	time := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        0,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	})
	time.SetSecond(20)
	if time.GetSecond() != 20 {
		t.Errorf("Expected second to be 0, got %d", time.GetSecond())
	}
}

func TestTTTimeSetMillisecond(t *testing.T) {
	time := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        0,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	})
	time.SetMillisecond(30)
	if time.GetMillisecond() != 30 {
		t.Errorf("Expected millisecond to be 0, got %d", time.GetMillisecond())
	}
}

func TestTTTimeAdd(t *testing.T) {
	time := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        0,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	})
	time.Add(NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        1,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	}))
	if time.GetYear() != 0 {
		t.Errorf("Expected year to be 0, got %d", time.GetYear())
	}
	if time.GetMonth() != 0 {
		t.Errorf("Expected month to be 0, got %d", time.GetMonth())
	}
	if time.GetDay() != 0 {
		t.Errorf("Expected day to be 0, got %d", time.GetDay())
	}
	if time.GetHour() != 1 {
		t.Errorf("Expected hour to be 1, got %d", time.GetHour())
	}
	if time.GetMinute() != 0 {
		t.Errorf("Expected minute to be 0, got %d", time.GetMinute())
	}
	if time.GetSecond() != 0 {
		t.Errorf("Expected second to be 0, got %d", time.GetSecond())
	}
	if time.GetMillisecond() != 0 {
		t.Errorf("Expected millisecond to be 0, got %d", time.GetMillisecond())
	}
}

func TestTTTimeSubtract(t *testing.T) {
	time := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        1,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	})
	time.Subtract(NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        0,
		minutes:      1,
		seconds:      0,
		milliseconds: 0,
	}))
	if time.GetYear() != 0 {
		t.Errorf("Expected year to be 0, got %d", time.GetYear())
	}
	if time.GetMonth() != 0 {
		t.Errorf("Expected month to be 0, got %d", time.GetMonth())
	}
	if time.GetDay() != 0 {
		t.Errorf("Expected day to be 0, got %d", time.GetDay())
	}
	if time.GetHour() != 0 {
		t.Errorf("Expected hour to be 0, got %d", time.GetHour())
	}
	if time.GetMinute() != 99 {
		t.Errorf("Expected minute to be 99, got %d", time.GetMinute())
	}
	if time.GetSecond() != 0 {
		t.Errorf("Expected second to be 0, got %d", time.GetSecond())
	}
	if time.GetMillisecond() != 0 {
		t.Errorf("Expected millisecond to be 0, got %d", time.GetMillisecond())
	}
}

func TestTTTimeIsBefore(t *testing.T) {
	time1 := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        0,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	})
	time2 := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        1,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	})
	if !time1.IsBefore(time2) {
		t.Errorf("Expected time1 to be before time2, got %t", time1.IsBefore(time2))
	}
	if time2.IsBefore(time1) {
		t.Errorf("Expected time2 to be before time1, got %t", time2.IsBefore(time1))
	}
}

func TestTTTimeIsAfter(t *testing.T) {
	time1 := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        1,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	})
	time2 := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        0,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	})
	if !time1.IsAfter(time2) {
		t.Errorf("Expected time1 to be after time2, got %t", time1.IsAfter(time2))
	}
	if time2.IsAfter(time1) {
		t.Errorf("Expected time2 to be after time1, got %t", time2.IsAfter(time1))
	}
}

func TestTTTimeIsEqual(t *testing.T) {
	time1 := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        0,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	})
	time2 := NewWithPartial(TTTimePartial{
		years:        0,
		months:       0,
		days:         0,
		hours:        0,
		minutes:      0,
		seconds:      0,
		milliseconds: 0,
	})
	if !time1.IsEqual(time2) {
		t.Errorf("Expected time1 to be equal to time2, got %t", time1.IsEqual(time2))
	}
}

func TestTTTimeToString(t *testing.T) {
	time := NewWithPartial(TTTimePartial{
		years:        1202,
		months:       3,
		days:         5,
		hours:        12,
		minutes:      23,
		seconds:      24,
		milliseconds: 0,
	})
	if time.ToString("%Y-%m-%d %H:%M:%S") != "1202-3-5 12:23:24" {
		t.Errorf("Expected string to be 1202-3-5 12:23:24, got %s", time.ToString("%Y-%m-%d %H:%M:%S"))
	}
}

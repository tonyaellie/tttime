package tttime

import (
	"fmt"
	"strings"
	"time"
)

const MILLISECOND_MULTIPLIER = 1
const SECOND_MULTIPLIER = MILLISECOND_MULTIPLIER * 1000
const MINUTE_MULTIPLIER = SECOND_MULTIPLIER * 100
const HOUR_MULTIPLIER = MINUTE_MULTIPLIER * 100
const DAY_MULTIPLIER = HOUR_MULTIPLIER * 100
const MONTH_MULTIPLIER = DAY_MULTIPLIER * 10
const YEAR_MULTIPLIER = MONTH_MULTIPLIER * 10

const ONE_DIGIT_MASK = 10
const TWO_DIGIT_MASK = 100
const THREE_DIGIT_MASK = 1000

var TTTimeEpoch = time.Date(2020, 10, 21, 18, 0, 0, 0, time.UTC)

type TTTimePartial struct {
	years        int
	months       int
	days         int
	hours        int
	minutes      int
	seconds      int
	milliseconds int
}

type TTTime struct {
	time int
}

func New() TTTime {
	return fromStandard(time.Now().UTC())
}

func NewWithDate(date time.Time) TTTime {
	return fromStandard(date)
}

func NewWithISO(iso string) TTTime {
	parsedTime, err := time.Parse(time.RFC3339, iso)
	if err != nil {
		// Handle the error appropriately, e.g., return a zero value or a default value
		return TTTime{
			time: 0,
		}
	}
	return fromStandard(parsedTime)
}

func NewWithPartial(partial TTTimePartial) TTTime {
	time := 0
	if partial.years != 0 {
		time += partial.years * YEAR_MULTIPLIER
	}
	if partial.months != 0 {
		time += partial.months * MONTH_MULTIPLIER
	}
	if partial.days != 0 {
		time += partial.days * DAY_MULTIPLIER
	}
	if partial.hours != 0 {
		time += partial.hours * HOUR_MULTIPLIER
	}
	if partial.minutes != 0 {
		time += partial.minutes * MINUTE_MULTIPLIER
	}
	if partial.seconds != 0 {
		time += partial.seconds * SECOND_MULTIPLIER
	}
	if partial.milliseconds != 0 {
		time += partial.milliseconds * MILLISECOND_MULTIPLIER
	}
	return TTTime{
		time: time,
	}
}

func fromStandard(standardTime time.Time) TTTime {
	hoursSinceEpoch := standardTime.Sub(TTTimeEpoch).Hours()
	return TTTime{
		time: int(hoursSinceEpoch * HOUR_MULTIPLIER),
	}
}

func (tttime TTTime) toStandard() time.Time {
	return TTTimeEpoch.Add(
		time.Duration(tttime.time/HOUR_MULTIPLIER) * time.Second * time.Minute * time.Hour,
	)
}

func (tttime TTTime) GetYear() int {
	return tttime.time / YEAR_MULTIPLIER
}

func (tttime TTTime) GetMonth() int {
	return (tttime.time / MONTH_MULTIPLIER) % ONE_DIGIT_MASK
}

func (tttime TTTime) GetDay() int {
	return (tttime.time / DAY_MULTIPLIER) % ONE_DIGIT_MASK
}

func (tttime TTTime) GetHour() int {
	return (tttime.time / HOUR_MULTIPLIER) % TWO_DIGIT_MASK
}

func (tttime TTTime) GetMinute() int {
	return (tttime.time / MINUTE_MULTIPLIER) % TWO_DIGIT_MASK
}

func (tttime TTTime) GetSecond() int {
	return (tttime.time / SECOND_MULTIPLIER) % TWO_DIGIT_MASK
}

func (tttime TTTime) GetMillisecond() int {
	return (tttime.time / MILLISECOND_MULTIPLIER) % THREE_DIGIT_MASK
}

func (tttime *TTTime) SetYear(year int) {
	tttime.time -= tttime.GetYear() * YEAR_MULTIPLIER
	tttime.time += year * YEAR_MULTIPLIER
}

func (tttime *TTTime) SetMonth(month int) {
	if month > 9 {
		panic("TTTime: Month can only be between 0-9")
	}
	tttime.time -= tttime.GetMonth() * MONTH_MULTIPLIER
	tttime.time += month * MONTH_MULTIPLIER
}

func (tttime *TTTime) SetDay(day int) {
	if day > 9 {
		panic("TTTime: Day can only be between 0-9")
	}
	tttime.time -= tttime.GetDay() * DAY_MULTIPLIER
	tttime.time += day * DAY_MULTIPLIER
}

func (tttime *TTTime) SetHour(hour int) {
	if hour > 99 {
		panic("TTTime: Hour can only be between 0-99")
	}
	tttime.time -= tttime.GetHour() * HOUR_MULTIPLIER
	tttime.time += hour * HOUR_MULTIPLIER
}

func (tttime *TTTime) SetMinute(minute int) {
	if minute > 99 {
		panic("TTTime: Minute can only be between 0-99")
	}
	tttime.time -= tttime.GetMinute() * MINUTE_MULTIPLIER
	tttime.time += minute * MINUTE_MULTIPLIER
}

func (tttime *TTTime) SetSecond(second int) {
	if second > 99 {
		panic("TTTime: Second can only be between 0-99")
	}
	tttime.time -= tttime.GetSecond() * SECOND_MULTIPLIER
	tttime.time += second * SECOND_MULTIPLIER
}

func (tttime *TTTime) SetMillisecond(millisecond int) {
	if millisecond > 999 {
		panic("TTTime: Millisecond can only be between 0-999")
	}
	tttime.time -= tttime.GetMillisecond() * MILLISECOND_MULTIPLIER
	tttime.time += millisecond * MILLISECOND_MULTIPLIER
}

func (tttime *TTTime) Add(time TTTime) {
	tttime.time += time.time
}

func (tttime *TTTime) Subtract(time TTTime) {
	tttime.time -= time.time
}

func (tttime TTTime) IsBefore(time TTTime) bool {
	return tttime.time < time.time
}

func (tttime TTTime) IsAfter(time TTTime) bool {
	return tttime.time > time.time
}

func (tttime TTTime) IsEqual(time TTTime) bool {
	return tttime.time == time.time
}

func (tttime TTTime) ToString(format string) string {
	format = strings.Replace(format, "%Y", fmt.Sprintf("%04d", tttime.GetYear()), -1)
	format = strings.Replace(format, "%m", fmt.Sprintf("%01d", tttime.GetMonth()), -1)
	format = strings.Replace(format, "%d", fmt.Sprintf("%01d", tttime.GetDay()), -1)
	format = strings.Replace(format, "%H", fmt.Sprintf("%02d", tttime.GetHour()), -1)
	format = strings.Replace(format, "%M", fmt.Sprintf("%02d", tttime.GetMinute()), -1)
	format = strings.Replace(format, "%S", fmt.Sprintf("%02d", tttime.GetSecond()), -1)
	format = strings.Replace(format, "%f", fmt.Sprintf("%03d", tttime.GetMillisecond()), -1)
	return format
}

func (tttime TTTime) ToStandardMilliseconds() int {
	time := tttime.toStandard()
	return int(time.UnixMilli())
}

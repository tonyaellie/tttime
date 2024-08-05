# TTTime Libraries

TTTime is a base 10 time representation. It is used to represent time in a way that is easy to understand and work with. Time is measured from 18:00:00 - 21/10/2020 UTC Standard Time (Day Zero).

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

## Packages

- [JavaScript/TypeScript](https://www.npmjs.com/package/tttime)
- [Python](https://pypi.org/project/tttime/)
- [Rust](https://crates.io/crates/tttime) - Implementation is missing some features, mainly interactions with standard time, it also is slightly buggy.
- [Go](https://github.com/tonyaellie/tttime/tree/main/go)
- More coming soon!

## API

This is the standard API for TTTime, some of the packages may lack some features or have slightly different implementations.

```python
class TTTime:
    # Constructor
    method constructor(input: Date | TTTimePartial | string):
        if input is TTTimePartial:
            for each field in input:
                set field using setter methods
        else if input is string:
            convert string to Date and store as time since epoch
        else:
            store current time since epoch

    method toStandard() -> Date:
        convert TTTime milliseconds to standard Date object

    # Getter Methods
    method getYear() -> number:
        return year component from stored time

    method getMonth() -> number:
        return month component from stored time

    method getDay() -> number:
        return day component from stored time

    method getHour() -> number:
        return hour component from stored time

    method getMinute() -> number:
        return minute component from stored time

    method getSecond() -> number:
        return second component from stored time

    method getMillisecond() -> number:
        return millisecond component from stored time

    # Setter Methods
    method setYear(year: number):
        update year component in stored time

    method setMonth(month: number):
        validate month and update month component in stored time

    method setDay(day: number):
        validate day and update day component in stored time

    method setHour(hour: number):
        validate hour and update hour component in stored time

    method setMinute(minute: number):
        validate minute and update minute component in stored time

    method setSecond(second: number):
        validate second and update second component in stored time

    method setMillisecond(millisecond: number):
        validate millisecond and update millisecond component in stored time

    # Arithmetic Methods
    method add(time: TTTime | Date | TTTimePartial):
        add given time to stored time

    method subtract(time: TTTime | Date | TTTimePartial):
        subtract given time from stored time

    # Comparison Methods
    method isBefore(time: TTTime | Date | TTTimePartial) -> boolean:
        compare stored time with given time

    method isAfter(time: TTTime | Date | TTTimePartial) -> boolean:
        compare stored time with given time

    method isEqual(time: TTTime | Date | TTTimePartial) -> boolean:
        compare stored time with given time

    # Formatting Method
    method toString(format: string = '%Y-%m-%d %H:%M:%S') -> string:
        format stored time based on given format

    # Utility Method
    method toStandardMilliseconds() -> number:
        return stored time in milliseconds since epoch
```

## Contributing

Contributions are welcome! Please open an issue or pull request if you have any suggestions or want to contribute at [https://github.com/tonyaellie/tttime](https://github.com/tonyaellie/tttime).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.
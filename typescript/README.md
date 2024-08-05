# TTTime TypeScript Library

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
npm install tttime
```

or with pnpm

```bash
pnpm install tttime
```

```typescript
import { TTTime } from 'tttime';

const time = new TTTime();
time.setYear(2020);
time.setMonth(9);
time.setDay(2);
time.setHour(18);
time.setMinute(0);
time.setSecond(0);
time.setMillisecond(0);

console.log(time.toString('%Y-%m-%d %H:%M:%S'));
// 2020-09-02 18:00:00
```

## API

```typescript
class TTTime {
  constructor(input?: Date | TTTimePartial | string);
  toStandard(): Date;

  getYear(): number;
  getMonth(): number;
  getDay(): number;
  getHour(): number;
  getMinute(): number;
  getSecond(): number;
  getMillisecond(): number;

  setYear(year: number): void;
  setMonth(month: number): void;
  setDay(day: number): void;
  setHour(hour: number): void;
  setMinute(minute: number): void;
  setSecond(second: number): void;
  setMillisecond(millisecond: number): void;

  add(time: TTTime | Date | TTTimePartial): void;
  subtract(time: TTTime | Date | TTTimePartial): void;

  isBefore(time: TTTime | Date | TTTimePartial): boolean;
  isAfter(time: TTTime | Date | TTTimePartial): boolean;
  isEqual(time: TTTime | Date | TTTimePartial): boolean;

  toString(format: string): string;
  toStandardMilliseconds(): number;
}
```

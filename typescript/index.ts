/**
 * A partial representation of a TTTime object.
 */
type TTTimePartial = {
  years?: number;
  months?: number;
  days?: number;
  hours?: number;
  minutes?: number;
  seconds?: number;
  milliseconds?: number;
};

const MILLISECOND_MULTIPLIER = 1;
const SECOND_MULTIPLIER = MILLISECOND_MULTIPLIER * 1000;
const MINUTE_MULTIPLIER = SECOND_MULTIPLIER * 100;
const HOUR_MULTIPLIER = MINUTE_MULTIPLIER * 100;
const DAY_MULTIPLIER = HOUR_MULTIPLIER * 100;
const MONTH_MULTIPLIER = DAY_MULTIPLIER * 10;
const YEAR_MULTIPLIER = MONTH_MULTIPLIER * 10;

const ONE_DIGIT_MASK = 10;
const TWO_DIGIT_MASK = 100;
const THREE_DIGIT_MASK = 1000;

export class TTTime {
  /**
   * The epoch is 18:00:00 - 21/10/2020 UTC.
   * This is to make it easier to work with standard time.
   */
  private static epoch = new Date(2020, 9, 21, 18, 0, 0);
  /**
   * The time is stored as the number of milliseconds since the epoch.
   * This is to make it easier to work with the time.
   */
  private time: number;

  /**
   * Creates a new TTTime object.
   * @param input The input to create the TTTime object from, either a Date, TTTimePartial, or an ISO 8601 string, by default it will be set to the current time.
   */
  constructor(input?: Date | TTTimePartial | string) {
    this.time = 0;
    if (typeof input === 'object' && !(input instanceof Date)) {
      if (input.years) {
        this.setYear(input.years);
      }
      if (input.months) {
        this.setMonth(input.months);
      }
      if (input.days) {
        this.setDay(input.days);
      }
      if (input.hours) {
        this.setHour(input.hours);
      }
      if (input.minutes) {
        this.setMinute(input.minutes);
      }
      if (input.seconds) {
        this.setSecond(input.seconds);
      }
      if (input.milliseconds) {
        this.setMillisecond(input.milliseconds);
      }
    } else if (typeof input === 'string') {
      this.time = this.fromStandard(new Date(input));
    } else {
      this.time = this.fromStandard(input || new Date());
    }
  }

  /**
   * Converts a Date object to a TTTime object.
   * @param standardTime The Date object to convert.
   * @returns The TTTime in milliseconds since the epoch.
   */
  private fromStandard(standardTime: Date): number {
    const hoursSinceEpoch =
      (standardTime.getTime() - TTTime.epoch.getTime()) / 1000 / 60 / 60;
    return Math.floor(hoursSinceEpoch * HOUR_MULTIPLIER);
  }

  /**
   * Converts a TTTime object to a Date object.
   * @returns The Date object.
   */
  toStandard(): Date {
    return new Date(
      (this.time / HOUR_MULTIPLIER) * 1000 * 60 * 60 + TTTime.epoch.getTime()
    );
  }

  /**
   * Gets the year of the TTTime object.
   * @returns The year.
   */
  getYear(): number {
    return Math.floor(this.time / YEAR_MULTIPLIER);
  }
  /**
   * Gets the month of the TTTime object.
   * @returns The month.
   */
  getMonth(): number {
    return Math.floor(this.time / MONTH_MULTIPLIER) % ONE_DIGIT_MASK;
  }
  /**
   * Gets the day of the TTTime object.
   * @returns The day.
   */
  getDay(): number {
    return Math.floor(this.time / DAY_MULTIPLIER) % ONE_DIGIT_MASK;
  }
  /**
   * Gets the hour of the TTTime object.
   * @returns The hour.
   */
  getHour(): number {
    return Math.floor(this.time / HOUR_MULTIPLIER) % TWO_DIGIT_MASK;
  }
  /**
   * Gets the minute of the TTTime object.
   * @returns The minute.
   */
  getMinute(): number {
    return Math.floor(this.time / MINUTE_MULTIPLIER) % TWO_DIGIT_MASK;
  }
  /**
   * Gets the second of the TTTime object.
   * @returns The second.
   */
  getSecond(): number {
    return Math.floor(this.time / SECOND_MULTIPLIER) % TWO_DIGIT_MASK;
  }
  /**
   * Gets the millisecond of the TTTime object.
   * @returns The millisecond.
   */
  getMillisecond(): number {
    return Math.floor(this.time / MILLISECOND_MULTIPLIER) % THREE_DIGIT_MASK;
  }

  /**
   * Sets the year of the TTTime object.
   * @param year The year to set.
   */
  setYear(year: number): void {
    // remove the current year and add the new year
    this.time -= this.getYear() * YEAR_MULTIPLIER;
    this.time += year * YEAR_MULTIPLIER;
  }
  /**
   * Sets the month of the TTTime object.
   * @param month The month to set.
   */
  setMonth(month: number): void {
    // if more than 1 digit throw error
    if (month > 9) {
      throw new Error('TTTime: Month can only be between 0-9');
    }
    // remove the current month and add the new month
    this.time -= this.getMonth() * MONTH_MULTIPLIER;
    this.time += month * MONTH_MULTIPLIER;
  }
  /**
   * Sets the day of the TTTime object.
   * @param day The day to set.
   */
  setDay(day: number): void {
    // if more than 1 digit throw error
    if (day > 9) {
      throw new Error('TTTime: Day can only be between 0-9');
    }
    // remove the current day and add the new day
    this.time -= this.getDay() * DAY_MULTIPLIER;
    this.time += day * DAY_MULTIPLIER;
  }
  /**
   * Sets the hour of the TTTime object.
   * @param hour The hour to set.
   */
  setHour(hour: number): void {
    // if more than 2 digit throw error
    if (hour > 99) {
      throw new Error('TTTime: Hour can only be between 0-99');
    }
    // remove the current hour and add the new hour
    this.time -= this.getHour() * HOUR_MULTIPLIER;
    this.time += hour * HOUR_MULTIPLIER;
  }
  /**
   * Sets the minute of the TTTime object.
   * @param minute The minute to set.
   */
  setMinute(minute: number): void {
    // if more than 2 digit throw error
    if (minute > 99) {
      throw new Error('TTTime: Minute can only be between 0-99');
    }
    // remove the current minute and add the new minute
    this.time -= this.getMinute() * MINUTE_MULTIPLIER;
    this.time += minute * MINUTE_MULTIPLIER;
  }
  /**
   * Sets the second of the TTTime object.
   * @param second The second to set.
   */
  setSecond(second: number): void {
    // if more than 2 digit throw error
    if (second > 99) {
      throw new Error('TTTime: Second can only be between 0-99');
    }
    // remove the current second and add the new second
    this.time -= this.getSecond() * SECOND_MULTIPLIER;
    this.time += second * SECOND_MULTIPLIER;
  }
  /**
   * Sets the millisecond of the TTTime object.
   * @param millisecond The millisecond to set.
   */
  setMillisecond(millisecond: number): void {
    // if more than 3 digit throw error
    if (millisecond > 999) {
      throw new Error('TTTime: Millisecond can only be between 0-999');
    }
    // remove the current millisecond and add the new millisecond
    this.time -= this.getMillisecond() * MILLISECOND_MULTIPLIER;
    this.time += millisecond * MILLISECOND_MULTIPLIER;
  }

  /**
   * Adds a TTTime object or a Date object or a TTTimePartial object to the TTTime object.
   * @param time The TTTime object or a Date object or a TTTimePartial object to add.
   */
  add(time: TTTime | Date | TTTimePartial): void {
    if (time instanceof TTTime) {
      this.time += time.time;
    } else {
      this.time += new TTTime(time).time;
    }
  }
  /**
   * Subtracts a TTTime object or a Date object or a TTTimePartial object from the TTTime object.
   * @param time The TTTime object or a Date object or a TTTimePartial object to subtract.
   */
  subtract(time: TTTime | Date | TTTimePartial): void {
    if (time instanceof TTTime) {
      this.time -= time.time;
    } else {
      this.time -= new TTTime(time).time;
    }
  }

  /**
   * Checks if the TTTime object is before a TTTime object, Date object, or a TTTimePartial object.
   * @param time The TTTime object, Date object, or a TTTimePartial object to compare against.
   * @returns True if the TTTime object is before the other object, false otherwise.
   */
  isBefore(time: TTTime | Date | TTTimePartial): boolean {
    if (time instanceof TTTime) {
      return this.time < time.time;
    } else {
      return this.time < new TTTime(time).time;
    }
  }
  /**
   * Checks if the TTTime object is after a TTTime object, Date object, or a TTTimePartial object.
   * @param time The TTTime object, Date object, or a TTTimePartial object to compare against.
   * @returns True if the TTTime object is after the other object, false otherwise.
   */
  isAfter(time: TTTime | Date | TTTimePartial): boolean {
    if (time instanceof TTTime) {
      return this.time > time.time;
    } else {
      return this.time > new TTTime(time).time;
    }
  }
  /**
   * Checks if the TTTime object is equal to a TTTime object, Date object, or a TTTimePartial object.
   * @param time The TTTime object, Date object, or a TTTimePartial object to compare against.
   * @returns True if the TTTime object is equal to the other object, false otherwise.
   */
  isEqual(time: TTTime | Date | TTTimePartial): boolean {
    if (time instanceof TTTime) {
      return this.time === time.time;
    } else {
      return this.time === new TTTime(time).time;
    }
  }

  /**
   * Formats the TTTime object as a string.
   * Use %Y for the year, %m for the month, %d for the day, %H for the hour, %M for the minute, %S for the second, and %f for the millisecond.
   * @example toString('%m/%d/%Y %H:%M:%S.%f') -> '10/21/2020 18:00:00.000'
   * @param format The format to use for the string. Default is '%Y-%m-%d %H:%M:%S'.
   * @returns The formatted string.
   */
  toString(format: string = '%Y-%m-%d %H:%M:%S'): string {
    // %Y - year
    // %m - month
    // %d - day
    // %H - hour
    // %M - minute
    // %S - second
    // %f - millisecond
    return format
      .replace('%Y', this.getYear().toString().padStart(4, '0'))
      .replace('%m', this.getMonth().toString())
      .replace('%d', this.getDay().toString())
      .replace('%H', this.getHour().toString().padStart(2, '0'))
      .replace('%M', this.getMinute().toString().padStart(2, '0'))
      .replace('%S', this.getSecond().toString().padStart(2, '0'))
      .replace('%f', this.getMillisecond().toString().padStart(3, '0'));
  }
  /**
   * Converts the TTTime object to a number representing the number of milliseconds since the epoch.
   * @example
   * const time = new TTTime({
   *   second: 3,
   * });
   * @returns The number of milliseconds since the epoch.
   */
  toStandardMilliseconds(): number {
    return this.toStandard().getTime();
  }
}

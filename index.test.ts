import { TTTime } from './index';
import { expect, test, describe } from 'vitest';

test('TTTime constructor with Date', () => {
  const time = new TTTime(new Date(2020, 9, 21, 18, 0, 0));
  expect(time.getYear()).toBe(0);
  expect(time.getMonth()).toBe(0);
  expect(time.getDay()).toBe(0);
  expect(time.getHour()).toBe(0);
  expect(time.getMinute()).toBe(0);
  expect(time.getSecond()).toBe(0);
  expect(time.getMillisecond()).toBe(0);
});

test('TTTime constructor with TTTimePartial', () => {
  const time = new TTTime({
    years: 1202,
    months: 3,
    days: 5,
    hours: 12,
    minutes: 23,
    seconds: 24,
  });
  expect(time.getYear()).toBe(1202);
  expect(time.getMonth()).toBe(3);
  expect(time.getDay()).toBe(5);
  expect(time.getHour()).toBe(12);
  expect(time.getMinute()).toBe(23);
  expect(time.getSecond()).toBe(24);
});

test('TTTime constructor with string', () => {
  const time = new TTTime('2020-10-21T18:00:00.000Z');
  expect(time.getYear()).toBe(0);
  expect(time.getMonth()).toBe(0);
  expect(time.getDay()).toBe(0);
  expect(time.getHour()).toBe(0);
  expect(time.getMinute()).toBe(0);
  expect(time.getSecond()).toBe(0);
  expect(time.getMillisecond()).toBe(0);
});

test('TTTime toStandard', () => {
  const time = new TTTime({});
  expect(time.toStandard().toISOString()).toBe('2020-10-21T18:00:00.000Z');
});

test('TTTime getYear', () => {
  const time = new TTTime({});
  expect(time.getYear()).toBe(0);
});

test('TTTime getMonth', () => {
  const time = new TTTime({});
  expect(time.getMonth()).toBe(0);
});

test('TTTime getDay', () => {
  const time = new TTTime({});
  expect(time.getDay()).toBe(0);
});

test('TTTime getHour', () => {
  const time = new TTTime({});
  expect(time.getHour()).toBe(0);
});

test('TTTime getMinute', () => {
  const time = new TTTime({});
  expect(time.getMinute()).toBe(0);
});

test('TTTime getSecond', () => {
  const time = new TTTime({});
  expect(time.getSecond()).toBe(0);
});

test('TTTime getMillisecond', () => {
  const time = new TTTime({});
  expect(time.getMillisecond()).toBe(0);
});

test('TTTime setYear', () => {
  const time = new TTTime();
  time.setYear(2020);
  expect(time.getYear()).toBe(2020);
});

test('TTTime setMonth', () => {
  const time = new TTTime();
  time.setMonth(9);
  expect(time.getMonth()).toBe(9);
});

test('TTTime setDay', () => {
  const time = new TTTime();
  time.setDay(2);
  expect(time.getDay()).toBe(2);
});

test('TTTime setHour', () => {
  const time = new TTTime();
  time.setHour(18);
  expect(time.getHour()).toBe(18);
});

test('TTTime setMinute', () => {
  const time = new TTTime();
  time.setMinute(0);
  expect(time.getMinute()).toBe(0);
});

test('TTTime setSecond', () => {
  const time = new TTTime();
  time.setSecond(0);
  expect(time.getSecond()).toBe(0);
});

test('TTTime setMillisecond', () => {
  const time = new TTTime();
  time.setMillisecond(0);
  expect(time.getMillisecond()).toBe(0);
});

test('TTTime toString', () => {
  const time = new TTTime({});
  expect(time.toString('%Y-%m-%d %H:%M:%S')).toBe('0000-0-0 00:00:00');
});

test('TTTime toStandardMilliseconds', () => {
  const time = new TTTime({});
  expect(time.toStandardMilliseconds()).toBe(
    new Date(2020, 9, 21, 18, 0, 0).getTime()
  );
});

test('TTTime add', () => {
  const time = new TTTime({});
  time.add(
    new TTTime({
      hours: 1,
      milliseconds: 23,
    })
  );
  expect(time.getYear()).toBe(0);
  expect(time.getMonth()).toBe(0);
  expect(time.getDay()).toBe(0);
  expect(time.getHour()).toBe(1);
  expect(time.getMinute()).toBe(0);
  expect(time.getSecond()).toBe(0);
  expect(time.getMillisecond()).toBe(23);
});

test('TTTime subtract', () => {
  const time = new TTTime({
    years: 2,
    hours: 20,
    milliseconds: 50,
  });
  time.subtract({
    months: 1,
    hours: 1,
    milliseconds: 23,
  });
  expect(time.getYear()).toBe(1);
  expect(time.getMonth()).toBe(9);
  expect(time.getDay()).toBe(0);
  expect(time.getHour()).toBe(19);
  expect(time.getMinute()).toBe(0);
  expect(time.getSecond()).toBe(0);
  expect(time.getMillisecond()).toBe(27);
});

test('TTTime isBefore', () => {
  const time = new TTTime({
    years: 2,
    hours: 20,
    milliseconds: 50,
  });
  const time2 = new TTTime({
    years: 2,
    days: 2,
    hours: 20,
    milliseconds: 50,
  });
  expect(time.isBefore(time2)).toBe(true);
  expect(time2.isBefore(time)).toBe(false);
});

test('TTTime isAfter', () => {
  const time = new TTTime({
    years: 2,
    hours: 20,
    milliseconds: 50,
  });
  const time2 = new TTTime({
    years: 2,
    days: 2,
    hours: 20,
    milliseconds: 50,
  });
  expect(time.isAfter(time2)).toBe(false);
  expect(time2.isAfter(time)).toBe(true);
});

test('TTTime isEqual', () => {
  const time = new TTTime({
    years: 2,
    hours: 20,
    milliseconds: 50,
  });
  const time2 = new TTTime({
    years: 2,
    days: 2,
    hours: 20,
    milliseconds: 50,
  });
  expect(time.isEqual(time2)).toBe(false);
  expect(time.isEqual(time)).toBe(true);
});

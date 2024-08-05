# TTTime Go Library

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
go get github.com/tonyaellie/tttime/go
```

```go
package main

import (
	"fmt"
	"github.com/tonyaellie/tttime/go"
)

func main() {
	time := tttime.NewWithPartial(tttime.TTTimePartial{
		years:        1202,
		months:       3,
		days:         5,
		hours:        12,
		minutes:      23,
		seconds:      24,
		milliseconds: 0,
	})
	fmt.Println(time.ToString("%Y-%m-%d %H:%M:%S"))
	// 1202-3-5 12:23:24
```

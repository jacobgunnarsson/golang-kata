package kata

import (
	"time"
)

func unluckydays(year int) (numUnluckyDays int) {
	startDate := time.Date(year, time.January, 13, 0, 0, 0, 0, time.UTC)

	for d := startDate; d.Year() == startDate.Year(); d = d.AddDate(0, 1, 0) {
		if d.Weekday() == time.Friday {
			numUnluckyDays++
		}
	}

	return numUnluckyDays
}

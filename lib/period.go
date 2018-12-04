package lib

import (
	"time"
)

// Period - Hold from and to timestamps for a given period
type Period struct {
	from int
	to   int
}

func (p *Period) String() (string, string) {
	return time.Unix(int64(p.from), 0).UTC().String(), time.Unix(int64(p.to), 0).UTC().String()
}

// GetLastMonthPeriod - Get Period for "LastMonth"
func GetLastMonthPeriod(now time.Time) Period {
	//now := time.Now().UTC()
	firstOfTheMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDayOfLastMonth := firstOfTheMonth.AddDate(0, 0, -1)

	// lastMonth := now.AddDate(0, -1, 0)

	fromTime := time.Date(lastDayOfLastMonth.Year(), lastDayOfLastMonth.Month(), 1, 0, 0, 0, 0, time.UTC)
	toTime := time.Date(lastDayOfLastMonth.Year(), lastDayOfLastMonth.Month(), lastDayOfLastMonth.Day(), 23, 59, 59, 0, time.UTC)

	// fmt.Printf("From: %s, To: %s\n", fromTime.String(), toTime.String())

	return Period{
		from: int(fromTime.Unix()),
		to:   int(toTime.Unix()),
	}
}

// GetTodayPeriod - Get Period for "Today"
func GetTodayPeriod(now time.Time) Period {
	//now := time.Now().UTC()
	fromTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	toTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)

	return Period{
		from: int(fromTime.Unix()),
		to:   int(toTime.Unix()),
	}
}

// GetThisYearPeriod - Get Period for "ThisYear"
func GetThisYearPeriod(now time.Time) Period {
	//now := time.Now().UTC()
	fromTime := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	toTime := time.Date(now.Year(), 12, 31, 23, 59, 59, 0, time.UTC)

	return Period{
		from: int(fromTime.Unix()),
		to:   int(toTime.Unix()),
	}
}

// GetLastYearPeriod - Get Period for "LastYear"
func GetLastYearPeriod(now time.Time) Period {
	//now := time.Now().UTC()
	fromTime := time.Date(now.Year()-1, 1, 1, 0, 0, 0, 0, time.UTC)
	toTime := time.Date(now.Year()-1, 12, 31, 23, 59, 59, 0, time.UTC)

	return Period{
		from: int(fromTime.Unix()),
		to:   int(toTime.Unix()),
	}
}

// GetPeriodByName - Get Period by name
func GetPeriodByName(name string, now int64) Period {
	// Initialize "now" in case it was not provided
	if now == 0 {
		now = time.Now().UTC().Unix()
	}
	date := time.Unix(now, 0)

	switch name {
	case "LastMonth":
		return GetLastMonthPeriod(date)
	case "ThisYear":
		return GetThisYearPeriod(date)
	case "LastYear":
		return GetLastYearPeriod(date)
	case "Today":
		return GetTodayPeriod(date)
	default:
		return GetTodayPeriod(date)
	}
}
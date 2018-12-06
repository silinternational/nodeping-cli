package lib

import (
	"time"
)


const periodThisMonth =  "ThisMonth"
const periodLastMonth =  "LastMonth"
const periodThisYear =  "ThisYear"
const periodLastYear =  "LastYear"

func GetValidPeriods() []string{
	return []string{
		periodThisMonth,
		periodLastMonth,
		periodThisYear,
		periodLastYear,
	}
}

func IsPeriodValid(pType string) bool {
	for _, validPeriod:= range GetValidPeriods() {
		if pType == validPeriod {
			return true
		}
	}

	return false
}

// Period - Hold From and To timestamps for a given period
type Period struct {
	From int64
	To   int64
}

func (p *Period) String() (string, string) {
	return time.Unix(int64(p.From), 0).UTC().String(), time.Unix(int64(p.To), 0).UTC().String()
}

// GetLastMonthPeriod - Get Period for "ThisMonth"
func GetThisMonthPeriod(now time.Time) Period {
	fromTime := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	toTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)

	// fmt.Printf("From: %s, To: %s\n", fromTime.String(), toTime.String())

	return Period{
		From: int64(fromTime.Unix()),
		To:   int64(toTime.Unix()),
	}
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
		From: int64(fromTime.Unix()),
		To:   int64(toTime.Unix()),
	}
}

// GetTodayPeriod - Get Period for "Today"
func GetTodayPeriod(now time.Time) Period {
	//now := time.Now().UTC()
	fromTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	toTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)

	return Period{
		From: int64(fromTime.Unix()),
		To:   int64(toTime.Unix()),
	}
}

// GetThisYearPeriod - Get Period for "ThisYear"
func GetThisYearPeriod(now time.Time) Period {
	//now := time.Now().UTC()
	fromTime := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	toTime := time.Date(now.Year(), 12, 31, 23, 59, 59, 0, time.UTC)

	return Period{
		From: int64(fromTime.Unix()),
		To:   int64(toTime.Unix()),
	}
}

// GetLastYearPeriod - Get Period for "LastYear"
func GetLastYearPeriod(now time.Time) Period {
	//now := time.Now().UTC()
	fromTime := time.Date(now.Year()-1, 1, 1, 0, 0, 0, 0, time.UTC)
	toTime := time.Date(now.Year()-1, 12, 31, 23, 59, 59, 0, time.UTC)

	return Period{
		From: int64(fromTime.Unix()),
		To:   int64(toTime.Unix()),
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
	case periodThisMonth:
		return GetThisMonthPeriod(date)
	case periodLastMonth:
		return GetLastMonthPeriod(date)
	case periodThisYear:
		return GetThisYearPeriod(date)
	case periodLastYear:
		return GetLastYearPeriod(date)
	default:
		return GetTodayPeriod(date)
	}
}

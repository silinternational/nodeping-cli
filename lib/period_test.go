package lib

import (
	"testing"
)

func TestGetLastMonthPeriod(t *testing.T) {
	period := GetPeriodByName("LastMonth", 1502980112)
	if period.from != 1498867200 {
		t.Errorf("Period 'from' time not correct. Expected 1498867200, got %d", period.from)
	}

	if period.to != 1501545599 {
		t.Errorf("Period 'to' time not correct. Expected 1501545599, got %d", period.to)
	}
}

func TestGetTodayPeriod(t *testing.T) {
	period := GetPeriodByName("Today", 1502980112)
	if period.from != 1502928000 {
		t.Errorf("Period 'from' time not correct. Expected 1502928000, got %d", period.from)
	}

	if period.to != 1503014399 {
		t.Errorf("Period 'to' time not correct. Expected 1503014399, got %d", period.to)
	}
}

func TestGetThisYearPeriod(t *testing.T) {
	period := GetPeriodByName("ThisYear", 1502980112)
	if period.from != 1483228800 {
		t.Errorf("Period 'from' time not correct. Expected 1483228800, got %d", period.from)
	}

	if period.to != 1514764799 {
		t.Errorf("Period 'to' time not correct. Expected 1514764799, got %d", period.to)
	}
}

func TestGetLastYearPeriod(t *testing.T) {
	period := GetPeriodByName("LastYear", 1502980112)
	if period.from != 1451606400 {
		t.Errorf("Period 'from' time not correct. Expected 1451606400, got %d", period.from)
	}

	if period.to != 1483228799 {
		t.Errorf("Period 'to' time not correct. Expected 1483228799, got %d", period.to)
	}
}
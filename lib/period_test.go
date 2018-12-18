package lib

import (
	"testing"
)

func TestGetThisMonthPeriod(t *testing.T) {
	period := GetPeriodByName("ThisMonth", 1502980112)  // Aug 17, 2017
	if period.From != 1501545600 {
		t.Errorf("Period 'From' time not correct. Expected 1501545600, got %d", period.From)
	}

	if period.To != 1503014399 {
		t.Errorf("Period 'To' time not correct. Expected 1503014399, got %d", period.To)
	}
}

func TestGetLastMonthPeriod(t *testing.T) {
	period := GetPeriodByName("LastMonth", 1502980112)
	if period.From != 1498867200 {
		t.Errorf("Period 'From' time not correct. Expected 1498867200, got %d", period.From)
	}

	if period.To != 1501545599 {
		t.Errorf("Period 'To' time not correct. Expected 1501545599, got %d", period.To)
	}
}

func TestGetTodayPeriod(t *testing.T) {
	period := GetPeriodByName("Today", 1502980112)
	if period.From != 1502928000 {
		t.Errorf("Period 'From' time not correct. Expected 1502928000, got %d", period.From)
	}

	if period.To != 1503014399 {
		t.Errorf("Period 'To' time not correct. Expected 1503014399, got %d", period.To)
	}
}

func TestGetThisYearPeriod(t *testing.T) {
	period := GetPeriodByName("ThisYear", 1502980112)
	if period.From != 1483228800 {
		t.Errorf("Period 'From' time not correct. Expected 1483228800, got %d", period.From)
	}

	if period.To != 1514764799 {
		t.Errorf("Period 'To' time not correct. Expected 1514764799, got %d", period.To)
	}
}

func TestGetLastYearPeriod(t *testing.T) {
	period := GetPeriodByName("LastYear", 1502980112)
	if period.From != 1451606400 {
		t.Errorf("Period 'From' time not correct. Expected 1451606400, got %d", period.From)
	}

	if period.To != 1483228799 {
		t.Errorf("Period 'To' time not correct. Expected 1483228799, got %d", period.To)
	}
}
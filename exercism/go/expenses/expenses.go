package expenses

import (
	"errors"
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var val []Record

	for i, v := range in {
		if predicate(v) {
			val = append(val, in[i])
		}
	}

	fmt.Println(val)

	return val
	panic("Please implement the Filter function")
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {

	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}

	panic("Please implement the ByDaysPeriod function")
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
	panic("Please implement the ByCategory function")
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var val float64 = 0.0
	for _, v := range in {
		if v.Day >= p.From && v.Day <= p.To {
			val += v.Amount
		}
	}
	return val
	panic("Please implement the TotalByPeriod function")
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var val float64 = 0.0

	for _, v := range in {

		if v.Day >= p.From && v.Day <= p.To {
			if v.Category != c && val == 0 {
				return 0, errors.New("unknown category")
			}
			if v.Category == c  {
			val += v.Amount
		}
		}

	}

	return val, nil
	panic("Please implement the CategoryExpenses function")
}

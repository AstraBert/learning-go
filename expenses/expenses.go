package expenses

import "errors"

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
	var records []Record
	for _,record := range in {
		if predicate(record) {
			records = append(records, record)
		}
	}
	return records
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		if r.Day <= p.To && r.Day >= p.From {
			return true
		} else {
			return false
		}
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		if r.Category == c {
			return true
		} else {
			return false
		}
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var total float64
	for _,record := range in {
		if record.Day <= p.To && record.Day >= p.From {
			total+=record.Amount
		}
	}
	return total
}

func isInCategories(c string, l []string) bool {
	var notInCat []string
	for _,r := range l {
		if c!=r {
			notInCat = append(notInCat, c)
		}
	}
	if len(notInCat) == len(l) {
		return false
	} else {
		return true
	}
}

var CategoryNotExistsError = errors.New("category does not exist")

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var categories []string
	for _,r := range in {
		categories = append(categories, r.Category)
	}
	if !isInCategories(c, categories) {
		return 0, CategoryNotExistsError
	} else {
		var total float64
		for _,record := range in {
			if record.Day <= p.To && record.Day >= p.From && record.Category == c {
				total+=record.Amount
			}
		}
		return total, nil
	}	
}

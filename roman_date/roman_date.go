package roman_date

import (
	"fmt"
)

type Month int

const (
	Jan Month = iota + 1
	Feb
	Mar
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
)

func (m Month) Index() int {
	return int(m) - 1
}

func (m Month) String() string {
	return [...]string {
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}[m.Index()]
}

type Date struct {
	Year  int
	Month Month
	Day   int
}

func (d Date) Print() {
	fmt.Print(d.Year, d.Month.String(), d.Day)
}

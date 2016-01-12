package main

import (
	"fmt"
	"math/rand"
	"time"
)

var days map[time.Month]int

func ask(d time.Time) int {
	var i int
	fmt.Scan(&i)

	for i > 6 {
		fmt.Printf("Reduced? ")
		fmt.Scan(&i)
	}

	if time.Weekday(i) != d.Weekday() {
		fmt.Printf("✗ not %s ", time.Weekday(i))
		return ask(d)
	}

	return i
}

func test() {
	month := time.Month(rand.Intn(12))
	date := time.Date(rand.Intn(347)+1753, month, rand.Intn(days[month]+1), 0, 0, 0, 0, time.UTC)
	fmt.Printf("%s: ", date.Format("January 2, 2006"))
	fmt.Printf("✓ %s\n", time.Weekday(ask(date)))
}

func init() {
	days = map[time.Month]int{
		time.January:   31,
		time.February:  29,
		time.March:     31,
		time.April:     30,
		time.May:       31,
		time.June:      30,
		time.July:      31,
		time.August:    31,
		time.September: 30,
		time.October:   31,
		time.November:  30,
		time.December:  31,
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	for {
		test()
	}
}

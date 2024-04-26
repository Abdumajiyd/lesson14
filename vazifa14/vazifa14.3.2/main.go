package main

import (
	"fmt"
	"time"
)

func Subtract(current, other time.Time) (days, hours, minutes, seconds int){
	diff := current.Sub(other)
	days = int(diff.Hours() / 24)
	hours = int(diff.Hours()) % 24
	minutes = int(diff.Minutes()) % 60
	seconds = int(diff.Seconds()) % 60
	return
}

func main() {
	now := time.Now()
	date := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	days, hours, minutes, seconds := Subtract(now, date)
	fmt.Printf("Kunlar: %d, Soatlar: %d, Daqiqalar: %d, Soniyalar: %d\n", days, hours, minutes, seconds)
}

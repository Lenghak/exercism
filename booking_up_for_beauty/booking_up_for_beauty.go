package booking

import (
	"fmt"
	"log"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	parsed, err := time.Parse("January _2, 2006 15:04:05", date)

	if err != nil {
		log.Fatal(err)
	}

	// return bool(t)
	return time.Now().After(parsed)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	parsed, err := time.Parse("Monday, January _2, 2006 15:04:05", date)

	if err != nil {
		log.Fatal(err)
	}

	return parsed.Hour() >= 12 && parsed.Hour() <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	d, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return fmt.Sprint("You have an appointment on ", d.Format("Monday, January 2, 2006, at 15:04."))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}

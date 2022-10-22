package booking

import (
	"time"
	"fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
t, err := time.Parse("1/02/2006 15:04:05", date)
if err != nil {
	panic(err)
}
	fmt.Println("t=", t)
		return t
	panic("Please implement the Schedule function")
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	    t, _ := time.Parse("January 2, 2006 15:04:05", date)
		fmt.Println("HasPassed t",t)
		return t.Unix() < time.Now().Unix()
	panic("Please implement the HasPassed function")
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	    t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
		fmt.Println("IsAfternoon t",t)
		hour := t.Hour()
		fmt.Println("Hour t", hour)
		return hour >=12 && hour < 18
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	    t, err := time.Parse("1/2/2006 15:04:05", date)
		if err != nil {
			fmt.Println("error", date, err)
		}
		return fmt.Sprintf(t.Format("You have an appointment on Monday, January 2, 2006, at 15:04."))
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	s := fmt.Sprintf("%d-09-15 00:00:00 +0000 UTC", time.Now().Year())
	t, _ := time.Parse("2006-01-2 15:04:05 -0700 MST", s)

		return t

	panic("Please implement the AnniversaryDate function")
}

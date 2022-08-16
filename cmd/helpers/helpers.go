package helpers

import (
	"strings"
	"time"
)

func DateFormat(d string) (f string) {
	s := strings.Split(d, "/")
	f = s[2] + "-" + s[1] + "-" + s[0]
	return
}

func DateFormatYear(d string) (f string) {
	s := strings.Split(d, "/")
	f = "20" + s[2] + "-" + s[1] + "-" + s[0]
	return
}

func MonthsCount(startDate time.Time, endDate time.Time) int {
	months := 0
	month := startDate.Month()
	for startDate.Before(endDate) {
		startDate = startDate.Add(time.Hour * 24)
		nextMonth := startDate.Month()
		if nextMonth != month {
			months++
		}
		month = nextMonth
	}
	return months + 1
}

func LongName(n string) (l string) {
	if n == "brad" {
		n = "bradley"
	}
	if n == "nat" {
		n = "natalie"
	}
	if n == "matt" {
		n = "matthew"
	}
	return n
}

func ChangeName(stylist string) (n string) {

	if stylist == "Michelle Railton" {
		n = "Michelle Stevenson"
	} else if stylist == "Jo Mahoney" {
		n = "Jo Birchall"
	} else if stylist == "Laura Crumplin" {
		n = "Laura Hall"
	} else if n == "Bradley Ryan" {
		n = "Brad Ryan"
	} else {
		n = stylist
	}
	return n
}

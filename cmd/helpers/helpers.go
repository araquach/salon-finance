package helpers

import "strings"

// date formatter for 01/02/2003 to 2003-02-01
func DateFormat(d string) (f string) {
	s := strings.Split(d, "/")
	f = s[2] + "-" + s[1] + "-" + s[0]
	return
}

// date formatter for 01/02/03 to 2003-02-01
func DateFormatYear(d string) (f string) {
	s := strings.Split(d, "/")
	f = "20" + s[2] + "-" + s[1] + "-" + s[0]
	return
}

// date formatter for 02-03 to 2003-02-01
func dateFormatDashes(d string) (f string) {
	s := strings.Split(d, "-")
	f = "20" + s[2] + "-" + s[1] + "-" + s[0]
	return
}

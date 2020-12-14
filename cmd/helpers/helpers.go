package helpers

import "strings"

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

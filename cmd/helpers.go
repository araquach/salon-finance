package main

import "strings"

func dateFormat(d string) (f string) {
	s := strings.Split(d, "/")
	f = s[2] + "-" + s[1] + "-" + s[0]
	return
}

func dateFormatYear(d string) (f string) {
	s := strings.Split(d, "/")
	f = "20" + s[2] + "-" + s[1] + "-" + s[0]
	return
}

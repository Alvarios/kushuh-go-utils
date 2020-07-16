package slack

import "regexp"

func EscapeMessage(m string) string {
	re1 := regexp.MustCompile("&")
	re2 := regexp.MustCompile("<")
	re3 := regexp.MustCompile(">")

	o := re1.ReplaceAllString(m, "&amp;")
	o = re2.ReplaceAllString(o, "&lt;")
	o = re3.ReplaceAllString(o, "&gt;")

	return o
}
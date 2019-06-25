package querybuilder

import (
	"fmt"
	"regexp"
	"time"
)

// Time helper for time string setter
func Time(t time.Time) string {
	str := t.Format("2006-01-02T15:04:05Z")

	return str
}

// Number helper for value setter
func Number(i int32) string {
	str := fmt.Sprint(i)

	return str
}

func escapeString(str string) string {
	re := regexp.MustCompile(`(\\)`)
	s := re.ReplaceAllString(str, "\\$1")

	re = regexp.MustCompile(`(')`)
	s = re.ReplaceAllString(s, "\\$1")

	return s
}

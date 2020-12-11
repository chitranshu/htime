package htime

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

var units = map[string]time.Duration{
	"second":  time.Second,
	"seconds": time.Second,
	"minute":  time.Minute,
	"minutes": time.Minute,
	"hour":    time.Hour,
	"hours":   time.Hour,
	"day":     time.Hour * 24,
	"days":    time.Hour * 24,
	"week":    time.Hour * 24 * 7,
	"weeks":   time.Hour * 24 * 7,
	"month":   time.Hour * 24 * 30,
	"months":  time.Hour * 24 * 30,
	"year":    time.Hour * 24 * 365,
	"years":   time.Hour * 24 * 365,
}

// Duration parses the 'text' and returns time.Duration
func Duration(text string) time.Duration {
	duration := time.Duration(0)
	if 0 == len(text) {
		return duration
	}
	text = strings.ToLower(text)

	exp := regexp.MustCompile(`([^a-z\d-]|and)+`)
	text = exp.ReplaceAllString(text, "")

	exp = regexp.MustCompile(`(second|minute|hour|day|week|month|year)s?`)

	for {
		match := exp.FindAllStringIndex(text, 1)
		if 0 == len(match) {
			return duration
		}
		start := match[0][0]
		end := match[0][1]

		value, err := strconv.Atoi(text[:start])
		if nil != err {
			return duration
		}
		unitKey := text[start:end]
		duration += (time.Duration(value) * units[unitKey])
		text = text[end:]
	}
}

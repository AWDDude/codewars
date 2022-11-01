package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(FormatDuration(1))
	fmt.Println(FormatDuration(62))
	fmt.Println(FormatDuration(120))
	fmt.Println(FormatDuration(3600))
	fmt.Println(FormatDuration(3662))
	fmt.Println(FormatDuration(3662000))
}

func FormatDuration(seconds int64) string {
	var a []string
	var s string
	minutes := seconds / 60
	seconds = seconds - (minutes * 60)
	if seconds > 0 {
		s = fmt.Sprintf("%v second", seconds)
		if seconds > 1 {
			s += "s"
		}
		a = append([]string{s}, a...)
	}

	hours := minutes / 60
	minutes = minutes - (hours * 60)
	if minutes > 0 {
		s = fmt.Sprintf("%v minute", minutes)
		if minutes > 1 {
			s += "s"
		}
		if len(a) > 0 {
			a = append([]string{", "}, a...)
		}
		a = append([]string{s}, a...)
	}

	days := hours / 24
	hours = hours - (days * 24)
	if hours > 0 {
		s = fmt.Sprintf("%v hour", hours)
		if hours > 1 {
			s += "s"
		}
		if len(a) > 0 {
			a = append([]string{", "}, a...)
		}
		a = append([]string{s}, a...)
	}

	years := days / 365
	days = days - (years * 365)
	if days > 0 {
		s = fmt.Sprintf("%v day", days)
		if days > 1 {
			s += "s"
		}
		if len(a) > 0 {
			a = append([]string{", "}, a...)
		}
		a = append([]string{s}, a...)
	}

	if years > 0 {
		s = fmt.Sprintf("%v year", years)
		if years > 1 {
			s += "s"
		}
		if len(a) > 0 {
			a = append([]string{", "}, a...)
		}
		a = append([]string{s}, a...)
	}

	if len(a) > 2 {
		a[len(a)-2] = " and "
	}
	return strings.Join(a, "")
}

package main

func HumanReadableTime(seconds int) string {
	var minutes, hours int
	minutes, seconds = divideWithRemainder(seconds, 60)
	hours, minutes = divideWithRemainder(minutes, 60)

	var timeStr string
	if hours < 10 {
		timeStr += "0"
	}
	timeStr += intToString(hours) + ":"

	if minutes < 10 {
		timeStr += "0"
	}
	timeStr += intToString(minutes) + ":"

	if seconds < 10 {
		timeStr += "0"
	}
	timeStr += intToString(seconds)

	return timeStr
}

func divideWithRemainder(dividend, divisor int) (quotient int, remainder int) {
	quotient = dividend / divisor
	remainder = dividend - (quotient * divisor)
	return quotient, remainder
}

func intToString(i int) (str string) {
	if i == 0 {
		return "0"
	}

	var neg bool
	if i < 0 {
		i *= -1
		neg = true
	}

	var c rune
	for i > 0 {
		c = rune('0' + (i - ((i / 10) * 10)))
		str = string(c) + str
		i = i / 10
	}

	if neg {
		str = "-" + str
	}

	return str
}

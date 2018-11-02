package entity

import (
	"fmt"
)

type Time struct {
	year, month, day, hour, minute int
}

func (time Time) isValidTime() bool {
	if time.year < 0 {
		return false
	}
	if time.month < 1 || time.month > 12 {
		return false
	}
	if time.day < 1 {
		return false
	}
	if (time.month == 1 || time.month == 3 || time.month == 5 || time.month == 7 || time.month == 8 || time.month == 10 || time.month == 12) && time.day > 31 {
		return false
	}
	if (time.month == 4 || time.month == 6 || time.month == 9 || time.month == 11) && time.day > 30 {
		return false
	}
	if time.month == 2 {
		if time.year % 4 == 0 && (time.year % 400 == 0 || time.year % 100 != 0) {
			if time.day > 29 {
				return false
			}
		} else {
			if time.day > 28 {
				return false
			}
		}
	}
	if time.hour < 0 || time.hour > 23 {
		return false
	}
	if time.minute < 0 || time.minute > 59 {
		return false
	}
	return true;
}

func stringToTime(str string) Time {
	if len(str) != 16 {
		return Time{0, 0, 0, 0, 0}
	}
	if str[5] != '-' || str[8] != '-' || str[11] != '/' || str[14] != ':' {
		return Time{0, 0, 0, 0, 0}
	}
	return Time{stringToInt(str[0:4]), stringToInt(str[5:7]), stringToInt(str[8:10]), stringToInt(str[11:13]), stringToInt(str[14:])}
}

func stringToInt(str string) int {
	sum := 0
	for i := 0; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			return 0
		} else {
			sum *= 10
			sum += int(str[i] - '0')
		}
	}
	return sum
}

func intToString(num int, size int) string {
	str := ""
	for ; len(str) < size; num /= 10 {
		str = fmt.Sprintf("%d%s", num % 10, str)
	}
	return str
}

func (time Time) toString() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s%s%s", intToString(time.year, 4), "-", intToString(time.month, 2), "-", intToString(time.day, 2), "/", intToString(time.hour, 2), ":", intToString(time.minute, 2))
}

func (a Time) equal(b Time) bool {
	return a.year == b.year && a.month == b.month && a.day == b.day && a.hour == b.hour && a.minute == b.minute
}

func (a Time) moreThan(b Time) bool {
	if a.year > b.year {
		return true
	} else if a.year == b.year {
		if a.month > b.month {
			return true
		} else if a.month == b.month {
			if a.day > b.day {
				return true
			} else if a.day == b.day {
				if a.hour > b.hour {
					return true
				} else if a.hour == b.hour {
					if a.minute > b.minute {
						return true
					} else {
						return false
					}
				} else {
					return false
				}
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}

func (a Time) lessThan(b Time) bool {
	if a.year < b.year {
		return true
	} else if a.year == b.year {
		if a.month < b.month {
			return true
		} else if a.month == b.month {
			if a.day < b.day {
				return true
			} else if a.day == b.day {
				if a.hour < b.hour {
					return true
				} else if a.hour == b.hour {
					if a.minute < b.minute {
						return true
					} else {
						return false
					}
				} else {
					return false
				}
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}
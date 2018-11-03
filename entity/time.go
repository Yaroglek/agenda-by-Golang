package entity

import (
	"fmt"
)

type Time struct {
	Year, Month, Day, Hour, Minute int
}

func (time Time) IsValidTime() bool {
	if time.Year < 0 {
		return false
	}
	if time.Month < 1 || time.Month > 12 {
		return false
	}
	if time.Day < 1 {
		return false
	}
	if (time.Month == 1 || time.Month == 3 || time.Month == 5 || time.Month == 7 || time.Month == 8 || time.Month == 10 || time.Month == 12) && time.Day > 31 {
		return false
	}
	if (time.Month == 4 || time.Month == 6 || time.Month == 9 || time.Month == 11) && time.Day > 30 {
		return false
	}
	if time.Month == 2 {
		if time.Year % 4 == 0 && (time.Year % 400 == 0 || time.Year % 100 != 0) {
			if time.Day > 29 {
				return false
			}
		} else {
			if time.Day > 28 {
				return false
			}
		}
	}
	if time.Hour < 0 || time.Hour > 23 {
		return false
	}
	if time.Minute < 0 || time.Minute > 59 {
		return false
	}
	return true;
}

func StringToTime(str string) Time {
	if len(str) != 16 {
		return Time{0, 0, 0, 0, 0}
	}
	if str[4] != '-' || str[7] != '-' || str[10] != '/' || str[13] != ':' {
		return Time{0, 0, 0, 0, 0}
	}
	return Time{StringToInt(str[0:4]), StringToInt(str[5:7]), StringToInt(str[8:10]), StringToInt(str[11:13]), StringToInt(str[14:])}
}

func StringToInt(str string) int {
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

func IntToString(num int, size int) string {
	str := ""
	for ; len(str) < size; num /= 10 {
		str = fmt.Sprintf("%d%s", num % 10, str)
	}
	return str
}

func (time Time) ToString() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s%s%s", IntToString(time.Year, 4), "-", IntToString(time.Month, 2), "-", IntToString(time.Day, 2), "/", IntToString(time.Hour, 2), ":", IntToString(time.Minute, 2))
}

func (a Time) Equal(b Time) bool {
	return a.Year == b.Year && a.Month == b.Month && a.Day == b.Day && a.Hour == b.Hour && a.Minute == b.Minute
}

func (a Time) MoreThan(b Time) bool {
	if a.Year > b.Year {
		return true
	} else if a.Year == b.Year {
		if a.Month > b.Month {
			return true
		} else if a.Month == b.Month {
			if a.Day > b.Day {
				return true
			} else if a.Day == b.Day {
				if a.Hour > b.Hour {
					return true
				} else if a.Hour == b.Hour {
					if a.Minute > b.Minute {
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

func (a Time) LessThan(b Time) bool {
	if a.Year < b.Year {
		return true
	} else if a.Year == b.Year {
		if a.Month < b.Month {
			return true
		} else if a.Month == b.Month {
			if a.Day < b.Day {
				return true
			} else if a.Day == b.Day {
				if a.Hour < b.Hour {
					return true
				} else if a.Hour == b.Hour {
					if a.Minute < b.Minute {
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
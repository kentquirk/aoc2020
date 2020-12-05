package main

import (
	"regexp"
	"strconv"
)

type validator func(string) bool

// ValidationPartB is used to hold field keys and references to validator functions
type ValidationPartB map[string]validator

func yearValidatorBuilder(min int, max int) validator {
	return func(s string) bool {
		year, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return year >= min && year <= max
	}
}

func heightValidator(s string) bool {
	pat := regexp.MustCompile("(^[[:digit:]]{2,3})(in|cm)$")
	m := pat.FindStringSubmatch(s)
	if m == nil {
		return false
	}
	value, _ := strconv.Atoi(m[1])
	unit := m[2]
	switch unit {
	case "in":
		return value >= 59 && value <= 76
	case "cm":
		return value >= 150 && value <= 193
	default:
		return false
	}
}

func hairValidator(s string) bool {
	pat := regexp.MustCompile("^#[0-9a-f]{6}$")
	return pat.MatchString(s)
}

func eyeValidator(s string) bool {
	pat := regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
	return pat.MatchString(s)
}

func passportIDValidator(s string) bool {
	pat := regexp.MustCompile("^[[:digit:]]{9}$")
	return pat.MatchString(s)
}

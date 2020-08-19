package internal

import (
	"regexp"
	"strings"
)

func IsNumeric(str string) bool {
	numRegex := regexp.MustCompile(`^[0-9]+(\.[0-9]+)?$`)
	str = strings.TrimSpace(str)
	return numRegex.Match([]byte(str))
}

func IsSet(slice []string, index int) bool {
	return len(slice) > index
}

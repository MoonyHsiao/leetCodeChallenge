package models

import (
	"fmt"
	"regexp"
	"strings"
)

func CheckRegexp(s string) {
	// pat := "[_,:.@#--?\";!` ]"
	pat := "[^0-9A-Za-z]"
	re, _ := regexp.Compile(pat)

	s = re.ReplaceAllString(s, "")
	s = strings.ToLower(s)
	fmt.Print(s)
}

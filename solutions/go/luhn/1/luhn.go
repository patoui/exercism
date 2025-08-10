package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	// check length, if less than 2 return false
	if len(id) < 2 {
		return false
	}

	reg := regexp.MustCompile("[^0-9]+")
	preparedId := reg.ReplaceAllString(id, "")

	// ensure after removing non-numeric values, the lengths are the same
	if len(id) != len(preparedId) {
		return false
	}

	total := 0
	count := 1
	for i := len(preparedId) - 1; i >= 0; i-- {
		itemInt, _ := strconv.Atoi(string(preparedId[i]))
		if count%2 == 0 {
			itemInt *= 2
			if itemInt > 9 {
				itemInt -= 9
			}
		}
		total += itemInt
		count++
	}

	// ensure id sum is divisible by 10
	return total%10 == 0
}

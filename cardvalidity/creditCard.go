package cardvalidity

import (
	"strings"
)

var beginningChars map[string]bool
var validChars map[string]bool

func InitializeLists() {
	beginningChars = map[string]bool{"4": false, "5": false, "6": false}
	validChars = map[string]bool{"0": false, "1": false, "2": false, "3": false, "4": false, "5": false, "6": false, "7": false, "8": false, "9": false}
}

// this string was called validCard in the hackerRank system
func ValidateCard(card string) string {

	if len(beginningChars) == 0 {
		InitializeLists()
	}

	parts := strings.Split(card, "-")

	if len(parts) > 4 {
		return "Invalid"
	}

	if len(parts) > 1 {
		for _, part := range parts {
			if len(part) != 4 {
				return "Invalid"
			}
		}
	}

	myCard := strings.Join(parts, "")

	if _, ok := beginningChars[string(myCard[0])]; !ok {
		return "Invalid"
	}

	if len(myCard) != 16 {
		return "Invalid"
	}
	prev, numTimes := "", 0

	for _, chrune := range myCard {
		ch := string(chrune)
		if _, ok := validChars[ch]; !ok {
			return "Invalid"
		}
		if ch == prev {
			numTimes++
		} else {
			numTimes = 0
		}
		if numTimes >= 3 {
			return "Invalid"
		}
		prev = ch
	}
	return "Valid"
}

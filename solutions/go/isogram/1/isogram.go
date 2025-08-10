package isogram

func IsIsogram(word string) bool {
	letter_counts := make(map[rune]uint8)
	for _, l := range word {
		if l != ' ' && l != '-' {
			if l >= 97 && l <= 122 {
				// lowercase letter, substract 32 to make it uppercase
				l -= 32
			}

			letter_counts[l]++
			if letter_counts[l] > 1 {
				return false
			}
		}
	}
	return true
}

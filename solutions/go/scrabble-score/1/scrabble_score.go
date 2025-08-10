package scrabble

func Score(word string) int {
	score := 0

	for _, l := range word {
		score += letterValue(l)
	}

	return score
}

func letterValue(letter rune) int {
	if letter >= 97 && letter <= 122 {
		// lowercase letter, substract 32 to make it uppercase
		letter -= 32
	}

	if letter >= 65 && letter <= 90 {
		switch letter {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			return 1
		case 'D', 'G':
			return 2
		case 'B', 'C', 'M', 'P':
			return 3
		case 'F', 'H', 'V', 'W', 'Y':
			return 4
		case 'K':
			return 5
		case 'J', 'X':
			return 8
		case 'Q', 'Z':
			return 10
		}
	}

	return 0
}

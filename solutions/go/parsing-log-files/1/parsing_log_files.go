package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*`)
	return re.Match([]byte(text))
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[(~\*\-=)]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password"`)
	passwordCount := 0

	for _, l := range lines {
		if len(re.FindAllStringIndex(l, 1)) > 0 {
			passwordCount++
		}
	}

	return passwordCount
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`(end-of-line\d*)`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+([a-zA-Z0-9]+)`)
	parsedLines := []string{}

	for _, l := range lines {
		mu := re.FindStringSubmatch(l)
		if len(mu) > 1 {
			parsedLines = append(parsedLines, fmt.Sprintf("[USR] %s %s", mu[1], l))
		} else {
			parsedLines = append(parsedLines, l)
		}
	}

	return parsedLines
}

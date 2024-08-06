package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

var _LABELS = []string{
	`\[TRC\]`,
	`\[DBG\]`,
	`\[INF\]`,
	`\[WRN\]`,
	`\[ERR\]`,
	`\[FTL\]`,
}

func IsValidLine(text string) bool {
	rex, _ := regexp.Compile(fmt.Sprintf(`^(%s)`, strings.Join(_LABELS, "|")))
	return rex.MatchString(text)
}

func SplitLogLine(text string) []string {
	rex, _ := regexp.Compile(`<(\*|~|=|\-|)*>`)
	return rex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	rex, _ := regexp.Compile(`(?i)\".*(password)+.*\"`)

	count := 0

	for _, line := range lines {
		if rex.MatchString(line) {
			count++
		}
	}

	return count

}

func RemoveEndOfLineText(text string) string {
	rex, _ := regexp.Compile(`end-of-line\d*`)
	return rex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	rex, _ := regexp.Compile(`User\s+(\w+)`)

	for index, line := range lines {
		if rex.MatchString(line) {
			submatches := rex.FindStringSubmatch(line)
			lines[index] = fmt.Sprintf("[USR] %s %s", submatches[1], line)
		}
	}

	return lines
}

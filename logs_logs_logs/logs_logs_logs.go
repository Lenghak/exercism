package logs

import (
	"strings"
	"unicode/utf8"
)

var _LOGS = map[rune]string{
	'❗': "recommendation",
	'🔍': "search",
	'☀': "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {

	var result rune = '_'

	for key := range _LOGS {
		if strings.ContainsRune(log, key) {
			if result == '_' || strings.IndexRune(log, result) > strings.IndexRune(log, key) {
				result = key
			}
		}
	}

	if result != '_' {
		return _LOGS[result]
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}

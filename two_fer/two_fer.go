package twofer

import "fmt"

func ShareWith(name string) string {
	out := "you"

	if name != "" {
		out = name
	}

	return fmt.Sprintf("One for %s, one for me.", out)
}

package twofer

import "fmt"

// ShareWith returns a string following the Two Fer rule.
func ShareWith(name string) string {
	var talkingTo string

	if name == "" {
		talkingTo = "you"
	} else {
		talkingTo = name
	}

	return fmt.Sprintf("One for %s, one for me.", talkingTo)
}
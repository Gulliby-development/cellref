package main

import "strconv"

// countText formats a count with a singular noun only when the count is one.
func countText(n int, noun string) string {
	if n == 1 {
		return strconv.Itoa(n) + " " + noun
	}
	return strconv.Itoa(n) + " " + noun + "s"
}

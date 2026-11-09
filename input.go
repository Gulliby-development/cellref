package main

import (
	"strconv"
	"strings"
)

// readItems parses meaningful input lines after checking the physical line limit.
func readItems(text string) ([]item, *failure) {
	if text == "" {
		return nil, nil
	}
	lines := strings.Split(text, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	if len(lines) > maxLines {
		return nil, fail(1, "input is too long (max 40 lines)")
	}

	items := make([]item, 0, len(lines))
	names := make(map[string]bool)
	for index, line := range lines {
		line = strings.TrimSuffix(line, "\r")
		words := inputWords(line)
		if len(words) == 0 || strings.HasPrefix(words[0], "#") {
			continue
		}
		lineNumber := index + 1
		if len(words) != 2 {
			return nil, fail(1, lineError(lineNumber, "expected NAME RANGE"))
		}
		name := words[0]
		if !validName(name) {
			return nil, fail(1, lineError(lineNumber, "bad name "+name))
		}
		r, ok := parseRange(words[1])
		if !ok {
			return nil, fail(1, lineError(lineNumber, "bad range "+words[1]))
		}
		if r.c1 > r.c2 || r.r1 > r.r2 {
			return nil, fail(1, lineError(lineNumber, "range runs backwards"))
		}
		if names[name] {
			return nil, fail(1, lineError(lineNumber, "duplicate name "+name))
		}
		names[name] = true
		r.name = name
		r.line = lineNumber
		items = append(items, r)
	}
	return items, nil
}

func inputWords(line string) []string {
	fields := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' || r == '\t' })
	return fields
}

func validName(name string) bool {
	if len(name) == 0 || len(name) > maxName || name[0] < 'a' || name[0] > 'z' {
		return false
	}
	for i := 1; i < len(name); i++ {
		c := name[i]
		if (c < 'a' || c > 'z') && (c < '0' || c > '9') && c != '-' {
			return false
		}
	}
	return true
}

func lineError(line int, reason string) string {
	return "line " + strconv.Itoa(line) + ": " + reason
}

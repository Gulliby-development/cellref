package main

import "strings"

// parseNumber accepts canonical signed decimal values within the shift limit.
func parseNumber(value string) (int, bool) {
	if value == "" {
		return 0, false
	}

	negative := value[0] == '-'
	digits := value
	if negative {
		digits = value[1:]
	}
	if digits == "" || (len(digits) > 1 && digits[0] == '0') {
		return 0, false
	}
	if negative && digits == "0" {
		return 0, false
	}

	n := 0
	for i := 0; i < len(digits); i++ {
		if digits[i] < '0' || digits[i] > '9' {
			return 0, false
		}
		n = n*10 + int(digits[i]-'0')
		if n > maxShift {
			return 0, false
		}
	}
	if negative {
		n = -n
	}
	return n, true
}

// parseOptions validates words following a command from left to right.
func parseOptions(args []string) (options, *failure) {
	var opts options
	seenRows, seenCols := false, false

	for i := 0; i < len(args); i++ {
		word := args[i]
		if word == "" || word == "-" || !strings.HasPrefix(word, "-") {
			return options{}, fail(2, "unexpected argument "+word)
		}

		var seen *bool
		switch word {
		case "--rows":
			seen = &seenRows
		case "--cols":
			seen = &seenCols
		default:
			return options{}, fail(2, "unknown option "+word)
		}
		if *seen {
			return options{}, fail(2, "given twice: "+word)
		}
		*seen = true
		if i+1 >= len(args) {
			return options{}, fail(2, "missing value for "+word)
		}
		i++
		value, ok := parseNumber(args[i])
		if !ok {
			return options{}, fail(2, "bad value for "+word)
		}
		if word == "--rows" {
			opts.rows = value
		} else {
			opts.cols = value
		}
	}
	return opts, nil
}

func readsStdin(args []string) bool {
	if len(args) == 0 || findCommand(args[0]) == nil {
		return false
	}
	_, failure := parseOptions(args[1:])
	return failure == nil
}

func run(args []string, text string) (int, string, string) {
	if len(args) == 0 {
		return 2, "", usage
	}

	switch args[0] {
	case "help", "-h", "--help":
		return 0, usage, ""
	case "--version":
		return 0, version + "\n", ""
	}

	cmd := findCommand(args[0])
	if cmd == nil {
		return 2, "", "cellref: unknown command " + args[0] + "\n"
	}
	opts, failure := parseOptions(args[1:])
	if failure != nil {
		return failure.status, "", "cellref: " + failure.msg + "\n"
	}

	items, failure := readItems(text)
	if failure != nil {
		return failure.status, "", "cellref: " + failure.msg + "\n"
	}
	out, failure := cmd.run(opts, items)
	if failure != nil {
		return failure.status, "", "cellref: " + failure.msg + "\n"
	}
	return 0, out, ""
}

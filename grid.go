package main

import (
	"sort"
	"strconv"
	"strings"
)

// parseCell reads a canonical spreadsheet cell and returns its 1-based column
// and row.
func parseCell(value string) (int, int, bool) {
	if len(value) < 2 {
		return 0, 0, false
	}

	i := 0
	column := 0
	for i < len(value) && value[i] >= 'A' && value[i] <= 'Z' {
		if i == 2 {
			return 0, 0, false
		}
		column = column*26 + int(value[i]-'A'+1)
		i++
	}
	if i == 0 || column > maxCol || i == len(value) {
		return 0, 0, false
	}

	digits := value[i:]
	if digits[0] == '0' || len(digits) > 4 {
		return 0, 0, false
	}
	for j := 0; j < len(digits); j++ {
		if digits[j] < '0' || digits[j] > '9' {
			return 0, 0, false
		}
	}
	row, err := strconv.Atoi(digits)
	if err != nil || row < 1 || row > maxRow {
		return 0, 0, false
	}
	return column, row, true
}

// parseRange parses range syntax without deciding whether the corners run
// backwards; input validation reports that as a separate error.
func parseRange(value string) (item, bool) {
	if strings.Count(value, ":") > 1 {
		return item{}, false
	}
	parts := strings.Split(value, ":")
	c1, r1, ok := parseCell(parts[0])
	if !ok {
		return item{}, false
	}
	c2, r2 := c1, r1
	if len(parts) == 2 {
		c2, r2, ok = parseCell(parts[1])
		if !ok {
			return item{}, false
		}
	}
	return item{c1: c1, r1: r1, c2: c2, r2: r2}, true
}

// columnName returns the canonical letters for a 1-based column number.
func columnName(column int) string {
	var letters [2]byte
	i := len(letters)
	for column > 0 {
		column--
		i--
		letters[i] = byte('A' + column%26)
		column /= 26
	}
	return string(letters[i:])
}

func formatCell(column, row int) string {
	return columnName(column) + strconv.Itoa(row)
}

// formatRange writes a singleton as one cell and other ranges as two corners.
func formatRange(r item) string {
	start := formatCell(r.c1, r.r1)
	if r.c1 == r.c2 && r.r1 == r.r2 {
		return start
	}
	return start + ":" + formatCell(r.c2, r.r2)
}

func rangeArea(r item) int {
	return (r.r2 - r.r1 + 1) * (r.c2 - r.c1 + 1)
}

// intersectRanges returns the inclusive intersection of two ranges, if any.
func intersectRanges(a, b item) (item, bool) {
	r := item{
		c1: max(a.c1, b.c1),
		r1: max(a.r1, b.r1),
		c2: min(a.c2, b.c2),
		r2: min(a.r2, b.r2),
	}
	if r.c1 > r.c2 || r.r1 > r.r2 {
		return item{}, false
	}
	return r, true
}

// boundsOf returns the smallest range enclosing all items.
func boundsOf(items []item) (item, bool) {
	if len(items) == 0 {
		return item{}, false
	}
	bounds := items[0]
	for _, r := range items[1:] {
		bounds.c1 = min(bounds.c1, r.c1)
		bounds.r1 = min(bounds.r1, r.r1)
		bounds.c2 = max(bounds.c2, r.c2)
		bounds.r2 = max(bounds.r2, r.r2)
	}
	return bounds, true
}

type rowInterval struct {
	start, end int
}

// coveredCells counts the distinct cells using merged row intervals for each
// column, avoiding allocation proportional to the sheet's area.
func coveredCells(items []item) int {
	total := 0
	for column := 1; column <= maxCol; column++ {
		intervals := make([]rowInterval, 0, len(items))
		for _, r := range items {
			if r.c1 <= column && column <= r.c2 {
				intervals = append(intervals, rowInterval{start: r.r1, end: r.r2})
			}
		}
		if len(intervals) == 0 {
			continue
		}
		sort.Slice(intervals, func(i, j int) bool {
			if intervals[i].start == intervals[j].start {
				return intervals[i].end < intervals[j].end
			}
			return intervals[i].start < intervals[j].start
		})
		start, end := intervals[0].start, intervals[0].end
		for _, interval := range intervals[1:] {
			if interval.start > end+1 {
				total += end - start + 1
				start, end = interval.start, interval.end
			} else if interval.end > end {
				end = interval.end
			}
		}
		total += end - start + 1
	}
	return total
}

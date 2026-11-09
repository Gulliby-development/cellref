package main

import (
	"strconv"
	"strings"
)

func cmdCheck(_ options, items []item) (string, *failure) {
	return "ok " + countText(len(items), "range") + "\n", nil
}

func cmdSize(_ options, items []item) (string, *failure) {
	var out strings.Builder
	for _, r := range items {
		rows := r.r2 - r.r1 + 1
		cols := r.c2 - r.c1 + 1
		area := rangeArea(r)
		out.WriteString(r.name)
		out.WriteByte(' ')
		out.WriteString(formatRange(r))
		out.WriteByte(' ')
		out.WriteString(strconv.Itoa(rows))
		out.WriteByte('x')
		out.WriteString(strconv.Itoa(cols))
		out.WriteByte(' ')
		out.WriteString(countText(area, "cell"))
		out.WriteByte('\n')
	}
	out.WriteString("total ")
	out.WriteString(countText(coveredCells(items), "cell"))
	out.WriteByte('\n')
	return out.String(), nil
}

func cmdUnion(_ options, items []item) (string, *failure) {
	bounds, ok := boundsOf(items)
	if !ok {
		return "bounds none\n", nil
	}
	var out strings.Builder
	out.WriteString("bounds ")
	out.WriteString(formatRange(bounds))
	out.WriteByte('\n')
	out.WriteString("covered ")
	out.WriteString(strconv.Itoa(coveredCells(items)))
	out.WriteString(" of ")
	out.WriteString(strconv.Itoa(rangeArea(bounds)))
	out.WriteString(" cells\n")
	return out.String(), nil
}

func cmdOverlap(_ options, items []item) (string, *failure) {
	var out strings.Builder
	pairs := 0
	for i := 0; i < len(items); i++ {
		for j := i + 1; j < len(items); j++ {
			intersection, ok := intersectRanges(items[i], items[j])
			if !ok {
				continue
			}
			pairs++
			out.WriteString(items[i].name)
			out.WriteByte(' ')
			out.WriteString(items[j].name)
			out.WriteByte(' ')
			out.WriteString(formatRange(intersection))
			out.WriteByte(' ')
			out.WriteString(countText(rangeArea(intersection), "cell"))
			out.WriteByte('\n')
		}
	}
	out.WriteString("overlaps ")
	out.WriteString(strconv.Itoa(pairs))
	out.WriteByte('\n')
	return out.String(), nil
}

func cmdShift(opts options, items []item) (string, *failure) {
	shifted := make([]item, len(items))
	for i, r := range items {
		moved := r
		moved.c1 += opts.cols
		moved.c2 += opts.cols
		moved.r1 += opts.rows
		moved.r2 += opts.rows
		if moved.c1 < 1 || moved.c2 > maxCol || moved.r1 < 1 || moved.r2 > maxRow {
			return "", fail(1, "line "+strconv.Itoa(r.line)+": "+r.name+" moves outside the sheet")
		}
		shifted[i] = moved
	}

	var out strings.Builder
	for i, r := range items {
		out.WriteString(r.name)
		out.WriteByte(' ')
		out.WriteString(formatRange(r))
		out.WriteByte(' ')
		out.WriteString(formatRange(shifted[i]))
		out.WriteByte('\n')
	}
	out.WriteString("shifted ")
	out.WriteString(countText(len(items), "range"))
	out.WriteByte('\n')
	return out.String(), nil
}

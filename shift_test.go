package main

import "testing"

var rowsShift = []row{
	{[]string{"shift", "--rows", "1", "--cols", "1"}, "a A1:C3\nb B2:D4\nc F6\n", 0, "a A1:C3 B2:D4\nb B2:D4 C3:E5\nc F6 G7\nshifted 3 ranges\n", ""},
	{[]string{"shift"}, "a A1:C3\nb B2:D4\nc F6\n", 0, "a A1:C3 A1:C3\nb B2:D4 B2:D4\nc F6 F6\nshifted 3 ranges\n", ""},
	{[]string{"shift", "--rows", "-1"}, "a A2:B3\n", 0, "a A2:B3 A1:B2\nshifted 1 range\n", ""},
	{[]string{"shift", "--cols", "-1"}, "a B1\nb C2:D3\n", 0, "a B1 A1\nb C2:D3 B2:C3\nshifted 2 ranges\n", ""},
	{[]string{"shift", "--rows", "-1"}, "a A1:C3\nb B2:D4\nc F6\n", 1, "", "cellref: line 1: a moves outside the sheet\n"},
	{[]string{"shift", "--cols", "-1"}, "a A1:C3\nb B2:D4\nc F6\n", 1, "", "cellref: line 1: a moves outside the sheet\n"},
	{[]string{"shift", "--rows", "3", "--cols", "27"}, "a Z1:AA2\n", 0, "a Z1:AA2 BA4:BB5\nshifted 1 range\n", ""},
	{[]string{"shift", "--cols", "26"}, "a A1\n", 0, "a A1 AA1\nshifted 1 range\n", ""},
	{[]string{"shift", "--cols", "702"}, "a A1\n", 1, "", "cellref: line 1: a moves outside the sheet\n"},
	{[]string{"shift", "--cols", "701"}, "a A1\n", 0, "a A1 ZZ1\nshifted 1 range\n", ""},
	{[]string{"shift", "--cols", "1"}, "a ZZ1\n", 1, "", "cellref: line 1: a moves outside the sheet\n"},
	{[]string{"shift", "--rows", "1"}, "a A9999\n", 1, "", "cellref: line 1: a moves outside the sheet\n"},
	{[]string{"shift", "--rows", "-9998"}, "a A9999\n", 0, "a A9999 A1\nshifted 1 range\n", ""},
	{[]string{"shift", "--rows", "-9999"}, "a A9999\n", 1, "", "cellref: line 1: a moves outside the sheet\n"},
	{[]string{"shift", "--rows", "9998"}, "a A1:B2\n", 1, "", "cellref: line 1: a moves outside the sheet\n"},
	{[]string{"shift", "--rows", "9997"}, "a A1:B2\n", 0, "a A1:B2 A9998:B9999\nshifted 1 range\n", ""},
	{[]string{"shift", "--rows", "1", "--cols", "1"}, "a B2\nb ZZ9999\n", 1, "", "cellref: line 2: b moves outside the sheet\n"},
	{[]string{"shift", "--rows", "-1"}, "ok A5\nbad A1\nlate A1\n", 1, "", "cellref: line 2: bad moves outside the sheet\n"},
	{[]string{"shift", "--cols", "1"}, "# c\na A1\n\nb ZZ1\n", 1, "", "cellref: line 4: b moves outside the sheet\n"},
	{[]string{"shift", "--rows", "2", "--cols", "-1"}, "p B1:C2\n", 0, "p B1:C2 A3:B4\nshifted 1 range\n", ""},
	{[]string{"shift", "--cols", "-26"}, "a AA1:AB2\n", 0, "a AA1:AB2 A1:B2\nshifted 1 range\n", ""},
	{[]string{"shift", "--rows", "1"}, "", 0, "shifted 0 ranges\n", ""},
}

func TestRowsShift(t *testing.T) { checkRows(t, rowsShift) }

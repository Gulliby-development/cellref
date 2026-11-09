package main

import "testing"

var rowsInput = []row{
	{[]string{"check"}, "", 0, "ok 0 ranges\n", ""},
	{[]string{"check"}, "a A1\n", 0, "ok 1 range\n", ""},
	{[]string{"check"}, "a A1:C3\nb B2:D4\nc F6\n", 0, "ok 3 ranges\n", ""},
	{[]string{"check"}, "# note\n\n  \na A1\n\t\n#b B2\n", 0, "ok 1 range\n", ""},
	{[]string{"check"}, "a A1\r\nb B2\r\n", 0, "ok 2 ranges\n", ""},
	{[]string{"check"}, "a A1", 0, "ok 1 range\n", ""},
	{[]string{"check"}, "\ta   A1:B2 \t\n", 0, "ok 1 range\n", ""},
	{[]string{"check"}, "a A1 # note\n", 1, "", "cellref: line 1: expected NAME RANGE\n"},
	{[]string{"check"}, "a\n", 1, "", "cellref: line 1: expected NAME RANGE\n"},
	{[]string{"check"}, "a A1 B2\n", 1, "", "cellref: line 1: expected NAME RANGE\n"},
	{[]string{"check"}, "A a1\n", 1, "", "cellref: line 1: bad name A\n"},
	{[]string{"check"}, "Ab A1\n", 1, "", "cellref: line 1: bad name Ab\n"},
	{[]string{"check"}, "1a A1\n", 1, "", "cellref: line 1: bad name 1a\n"},
	{[]string{"check"}, "a_b A1\n", 1, "", "cellref: line 1: bad name a_b\n"},
	{[]string{"check"}, "-a A1\n", 1, "", "cellref: line 1: bad name -a\n"},
	{[]string{"check"}, "a-b-2 A1\n", 0, "ok 1 range\n", ""},
}

func TestRowsInput(t *testing.T) { checkRows(t, rowsInput) }

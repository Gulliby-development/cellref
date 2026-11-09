package main

import "testing"

var rowsOptions = []row{
	{[]string{"check", "x"}, "", 2, "", "cellref: unexpected argument x\n"},
	{[]string{"check", ""}, "", 2, "", "cellref: unexpected argument \n"},
	{[]string{"size", "-"}, "", 2, "", "cellref: unexpected argument -\n"},
	{[]string{"check", "--nope"}, "", 2, "", "cellref: unknown option --nope\n"},
	{[]string{"shift", "--rows"}, "", 2, "", "cellref: missing value for --rows\n"},
	{[]string{"shift", "--cols"}, "", 2, "", "cellref: missing value for --cols\n"},
	{[]string{"shift", "--rows", "x"}, "", 2, "", "cellref: bad value for --rows\n"},
	{[]string{"shift", "--rows", ""}, "", 2, "", "cellref: bad value for --rows\n"},
	{[]string{"shift", "--rows", "+1"}, "", 2, "", "cellref: bad value for --rows\n"},
	{[]string{"shift", "--rows", "-0"}, "", 2, "", "cellref: bad value for --rows\n"},
	{[]string{"shift", "--rows", "00"}, "", 2, "", "cellref: bad value for --rows\n"},
	{[]string{"shift", "--rows", "007"}, "", 2, "", "cellref: bad value for --rows\n"},
	{[]string{"shift", "--rows", "1.5"}, "", 2, "", "cellref: bad value for --rows\n"},
	{[]string{"shift", "--rows", "10000"}, "", 2, "", "cellref: bad value for --rows\n"},
	{[]string{"shift", "--rows", "-10000"}, "", 2, "", "cellref: bad value for --rows\n"},
	{[]string{"shift", "--cols", "9999"}, "only C3\n", 1, "", "cellref: line 1: only moves outside the sheet\n"},
	{[]string{"shift", "--rows", "-9999"}, "", 0, "shifted 0 ranges\n", ""},
	{[]string{"shift", "--rows", "1", "--rows", "2"}, "", 2, "", "cellref: given twice: --rows\n"},
	{[]string{"shift", "--rows", "x", "--rows", "2"}, "", 2, "", "cellref: bad value for --rows\n"},
	{[]string{"shift", "--rows", "--cols"}, "", 2, "", "cellref: bad value for --rows\n"},
	{[]string{"shift", "--rows", "-1", "--cols", "-1"}, "a A1:C3\nb B2:D4\nc F6\n", 1, "", "cellref: line 1: a moves outside the sheet\n"},
	{[]string{"shift", "--cols", "1", "extra"}, "", 2, "", "cellref: unexpected argument extra\n"},
	{[]string{"check", "--rows", "3"}, "a A1:C3\nb B2:D4\nc F6\n", 0, "ok 3 ranges\n", ""},
	{[]string{"size", "--cols", "-4", "--rows", "7"}, "a A1:C3\nb B2:D4\nc F6\n", 0, "a A1:C3 3x3 9 cells\nb B2:D4 3x3 9 cells\nc F6 1x1 1 cell\ntotal 15 cells\n", ""},
	{[]string{"shift", "--rows=1"}, "", 2, "", "cellref: unknown option --rows=1\n"},
	{[]string{"shift", "--bogus", "--rows"}, "", 2, "", "cellref: unknown option --bogus\n"},
}

func TestRowsOptions(t *testing.T) { checkRows(t, rowsOptions) }

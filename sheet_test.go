package main

import "testing"

var rowsSheet = []row{
	{[]string{"check"}, "abcdefghijkl A1\n", 0, "ok 1 range\n", ""},
	{[]string{"check"}, "abcdefghijklm A1\n", 1, "", "cellref: line 1: bad name abcdefghijklm\n"},
	{[]string{"check"}, "a ZZ9999\n", 0, "ok 1 range\n", ""},
	{[]string{"check"}, "a AAA1\n", 1, "", "cellref: line 1: bad range AAA1\n"},
	{[]string{"check"}, "a A10000\n", 1, "", "cellref: line 1: bad range A10000\n"},
	{[]string{"check"}, "a ZZ9999:ZZ9999\n", 0, "ok 1 range\n", ""},
	{[]string{"check"}, "a A9999:ZZ9999\n", 0, "ok 1 range\n", ""},
	{[]string{"shift", "--rows", "9999"}, "a A1\n", 1, "", "cellref: line 1: a moves outside the sheet\n"},
	{[]string{"shift", "--rows", "10000"}, "a A1\n", 2, "", "cellref: bad value for --rows\n"},
	{[]string{"shift", "--cols", "9999"}, "a A1\n", 1, "", "cellref: line 1: a moves outside the sheet\n"},
	{[]string{"shift", "--cols", "-9999"}, "a ZZ1\n", 1, "", "cellref: line 1: a moves outside the sheet\n"},
	{[]string{"union"}, "a A1:ZZ9999\n", 0, "bounds A1:ZZ9999\ncovered 7019298 of 7019298 cells\n", ""},
	{[]string{"size"}, "a A1:ZZ9999\n", 0, "a A1:ZZ9999 9999x702 7019298 cells\ntotal 7019298 cells\n", ""},
	{[]string{"overlap"}, "a A1:ZZ9999\nb ZZ9999\n", 0, "a b ZZ9999 1 cell\noverlaps 1\n", ""},
}

func TestRowsSheet(t *testing.T) { checkRows(t, rowsSheet) }

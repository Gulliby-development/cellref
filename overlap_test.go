package main

import "testing"

var rowsOverlap = []row{
	{[]string{"overlap"}, "a A1:C3\nb B2:D4\nc F6\n", 0, "a b B2:C3 4 cells\noverlaps 1\n", ""},
	{[]string{"overlap"}, "", 0, "overlaps 0\n", ""},
	{[]string{"overlap"}, "a A1:B2\nb C1:D2\nc A3:B4\n", 0, "overlaps 0\n", ""},
	{[]string{"overlap"}, "big A1:J10\nmid C3:E5\ndot D4\n", 0, "big mid C3:E5 9 cells\nbig dot D4 1 cell\nmid dot D4 1 cell\noverlaps 3\n", ""},
	{[]string{"overlap"}, "a B2:C3\nb B2:C3\nc B2:C3\n", 0, "a b B2:C3 4 cells\na c B2:C3 4 cells\nb c B2:C3 4 cells\noverlaps 3\n", ""},
	{[]string{"overlap"}, "row A5:Z5\ncol C1:C20\ncell C5\n", 0, "row col C5 1 cell\nrow cell C5 1 cell\ncol cell C5 1 cell\noverlaps 3\n", ""},
	{[]string{"overlap"}, "a A1\nb ZZ9999\n", 0, "overlaps 0\n", ""},
	{[]string{"overlap"}, "only C3\n", 0, "overlaps 0\n", ""},
	{[]string{"overlap"}, "a A1:E5\nb C3:G7\nc E5:I9\nd A5:I5\ne E1:E9\n", 0, "a b C3:E5 9 cells\na c E5 1 cell\na d A5:E5 5 cells\na e E1:E5 5 cells\nb c E5:G7 9 cells\nb d C5:G5 5 cells\nb e E3:E7 5 cells\nc d E5:I5 5 cells\nc e E5:E9 5 cells\nd e E5 1 cell\noverlaps 10\n", ""},
	{[]string{"overlap"}, "a A1:C3\nb C3:E5\n", 0, "a b C3 1 cell\noverlaps 1\n", ""},
	{[]string{"overlap"}, "a A1:C3\nb D1:F3\n", 0, "overlaps 0\n", ""},
	{[]string{"overlap"}, "a A1:C3\nb A4:C6\n", 0, "overlaps 0\n", ""},
	{[]string{"overlap"}, "a A1:C3\nb B3:B5\nc B1\n", 0, "a b B3 1 cell\na c B1 1 cell\noverlaps 2\n", ""},
	{[]string{"overlap"}, "z A1:B2\ny A1:B2\nx B2\n", 0, "z y A1:B2 4 cells\nz x B2 1 cell\ny x B2 1 cell\noverlaps 3\n", ""},
	{[]string{"overlap"}, "a ZZ9999\nb ZZ9999\n", 0, "a b ZZ9999 1 cell\noverlaps 1\n", ""},
	{[]string{"overlap"}, "a A1:ZZ1\nb ZZ1:ZZ9999\n", 0, "a b ZZ1 1 cell\noverlaps 1\n", ""},
	{[]string{"overlap"}, "a C3\nb C3\nc C3\nd C3\n", 0, "a b C3 1 cell\na c C3 1 cell\na d C3 1 cell\nb c C3 1 cell\nb d C3 1 cell\nc d C3 1 cell\noverlaps 6\n", ""},
}

func TestRowsOverlap(t *testing.T) { checkRows(t, rowsOverlap) }

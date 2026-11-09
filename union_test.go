package main

import "testing"

var rowsUnion = []row{
	{[]string{"union"}, "a A1:C3\nb B2:D4\nc F6\n", 0, "bounds A1:F6\ncovered 15 of 36 cells\n", ""},
	{[]string{"union"}, "", 0, "bounds none\n", ""},
	{[]string{"union"}, "a A1:B2\nb C1:D2\nc A3:B4\n", 0, "bounds A1:D4\ncovered 12 of 16 cells\n", ""},
	{[]string{"union"}, "big A1:J10\nmid C3:E5\ndot D4\n", 0, "bounds A1:J10\ncovered 100 of 100 cells\n", ""},
	{[]string{"union"}, "a B2:C3\nb B2:C3\nc B2:C3\n", 0, "bounds B2:C3\ncovered 4 of 4 cells\n", ""},
	{[]string{"union"}, "row A5:Z5\ncol C1:C20\ncell C5\n", 0, "bounds A1:Z20\ncovered 45 of 520 cells\n", ""},
	{[]string{"union"}, "a A1\nb ZZ9999\n", 0, "bounds A1:ZZ9999\ncovered 2 of 7019298 cells\n", ""},
	{[]string{"union"}, "only C3\n", 0, "bounds C3\ncovered 1 of 1 cells\n", ""},
	{[]string{"union"}, "tl A1\nbr ZZ9999\ntr ZZ1\nbl A9999\n", 0, "bounds A1:ZZ9999\ncovered 4 of 7019298 cells\n", ""},
	{[]string{"union"}, "w A1:ZZ1\n", 0, "bounds A1:ZZ1\ncovered 702 of 702 cells\n", ""},
	{[]string{"union"}, "a A1:E5\nb C3:G7\nc E5:I9\nd A5:I5\ne E1:E9\n", 0, "bounds A1:I9\ncovered 57 of 81 cells\n", ""},
	{[]string{"union"}, "a Z1\nb AA1\n", 0, "bounds Z1:AA1\ncovered 2 of 2 cells\n", ""},
	{[]string{"union"}, "a AZ5:BA6\n", 0, "bounds AZ5:BA6\ncovered 4 of 4 cells\n", ""},
	{[]string{"union"}, "a B2:B2\nb B2\n", 0, "bounds B2\ncovered 1 of 1 cells\n", ""},
	{[]string{"union"}, "# none\n\n", 0, "bounds none\n", ""},
	{[]string{"union", "--cols", "3"}, "a A1:C3\nb B2:D4\nc F6\n", 0, "bounds A1:F6\ncovered 15 of 36 cells\n", ""},
}

func TestRowsUnion(t *testing.T) { checkRows(t, rowsUnion) }

package main

import "testing"

var rowsSize = []row{
	{[]string{"size"}, "a A1:C3\nb B2:D4\nc F6\n", 0, "a A1:C3 3x3 9 cells\nb B2:D4 3x3 9 cells\nc F6 1x1 1 cell\ntotal 15 cells\n", ""},
	{[]string{"size"}, "", 0, "total 0 cells\n", ""},
	{[]string{"size"}, "a A1:B2\nb C1:D2\nc A3:B4\n", 0, "a A1:B2 2x2 4 cells\nb C1:D2 2x2 4 cells\nc A3:B4 2x2 4 cells\ntotal 12 cells\n", ""},
	{[]string{"size"}, "big A1:J10\nmid C3:E5\ndot D4\n", 0, "big A1:J10 10x10 100 cells\nmid C3:E5 3x3 9 cells\ndot D4 1x1 1 cell\ntotal 100 cells\n", ""},
	{[]string{"size"}, "a B2:C3\nb B2:C3\nc B2:C3\n", 0, "a B2:C3 2x2 4 cells\nb B2:C3 2x2 4 cells\nc B2:C3 2x2 4 cells\ntotal 4 cells\n", ""},
	{[]string{"size"}, "row A5:Z5\ncol C1:C20\ncell C5\n", 0, "row A5:Z5 1x26 26 cells\ncol C1:C20 20x1 20 cells\ncell C5 1x1 1 cell\ntotal 45 cells\n", ""},
	{[]string{"size"}, "a A1\nb ZZ9999\n", 0, "a A1 1x1 1 cell\nb ZZ9999 1x1 1 cell\ntotal 2 cells\n", ""},
	{[]string{"size"}, "only C3\n", 0, "only C3 1x1 1 cell\ntotal 1 cell\n", ""},
	{[]string{"size"}, "tl A1\nbr ZZ9999\ntr ZZ1\nbl A9999\n", 0, "tl A1 1x1 1 cell\nbr ZZ9999 1x1 1 cell\ntr ZZ1 1x1 1 cell\nbl A9999 1x1 1 cell\ntotal 4 cells\n", ""},
	{[]string{"size"}, "w A1:ZZ1\n", 0, "w A1:ZZ1 1x702 702 cells\ntotal 702 cells\n", ""},
	{[]string{"size"}, "t A1:A9999\n", 0, "t A1:A9999 9999x1 9999 cells\ntotal 9999 cells\n", ""},
	{[]string{"size"}, "a A1:A1\n", 0, "a A1 1x1 1 cell\ntotal 1 cell\n", ""},
	{[]string{"size"}, "x A9:C11\ny B10\n", 0, "x A9:C11 3x3 9 cells\ny B10 1x1 1 cell\ntotal 9 cells\n", ""},
	{[]string{"size"}, "z AA1:AB2\ny Z1:AA1\n", 0, "z AA1:AB2 2x2 4 cells\ny Z1:AA1 1x2 2 cells\ntotal 5 cells\n", ""},
	{[]string{"size"}, "# none\n", 0, "total 0 cells\n", ""},
	{[]string{"size"}, "a A1:E5\nb C3:G7\nc E5:I9\nd A5:I5\ne E1:E9\n", 0, "a A1:E5 5x5 25 cells\nb C3:G7 5x5 25 cells\nc E5:I9 5x5 25 cells\nd A5:I5 1x9 9 cells\ne E1:E9 9x1 9 cells\ntotal 57 cells\n", ""},
	{[]string{"size", "--rows", "5"}, "a A1:C3\nb B2:D4\nc F6\n", 0, "a A1:C3 3x3 9 cells\nb B2:D4 3x3 9 cells\nc F6 1x1 1 cell\ntotal 15 cells\n", ""},
}

func TestRowsSize(t *testing.T) { checkRows(t, rowsSize) }

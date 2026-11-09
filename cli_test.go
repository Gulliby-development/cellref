package main

import "testing"

var rowsCli = []row{
	{[]string{}, "", 2, "", "usage: cellref check [--rows N] [--cols N]\n       cellref size [--rows N] [--cols N]\n       cellref union [--rows N] [--cols N]\n       cellref overlap [--rows N] [--cols N]\n       cellref shift [--rows N] [--cols N]\n"},
	{[]string{"help"}, "", 0, "usage: cellref check [--rows N] [--cols N]\n       cellref size [--rows N] [--cols N]\n       cellref union [--rows N] [--cols N]\n       cellref overlap [--rows N] [--cols N]\n       cellref shift [--rows N] [--cols N]\n", ""},
	{[]string{"-h", "extra"}, "", 0, "usage: cellref check [--rows N] [--cols N]\n       cellref size [--rows N] [--cols N]\n       cellref union [--rows N] [--cols N]\n       cellref overlap [--rows N] [--cols N]\n       cellref shift [--rows N] [--cols N]\n", ""},
	{[]string{"--help", "--bogus"}, "", 0, "usage: cellref check [--rows N] [--cols N]\n       cellref size [--rows N] [--cols N]\n       cellref union [--rows N] [--cols N]\n       cellref overlap [--rows N] [--cols N]\n       cellref shift [--rows N] [--cols N]\n", ""},
	{[]string{"--version"}, "", 0, "cellref 0.1.0\n", ""},
	{[]string{"--version", "x"}, "", 0, "cellref 0.1.0\n", ""},
	{[]string{"frob"}, "", 2, "", "cellref: unknown command frob\n"},
	{[]string{""}, "", 2, "", "cellref: unknown command \n"},
	{[]string{"Check"}, "", 2, "", "cellref: unknown command Check\n"},
	{[]string{"--rows"}, "", 2, "", "cellref: unknown command --rows\n"},
	{[]string{"--rows", "5"}, "", 2, "", "cellref: unknown command --rows\n"},
	{[]string{"help", "shift"}, "", 0, "usage: cellref check [--rows N] [--cols N]\n       cellref size [--rows N] [--cols N]\n       cellref union [--rows N] [--cols N]\n       cellref overlap [--rows N] [--cols N]\n       cellref shift [--rows N] [--cols N]\n", ""},
	{[]string{"size"}, "a A1:C3\nb B2:D4\nc F6\n", 0, "a A1:C3 3x3 9 cells\nb B2:D4 3x3 9 cells\nc F6 1x1 1 cell\ntotal 15 cells\n", ""},
	{[]string{"check"}, "a A1:C3\nb B2:D4\nc F6\n", 0, "ok 3 ranges\n", ""},
	{[]string{"union"}, "", 0, "bounds none\n", ""},
}

func TestRowsCli(t *testing.T) { checkRows(t, rowsCli) }

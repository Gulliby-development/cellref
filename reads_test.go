package main

import "testing"

var rowsReads = []row{
	{[]string{"check"}, "", 1, "", ""},
	{[]string{"size"}, "", 1, "", ""},
	{[]string{"union"}, "", 1, "", ""},
	{[]string{"overlap"}, "", 1, "", ""},
	{[]string{"shift"}, "", 1, "", ""},
	{[]string{"shift", "--rows", "-5"}, "", 1, "", ""},
	{[]string{"shift", "--cols", "9999", "--rows", "-9999"}, "", 1, "", ""},
	{[]string{"help"}, "", 0, "", ""},
	{[]string{"-h"}, "", 0, "", ""},
	{[]string{"--help"}, "", 0, "", ""},
	{[]string{"--version"}, "", 0, "", ""},
	{[]string{}, "", 0, "", ""},
	{[]string{"frob"}, "", 0, "", ""},
	{[]string{""}, "", 0, "", ""},
	{[]string{"check", "x"}, "", 0, "", ""},
	{[]string{"shift", "--rows"}, "", 0, "", ""},
	{[]string{"shift", "--rows", "10000"}, "", 0, "", ""},
	{[]string{"shift", "--rows", "-0"}, "", 0, "", ""},
	{[]string{"shift", "--rows", "1", "--rows", "2"}, "", 0, "", ""},
	{[]string{"shift", "--cols", "7", "--bogus"}, "", 0, "", ""},
	{[]string{"Shift"}, "", 0, "", ""},
	{[]string{"--rows", "5"}, "", 0, "", ""},
	{[]string{"size", "-"}, "", 0, "", ""},
	{[]string{"union", ""}, "", 0, "", ""},
	{[]string{"overlap", "--cols", "-1", "--rows", "1"}, "", 1, "", ""},
	{[]string{"check", "--cols", "00"}, "", 0, "", ""},
	{[]string{"check", "--cols", "-1"}, "", 1, "", ""},
	{[]string{"size", "--rows=1"}, "", 0, "", ""},
}

func TestRowsReads(t *testing.T) { checkReadsRows(t, rowsReads) }

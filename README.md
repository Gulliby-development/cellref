<p align="center"><img src="docs/logo.svg" width="96" alt="" /></p>

# cellref

[![CI](https://github.com/Gulliby-development/cellref/actions/workflows/ci.yml/badge.svg)](https://github.com/Gulliby-development/cellref/actions/workflows/ci.yml) [![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

cellref is an offline, standard-library-only Go command-line tool for checking and
analyzing named spreadsheet ranges. It is intended for spreadsheet users and for
people who automate spreadsheet-related work from a terminal, scripts, or CI.
Developers can build and test it with Go; it makes no network requests.

The program supports Windows, macOS, and Linux and requires Go 1.27.0 or newer.

## Build and run

From the repository directory, build the executable:

```sh
go build -o cellref .
```

Run it as `./cellref` on POSIX shells, or invoke the built executable using the
appropriate platform command. The command reads range data from standard input
only when a valid command and its options have been supplied.

## Commands

```text
usage: cellref check [--rows N] [--cols N]
       cellref size [--rows N] [--cols N]
       cellref union [--rows N] [--cols N]
       cellref overlap [--rows N] [--cols N]
       cellref shift [--rows N] [--cols N]
```

The first word selects a command. `help`, `-h`, and `--help` print the usage
to standard output and exit with status 0; `--version` prints exactly
`cellref 0.1.0` followed by a line feed and exits with status 0. These take
precedence over any following words and do not read standard input. With no
command, usage is written to standard error and the program exits with status 2.
Command names are case-sensitive. Any other first word, including an empty one,
produces `cellref: unknown command WORD` and status 2.

All commands accept `--rows N` and `--cols N`; both default to zero. They are
validated in their order on the command line, and only `shift` applies the
values. Each option consumes the next word as its value, even when that word
looks like an option. For each encountered option, validation checks in this
order: a repeated option reports `cellref: given twice: OPTION`; otherwise, an
option without a following word reports `cellref: missing value for OPTION`;
otherwise, an invalid value reports `cellref: bad value for OPTION`. An
argument that is empty, exactly `-`, or does not start with `-` reports
`cellref: unexpected argument WORD`; an unsupported word starting with `-`
reports `cellref: unknown option WORD`. The first command-line problem wins,
exits with status 2, and input is not read.

Values are canonical decimal integers from -9999 through 9999. A minus sign is
allowed, but a plus sign, leading zero, and negative zero are not. Thus 0 is
valid, while 007 and -0 are invalid.

## Input

Each meaningful input line contains a unique name and one range, separated by
one or more spaces or tabs. Leading and trailing spaces or tabs are ignored.
Blank lines and lines whose first word starts with `#` are skipped; inline
comments are not supported. Input is split at line feeds; empty text has zero
physical lines, and a final line without a line feed still counts as a line.
One carriage return at the end of each line is removed.

At most 40 physical lines are accepted. The limit is checked before line
contents, and skipped lines count toward it. Names contain 1 to 12 characters:
a lowercase letter followed by lowercase letters, digits, or hyphens.

A cell has one or two uppercase column letters followed by a row written in
decimal digits without a leading zero, from 1 to 9999. Columns run from A
through ZZ (1 through 702); for example, AA is column 27. A range is one cell
or two cells separated by a colon, such as `B2:D4`.
The ending cell cannot be to the left of, or above, the starting cell. A
single-cell range may be written as `C3` or `C3:C3`.

Diagnostic line number N is the one-based physical line number, including
skipped lines. For each non-skipped line, validation checks the number of words first
(`cellref: line N: expected NAME RANGE`), then the name
(`cellref: line N: bad name WORD`), then the range syntax
(`cellref: line N: bad range WORD`), then backward order
(`cellref: line N: range runs backwards`), and finally duplicate names
(`cellref: line N: duplicate name NAME`). The first invalid line in input
order is reported. An input error exits with status 1.

A command that succeeds writes its result to standard output and nothing to
standard error. A failure writes no standard output and one diagnostic line to
standard error, terminated by a line feed. Input and shift failures exit 1;
command-line failures exit 2. An input exceeding the line limit reports
`cellref: input is too long (max 40 lines)`. When shifting would move a range
outside the sheet, the diagnostic is
`cellref: line N: NAME moves outside the sheet`.

## Reports

All output uses canonical range notation: a single cell is printed as one cell
(for example, `C3`, not `C3:C3`), and a larger range is printed as
`START:END`. Results that list ranges preserve input order. Counts use a
singular noun only for one: `1 range` and `1 cell`; zero and other counts
use plurals.

- `check` prints `ok N ranges`, counting meaningful input lines. With no
  ranges it prints `ok 0 ranges`.
- `size` prints one line per range in input order as
  `NAME RANGE RxC N cells`, where R and C are inclusive row and column
  dimensions. The cell noun is singular for a one-cell range. Its final line
  gives the number of distinct cells covered by all ranges; overlapping cells
  count once. With no ranges it prints `total 0 cells`.
- `union` prints the smallest enclosing bounds and the number of distinct
  covered cells compared with the area of those bounds. The covered line
  always says `cells`, including when either count is one. With no ranges it
  prints `bounds none`.
- `overlap` prints each pair that shares one or more cells, in input order:
  earlier range first, then each later range in order. Each line gives both
  names, the inclusive intersection, and its cell count. Merely touching at
  an edge is not an overlap. With no matching pairs, including empty input, it
  prints `overlaps 0`.
- `shift` moves ranges down/right for positive row/column values and
  up/left for negative values. For each range it prints its original and new
  canonical range, then the number shifted. A move is valid only if every cell
  remains within columns A through ZZ and rows 1 through 9999. The first
  out-of-sheet range in input order reports its original line number; the
  command emits no partial report. With no ranges it prints
  `shifted 0 ranges`.

A command-line error exits with status 2; invalid input and out-of-sheet shifts
exit with status 1. Every failure leaves standard output empty and writes one
line to standard error, prefixed by `cellref: ` and ending with a line feed.
Command-line diagnostics are `unknown command WORD`, `unexpected argument
WORD`, `unknown option WORD`, `given twice: OPTION`, `missing value for
OPTION`, and `bad value for OPTION`, each after the prefix. Input diagnostics
are `input is too long (max 40 lines)`, or
`line N: expected NAME RANGE`, `line N: bad name WORD`,
`line N: bad range WORD`, `line N: range runs backwards`, and
`line N: duplicate name NAME`. A failed shift reports
`line N: NAME moves outside the sheet`.

## Examples

The following examples use a POSIX shell in the repository directory, after
building the program with `go build -o cellref .`. Here, `cellref` means
that built program, invoked as `./cellref`. The final example uses `2>&1`
to show diagnostic output alongside standard output.

```text
$ printf 'a A1:C3\nb B2:D4\n' | cellref overlap
a b B2:C3 4 cells
overlaps 1
$ printf 'a A1:C3\nb B2:D4\nc F6\n' | cellref size; printf 'a A1:C3\nb B2:D4\nc F6\n' | cellref union
a A1:C3 3x3 9 cells
b B2:D4 3x3 9 cells
c F6 1x1 1 cell
total 15 cells
bounds A1:F6
covered 15 of 36 cells
$ printf 'a Z1:AA2\nb C3\n' | cellref shift --cols 3 --rows 1
a Z1:AA2 AC2:AD3
b C3 F4
shifted 2 ranges
$ printf 'a B2:A1\n' | cellref check 2>&1; printf 'a A1\n' | cellref shift --rows -1 2>&1
cellref: line 1: range runs backwards
cellref: line 1: a moves outside the sheet
```

## Development

Run the three project checks from the repository directory:

```sh
go vet ./...
gofmt -l .
go test ./...
```

The formatting check should print no file names.

See [LICENSE](LICENSE) for the license.

---

Built with [BAKED](https://github.com/Gulliby-development/baked), an open-source coding agent (bench round R96).

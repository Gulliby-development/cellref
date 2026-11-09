package main

const (
	version  = "cellref 0.1.0"
	maxLines = 40
	maxName  = 12
	maxCol   = 702
	maxRow   = 9999
	maxShift = 9999
)

const usage = "usage: cellref check [--rows N] [--cols N]\n" +
	"       cellref size [--rows N] [--cols N]\n" +
	"       cellref union [--rows N] [--cols N]\n" +
	"       cellref overlap [--rows N] [--cols N]\n" +
	"       cellref shift [--rows N] [--cols N]\n"

// failure is an error that carries the exit status of the program: 2 for the
// command line, 1 for the input. fail builds one; run turns it into the reply
// (status, "", "cellref: " + msg + "\n").
type failure struct {
	status int
	msg    string
}

func (f *failure) Error() string { return f.msg }

func fail(status int, msg string) *failure { return &failure{status: status, msg: msg} }

// options holds the values of --rows and --cols (0 when the option is absent).
type options struct {
	rows, cols int
}

// item is one meaningful line of the input: the name, the line number (counted
// from 1, skipped lines included) and the corners of the range, top left first.
type item struct {
	name           string
	line           int
	c1, r1, c2, r2 int
}

// command is one entry of the command table; run returns the whole standard
// output of the command or a failure.
type command struct {
	name string
	run  func(opts options, items []item) (string, *failure)
}

var commands = []command{
	{"check", cmdCheck},
	{"size", cmdSize},
	{"union", cmdUnion},
	{"overlap", cmdOverlap},
	{"shift", cmdShift},
}

func findCommand(name string) *command {
	for i := range commands {
		if commands[i].name == name {
			return &commands[i]
		}
	}
	return nil
}

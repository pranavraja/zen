package zen

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"strings"
)

type failingLine struct {
	content  string
	filename string
	next     string
	number   int
	prev     string
}

// TODO move these to a per-scenerio structure to allow
// for concurrent text execution in future Go versions (1.4-ish?)
var lastContext string
var lastWhen string
var lastTitle string

var (
	reset  = "\033[0m"
	white  = "\033[37m\033[1m"
	grey   = "\x1B[90m"
	red    = "\033[31m\033[1m"
	yellow = "\033[43m"
)

func (test *Test) PrintContext() {
	if lastContext == test.Context {
		return
	}
	fmt.Printf("\033[37m \033[1mGiven %s \n", test.Context)
	lastContext = test.Context
}

func (test *Test) PrintWhen() {
	if lastWhen == test.When {
		return
	}
	fmt.Printf("\033[37m \033[1m  When %s \n", test.When)
	lastWhen = test.When
}

func (test *Test) PrintTitle() {
	if lastTitle == test.Title {
		return
	}
	fmt.Printf("\033[37m \033[1m    It %s %s\n", test.Title, reset)
	lastTitle = test.Title
}

func (test *Test) PrintTitleNotImplemented() {
	if lastTitle == test.Title {
		return
	}
	fmt.Printf("\033[37m \033[1m    It %s %s<<-- NOT IMPLEMENTED%s\n", test.Title, yellow, reset)
	lastTitle = test.Title
}

func (test *Test) PrintError(message string) {
	test.PrintTitle()
	failingLine, err := getFailingLine()

	if err != nil {
		return
	}

	fmt.Printf("%s        %s %s %s %s\n", red, message, grey, path.Base(failingLine.filename), reset)
	test.PrintFailingLine(&failingLine)
	test.T.Fail()
}

func (test *Test) PrintFailingLine(failingLine *failingLine) {
	fmt.Printf("%s             %d. %s\n", grey, failingLine.number-1, failingLine.prev)
	fmt.Printf("%s             %d. %s %s\n", white, failingLine.number, failingLine.content, reset)
	fmt.Printf("%s             %d. %s\n", grey, failingLine.number+1, failingLine.next)
	fmt.Println(reset)
}

func getFailingLine() (failingLine, error) {
	_, filename, ln, _ := runtime.Caller(3)
	// TODO: this is really hacky, need to find a way of not using magic numbers for runtime.Caller
	// If we are not in a test file, we must still be inside this package,
	// so we need to go up one more stack frame to get to the test file
	if !strings.HasSuffix(filename, "_test.go") {
		_, filename, ln, _ = runtime.Caller(4)
	}

	bf, err := ioutil.ReadFile(filename)

	if err != nil {
		return failingLine{}, fmt.Errorf("Failed to open %s", filename)
	}

	lines := strings.Split(string(bf), "\n")[ln-2 : ln+2]

	return failingLine{
		softTabs(lines[1]),
		filename,
		softTabs(lines[2]),
		int(ln),
		softTabs(lines[0]),
	}, nil

}

func softTabs(text string) string {
	return strings.Replace(text, "\t", "  ", -1)
}

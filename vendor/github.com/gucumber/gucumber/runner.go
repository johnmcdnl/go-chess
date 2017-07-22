package gucumber

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/shiena/ansicolor"
	"github.com/cucumber/gherkin-go"
)

const (
	clrWhite  = "0"
	clrRed    = "31"
	clrGreen  = "32"
	clrYellow = "33"
	clrCyan   = "36"

	txtUnmatchInt   = `(\d+)`
	txtUnmatchFloat = `(-?\d+(?:\.\d+)?)`
	txtUnmatchStr   = `"(.+?)"`
)

var (
	reUnmatchInt   = regexp.MustCompile(txtUnmatchInt)
	reUnmatchFloat = regexp.MustCompile(txtUnmatchFloat)
	reUnmatchStr   = regexp.MustCompile(`(<|").+?("|>)`)
	reOutlineVal   = regexp.MustCompile(`<(.+?)>`)
)

type Runner struct {
	*Context
	GherkinDocuments []*gherkin.GherkinDocument
	Results          []RunnerResult
	Unmatched        []*gherkin.Step
	FailCount        int
	SkipCount        int
}

type RunnerResult struct {
	*TestingT
	*gherkin.Pickle
}

func (c *Context) RunDir(dir string) (*Runner, error) {
	var featureFiles []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".feature") {
			featureFiles = append(featureFiles, path)
		}
		return nil
	})

	runner, err := c.RunFiles(featureFiles)
	if err != nil {
		panic(err)
	}

	if len(runner.Unmatched) > 0 {
		fmt.Println("Some steps were missing, you can add them by using the following step definition stubs: ")
		fmt.Println("")
		fmt.Print(runner.MissingMatcherStubs())
	}

	os.Exit(runner.FailCount)
	return runner, err
}

func parseFile(filename string) (*gherkin.GherkinDocument, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return gherkin.ParseGherkinDocument(f)
}

func (c *Context) RunFiles(featurePaths []string) (*Runner, error) {
	r := Runner{
		Context:          c,
		GherkinDocuments: []*gherkin.GherkinDocument{},
		Results:          []RunnerResult{},
		Unmatched:        []*gherkin.Step{},
	}

	for _, filePath := range featurePaths {
		gd, err := parseFile(filePath)
		if err != nil {
			return nil, err
		}
		r.GherkinDocuments = append(r.GherkinDocuments, gd)
	}

	r.run()
	return &r, nil
}

func (r *Runner) MissingMatcherStubs() string {
	var buf bytes.Buffer
	matches := map[string]bool{}

	buf.WriteString(`import . "github.com/gucumber/gucumber"` + "\n\n")
	buf.WriteString("func init() {\n")

	for _, m := range r.Unmatched {
		numInts, numFloats, numStrs := 1, 1, 1
		str, args := m.Text, []string{}
		str = reUnmatchInt.ReplaceAllStringFunc(str, func(s string) string {
			args = append(args, fmt.Sprintf("i%d int", numInts))
			numInts++
			return txtUnmatchInt
		})
		str = reUnmatchFloat.ReplaceAllStringFunc(str, func(s string) string {
			args = append(args, fmt.Sprintf("s%d float64", numFloats))
			numFloats++
			return txtUnmatchFloat
		})
		str = reUnmatchStr.ReplaceAllStringFunc(str, func(s string) string {
			args = append(args, fmt.Sprintf("s%d string", numStrs))
			numStrs++
			return txtUnmatchStr
		})

		//TODO
		//if len(m.Argument) > 0 {
		//	if m.Argument.IsTabular() {
		//		args = append(args, "table [][]string")
		//	} else {
		//		args = append(args, "data string")
		//	}
		//}

		// Don't duplicate matchers. This is mostly for scenario outlines.
		if matches[str] {
			continue
		}
		matches[str] = true

		fmt.Fprintf(&buf, "\t%s(`^%s$`, func(%s) {\n\t\tT.Skip() // pending\n\t})\n\n",
			m.Type, str, strings.Join(args, ", "))
	}

	buf.WriteString("}\n")
	return buf.String()
}

func (r *Runner) run() {
	if r.BeforeAllFilter != nil {
		r.BeforeAllFilter()
	}
	for _, g := range r.GherkinDocuments {
		r.runGherkinDoc(g)
	}
	if r.AfterAllFilter != nil {
		r.AfterAllFilter()
	}

	r.line("0;1", "Finished (%d passed, %d failed, %d skipped).\n",
		len(r.Results)-r.FailCount-r.SkipCount, r.FailCount, r.SkipCount)
}

func (r *Runner) runGherkinDoc(g *gherkin.GherkinDocument) {
	for _, p := range g.Pickles() {
		r.runPickle(p)
	}

}

func (r *Runner) runPickle(p *gherkin.Pickle) {
	//TODO
	//if !s.FilterMatched(f, c.Filters...) {
	//	return
	//
	//}

	for k, fn := range r.BeforeFilters {
		//TODO
		//if s.FilterMatched(f, strings.Split(k, "|")...) {
		//	fn()
		//}
		fmt.Sprint(k, fn)
	}

	t := &TestingT{}
	var skipping bool
	var clr = clrGreen

	for _, step := range p.Steps {
		errCount := len(t.errors)
		found := false
		if !skipping && !t.Failed() {
			done := make(chan bool)
			go func() {
				defer func() {
					r.Results = append(r.Results, RunnerResult{t, p})

					if t.Skipped() {
						r.SkipCount++
						skipping = true
						clr = clrYellow
					} else if t.Failed() {
						r.FailCount++
						clr = clrRed
					}

					done <- true
				}()

				//TODO
				args := "TODO"
				if len(step.Arguments) == 0 {
					args = ""
				}

				f, err := r.Execute(t, step.Text, string(args))
				if err != nil {
					t.Error(err)
				}
				found = f

				if !f {
					t.Skip("no match function for step")
				}
			}()
			<-done
		}

		if len(t.errors) > errCount {
			r.line(clrRed, "\n"+t.errors[len(t.errors)-1].message)
		}
	}

	for k, fn := range r.AfterFilters {
		//if s.FilterMatched(f, strings.Split(k, "|")...) {
		//	fn()
		//}
		//TODO
		fmt.Sprint(k, fn)
	}
}

var writer = ansicolor.NewAnsiColorWriter(os.Stdout)

func (r *Runner) line(clr, text string, args ...interface{}) {
	fmt.Fprintf(writer, "\033[%sm%s\033[0;0m\n", clr, fmt.Sprintf(text, args...))
}

func (r *Runner) fileLine(clr, text, filename string, line int, max int, args ...interface{}) {
	space, str := "", fmt.Sprintf(text, args...)
	if l := max + 5 - len(str); l > 0 {
		space = strings.Repeat(" ", l)
	}
	comment := fmt.Sprintf("%s \033[39;0m# %s:%d", space, filename, line)
	r.line(clr, "%s%s", str, comment)
}

type Tester interface {
	Errorf(format string, args ...interface{})
	Skip(args ...interface{})
}

type TestingT struct {
	skipped bool
	errors  []TestError
}

type TestError struct {
	message string
	stack   []byte
}

func (t *TestingT) Errorf(format string, args ...interface{}) {
	var buf bytes.Buffer

	str := fmt.Sprintf(format, args...)
	sbuf := make([]byte, 8192)
	for {
		size := runtime.Stack(sbuf, false)
		if size < len(sbuf) {
			break
		}
		buf.Write(sbuf[0:size])
	}

	t.errors = append(t.errors, TestError{message: str, stack: buf.Bytes()})
}

func (t *TestingT) Skip(args ...interface{}) {
	t.skipped = true
}

func (t *TestingT) Skipped() bool {
	return t.skipped
}

func (t *TestingT) Failed() bool {
	return len(t.errors) > 0
}

func (t *TestingT) Error(err error) {
	t.errors = append(t.errors, TestError{message: err.Error()})
}

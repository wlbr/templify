package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var bad = `package main
import "fmt"
func main() {
fmt.Printf("Hello, world.\n")
}`

var good = `package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
}
`

func TestFormatFile(t *testing.T) {

	var f, _ = ioutil.TempFile(".", "templifytest")
	var fname = f.Name()
	f.Close()

	ioutil.WriteFile(fname, []byte(bad), 0644)

	formatFile(fname)

	uglytmp, _ := ioutil.ReadFile(fname)
	ugly := string(uglytmp)

	defer os.Remove(fname)

	if good != ugly {
		fmt.Printf("Formatted file '%s' differs from gold standard.\n", fname)
		t.Fail()
	}
}

func TestFlagging(t *testing.T) {

	args := []string{"-p", "testpackage", "/my/path/testfile.tpl"}
	os.Args = append(os.Args, args...)
	flagging()
	if pckg != "testpackage" {
		fmt.Printf("package flag was not as expected.  'pckg=%s'", pckg)
		t.Fail()
	}

	if inputfile != "/my/path/testfile.tpl" {
		fmt.Printf("inputfile flag was not as expected.  'inputfile=%s'", inputfile)
		t.Fail()
	}
	if outfilename != "/my/path/testfile.go" {
		fmt.Printf("output filename not correctly determined.  'outfilename=%s'", inputfile)
		t.Fail()
	}

	if functionname != "testfile" {
		fmt.Printf("functionname not correctly determined.  'functionname=%s'", inputfile)
		t.Fail()
	}
	fmt.Println()
}

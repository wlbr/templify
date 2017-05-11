package main

import (
	"flag"
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

	flag.Set("p", "test")
	flagging()
}

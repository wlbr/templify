package main

//go:generate templify -p main -o embed.go embed.tpl

import (
	"errors"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

var pckg string
var tmpl string
var out string
var frmt bool

const embedtpl = "embed.tpl"

func init() {

}

func flagging() {
	flag.StringVar(&pckg, "p", "main", "name of package to be used in generated code")
	flag.StringVar(&out, "o", "", "name of output file. Defaults to name of template file + '.go'")
	flag.BoolVar(&frmt, "n", false, "do not format the generated source. Default false means source will be formatted.")
	flag.Parse()

	tmpl = flag.Arg(0)
	if tmpl == "" {
		fmt.Println(errors.New("No template file given as argument."))
		os.Exit(1)
	}
}

func readTemplifyTemplate(tplname string) (*template.Template, error) {
	tpl, err := template.ParseFiles(tplname)
	if err != nil {
		fmt.Printf("Error reading templifytemplate file '%s'\n%v", tplname, err)
	}
	return tpl, err
}

const internalTemplate = "/*\n" +
	"* CODE GENERATED AUTOMATICALLY WITH github.com/wlbr/templify\n" +
	"* using the internal, hardcoded template.\n" +
	"* THIS FILE SHOULD NOT BE EDITED BY HAND\n" +
	"*/\n" +
	"\n" +
	"package {{.Pckg}}\n" +
	"\n" +
	"var {{.Tmplname}}tmpl = `{{.Tmplcontent}}`\n" +
	"\n" +
	"func {{.Tmplname}}Template() string {\n" +
	"return {{.Tmplname}}tmpl\n" +
	"}\n"

func readTargetTemplate(tplname string) string {
	tpl, err := ioutil.ReadFile(tplname)
	if err != nil {
		fmt.Printf("Error reading target template file '%s'\n%v", tplname, err)
		os.Exit(1)
	}

	r := strings.NewReplacer("\"", "\\\"", "\n", "\\n\" +\n\t\"", "\t", "\\t")
	return r.Replace(string(tpl))

}

func formatFile(fname string) {
	fstr, err := ioutil.ReadFile(out)
	if err != nil {
		fmt.Printf("Error reading generated file %s before passing it to gofmt.\n%v\n", out, err)
		os.Exit(1)
	} else {
		fstr, err = format.Source(fstr)
		if err != nil {
			fmt.Printf("Error running gofmt on the generated file '%s'\n%v\n", out, err)
			os.Exit(1)
		} else {
			foutfile, err := os.Create(out)
			if err != nil {
				fmt.Printf("Error creating formatted target file '%s'\n%v\n", out, err)
				os.Exit(1)
			} else {
				defer foutfile.Close()
				fmt.Fprintf(foutfile, "%s", fstr)
			}
		}
	}
}

func main() {
	flagging()
	// tpl, err := readTemplifyTemplate(embedtpl)
	tpl, err := template.New("embed").Parse(embedTemplate())

	//tpl, err := template.New("embed").Parse(internalTemplate)
	if err != nil {
		fmt.Printf("Error parsing code generation template\n%v", err)
		os.Exit(1)
	}

	data := struct {
		Pckg        string
		Tmplname    string
		Tmplcontent string
	}{
		Pckg:     pckg,
		Tmplname: strings.Split(tmpl, ".")[0],
	}
	data.Tmplcontent = readTargetTemplate(tmpl)

	if out == "" {
		out = strings.Split(tmpl, ".")[0] + ".go"
	}
	outfile, err := os.Create(out)
	if err != nil {
		fmt.Printf("Error creating target file '%s'\n%v\n", out, err)
		os.Exit(1)
	}
	defer outfile.Close()
	tpl.Execute(outfile, data)

	if !frmt {
		formatFile(out)
	}
}

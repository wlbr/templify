package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

var pckg string
var tmpl string
var out string

const embedtpl = "embed.tpl"

func init() {

}

func flagging() {
	flag.StringVar(&pckg, "p", "main", "name of package to be used in generated code")
	flag.StringVar(&out, "o", "", "name of output file. Defaults to name of template file + '.go'")

	flag.Parse()

	tmpl = flag.Arg(0)
	if tmpl == "" {
		fmt.Println(errors.New("No template file given as argument."))
		os.Exit(1)
	}
}

func readTemplifyTemplate(tplname string) *template.Template {
	tpl, err := template.ParseFiles(tplname)
	if err != nil {
		fmt.Printf("Error reading templifytemplate file '%s'\n%v", tplname, err)
		os.Exit(1)
	}
	return tpl
}

const embedTemplate = "/*\n" +
	"* CODE GENERATED AUTOMATICALLY WITH github.com/wlbr/templify\n" +
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
	/*stpl := string(tpl)
	stpl = strings.Replace(stpl, "\"", "\\\"", -1)
	stpl = strings.Replace(stpl, "\n", "\\n\" +\n\"", -1)*/

	return string(tpl)
}

func main() {
	flagging()
	//tpl := readTemplifyTemplate(embedtpl)
	tpl, err := template.New("embed").Parse(embedTemplate)
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

}

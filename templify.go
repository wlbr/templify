package main

//go:generate templify -p main -o embed.go embed.tpl

import (
	"errors"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

var pckg string
var inputfile string
var outfilename string
var functionname string
var tmplname string
var frmt bool
var exp bool

const embedtpl = "embed.tpl"

func flagging() {
	flag.StringVar(&pckg, "p", "main", "name of package to be used in generated code.")
	flag.StringVar(&outfilename, "o", "", "name of output file. Defaults to name of template file + '.go'.")
	flag.StringVar(&functionname, "f", "", "name of generated, the template returning function. Its name will "+
		"have 'Template' attached. Will be set to $(basename -s .ext outputfile) if empty (default).")
	flag.StringVar(&tmplname, "t", "", "name of alternate code generation template file. If empty (default), "+
		"then the embedded template will be used. Template variables supplied are: .Name, .Package, .Content")
	flag.BoolVar(&frmt, "n", false, "do not format the generated source. Default false means source will be formatted.")
	flag.BoolVar(&exp, "e", false, "export the generated, the template returning function. "+
		"Default (false) means the function will not be exported.")
	flag.Parse()

	inputfile = flag.Arg(0)
	if inputfile == "" {
		fmt.Println(errors.New("No template file given as argument."))
		os.Exit(1)
	}

	if outfilename == "" {
		indir := path.Dir(inputfile)
		inext := path.Ext(path.Base(inputfile))
		inname := strings.TrimSuffix(path.Base(inputfile), inext)
		outfilename = fmt.Sprintf("%s/%s.go", indir, inname)
	}

	if functionname == "" {
		ext := path.Ext(path.Base(outfilename))
		functionname = strings.TrimSuffix(path.Base(outfilename), ext)
	}

	if exp {
		functionname = strings.ToUpper(functionname[0:1]) + functionname[1:]
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
	fstr, err := ioutil.ReadFile(outfilename)
	if err != nil {
		fmt.Printf("Error reading generated file %s before passing it to gofmt.\n%v\n", outfilename, err)
		os.Exit(1)
	} else {
		fstr, err = format.Source(fstr)
		if err != nil {
			fmt.Printf("Error running gofmt on the generated file '%s'\n%v\n", outfilename, err)
			os.Exit(1)
		} else {
			foutfile, err := os.Create(outfilename)
			if err != nil {
				fmt.Printf("Error creating formatted target file '%s'\n%v\n", outfilename, err)
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

	var tpl *template.Template
	var err error

	if tmplname != "" {
		tpl, err = readTemplifyTemplate(tmplname)

	} else {
		tpl, err = template.New("embed").Parse(embedTemplate())
	}

	//tpl, err := template.New("embed").Parse(internalTemplate)

	if err != nil {
		fmt.Printf("Error parsing code generation template\n%v", err)
		os.Exit(1)
	}

	data := struct {
		Package string
		Name    string
		Content string
	}{
		Package: pckg,
		Name:    strings.Split(functionname, ".")[0],
	}
	data.Content = readTargetTemplate(inputfile)

	if outfilename == "" {
		outfilename = strings.Split(inputfile, ".")[0] + ".go"
	}
	outfile, err := os.Create(outfilename)
	if err != nil {
		fmt.Printf("Error creating target file '%s'\n%v\n", outfilename, err)
		os.Exit(1)
	}
	defer outfile.Close()
	tpl.Execute(outfile, data)

	if !frmt {
		formatFile(outfilename)
	}
}

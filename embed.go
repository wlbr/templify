/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

// embed is a generated function returning the template as a string.
// That string should be parsed by the functions of the template package.
func embedTemplate() string {
	var tmpl = "/*\n" +
		" * CODE GENERATED AUTOMATICALLY WITH\n" +
		" *    github.com/wlbr/templify\n" +
		" * THIS FILE SHOULD NOT BE EDITED BY HAND\n" +
		" */\n" +
		"\n" +
		"package {{.Package}}\n" +
		"\n" +
		"// {{.Name}} is a generated function returning the template as a string.\n" +
		"// That string should be parsed by the functions of the template package.\n" +
		"func {{.Name}}Template() string {\n" +
		"\tvar tmpl = \"{{.Content}}\"\n" +
		"\treturn tmpl\n" +
		"}\n" +
		""
	return tmpl
}

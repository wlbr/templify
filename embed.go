/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

func embedTemplate() string {
	var tmpl = "/*\n" +
		" * CODE GENERATED AUTOMATICALLY WITH \n" +
		" *    github.com/wlbr/templify \n" +
		" * THIS FILE SHOULD NOT BE EDITED BY HAND\n" +
		" */\n" +
		"\n" +
		"package {{.Package}}\n" +
		"\n" +
		"func {{.Name}}Template() string {\n" +
		"\tvar tmpl = \"{{.Content}}\"\n" +
		"\treturn tmpl\n" +
		"}\n" +
		""
	return tmpl
}

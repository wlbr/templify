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
		"package {{.Pckg}}\n" +
		"\n" +
		"func {{.Tmplname}}Template() string {\n" +
		"\tvar tmpl = \"{{.Tmplcontent}}\"\n" +
		"\treturn tmpl\n" +
		"}\n" +
		""
	return tmpl
}

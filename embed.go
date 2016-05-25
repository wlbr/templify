/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

var embedtmpl = "/*\n" +
	" * CODE GENERATED AUTOMATICALLY WITH \n" +
	" *    github.com/wlbr/templify \n" +
	" * THIS FILE SHOULD NOT BE EDITED BY HAND\n" +
	" */\n" +
	"\n" +
	"package {{.Pckg}}\n" +
	"\n" +
	"var {{.Tmplname}}tmpl = \"{{.Tmplcontent}}\"\n" +
	"\n" +
	"func {{.Tmplname}}Template() string {\n" +
	"\treturn {{.Tmplname}}tmpl\n" +
	"}\n" +
	""

func embedTemplate() string {
	return embedtmpl
}

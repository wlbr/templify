/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

func exTemplate() string {
	var tmpl = "<!DOCTYPE html PUBLIC \"-//W3C//DTD HTML 4.01 Transitional//EN\" \"http://www.w3.org/TR/html4/loose.dtd\">\n" +
		"\n" +
		"<html itemscope itemtype=\"http://schema.org/\">\n" +
		"<head>\n" +
		"  <meta http-equiv=\"Content-Type\" content=\"text/html; charset=us-ascii\">\n" +
		"  {{.Head}}\n" +
		"  <title>{{.Title}}</title>\n" +
		"</head>\n" +
		"<body>\n" +
		"  <div id=\"Center\">\n" +
		"    <div id=\"Header\">\n" +
		"      <h1>Firstline</h1>\n" +
		"    </div>\n" +
		"    <div id=\"Content\">\n" +
		"      <!-- =============================================== -->\n" +
		"\n" +
		"      {{.Content}}\n" +
		"\n" +
		"      <!-- =============================================== -->\n" +
		"    </div>\n" +
		"  </div>\n" +
		"</body>\n" +
		"</html>"
	return tmpl
}

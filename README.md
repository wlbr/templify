#templify
A tool to be used with 'go generate' to embed external template files into Go code.

##Scenario
An often used scenario in developing go applications is to generate text (e.g. html pages) 
using the Go template packages. You have two choices: either editing the template in a
dedicated file and to read this file at runtime from your application. Or you add a 
"big" string containing the template to the source of your app and parse this string to 
create the template.

If you want to create a single file application, than the more handy solution with a 
dedicated file is not desirable. On the other hand to edit a complex html template within 
a Go string in the source is very unfomfortable. 

templify allows you to edit your template in an extra file and to generate an Go source
file containing the embedded string.

##Installation
   `go get github.com/wlbr/templify`

##Usage

Simply add a line 

   `//go:generate templify mytemplate.file`

for each template you want to embed. Every time you run a `go generate` in the 
corresponding folder, the file `mytemplate.go` will be created. It contains a 
function `mytemplateTemplate` returning the template string.

You may use `templify mytemplate.file` directly on the command line.

##Switches
Will be documented later.



all: build generate

.PHONY: clean
clean: 
	echo clean 
	rm -f embed.go
	rm -f templify
	rm -f example/ex.go


generate: embed.tpl
	echo generate
	go generate templify.go

build: templify.go embed.go 
	echo build
	go build templify.go embed.go

examples: example/ex.tpl example/templifying.go
	go generate example/templifying.go
	go build example/templifying.go example/ex.go

	


	

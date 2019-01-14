
all: test build generate

.PHONY: clean
clean:
	echo clean
	rm -f embed.go
	rm -f templify
	rm -f example/ex.go
	rm -f examplte/templify


generate: embed.tpl
	echo generate
	go generate templify.go

build: templify.go
	echo build
	go build templify.go embed.go

test: templify.go embed.tpl embed.go templify_test.go
	echo test
	go test

examples: example/ex.tpl example/templifying.go
	go generate example/templifying.go
	go build example/templifying.go example/ex.go







#export GOOS=linux
#export GOARCH=amd64

	
all: clean generate build 

install-go-deps:
	go get -u github.com/govend/govend
	
.PHONY: clean
clean: 
	echo clean 
	rm -f embed.go
	#rm -f templify


generate: embed.tpl
	go generate templify.go

build: templify.go generate
	go build templify.go embed.go

run:  
	go run templify.go -p bla -o bla/supertest.go test.tpl
	


	

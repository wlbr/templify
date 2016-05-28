
all: generate build

.PHONY: clean
clean: 
	echo clean 
	rm -f embed.go
	rm -f templify
	rm -f example/ex.go


generate: embed.tpl
	go generate templify.go

build: templify.go embed.go generate
	go build templify.go embed.go

	


	

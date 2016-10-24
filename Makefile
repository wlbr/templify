
all: build generate

.PHONY: clean
clean: 
	echo clean 
	rm -f embed.go
	rm -f templify
	rm -f example/ex.go


generate: embed.tpl
	go generate templify.go

build: templify.go embed.go 
	go build templify.go embed.go

	


	

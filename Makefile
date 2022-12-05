compile:
	GOOS=darwin GOARCH=amd64 go build -o bin/main-darwin-386 main.go
	#GOOS=linux GOARCH=amd64 go build -o bin/main-linux-386 main.go
all:
	compile
clean:
	rm bin/*
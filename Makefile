BINARY = "light-mode"

build:
	go build -buildmode=plugin -o $(BINARY) main.go

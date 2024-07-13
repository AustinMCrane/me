all: pre build

pre:
	go build
build: 
	./me src public

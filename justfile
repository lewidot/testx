build:
	go build -o bin/testx

run: build
	./bin/testx

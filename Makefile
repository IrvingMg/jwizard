BINARY_NAME=jwizard

clean:
	rm -rf bin/

build: clean
	go build -o bin/${BINARY_NAME} main.go

run: build
	./bin/${BINARY_NAME}

test:
	go test ./... -v -count=1

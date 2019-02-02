.PHONY: build
build: 
	go build -o fake-server

.PHONY: run
run: 
	./fake-server

.PHONY: clean
clean:
	go clean
	rm -f fake-server

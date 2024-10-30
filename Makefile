.PHONY: build clean

build:
	@mkdir -p bin
	@go build -o bin/lsp main.go

clean:
	@rm -rf bin/

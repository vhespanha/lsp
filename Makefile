.PHONY: build clean install

build:
	@mkdir -p bin
	@go build -o bin/lsp main.go

clean:
	@rm -rf bin/

install:
	@mkdir -p $(HOME)/.local/bin
	@cp bin/lsp $(HOME)/.local/bin/lsp

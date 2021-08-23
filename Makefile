cmd ?= gen
build:
	@GOOS=darwin GOARCH=amd64 go build -o ./bin/$(cmd) ./cmd/$(cmd)
.PHONY: build

bin:
	./bin/$(cmd)
.PHONY: bin

run:
	@go run ./cmd/$(cmd)
.PHONY: run

gen:
	@./bin/gen
.PHONY: gen

commit:
	@./bin/commit
.PHONY: commit

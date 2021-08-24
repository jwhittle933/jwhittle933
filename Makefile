cmd ?= gen
build:
	@go build -o ./bin/gen ./cmd/gen
	@go build -o ./bin/commit ./cmd/commit
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
	./bin/commit
.PHONY: commit

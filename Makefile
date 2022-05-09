.PHONY:	all clean code-vet code-fmt test

DEPS := $(shell find . -type f -name "*.go" -printf "%p ")

all: code-vet code-fmt test build/cog-helper

clean:
	$(RM) -rf build

get: $(DEPS)
	go get ./...

test: get
	go test ./...

test_verbose: get
	go test -v ./...

build/cog-helper: $(DEPS)
	mkdir -p build
	go build -o build ./...

code-vet: get $(DEPS)
## Run go vet for this project. More info: https://golang.org/cmd/vet/
	@echo go vet
	go vet $$(go list ./... )

code-fmt: get $(DEPS)
## Run go fmt for this project
	@echo go fmt
	go fmt $$(go list ./... )

lint: $(DEPS)
## Run golint for this project
	@echo golint
	golint $$(go list ./... )

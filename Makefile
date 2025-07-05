SOCIAL_APP=cmd/social/main.go
SOCIAL_BIN=bin/social

generate:
	go generate ./api/generate.go

build: generate
	go build -o ${SOCIAL_BIN} $(SOCIAL_APP)

run: build
	./${SOCIAL_BIN}

test:
	go test -v ./...

all: build
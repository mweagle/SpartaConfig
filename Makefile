.DEFAULT_GOAL=build
.PHONY: build test get run

clean:
	go clean .

generate:
	go generate -x
	@echo "Generate complete: `date`"

vet: generate
	go tool vet -composites=false *.go

get: clean
	go get -u -v ./...

build: get generate vet
	go build .

test:
	go test ./test/...

delete:
	go run main.go delete

provision-default: generate vet
	go run main.go provision --level info --s3Bucket $(S3_BUCKET) --noop

provision-staging: generate vet
	go run main.go provision --level info --s3Bucket $(S3_BUCKET) --tags staging --noop

provision-production: generate vet
	go run main.go provision --level info --s3Bucket $(S3_BUCKET) --tags production --noop

describe-staging: generate vet
	go run main.go --level info describe --out ./graph.html --tags staging

describe-production: generate vet
	go run main.go --level info describe --out ./graph.html --tags production

explore:
	go run main.go --level info explore
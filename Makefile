default: setup build.compose up

setup: setup.deps

setup.deps: setup.deps.go

setup.deps.go:
	go mod tidy

generate:
	go generate ./...

test: test.go

test.go:
	go test -v -shuffle=on ./...

build.compose:
	docker compose build

build.docker:
	docker build . -t fremantle-industries/tabletop

build.bin:
	go build -ldflags="-w -s" -o ./bin/tabletop ./main.go

up:
	docker compose up

down:
	docker compose down

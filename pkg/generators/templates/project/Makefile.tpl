default: setup.deps test build.compose up

setup.deps:
	go mod tidy

test: 
	go test -v -shuffle=on ./...

build.compose:
	docker compose build

build.docker:
	docker build . -t {{.ProjectName}}

build.bin:
	go build -ldflags="-w -s" -o ./bin/{{.ProjectName}} ./main.go

start.bin:
	./bin/{{.ProjectName}}

up:
	docker compose up

down:
	docker compose down

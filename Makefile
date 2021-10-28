default: setup build.compose up

configure:
	./tools/configure

setup: setup.deps

setup.deps: setup.deps.go setup.deps.controlui

setup.deps.go:
	go mod tidy

setup.deps.controlui:
	cd pkg/control/ui && npm install

generate:
	go generate ./...

test: test.go test.controlui

test.go:
	go test -v -shuffle=on ./...

test.controlui:
	cd pkg/control/ui && npm run test

build.compose:
	docker compose build

build.docker:
	docker build . -t fremantle-industries/tabletop

build.bin:
	go build -ldflags="-w -s" -o ./bin/tabletop ./main.go

build.controlui:
	cd pkg/control/ui && npm run build

start.controlui:
	cd pkg/control/ui && npm run dev

start.controlapi:
	go run ./main.go control start

publish.docker:
	docker push fremantle-industries/tabletop

k8s.apply:
	kubectl apply -f deploy/k8s

up:
	docker compose up

down:
	docker compose down

.PHONY = all \
	setup setup/agent setup/worker setup/console setup/explore \
	test test/agent test/worker test/console test/explore \
	run run/agent run/worker run/console run/explore \
	compose.build compose.build/agent compose.build/worker compose.build/console compose.build/explore \
	k3s.cluster.create k3s.cluster.delete \
	k3s.apply k3s.apply/agent k3s.apply/worker k3s.apply/console k3s.apply/explore \
	k3s.delete k3s.delete/agent k3s.delete/worker k3s.delete/console k3s.delete/explore \
	push push/nix.console push/nix.agent push/nix.explore push/nix.worker push/docker.console push/docker.agent push/docker.explore push/docker.worker push/git

all: setup test compose.build k3s.apply

APPS = agent worker console explore

# install
setup: $(addprefix setup/, $(APPS))
setup/%:
	cd $* && bun install

# test
test: $(addprefix test/, $(APPS))
test/%: 
	bun test $*

# run
run: $(addprefix run/, $(APPS))
run/%: 
	bun run $*/index.ts

# docker
compose.build: $(addprefix compose.build/, $(APPS))
compose.build/%:
	docker compose build $*

# gui
open%: OPEN_HOST = 127.0.0.1.traefik.me
open%: GUIS = console explore agent worker grafana qryn redpanda minio-console traefik
open: open/console
open/%:
	@open https://$*.${OPEN_HOST}

# k3s
APP_RESOURCES = namespace traefik
INFRA_RESOURCES = minio

k3s.cluster.create:
	# k3d cluster create tabletop --config k3s/cluster.yaml --registry-config k3s/registries.yaml
	k3d cluster create tabletop --config k3s/cluster.yaml
k3s.cluster.delete:
	k3d cluster delete tabletop

k3s.apply: $(addprefix k3s.apply/, $(APP_RESOURCES)) $(addprefix k3s.apply/, $(INFRA_RESOURCES)) $(addprefix k3s.apply/, $(APPS))
k3s.apply/namespace:
	kubectl apply -f k3s/resources/namespace.yaml
k3s.apply/traefik:
	kubectl apply -f k3s/resources/traefik.yaml
k3s.apply/%:
	kubectl apply -f k3s/resources/$*/*.yaml

k3s.delete: $(addprefix k3s.delete/, $(APPS)) $(addprefix k3s.delete/, $(INFRA_RESOURCES)) $(addprefix k3s.delete/, $(APP_RESOURCES))
k3s.delete/namespace:
	kubectl delete -f k3s/namespace.yaml
k3s.delete/traefik:
	kubectl delete -f k3s/traefik.yaml
k3s.delete/%:
	kubectl delete -f k3s/$*/*.yaml

# artifacts
BUILD_ARTIFACTS = console agent explore worker

build: $(addprefix build/, $(BUILD_ARTIFACTS))
build/%:
	cd $* && bun build --outfile $* index.ts --compile

# PUSH_ARTIFACTS = nix.console nix.agent nix.explore nix.worker docker.console docker.agent docker.explore docker.worker git
#
# push: $(addprefix push/, $(PUSH_ARTIFACTS))
# push/nix.console:
# 	echo "TODO: push/nix.console"
# push/nix.agent:
# 	echo "TODO: push/nix.agent"
# push/nix.explore:
# 	echo "TODO: push/nix.explore"
# push/nix.worker:
# 	echo "TODO: push/nix.worker"
# push/docker.console:
# 	echo "TODO: push/docker.console"
# push/docker.agent:
# 	echo "TODO: push/docker.agent"
# push/docker.explore:
# 	echo "TODO: push/docker.explore"
# push/docker.worker:
# 	echo "TODO: push/docker.worker"
# push/git:
# 	git push origin HEAD

.DEFAULT_GOAL := build-all

export PROJECT := "girc-ui"
export PACKAGE := "github.com/lrstanley/girc-ui"
export COMPOSE_PROJECT := "girc-ui"
export USER := $(shell id -u)
export GROUP := $(shell id -g)
export LICENSE_IGNORE := "wailsjs"

license:
	curl -sL https://liam.sh/-/gh/g/license-header.sh | bash -s

prepare: clean node-prepare go-prepare
	@echo

build-all: prepare
	wails build \
		-clean \
		-platform windows/amd64,linux/amd64 \
		-ldflags "-s -w" \
		-tags netgo,osusergo,static_build \
		-nsis \
		-webview2 embed
		# -upx \
		# -upxflags "--best --lzma" \

up: node-upgrade-deps go-upgrade-deps
	@echo

clean:
	/bin/rm -rfv "frontend/dist/*" "build/"

docker:
	docker compose \
		--project-name ${COMPOSE_PROJECT} \
		--file docker-compose.yaml \
		up \
		--remove-orphans \
		--build \
		--timeout 0 ${COMPOSE_ARGS}

docker-clean:
	docker compose \
		--project-name ${COMPOSE_PROJECT} \
		--file docker-compose.yaml \
		down \
		--volumes \
		--remove-orphans \
		--rmi local --timeout 1

# frontend
node-fetch:
	command -v pnpm >/dev/null >&2 || npm install \
		--no-audit \
		--no-fund \
		--quiet \
		--global pnpm
	cd frontend/ && pnpm install --silent

node-upgrade-deps:
	cd frontend/ && pnpm up -i

node-prepare: license node-fetch
	@echo

# backend
go-prepare: license go-fetch
	go generate -x ./...

go-fetch:
	go mod download
	go mod tidy

go-upgrade-deps:
	go get -u ./...
	go mod tidy

go-upgrade-deps-patch:
	go get -u=patch ./...
	go mod tidy

debug: license
	wails dev

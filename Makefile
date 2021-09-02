include linting.mk

.PHONY: build_ui
build_ui:
	cd ui; \
	yarn install; \
	NEXT_TELEMETRY_DISABLED=1 yarn run export

local_run:
	go run main.go

deps:
	go mod tidy
	go mod vendor
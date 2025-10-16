REGISTRY =  975050366595.dkr.ecr.ap-southeast-3.amazonaws.com/provider_mock
BUILD_DIR = build
VERSION = $(shell git rev-parse --short HEAD)
COMMIT_HASH ?= $(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE ?= $(shell date +%FT%T%z)
ARGS ?=
# remove debug info from the binary & make it smaller
LDFLAGS += -s -w

.PHONY: build
build:
	env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags "${LDFLAGS}" -o ${BUILD_DIR}/provider_mock .


.PHONY: build-container-image
build-container-image:
	docker build -f Dockerfile -t $(REGISTRY):latest .
	echo "::set-output name=image::$(REGISTRY):latest"


.PHONY: push-container-image
push-container-image:
	docker push  $(REGISTRY):latest

.PHONY: run-local stop-local

run-local:
	@echo "Building binary..."
	go build -o app
	@echo "Starting app on 0.0.0.0..."
	nohup ./app > app_log.log 2>&1 & echo $$! > app.pid
	@echo "App running. PID: $$(cat app.pid), Logs: app_log.log"

stop-local:
	@if [ -f app.pid ]; then \
		echo "Stopping app (PID: $$(cat app.pid))..."; \
		kill $$(cat app.pid) && rm app.pid; \
	else \
		echo "No app.pid file found. Is the app running?"; \
	fi


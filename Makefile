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


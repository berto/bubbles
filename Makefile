PID      = /tmp/bubbles.pid
GO_FILES = $(wildcard *.go)
dev:
	@make restart
	@fswatch -o . | xargs -n1 -I{}  make restart || make kill
kill:
	@kill `cat $(PID)` || true
restart:
	@make kill
	@go run $(GO_FILES) & echo $$! > $(PID)

.PHONY: serve dev restart kill

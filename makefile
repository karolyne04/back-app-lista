launch:
	set -a && . $(PWD)/.env && set +a && go run cmd/api/main.go

.PHONY: launch
## test: run unit tests
test:
	@go test -v -vet=off ./...

## lint: run linters checking
lint:
	@golangci-lint run --print-issued-lines=false --out-format=colored-line-number --issues-exit-code=1


.DEFAULT_GOAL=help
help: Makefile
	@echo "Usage: make [command]"
	@echo "Commands:"
	@sed -n "s/^##//p" $< | column -t -s ":" |  sed -e "s/^/ /"

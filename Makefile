MOCK_SERVER_ADDRESS=localhost:1080
MOCK_SERVER_RESET_URL=http://$(MOCK_SERVER_ADDRESS)/mockserver/reset
swag:
	swag init -g ./cmd/api/main.go -o ./docs

.PHONY: swag

wire:
	wire ./internal/wired/wired.go

.PHONY: wire

unit-test:
	 go test ./internal/... -coverpkg=./internal/...  -covermode=atomic -coverprofile coverage.out; \
 	 go tool cover -func coverage.out | grep total; \
 	 rm -r coverage.out

.PHONY: unit-test

test:
	@docker-compose -f docker-compose.yml up -d
	@sleep 1 && \
	go test ./... -coverpkg=./internal/...  -covermode=atomic -coverprofile coverage.out;\
	go tool cover -func coverage.out | grep total; \
    rm -r coverage.out
	@docker-compose -f docker-compose.yml down

.PHONY: test

lint:
	golangci-lint run --fix -v

.PHONY: lint

format:
	go fmt ./internal/...

.PHONY: format

reset-mock:
	curl -X PUT ${MOCK_SERVER_RESET_URL}

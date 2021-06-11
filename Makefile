test:
	@docker-compose -f docker-compose.yml up --build -d
	@sleep 7 && \
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

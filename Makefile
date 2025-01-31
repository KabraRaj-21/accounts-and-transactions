.PHONY: test-unit test-functional build-and-run

# Run unit tests with race detector, verbose output, and generate coverage report
test-unit:
	go test -v -race -cover ./... -tags=unit -coverpkg=./... -coverprofile=coverage-unit.out
	go tool cover -html=coverage-unit.out -o coverage-unit.html
	@echo "Unit test coverage report: coverage-unit.html"

# Run functional tests (SLT) with race detector, verbose output, and generate coverage report
test-functional:
	go test -v -race -cover ./... -tags=slt -coverpkg=./... -coverprofile=coverage-functional.out
	go tool cover -html=coverage-functional.out -o coverage-functional.html
	@echo "Functional test coverage report: coverage-functional.html"

# Build and start the service using Docker Compose
start-service:
	docker-compose up --build

stop-service:
	docker-compose down


ALL_PACKAGES= $(shell go list ./... | grep -v -e "repository" -e "mocks" -e "config")
ALL_PACKAGES_TESTS=$(shell go list ./...| grep -v "vendor")

test:
	$(foreach pkg, $(ALL_PACKAGES),\
	go test -v -race $(pkg);)
	
# Coverage with HTML output
cover-html:
	@echo "mode: count" > coverage-all.out

	$(foreach pkg, $(ALL_PACKAGES),\
	go test -coverprofile=coverage.out -covermode=count $(pkg);\
	tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out
	rm coverage.out coverage-all.out

run-api:
	go run cmd/api/*.go

run-consumer:
	go run cmd/consumer/*.go

run-scheduler:
	go run cmd/scheduler/*.go


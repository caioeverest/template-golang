envfile:=.env.local
ifeq ($(shell test ! -f .env.local && echo -n yes),yes)
    envfile=env.example
endif

include $(envfile)
export $(shell sed 's/=.*//' $(envfile))
export ENVIRONMENT=development
export APP_NAME={{.RepositoryName}}
export APP_VERSION=$(shell git branch --show-current | cut -d '/' -f2)

test: test-base test-clean
test-report: test-report.html test-clean
test-coverage: test-report.cover test-clean

test-base:
	@go test -race -failfast -cover ./app/...

test-report.html:
	@go test -covermode=count -coverprofile=report.out.tmp ./app/...
	cat report.out.tmp | grep -v "mock_*" > report.out
	@go tool cover -html=report.out

test-report.cover:
	@go test -race -coverpkg=./app/... -coverprofile=profile.cov.tmp ./app/...
	cat profile.cov.tmp | grep -v "mock_*" > profile.cov
	@go tool cover -func profile.cov

.PHONY: test-clean
test-clean:
	if [ -f .app/repository/.*.sqlite.db ]; then rm ./app/repository/.*.sqlite.db; fi
	if [ -f profile.cov ]; then rm profile.cov*; fi
	if [ -f report.out ]; then rm report.out*; fi

#To use go lint you must install `go get -u github.com/golangci/golangci-lint/cmd/golangci-lint`
lint:
	@golangci-lint run -E golint -E bodyclose ./...

.PHONY: clean
clean:
	@rm -rf bin/

build: clean
	@go build -o bin/${APP_NAME} main.go

run:
	@go run main.go

docker-clean:
	@docker rm {{.Author}}/${APP_NAME}:${APP_VERSION}
	@docker rmi {{.Author}}/${APP_NAME}:${APP_VERSION}

docker-build:
	@docker build -t {{.Author}}/${APP_NAME}:${APP_VERSION} .

# To use mockery you must install `go get -u go get github.com/vektra/mockery/v2/.../`
update-mocks:
	@mockery --all --inpackage --case=underscore

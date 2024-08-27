.EXPORT_ALL_VARIABLES:

GOBIN          = $(PWD)/.bin
GO 		       = go
LINT_VERSION   =1.51.1
LINT_IMAGE     =golangci/golangci-lint:v${LINT_VERSION}-alpine
LINT_FLAGS     =--timeout=10m0s

service_name=go-safe

.PHONY: help
help: ## - Show this help message
	@printf "\033[32m\xE2\x9c\x93 usage: make [target]\n\n\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: gen-abis
gen-abis:
	abigen --abi=./contracts/abis/safe.json --pkg=safe --out=contracts/safe/safe.go
	abigen --abi=./contracts/abis/wallet_registry.json --pkg=walletregistry --out=contracts/walletregistry/wallet_registry.go
	abigen --abi=./contracts/abis/multicall3.json --pkg=multicall --out=contracts/multicall/multicall.go

build: build-common ## - build
	@ $(GO) build ./...

.PHONY: build-common
build-common: ## - execute build common tasks clean and mod tidy
	@ $(GO) version
	@ $(GO) clean
	@ $(GO) mod tidy && $(GO) mod download
	@ $(GO) mod verify

.PHONY: test
test: ## - execute go test command
	@ go test -v -cover ./...

.PHONY: scan
scan: ## - execute GOSEC static code analysis
	@ gosec -fmt=sarif -out=$(service_name).sarif -exclude-dir=test -exclude-dir=bin -severity=medium ./... | 2>&1
	@ echo ""
	@ cat ./$(service_name).sarif

test-coverage: ## - execute go test command with coverage
	@ mkdir -p .coverage && mkdir -p .report
	@ go test -json -v -cover -covermode=atomic -coverprofile=.coverage/cover.out ./... > .report/report.out

.PHONY: sonar-scan-local
sonar-scan-local: test-coverage ## - start sonar qube locally with docker (you will need docker installed in your machine)
	@ $(SHELL) .scripts/sonar-start.sh
	@ echo "login with user: admin pwd: 1234"

.PHONY: sonar-stop
sonar-stop: ## - stop sonar qube docker container
	@ docker stop sonarqube

.PHONY: ci-lint
ci-lint: ## - runs golangci-lint
	@ golangci-lint run -v ${LINT_FLAGS}

.PHONY: ci-lint-docker
ci-lint-docker: ## - runs golangci-lint with docker container
	@ docker run --rm -v "$(shell pwd)":/app -w /app ${LINT_IMAGE} golangci-lint run ${LINT_FLAGS}

.PHONY: lint
lint: ## Run linters
	$(info $(M) running linters...)
	golangci-lint run --timeout 5m0s ./...
test:
	go test -short

e2e-test:
	./scripts/e2e-test $(authKey)

.PHONY: test e2e-test docs

.PHONY: lint
lint:
	go mod tidy -v
	goimports-reviser -format -set-alias -project-name github.com/solarhell/deepl -rm-unused ./...
	go vet ./...
	staticcheck ./...
	nilaway ./...
	golangci-lint run

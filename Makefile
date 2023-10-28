go.mod:
	go mod tidy

.PHONY: test
test:
	go test -cover -count 1 ./...

.PHONY: fuzz
fuzz:
	go test -fuzztime 10s  -fuzz .

.PHONY: vet
vet:
	go vet ./...

.PHONY: errcheck
errcheck:
	errcheck ./...

.PHONY: staticcheck
staticcheck:
	staticcheck -checks="all,-ST1000" ./...

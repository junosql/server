.PHONY: test

test:
	go test -v $(shell find . -name "*_test.go")

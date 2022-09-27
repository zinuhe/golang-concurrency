# .PHONY: testConcurrency01
# testConcurrency01:
# 	go test ./... -test.run=TestMainConcurrency01

test02:
	go test -run TestMainConcurrency02

test022:
	go test example02_test #GOROOT calling package

test:
	go test ./...

test_full:
	go test -v ./...

test_coverage:
	go test ./... -cover -v

clean:
	go clean

dep:
	go mod download

vet:
	go vet

lint:
	golangci-lint run --enable-all
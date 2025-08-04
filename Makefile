PKG=./...

.PHONY: all build run test fmt tidy clean

all: test

test:
	go test -v $(PKG)

testname:
	go test -v -run '$(NAME)' $(PKG)

fmt:
	go fmt $(PKG)

tidy:
	go mod tidy

clean:
	go clean -testcache

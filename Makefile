default: test

ci: test test_coveralls

deps:
    go get -v -d ./...
    # For syntax check
    go get golang.org/x/tools/cmd/vet

test: deps
    go test -v ./... -timeout=15s
    # Syntax check
    go vet -v ./...


# For Coveralls.io
deps_coveralls:
    go get github.com/axw/gocov/gocov
    go get github.com/mattn/goveralls
    if ! go get github.com/golang/tools/cmd/cover; then go get golang.org/x/tools/cmd/cover; fi

test_coveralls: desp_coveralls
    $HOME/gopath/bin/goveralls -service=travis-ci


.PHONY: default ci before_deps deps test
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
	go get github.com/mattn/goveralls
	go get golang.org/x/tools/cmd/cover
	go get github.com/wadey/gocovmerge

test_coveralls: deps_coveralls
	find . -type d -regex "^./[^.].*" | xargs -I {} go test -v -covermode=count -coverprofile=.profiles/{}_coverage.out {}
	gocovmerge .profiles/* > coverage.out
	@${HOME}/gopath/bin/goveralls -coverprofile=coverage.out -service=travis-ci -repotoken ${COVERALLS_TOKEN}
	@#${HOME}/gopath/bin/goveralls -service=travis-ci -repotoken ${COVERALLS_TOKEN} ./...
	@#go test -v -covermode=count -coverprofile=coverage.out ./...

.PHONY: default ci before_deps deps test

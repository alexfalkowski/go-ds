.PHONY: vendor

lint: dep
	golangci-lint run

fix-lint: dep
	golangci-lint run --fix

specs: dep
	go test -race -mod vendor -v -covermode=atomic -coverpkg=./... -coverprofile=test/profile.cov ./...

coverage:
	go tool cover -html=test/profile.cov

download:
	go mod download

tidy:
	go mod tidy

vendor:
	go mod vendor

dep: download tidy vendor

goveralls:
	goveralls -coverprofile=test/profile.cov -service=circle-ci -repotoken=vuNtS4222HStiAwy29cxZ3nDHElAsk65s

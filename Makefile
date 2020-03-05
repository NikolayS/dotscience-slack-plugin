install:
	cd cmd/ds-slack-plugin && go install

run: install
	ds-slack-plugin

test:
	go get github.com/mfridman/tparse
	go test -json -v `go list ./... | egrep -v /tests` -cover | tparse -all -smallscreen

image-push:
	docker build -t dotscience/dotscience-slack-plugin:latest -f Dockerfile .
	docker push dotscience/dotscience-slack-plugin:latest
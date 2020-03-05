install:
	cd cmd/ds-slack-plugin && go install

run: install
	ds-slack-plugin

image-push:
	docker build -t dotscience/dotscience-slack-plugin:latest -f Dockerfile .
	docker push dotscience/dotscience-slack-plugin:latest
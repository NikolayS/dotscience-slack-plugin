FROM golang:1.13.8-alpine3.11 AS build-env

WORKDIR /usr/local/go/src/github.com/dotmesh-io/dotscience-slack-plugin
COPY . /usr/local/go/src/github.com/dotmesh-io/dotscience-slack-plugin

RUN apk update && apk upgrade
RUN cd cmd/ds-slack-plugin && go install

FROM alpine:latest
LABEL "com.dotscience.dotscience-slack-plugin"="true"
COPY --from=build-env /usr/local/go/bin/ds-slack-plugin /bin/ds-slack-plugin

CMD ["ds-slack-plugin"]

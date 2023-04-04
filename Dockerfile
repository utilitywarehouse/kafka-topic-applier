FROM golang:1.20-alpine AS build
WORKDIR /go/src/github.com/utilitywarehouse/kafka-topic-applier
COPY . /go/src/github.com/utilitywarehouse/kafka-topic-applier
RUN apk --no-cache add git \
      && go get ./... \
      && CGO_ENABLED=0 go build -o /kafka-topic-applier .

FROM alpine:3.17
RUN apk add --no-cache ca-certificates
COPY --from=build /kafka-topic-applier /kafka-topic-applier
CMD [ "/kafka-topic-applier" ]

FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/my-user-settings-service
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/my-user-settings-service /go/src/my-user-settings-service


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/my-user-settings-service /usr/local/bin/my-user-settings-service
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["my-user-settings-service"]

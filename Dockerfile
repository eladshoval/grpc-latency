ARG GO_VERSION=1.18
# STAGE 1 - Build Go artifact
FROM golang:${GO_VERSION}-buster AS builder
#ARG user
#ARG password
#ENV GOPROXY="https://${user}:${password}@identiq.jfrog.io/artifactory/api/go/go-modules"
ENV GOSUMDB=off

WORKDIR /go/src/app

COPY . /go/src/app
RUN go get -d -v ./...
RUN mkdir -p /go/bin/app
RUN go build -o /go/bin/app ./...

# STAGE 2 - Copy the artifact from builder stage into minimal runnable image
FROM debian:buster

#Add tcpdump for debug purposes
RUN apt update && \
    apt install -yqq tcpdump
# Copy artifact
COPY --from=builder /go/bin/app /
# Copy configuration
#COPY --from=builder /go/src/app/conf /conf

# adjust environment for production purposes
ENV ENVIRONMENT=production
ENTRYPOINT ["/pokemon"]

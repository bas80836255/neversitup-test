############################################################################################
# Build executable binary
############################################################################################
# Always pull image by digest
# Always pull trusted image
# golang:1.19.2-alpine3.16
FROM golang@sha256:f3e683657ddf73726b5717c2ff80cdcd9e9efb7d81f77e4948fada9a10dc7257 as build-stage

# Install git and mercurial
RUN apk update && \
    apk --no-cache add git mercurial gcc g++ && \
    rm -rf /var/cache/apk/*

# Define environment variables
ENV APPNAME=service \
    CGO_ENABLED=0 \
    WRKDIR=/app \
    GO111MODULE=on \
    USER=appuser \
    UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR $WRKDIR

# copy and fetch external module
COPY go.mod go.sum ./
RUN go mod download && \
    go mod verify

# copy and build application
# removing debug informations and compile only for linux target and disabling cross compilation
COPY . .
RUN GOOS=linux GOARCH=amd64 go test ./...
RUN GOOS=linux GOARCH=amd64 go build -buildvcs=false -ldflags="-w -s" -o $APPNAME

############################################################################################
# Build image
############################################################################################
# alpine:3.16
FROM alpine:3.16 as server-stage

# Define environment variables
ENV APPNAME=service
ENV APPPATH=/app/$APPNAME
ENV WRKDIR=/build

WORKDIR $WRKDIR

# Copy builded code from the builder
COPY --from=build-stage $APPPATH $WRKDIR

# Copy application config
COPY config.yaml $WRKDIR/config.yaml

# Copy package.json file
COPY package.json $WRKDIR/package.json

USER 1001

EXPOSE 8080

ENTRYPOINT ["/build/service"]
FROM golang:1.22.0-bullseye AS dev

ENV CGO_ENABLED 0
WORKDIR /src/app
COPY . .

RUN apt-get update && apt-get upgrade -y

RUN go install github.com/cosmtrek/air@v1.51.0
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/go-delve/delve/cmd/dlv@v1.22.1
RUN go install github.com/google/wire/cmd/wire@v0.6.0
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2
RUN go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
RUN go install go.uber.org/mock/mockgen@latest

EXPOSE 9000
EXPOSE 2345

CMD ["air", "-c", ".air.toml"]

FROM golang:1.22.0-bullseye AS build
WORKDIR /go/src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN apt-get update && apt-get upgrade -y
RUN go install github.com/google/wire/cmd/wire@v0.6.0
RUN go generate -x -tags wireinject wire/wire.go
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o app .

FROM gcr.io/distroless/base
WORKDIR /go/src

COPY --from=build /go/src/app .

USER nonroot:nonroot

EXPOSE 9000
ENV SERVER_PORT=9000

CMD ["./app"]


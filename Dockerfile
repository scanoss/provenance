FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go generate ./pkg/cmd/server.go
RUN go build -o ./scanoss-provenance ./cmd/server

FROM debian:buster-slim

WORKDIR /app
 
COPY --from=build /app/scanoss-provenance /app/scanoss-provenance

EXPOSE 50051

ENTRYPOINT ["./scanoss-provenance"]
#CMD ["--help"]

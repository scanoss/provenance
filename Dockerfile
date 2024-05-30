FROM golang:1.19 as build

# UNPUBLISHED PAPI
RUN apt update && apt install -y protobuf-compiler

RUN git clone -b SCP118-AddProvenanze https://github.com/scanoss/papi /papi

WORKDIR /papi

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
     go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

RUN go get github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor@v2.18.0

RUN go install \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc

RUN go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

RUN make build_go
# FINISH UNPUBLISHED API

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

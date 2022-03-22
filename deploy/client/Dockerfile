FROM golang:1.17-buster as build
WORKDIR /app

RUN apt update -y && apt install -y protobuf-compiler
COPY . ./
RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    swag init --output docs/client --generalInfo cmd/client/main.go  && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest  && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest  && \
    protoc --go_out=internal --go-grpc_out=internal proto/storage.proto  && \
    CGO_ENABLED=0 go mod tidy && go build -o bin/client cmd/client/main.go

FROM gcr.io/distroless/base-debian10 AS app
ENV TZ=Europe/London
WORKDIR /
COPY --from=build /app/bin/client /client
COPY configs/client.yaml ./configs/client.yaml
ENTRYPOINT ["/client"]

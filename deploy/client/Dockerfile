FROM golang:1.17-buster as build
WORKDIR /app
COPY . ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --output docs/client --generalInfo cmd/client/main.go
RUN CGO_ENABLED=0 go mod tidy && go build -o bin/client cmd/client/main.go

FROM gcr.io/distroless/base-debian10 AS app
ENV TZ=Europe/London
WORKDIR /
COPY --from=build /app/bin/client /client
COPY configs/client.yaml ./configs/client.yaml
ENTRYPOINT ["/client"]

FROM golang:1.17-buster as build
WORKDIR /app
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 go mod tidy && go build -o bin/server cmd/server/main.go

FROM gcr.io/distroless/base-debian10 AS app
ENV TZ=Europe/London
WORKDIR /
COPY --from=build /app/bin/server /server
COPY configs/server.yaml ./configs/server.yaml
ENTRYPOINT ["/server"]

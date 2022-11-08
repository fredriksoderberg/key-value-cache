# Build
FROM golang:1.16-alpine AS build

WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN CGO_ENABLED=0 go build -o /key-value-cache

# Deploy

FROM alpine:3.14

WORKDIR /

COPY --from=build /key-value-cache /key-value-cache

EXPOSE 8080

ENTRYPOINT ["/key-value-cache"]

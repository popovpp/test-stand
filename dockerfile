FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o main main.go resolvers.go

FROM alpine

# RUN apk add --no-cache --upgrade bash

COPY --from=builder /build/docs /build/docs

WORKDIR /build

COPY --from=builder /build/main /build/main

CMD ["./main"]

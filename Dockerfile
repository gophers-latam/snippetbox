FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum /app/

COPY cmd /app/cmd

COPY internal /app/internal

COPY ui /app/ui

RUN go mod download 

RUN go build -o snippetbox /app/cmd/web/

FROM ubuntu:24.10 AS produccion

WORKDIR /app

COPY --from=builder /app/snippetbox .

COPY --from=builder /app/ui /app/ui

EXPOSE 4000

CMD ["./snippetbox"]

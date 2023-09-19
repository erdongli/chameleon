FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

ARG app
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -a -trimpath -o chameleon ./cmd/chameleon/main.go

FROM alpine

ARG app
COPY --from=builder /app/chameleon ./chameleon

ENTRYPOINT ["./chameleon"]

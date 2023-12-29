FROM golang:alpine AS builder

WORKDIR /build
COPY ./ ./
RUN go build cmd/main.go

FROM alpine

WORKDIR /build
COPY --from=builder /build/main /build/main
COPY ./config/cfg.yml /build/config/cfg.yml
RUN chmod +x main
CMD ["/build/main"]
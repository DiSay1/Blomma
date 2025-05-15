FROM golang:1.18-alpine3.17 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build

FROM alpine:3.17
LABEL org.opencontainers.image.source=https://github.com/DiSay1/Blomma
WORKDIR /app
COPY --from=builder /app/Blomma /usr/local/bin/
COPY --from=builder /app/web /app/web 
EXPOSE 8080
ENTRYPOINT [ "/usr/local/bin/Blomma" ]

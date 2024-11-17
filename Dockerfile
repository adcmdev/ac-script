FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ac

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/ac .

ENTRYPOINT ["./ac"]

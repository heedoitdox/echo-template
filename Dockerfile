
FROM golang:alpine as builder

RUN apk add --no-cache ca-certificates

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY echo-template .
COPY .env .

EXPOSE 80

ENTRYPOINT [ "./echo-template" ]
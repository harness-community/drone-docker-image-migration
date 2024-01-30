FROM alpine:latest as builder

RUN apk --no-cache add ca-certificates

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV GODEBUG netdns=go

ADD release/linux/amd64/drone-migrate /bin/

ENTRYPOINT ["/bin/drone-migrate"]
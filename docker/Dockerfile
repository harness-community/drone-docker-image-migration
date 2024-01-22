FROM alpine:latest

ENV GODEBUG netdns=go
# Install Skopeo
RUN apk --no-cache add ca-certificates curl && \
    apk add --no-cache skopeo

RUN apk add --no-cache skopeo

ADD release/linux/amd64/drone-migrate /bin/
ENTRYPOINT ["/bin/drone-migrate"]
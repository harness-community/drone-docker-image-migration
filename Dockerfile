FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates curl && \
    apk add --no-cache docker

COPY script.sh /app/

CMD ["sh", "/app/script.sh"]
FROM alpine:3.6

RUN apk --no-cache add bash jq

COPY j2y /usr/local/bin/j2y
COPY docker-entrypoint.sh /docker-entrypoint.sh

ENTRYPOINT ["/docker-entrypoint.sh"]
FROM alpine:3.5

RUN apk --no-cache add bash jq

COPY json2yaml /usr/local/bin/json2yaml

ENTRYPOINT ["/docker-entrypoint.sh"]
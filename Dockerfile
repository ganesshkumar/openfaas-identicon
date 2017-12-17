# Builder
FROM golang:1.9rc2-alpine as BUILDER

RUN apk --no-cache add make
WORKDIR /go/src/github.com/ganesshkumar/identicon
COPY . /go/src/github.com/ganesshkumar/identicon

RUN make

# Runner
FROM alpine

COPY --from=builder /go/bin/identicon /usr/bin

RUN apk --no-cache add curl \
    && curl -sSL https://github.com/alexellis/faas/releases/download/0.6.0/fwatchdog > /usr/bin/fwatchdog \
    && chmod +x /usr/bin/fwatchdog \
    && chmod +x /usr/bin/identicon \
    && apk del curl --no-cache

ENV afterburn=true
ENV mode=afterburn
ENV fprocess "/usr/bin/identicon"

HEALTHCHECK --interval=1s CMD [ -e /tmp/.lock ] || exit 1

CMD [ "/usr/bin/fwatchdog"]

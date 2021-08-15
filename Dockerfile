FROM node:lts-alpine as node-builder

WORKDIR /build

COPY . /build

RUN cd /build/assets \
    && yarn && yarn build

FROM golang:alpine as go-builder

WORKDIR /build

COPY --from=node-builder /build /build

RUN go build -o /build/picbed

FROM alpine:latest

COPY --from=go-builder /build/picbed /app/picbed

RUN mkdir -p /etc/picbed && touch /etc/picbed/config.yml \
    && echo "imgSrc: /data" >> /etc/picbed/config.yml

VOLUME /etc/picbed/config.yml

VOLUME /data

EXPOSE 8080

ENTRYPOINT [ "/app/picbed" ]

CMD ["-c", "/etc/picbed/config.yml" ]
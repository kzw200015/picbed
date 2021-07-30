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

COPY --from=go-builder /build/config.example.yml

VOLUME /data

ENTRYPOINT [ "/app/picbed" ]

CMD [ "-s", "/data" ]
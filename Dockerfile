# Build backend binary file
FROM golang:1.21.0-alpine3.17 AS be-builder
ARG RELEASE_BUILD
ENV RELEASE_BUILD=$RELEASE_BUILD
WORKDIR /go/src/github.com/f97/n
COPY . .
RUN docker/backend-build-pre-setup.sh
RUN apk add git gcc g++ libc-dev
RUN ./build.sh backend

# Build frontend files
FROM node:20.5.1-alpine3.17 AS fe-builder
ARG RELEASE_BUILD
ENV RELEASE_BUILD=$RELEASE_BUILD
WORKDIR /go/src/github.com/f97/n
COPY . .
RUN docker/frontend-build-pre-setup.sh
RUN apk add git
RUN ./build.sh frontend

# Package docker image
FROM alpine:3.18.3
LABEL maintainer="MaysWind <i@mayswind.net>"
RUN addgroup -S -g 1000 gofire && adduser -S -G gofire -u 1000 gofire
RUN apk --no-cache add tzdata
COPY docker/docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh
RUN mkdir -p /gofire && chown 1000:1000 /gofire \
  && mkdir -p /gofire/data && chown 1000:1000 /gofire/data \
  && mkdir -p /gofire/log && chown 1000:1000 /gofire/log
WORKDIR /gofire
COPY --from=be-builder --chown=1000:1000 /go/src/github.com/mayswind/gofire/gofire /gofire/gofire
COPY --from=fe-builder --chown=1000:1000 /go/src/github.com/mayswind/gofire/dist /gofire/public
COPY --chown=1000:1000 conf /gofire/conf
COPY --chown=1000:1000 templates /gofire/templates
COPY --chown=1000:1000 LICENSE /gofire/LICENSE
USER 1000:1000
EXPOSE 8080
ENTRYPOINT ["/docker-entrypoint.sh"]

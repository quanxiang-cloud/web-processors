FROM alpine as certs
RUN apk update && apk add ca-certificates


FROM node:16.13.2 as web
WORKDIR /stc
COPY ./packages/stc .
RUN npm run prepkg && npm run pkg


FROM golang:1.16.6-alpine3.14 AS builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build  -o ./bin/dispatch -mod=vendor -ldflags='-s -w'  -installsuffix cgo ./dispatch/cmd/. && \
    CGO_ENABLED=0 go build  -o ./bin/scripts/fileserver -mod=vendor -ldflags='-s -w'  -installsuffix cgo ./dispatch/pkg/command/fileserver/. && \
    CGO_ENABLED=0 go build  -o ./bin/scripts/persona -mod=vendor -ldflags='-s -w'  -installsuffix cgo ./dispatch/pkg/command/persona/.



FROM scratch
COPY --from=certs /etc/ssl/certs /etc/ssl/certs

WORKDIR /dispatch
COPY --from=builder ./build/bin .
COPY --from=web  ./stc/dist/evolution ./scripts/evolution

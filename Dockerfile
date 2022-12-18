# Builder
FROM golang:1.19.4-alpine3.17 as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /cli

COPY . .

RUN make clean engine

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /cli 

WORKDIR /cli 

EXPOSE 9090

COPY --from=builder /cli/engine /cli

CMD /cli/engine
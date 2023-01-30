ARG BUILDER_BASE=golang:1.19.5-alpine3.17
FROM ${BUILDER_BASE} AS builder
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache git make gcc build-base

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /irisscaffold

COPY . .

RUN go mod tidy

# Build the SMA executable.
RUN make build_greet


FROM alpine3.17

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata

ENV APP_PORT=8080

EXPOSE $APP_PORT

COPY --from=builder /irisscaffold/greet /bin/
COPY --from=builder /irisscaffold/configs/greet/config.toml.template /etc/greet/config.toml

ENTRYPOINT ["/bin/greet"]
CMD ["-c", "/etc/greet"]

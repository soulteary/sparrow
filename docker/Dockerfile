FROM golang:1.20.0 AS Builder
ENV TZ=Asia/Shanghai
WORKDIR /app
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn
COPY ./pkg ./pkg
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build .


FROM debian:bullseye
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=Builder /app/sparrow /bin/sparrow
ENV TZ=Asia/Shanghai
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
EXPOSE 8080
CMD ["sparrow"]
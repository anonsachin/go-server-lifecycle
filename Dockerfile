ARG GO_VER=1.14.3
ARG ALPINE_VER=3.11

FROM golang:${GO_VER}-alpine${ALPINE_VER}

WORKDIR /go/src/server/
COPY . .

RUN go mod vendor
RUN go build -o server
RUN mv server /go/bin/

EXPOSE 3000
CMD ["server"]
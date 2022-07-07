FROM golang:1.18.3-buster AS builder

WORKDIR /go/src/app
ENV UID=1001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "apps"

RUN mv /etc/localtime localtime.backup
RUN ln -s /usr/share/zoneinfo/Asia/Jakarta /etc/localtime

COPY go.mod \
     go.sum ./ 
RUN go mod download
COPY . ./
RUN go build -o backend-service

FROM ubuntu:20.04
WORKDIR /app
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /go/src/app .
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Jakarta

USER apps:apps
ENTRYPOINT ["./backend-service"]
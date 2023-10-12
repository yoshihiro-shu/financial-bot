FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY apps/notification ./apps/notification
COPY internal ./internal
RUN go build -ldflags="-s -w" -o /app/notification /build/apps/notification/notification.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Tokyo /usr/share/zoneinfo/Asia/Tokyo
ENV TZ Asia/Tokyo

WORKDIR /app
COPY --from=builder /app/notification /app/notification
COPY --from=builder /build/apps/notification/etc/notification.yaml /app/etc/notification.yaml

CMD ["./notification"]
FROM golang:1.16.5-alpine AS builder
RUN mkdir /build
ADD go.mod go.sum hello.go /build/
WORKDIR /build
RUN go build


FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/hello-go /app/
COPY views/ /app/views
WORKDIR /app
CMD ["./hello-go"]


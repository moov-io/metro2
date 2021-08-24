FROM golang:1.17-alpine as builder
WORKDIR /go/src/github.com/moov-io/metro2
RUN apk add -U make
RUN adduser -D -g '' --shell /bin/false moov
COPY . .
RUN go mod download
RUN make build
USER moov

FROM scratch
LABEL maintainer="Moov <support@moov.io>"

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/moov-io/metro2/bin/metro2 /bin/metro2
COPY --from=builder /etc/passwd /etc/passwd

USER moov
EXPOSE 8080
ENTRYPOINT ["/bin/metro2"]
CMD ["web"]

FROM golang:1.22.3-alpine as builder
COPY . /app
WORKDIR /app
EXPOSE 2137
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint

FROM gcr.io/distroless/static-debian11 AS release-stage
WORKDIR /
COPY --chown=nonroot --from=builder /entrypoint /entrypoint
COPY --chown=nonroot --from=builder /app/assets /assets
EXPOSE 2137
USER nonroot:nonroot
ENTRYPOINT ["/entrypoint"]

FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o app

FROM gcr.io/distroless/static-debian12
WORKDIR /
COPY --from=builder /app/app /app
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/app"]


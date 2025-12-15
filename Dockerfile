# Multi-stage build for the Go application
FROM golang:1.25.4 AS builder
WORKDIR /app

# cache deps
COPY go.mod go.sum ./
RUN go mod download

# copy source
COPY . .

# build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /songs .

FROM scratch
COPY --from=builder /songs /songs

EXPOSE 8080
ENV PORT=8080

USER 65532:65532
ENTRYPOINT ["/songs"]

# Builder
FROM golang:alpine as builder
WORKDIR /app

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && \
    apk add --no-cache git ca-certificates && \
    update-ca-certificates

# Add src files
ADD client .
ADD go.mod .
ADD go.sum .

# Fetch dependencies.
RUN go mod download
RUN go mod verify

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/dapr-client client.go

# Runner
FROM scratch

# Import the user and group files from the builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy our static executable.
COPY --from=builder /go/bin/dapr-client /go/bin/dapr-client

EXPOSE 4000
ENTRYPOINT ["/go/bin/dapr-client"]
############################
# STEP 1 build executable binary
############################
FROM golang:1.16.7 as builder

WORKDIR /app
COPY . ./

# Using go mod.
# RUN go mod download
# Build the binary

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -mod=vendor -o ./bin/svc

############################
# STEP 2 build a small image
############################
FROM scratch

# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy our static executable
COPY --from=builder /app/bin/svc /svc
COPY --from=builder /app/migrations ./migrations

EXPOSE $PORT

# Run the svc binary.
ENTRYPOINT ["./svc"]
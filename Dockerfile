############################
# STEP 1 build executable binary
############################
ARG DOCKER_IMAGE_GOLANG="mirror.gcr.io/library/golang:1.17-alpine"
FROM ${DOCKER_IMAGE_GOLANG} as builder
# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates
# Create appuser
RUN adduser -D -g '' appuser
COPY . $GOPATH/src/
WORKDIR $GOPATH/src/
RUN rm -rf $GOPATH/pkg/* $GOPATH/src/go.sum $GOPATH/.git /var/cache/apk/*
ENV GOBIN=$GOPATH/bin
ENV PATH=$GOBIN:$PATH
ENV GO111MODULE=on
# Fetch dependencies.
# Using go mod with go >= 1.11
RUN go mod download
# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' .
############################
# STEP 2 build a small image
############################
FROM scratch
# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable
COPY --from=builder /go/bin/versions-api /go/bin/
COPY --from=builder /go/src/sql/ /sql
# Use an unprivileged user.
USER appuser
# Run the APP_NAME binary.
ENTRYPOINT ["/go/bin/versions-api"]

### TLS root certs and non-root user
FROM ubuntu:focal-20220531@sha256:397930ce1f68f54e10df2fede8e268fe086e2a7e58d83e532fb3abacfa040da7 AS ubuntu

RUN \
  # Note that the lack of a "lock" mechanism for apt dependencies
  # currently prevents a fully reproducible build
  apt-get update \
  && apt-get install -y --no-install-recommends \
  # Install TLS root certificates
  ca-certificates \
  && rm -rf /var/lib/apt/lists/* \
  \
  # Add a non-root user
  && useradd app

# ---

### Release image
FROM scratch

LABEL org.opencontainers.image.source="https://github.com/patrickhoefler/ghtop"

# Copy the TLS certificates for encrypted network communication
COPY --from=ubuntu /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Run as non-root user
COPY --from=ubuntu /etc/passwd /etc/passwd
USER app

# This currently only works with goreleaser
# or if you manually copy the binary into the main project directory
COPY ghtop /

ENTRYPOINT ["/ghtop"]


### SSL certs
FROM ubuntu:latest@sha256:4e4bc990609ed865e07afc8427c30ffdddca5153fd4e82c20d8f0783a291e241 AS ubuntu

# Because there is no "lock" mechanism for apt dependencies,
# this step prevents a fully reproducible build
RUN apt-get update \
  && apt-get install -y --no-install-recommends \
  ca-certificates \
  && rm -rf /var/lib/apt/lists/*

# Needed to run as non-root user
RUN groupadd --gid 1000 ghtop \
  && useradd --uid 1000 --gid ghtop --shell /bin/bash ghtop

# ---

### Release image
FROM scratch

LABEL org.opencontainers.image.source="https://github.com/patrickhoefler/ghtop"

# Copy the TLS certificates for encrypted network communication
COPY --from=ubuntu /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Run as non-root user
COPY --from=ubuntu /etc/passwd /etc/passwd
USER ghtop

# This currently only works with goreleaser
# or if you manually copy the binary into the main project directory
COPY ghtop /
ENTRYPOINT ["/ghtop"]

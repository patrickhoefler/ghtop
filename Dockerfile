
### SSL certs
FROM ubuntu@sha256:1d7b639619bdca2d008eca2d5293e3c43ff84cbee597ff76de3b7a7de3e84956 AS ubuntu

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

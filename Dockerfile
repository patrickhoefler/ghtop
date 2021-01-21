
### SSL certs
FROM ubuntu:latest@sha256:3093096ee188f8ff4531949b8f6115af4747ec1c58858c091c8cb4579c39cc4e AS ubuntu

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

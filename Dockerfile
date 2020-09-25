FROM scratch

LABEL org.opencontainers.image.source="https://github.com/patrickhoefler/ghtop"

COPY ghtop /
ENTRYPOINT ["/ghtop"]

# Build plugin.
FROM kong/go-plugin-tool:2.0.4-alpine-latest AS builder
RUN mkdir -p /tmp/go-plugins/
COPY . /tmp/go-plugins/
RUN cd /tmp/go-plugins/ && \
    apk add make && \
    make all

# Add plugin to Kong.
FROM kong:2.3.3-alpine
RUN mkdir /tmp/go-plugins
COPY --from=builder  /tmp/go-plugins/key-checker /usr/local/bin/key-checker
COPY config.yml /tmp/config.yml
RUN /usr/local/bin/key-checker -dump

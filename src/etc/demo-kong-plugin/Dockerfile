# Build plugin.
FROM kong/go-plugin-tool:2.0.4-alpine-latest

ENV BUILD_DIR /tmp/go-plugins/

WORKDIR $BUILD_DIR

COPY . .
RUN apk add make
RUN make all

# Build Stage
FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-statocratic"
LABEL REPO="https://github.com/vaguelytitled/statocratic"

ENV PROJPATH=/go/src/github.com/vaguelytitled/statocratic

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/vaguelytitled/statocratic
WORKDIR /go/src/github.com/vaguelytitled/statocratic

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/vaguelytitled/statocratic"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/statocratic/bin

WORKDIR /opt/statocratic/bin

COPY --from=build-stage /go/src/github.com/vaguelytitled/statocratic/bin/statocratic /opt/statocratic/bin/
RUN chmod +x /opt/statocratic/bin/statocratic

# Create appuser
RUN adduser -D -g '' statocratic
USER statocratic

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/statocratic/bin/statocratic"]

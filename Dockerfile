FROM alpine:latest
MAINTAINER Telemetry Team <telemetry@heroku.com>

WORKDIR /
COPY dist/spew /spew

CMD ["/spew"]
FROM scratch
MAINTAINER Telemetry Team <telemetry@heroku.com>

COPY dist/spew /app/spew

CMD ["/app/spew"]
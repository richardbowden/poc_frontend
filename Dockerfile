FROM ubuntu:14.04

ADD https://github.com/richardbowden/poc_frontend/releases/download/v0.0.1/poc_frontend_linux_amd64 /poc/poc_front_linux_amd64

RUN chmod +x /poc/poc_front_linux_amd64

EXPOSE 8080 8080

ENTRYPOINT ["/poc/poc_front_linux_amd64"]

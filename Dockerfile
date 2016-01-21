FROM ubuntu:14.04


ADD frontend-svc-* /poc/frontend-svc

RUN chmod +x /poc/frontend-svc && apt-get update && \
apt-get install curl

ADD docker-entrypoint.sh /docker-entrypoint.sh

CMD /docker-entrypoint.sh

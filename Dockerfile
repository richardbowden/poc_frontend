FROM ubuntu:14.04

ENV PORT 8080

ADD frontend-svc-* /poc/frontend-svc

RUN chmod +x /poc/frontend-svc

EXPOSE ${PORT}

ADD docker-entrypoint.sh /docker-entrypoint.sh

CMD /docker-entrypoint.sh

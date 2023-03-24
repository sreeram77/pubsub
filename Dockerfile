FROM alpine:3.17

RUN mkdir -p /file

COPY build/main /main
COPY file /file

RUN chmod +x /main

ENTRYPOINT ["/main"]
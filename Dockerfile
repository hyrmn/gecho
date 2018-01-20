FROM scratch

ENV PORT 7

COPY build/echo-server /
ENTRYPOINT ["/echo-server"]

EXPOSE 7
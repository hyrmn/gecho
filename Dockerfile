FROM scratch

ENV PORT 7000

COPY build/echo-server /
ENTRYPOINT ["/echo-server"]

EXPOSE 7000
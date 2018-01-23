FROM golang as compiler 

RUN CGO_ENABLED=0 go get -a -ldflags '-s' github.com/hyrmn/GoTcpEchoServer

FROM scratch

ENV PORT 7
COPY --from=compiler /go/bin/GoTcpEchoServer ./echo-server

ENTRYPOINT ["./echo-server"]
EXPOSE 7
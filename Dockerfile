FROM golang as compiler
ENV GOOS=linux
ENV CGO_ENABLED=0

COPY *.go /go/src/echo-server/
WORKDIR /go/src/echo-server/
RUN go build -ldflags '-s' -o /bin/echo-server

FROM scratch
ENV PORT 7
COPY --from=compiler /bin/echo-server /
ENTRYPOINT ["./echo-server"]

EXPOSE 7
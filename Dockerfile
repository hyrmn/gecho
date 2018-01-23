FROM golang as compiler
ENV GOBIN=/go/bin
ENV GOOS=linux
ENV CGO_ENABLED=0
COPY main.go main.go
RUN go get -a && go build -ldflags '-s' -o /echo-server main.go

FROM scratch
ENV PORT 7
COPY --from=compiler /echo-server /
ENTRYPOINT ["./echo-server"]

EXPOSE 7
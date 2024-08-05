FROM golang:alpine AS compile
COPY hello-app.go /go
RUN go build hello-app.go

FROM registry.suse.com/suse/sle15:15.6
COPY --from=compile /go/hello-app /
EXPOSE 8080
ENTRYPOINT ["/hello-app"]

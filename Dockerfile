FROM golang:alpine AS builder
COPY src/demo.go /go/
RUN apk update && apk add git && \
# need git installed to run 'go get' on github
    go get gopkg.in/gorethink/gorethink.v3 && \
# statically compile demo.go with all libraries built in
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM scratch
COPY --from=builder /go/app /app
ENTRYPOINT ["/app"]
CMD ["54.174.187.41:28015", "test", "movies"]

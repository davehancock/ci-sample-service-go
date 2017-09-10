FROM scratch

COPY ./ci-sample-service-go /go/bin/ci-sample-service-go

EXPOSE 8080

ENTRYPOINT ["/go/bin/ci-sample-service-go"]

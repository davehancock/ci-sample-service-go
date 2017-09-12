# ci-sample-service-go

A Sample project demonstrating how to build and deploy (to Docker Hub) a simple Go app using Jenkins declarative pipeline.

To build the app locally from the command line:

```shell
env GOOS=linux GOARCH=386 go build && docker build -t test . && docker run -p 8081:8080 test
```

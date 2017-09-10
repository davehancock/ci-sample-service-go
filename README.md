# ci-sample-service-go

```shell
env GOOS=linux GOARCH=386 go build && docker build -t test . && docker run -p 8081:8080 test
```

# HelloCourse microservice


My new Go kit microservice project


This service has been tested to run on Golang v1.6 and v1.7.

To run locally, install the required dependencies into it:

```sh
$ go get ./...
```

You can also execute the sample tests using:

```sh
$ go test -v ./...
```

You can now execute the service as follows:

```sh
$ go run main.go
```

By default, the service will listen on port 8080.  You can test the
service using curl:

```sh
$ curl localhost:8080
{"v":"Hello, World"}
```

---
Created by Atomist. Need Help? <a href="https://join.atomist.com/">Join our Slack team</a>

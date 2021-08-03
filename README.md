# go-vs-js

We want to try out a simple example using Go and Javascript, so that we can compare both solutions on a small scale.

## Frameworks

Golang: [gin](https://gin-gonic.com/)
Javascript: [nest.js](https://nestjs.com/)

## Requirements

* One REST call to `localhost:<port>/getDocuments`.
* Without any parameter, will return all documents in the database `mydatabase` and collection `mycoll`.
* with an optional parameter `name`, the endpoint will filter for documents where the value of the attribute `name` is LIKE the parameter value.
* The REST endpoint should be tested using the frameworks testing method.

## Golang

To init we had to:
```bash
go mod init github.com/mjainta/go-vs-js
go get github.com/gin-gonic/gin
go build main.go
go run main.go
```

Then we can visit [localhost:8080/ping](localhost:8080/ping).

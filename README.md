# go-vs-js

We want to try out a simple example using Go and Javascript, so that we can compare both solutions on a small scale.

## Frameworks

Golang: [gin](https://github.com/gin-gonic/gin)
Javascript: [nest.js](https://nestjs.com/)

## Requirements

* One REST call to `localhost:<port>/getDocuments`.
* Without any parameter, will return all documents in the database `mydatabase` and collection `mycoll`.
* with an optional parameter `name`, the endpoint will filter for documents where the value of the attribute `name` is LIKE the parameter value.
* The REST endpoint should be tested using the frameworks testing method.

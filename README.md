# go-vs-js

We want to try out a simple example using the Gin (Go) and Nest.js (Javascript) frameworks, so that we can compare both solutions on a small scale.

## Frameworks

- [gin](https://gin-gonic.com/)
- ~~[nest.js](https://nestjs.com/)~~
- [express.js](http://expressjs.com/)

## Requirements

* One REST call to `localhost:<port>/documents`.
* Without any parameter, will return all documents in the database `mydatabase` and collection `mycoll`.
* with an optional parameter `name`, the endpoint will filter for documents where the value of the attribute `name` is LIKE the parameter value.
* The REST endpoint should be tested using the frameworks testing method.

# Gin (Golang)

To init we had to:
```bash
# Follow https://gin-gonic.com/docs/quickstart/ and use the curl command to get the main.go
go mod init github.com/mjainta/go-vs-js
go get github.com/gin-gonic/gin
# Builds and runs the app
make run
# Starts unittests
make test
```

Then we can visit [localhost:8080/ping](localhost:8080/ping).

## Using mongodb

```bash
go get go.mongodb.org/mongo-driver/mongo
```

# ExpressJs (JS)

To init we had to:
```bash
# Follow http://expressjs.com/en/starter/generator.html to get the application skeleton
yarn install
# Runs the app
yarn start
# Starts unittests
yarn test
```
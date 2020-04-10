# Go API using Gin-Gonic framework and MongoDB

## Installation

Follow these steps to install the api, and run it on your machine

### Prerequistes

You need to have Go 1.11 or above installed on your machine, as of MongoDB

### Installation Steps

```sh
$ git clone https://github.com/Dotaker/Go-Gin-RESTful-API.git/
$ cd Go-Gin-RESTful-API
$ go get github.com/gin-gonic/gin gopkg.in/mgo.v2/mgo
$ source .env && go build
$ ./CRUD-Operation
```

Go to `localhost:8081/api/users` on your favorite web browser, you should see a JSON response.

## Tests

This API has implented tests in the routes folder.

To run them, you need to follow some simple steps :

```sh
$ cd routes
$ go test
```

The output should print in the console, whether it has failed or passed.

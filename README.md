# zomato

[![GoDoc](https://godoc.org/github.com/go-india/zomato?status.svg)](https://godoc.org/github.com/go-india/zomato)
[![Build Status](https://travis-ci.org/go-india/zomato.svg?branch=master)](https://travis-ci.org/go-india/zomato)
[![Coverage Status](https://coveralls.io/repos/github/go-india/zomato/badge.svg?branch=master)](https://coveralls.io/github/go-india/zomato?branch=master)
[![Report card](https://goreportcard.com/badge/github.com/go-india/zomato)](https://goreportcard.com/report/github.com/go-india/zomato)

zomato is a [Go](http://golang.org/) client library for accessing [Zomato API](https://developers.zomato.com/api).

<img src="https://b.zmtcdn.com/images/logo/zomato_logo.svg" width="128">

> Zomato APIs give you access to the freshest and most exhaustive information for over 1.5 million **restaurants** across 10,000 cities globally. Power your content with the most exhaustive curated restaurant information.

### Installation

Requires Go version 1.7 or above.

```bash
$ go get -u github.com/go-india/zomato
```

### Usage

Construct a new zomato client, then use the various methods on the client to access different parts of the Zomato API.

For demonstration:

```go
package main

import (
  "context"
  "github.com/go-india/zomato"
)

var ctx = context.Background()

func main() {
  client := zomato.NewClient(API_KEY)

  // Gets restaurant details.
  res, err := client.Restaurant(ctx, 463)

  // Search for restaurants.
  restaurants, err := c.Search(ctx, zomato.SearchReq{
    Query: "delhi",
    Radius: 200,
  })

  // Gets reviews of the restaurant.
  reviews, err := c.Reviews(ctx, zomato.ReviewsReq{
    RestaurantID: 463,
    Count:        100,
  })
}
```

`Notes`
* Using the [context](https://godoc.org/context) package for passing context.  
* Make sure you have a valid API Key. If not, you can get a new one by registering at [Zomato Developers Portal](https://developers.zomato.com/api).

For complete usage of zomato, see the full [package docs](https://godoc.org/github.com/go-india/zomato).

#### Authentication

If you are using concrete Client, then you need to assign `client.Auth` field to make the client methods use authenticator for requests.

```go
client := zomato.Client{
  Auth: zomato.NewAuth(API_KEY),
}
```

This will add API Key to each request made by client methods.

#### Integration Tests

You can run integration tests from the directory.

```bash
$ go test -v
```

`Note`: Use `-update` flag to update the testdata. When using update flag, you will need to define `ZOMATO_TEST_API_KEY` in your environment for tests to use the API Key for testing.

### Contributing

We welcome pull requests, bug fixes and issue reports. Before proposing a change, please discuss your change by raising an issue.

### License

This library is distributed under the MIT license found in the [LICENSE](./LICENSE) file.

### Author

[Yash Raj Singh](http://yashrajsingh.net/)
/*
Package zomato provides a client for using Zomato API.

You can read the API server documentation at https://developers.zomato.com/api

Usage

Construct a new Zomato client, then use the various methods on the client to access different parts of the Zomato API.

For demonstration:

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

Notes:

* Using the https://godoc.org/context package for passing context.

* Look at tests(*_test.go) files for more sample usage.

Authentication

If you are using concrete Client, then you need to assign client.Auth field to make the client methods use authenticator for requests.

  client := zomato.Client{
    Auth: zomato.NewAuth(API_KEY),
  }

This will add API Key to each request made by client methods.
*/
package zomato

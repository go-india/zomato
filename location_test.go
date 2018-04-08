package zomato_test

import (
	"context"
	"testing"

	"github.com/go-india/zomato"
)

func TestLocationDetails(t *testing.T) {
	c := zomato.NewClient(getAPIKey())
	testClient(&c, t)

	resp, err := c.LocationDetails(context.Background(), 289, zomato.SubZone)
	if err != nil {
		t.Fatalf("LocationDetails failed: %+v", err)
	}

	if resp.BestRatedRestaurant == nil {
		t.Fatal("invalid response length")
	}
}

func TestLocations(t *testing.T) {
	c := zomato.NewClient(getAPIKey())
	testClient(&c, t)

	req := zomato.LocationsReq{
		Query:     "delhi",
		Latitude:  28.7041,
		Longitude: 77.1025,
		Count:     100,
	}

	resp, err := c.Locations(context.Background(), req)
	if err != nil {
		t.Fatalf("Locations failed: %+v", err)
	}

	if resp.LocationSuggestions == nil {
		t.Fatal("invalid response length")
	}
}

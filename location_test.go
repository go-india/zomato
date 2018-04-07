package zomato_test

import (
	"testing"

	"github.com/go-india/zomato"
)

func TestLocationDetails(t *testing.T) {
	c := zomato.Client{
		Auth: zomato.NewAuth(getAPIKey()),
	}
	testClient(&c, t)

	req := zomato.LocationDetailsReq{
		EntityID:   289,
		EntityType: zomato.SubZone,
	}

	var resp zomato.LocationDetailsResp
	err := c.Do(c.Auth(req), &resp)
	if err != nil {
		t.Fatalf("Do failed: %+v", err)
	}

	if resp.BestRatedRestaurant == nil {
		t.Fatal("invalid response length")
	}
}

func TestLocations(t *testing.T) {
	c := zomato.Client{
		Auth: zomato.NewAuth(getAPIKey()),
	}
	testClient(&c, t)

	req := zomato.LocationsReq{
		Query:     "delhi",
		Latitude:  28.7041,
		Longitude: 77.1025,
		Count:     100,
	}

	var resp zomato.LocationsResp
	err := c.Do(c.Auth(req), &resp)
	if err != nil {
		t.Fatalf("Do failed: %+v", err)
	}

	if resp.LocationSuggestions == nil {
		t.Fatal("invalid response length")
	}
}

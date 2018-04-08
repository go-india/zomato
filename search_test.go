package zomato_test

import (
	"context"
	"testing"

	"github.com/go-india/zomato"
)

func TestSearch(t *testing.T) {
	c := zomato.NewClient(getAPIKey())
	testClient(&c, t)

	req := zomato.SearchReq{
		Query:         "delhi",
		EntityType:    zomato.CityEntity,
		Start:         1,
		Count:         100,
		Latitude:      28.7041,
		Longitude:     77.1025,
		Radius:        100,
		Sort:          zomato.Cost,
		Order:         zomato.Ascending,
		Establishment: "new",
		Cuisines:      []string{"north indian"},
		Collection:    "north",
		Category:      "north",
	}

	resp, err := c.Search(context.Background(), req)
	if err != nil {
		t.Fatalf("Search failed: %+v", err)
	}

	if resp.Restaurants == nil {
		t.Fatal("invalid response length")
	}
}

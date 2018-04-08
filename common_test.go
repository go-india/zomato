package zomato_test

import (
	"context"
	"testing"

	"github.com/go-india/zomato"
)

func TestCategories(t *testing.T) {
	c := zomato.NewClient(getAPIKey())
	testClient(&c, t)

	resp, err := c.Categories(context.Background())
	if err != nil {
		t.Fatalf("Categories failed: %+v", err)
	}

	if resp.Categories == nil {
		t.Fatal("invalid response length")
	}
}

func TestCities(t *testing.T) {
	c := zomato.NewClient(getAPIKey())
	testClient(&c, t)

	req := zomato.CitiesReq{
		Query:     "delhi",
		Latitude:  28.7041,
		Longitude: 77.1025,
		CityIDs:   []int64{1, 2, 3},
		Count:     100,
	}

	resp, err := c.Cities(context.Background(), req)
	if err != nil {
		t.Fatalf("Cities failed: %+v", err)
	}

	if resp.LocationSuggestions == nil {
		t.Fatal("invalid response length")
	}
}

func TestCollections(t *testing.T) {
	c := zomato.NewClient(getAPIKey())
	testClient(&c, t)

	req := zomato.CollectionsReq{
		CityID:    1,
		Latitude:  28.7041,
		Longitude: 77.1025,
		Count:     100,
	}

	resp, err := c.Collections(context.Background(), req)
	if err != nil {
		t.Fatalf("Collections failed: %+v", err)
	}

	if resp.Collections == nil {
		t.Fatal("invalid response length")
	}
}

func TestCuisines(t *testing.T) {
	c := zomato.NewClient(getAPIKey())
	testClient(&c, t)

	req := zomato.CuisinesReq{
		CityID:    1,
		Latitude:  28.7041,
		Longitude: 77.1025,
	}

	resp, err := c.Cuisines(context.Background(), req)
	if err != nil {
		t.Fatalf("Cuisines failed: %+v", err)
	}

	if resp.Cuisines == nil {
		t.Fatal("invalid response length")
	}
}

func TestEstablishments(t *testing.T) {
	c := zomato.NewClient(getAPIKey())
	testClient(&c, t)

	req := zomato.EstablishmentsReq{
		CityID:    1,
		Latitude:  28.7041,
		Longitude: 77.1025,
	}

	resp, err := c.Establishments(context.Background(), req)
	if err != nil {
		t.Fatalf("Establishments failed: %+v", err)
	}

	if resp.Establishments == nil {
		t.Fatal("invalid response length")
	}
}

func TestGeoCode(t *testing.T) {
	c := zomato.NewClient(getAPIKey())
	testClient(&c, t)

	resp, err := c.GeoCode(context.Background(), 28.7041, 77.1025)
	if err != nil {
		t.Fatalf("GetCode failed: %+v", err)
	}

	if resp.NearbyRestaurants == nil {
		t.Fatal("invalid response length")
	}
}

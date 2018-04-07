package zomato_test

import (
	"testing"

	"github.com/go-india/zomato"
)

func TestCategories(t *testing.T) {
	c := zomato.Client{
		Auth: zomato.NewAuth(getAPIKey()),
	}
	testClient(&c, t)

	var resp zomato.CategoriesResp
	err := c.Do(c.Auth(zomato.CategoriesReq{}), &resp)
	if err != nil {
		t.Fatalf("Do failed: %+v", err)
	}

	if resp.Categories == nil {
		t.Fatal("invalid response length")
	}
}

func TestCities(t *testing.T) {
	c := zomato.Client{
		Auth: zomato.NewAuth(getAPIKey()),
	}
	testClient(&c, t)

	req := zomato.CitiesReq{
		Query:     "delhi",
		Latitude:  28.7041,
		Longitude: 77.1025,
		CityIDs:   []int64{1, 2, 3},
		Count:     100,
	}

	var resp zomato.CitiesResp
	err := c.Do(c.Auth(req), &resp)
	if err != nil {
		t.Fatalf("Do failed: %+v", err)
	}

	if resp.LocationSuggestions == nil {
		t.Fatal("invalid response length")
	}
}

func TestCollections(t *testing.T) {
	c := zomato.Client{
		Auth: zomato.NewAuth(getAPIKey()),
	}
	testClient(&c, t)

	req := zomato.CollectionsReq{
		CityID:    1,
		Latitude:  28.7041,
		Longitude: 77.1025,
		Count:     100,
	}

	var resp zomato.CollectionsResp
	err := c.Do(c.Auth(req), &resp)
	if err != nil {
		t.Fatalf("Do failed: %+v", err)
	}

	if resp.Collections == nil {
		t.Fatal("invalid response length")
	}
}

func TestCuisines(t *testing.T) {
	c := zomato.Client{
		Auth: zomato.NewAuth(getAPIKey()),
	}
	testClient(&c, t)

	req := zomato.CuisinesReq{
		CityID:    1,
		Latitude:  28.7041,
		Longitude: 77.1025,
	}

	var resp zomato.CuisinesResp
	err := c.Do(c.Auth(req), &resp)
	if err != nil {
		t.Fatalf("Do failed: %+v", err)
	}

	if resp.Cuisines == nil {
		t.Fatal("invalid response length")
	}
}

func TestEstablishments(t *testing.T) {
	c := zomato.Client{
		Auth: zomato.NewAuth(getAPIKey()),
	}
	testClient(&c, t)

	req := zomato.EstablishmentsReq{
		CityID:    1,
		Latitude:  28.7041,
		Longitude: 77.1025,
	}

	var resp zomato.EstablishmentsResp
	err := c.Do(c.Auth(req), &resp)
	if err != nil {
		t.Fatalf("Do failed: %+v", err)
	}

	if resp.Establishments == nil {
		t.Fatal("invalid response length")
	}
}

func TestGeoCode(t *testing.T) {
	c := zomato.Client{
		Auth: zomato.NewAuth(getAPIKey()),
	}
	testClient(&c, t)

	req := zomato.GeoCodeReq{ // delhi
		Latitude:  28.7041,
		Longitude: 77.1025,
	}

	var resp zomato.GeoCodeResp
	err := c.Do(c.Auth(req), &resp)
	if err != nil {
		t.Fatalf("Do failed: %+v", err)
	}

	if resp.NearbyRestaurants == nil {
		t.Fatal("invalid response length")
	}
}

package zomato_test

import (
	"testing"

	"github.com/go-india/zomato"
)

func TestDailyMenu(t *testing.T) {
	c := zomato.Client{
		Auth: zomato.NewAuth(getAPIKey()),
	}
	testClient(&c, t)

	req := zomato.DailyMenuReq{
		RestaurantID: 16514301,
	}

	var resp zomato.DailyMenuResp
	err := c.Do(c.Auth(req), &resp)
	if err != nil {
		t.Fatalf("Do failed: %+v", err)
	}

	if resp.DailyMenus == nil {
		t.Fatal("invalid response length")
	}
}

func TestRestaurant(t *testing.T) {
	c := zomato.Client{
		Auth: zomato.NewAuth(getAPIKey()),
	}
	testClient(&c, t)

	req := zomato.RestaurantReq{
		RestaurantID: 463,
	}

	var resp zomato.Restaurant
	err := c.Do(c.Auth(req), &resp)
	if err != nil {
		t.Fatalf("Do failed: %+v", err)
	}

	if resp.ID == nil {
		t.Fatal("invalid response length")
	}
}

func TestReviews(t *testing.T) {
	c := zomato.Client{
		Auth: zomato.NewAuth(getAPIKey()),
	}
	testClient(&c, t)

	req := zomato.ReviewsReq{
		RestaurantID: 463,
		Start:        1,
		Count:        100,
	}

	var resp zomato.ReviewsResp
	err := c.Do(c.Auth(req), &resp)
	if err != nil {
		t.Fatalf("Do failed: %+v", err)
	}

	if resp.UserReviews == nil {
		t.Fatal("invalid response length")
	}
}

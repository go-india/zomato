package zomato_test

import (
	"context"
	"testing"

	"github.com/go-india/zomato"
)

func TestDailyMenu(t *testing.T) {
	c := zomato.NewClient(getAPIKey())
	testClient(&c, t)

	resp, err := c.DailyMenu(context.Background(), 16514301)
	if err != nil {
		t.Fatalf("DailyMenu failed: %+v", err)
	}

	if resp.DailyMenus == nil {
		t.Fatal("invalid response length")
	}
}

func TestRestaurant(t *testing.T) {
	c := zomato.NewClient(getAPIKey())
	testClient(&c, t)

	resp, err := c.Restaurant(context.Background(), 463)
	if err != nil {
		t.Fatalf("Restaurant failed: %+v", err)
	}

	if resp.ID == nil {
		t.Fatal("invalid response length")
	}
}

func TestReviews(t *testing.T) {
	c := zomato.NewClient(getAPIKey())
	testClient(&c, t)

	req := zomato.ReviewsReq{
		RestaurantID: 463,
		Start:        1,
		Count:        100,
	}

	resp, err := c.Reviews(context.Background(), req)
	if err != nil {
		t.Fatalf("Reviews failed: %+v", err)
	}

	if resp.UserReviews == nil {
		t.Fatal("invalid response length")
	}
}

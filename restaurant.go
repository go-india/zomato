package zomato

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// DailyMenuReq parameters
type DailyMenuReq struct {
	// ID of restaurant whose details are requested
	RestaurantID int64 `url:"res_id" validate:"required"`
}

// Request encodes DailyMenuReq parameters returning a new http.Request
func (r DailyMenuReq) Request() (*http.Request, error) {
	err := validate.Struct(r)
	if err != nil {
		return nil, errors.Wrap(err, "invalid request")
	}

	urlStr := DefaultBaseURL + "/v2.1/dailymenu"

	values, err := query.Values(r)
	if err != nil {
		return nil, errors.Wrap(err, "encoding query params failed")
	}
	if params := values.Encode(); params != "" {
		urlStr += "?" + params
	}

	return http.NewRequest(http.MethodGet, urlStr, nil)
}

// DailyMenuResp holds daily menus of a restaurant
type DailyMenuResp struct {
	Status     *string `json:"status,omitempty"`
	DailyMenus []struct {
		DailyMenu *DailyMenu `json:"daily_menu,omitempty"`
	} `json:"daily_menus,omitempty"` // List of restaurant's menu details
}

// Dish holds dish/menu item details
type Dish struct {
	ID    *int64  `json:"dish_id,string,omitempty"` // Menu Item ID
	Name  *string `json:"name,omitempty"`           // Menu Item Title
	Price *string `json:"price,omitempty"`          // Menu Item Price
}

// DailyMenu holds daily menu
type DailyMenu struct {
	ID        *int64     `json:"daily_menu_id,string,omitempty"` // ID of the restaurant
	Name      *string    `json:"name,omitempty"`                 // Name of the restaurant
	StartDate *time.Time // `json:"start_date,omitempty"`           // Daily Menu start timestamp
	EndDate   *time.Time // `json:"end_date,omitempty"`             // Daily Menu end timestamp
	Dishes    []struct {
		Dish *Dish `json:"dish,omitempty"`
	} `json:"dishes,omitempty"` // Menu item in the category
}

// UnmarshalJSON convert JSON data to struct
func (d *DailyMenu) UnmarshalJSON(data []byte) error {
	type Alias DailyMenu
	t := struct {
		Alias

		StartDate string `json:"start_date,omitempty"`
		EndDate   string `json:"end_date,omitempty"`
	}{}
	if err := json.Unmarshal(data, &t); err != nil {
		return errors.Wrap(err, "UnmarshalJSON failed")
	}

	*d = DailyMenu(t.Alias)

	if len(t.StartDate) > 1 {
		sd, err := time.Parse("2006-01-02 15:04:05", t.StartDate)
		if err != nil {
			return errors.Wrap(err, "parse StartDate failed")
		}
		d.StartDate = &sd
	}

	if len(t.EndDate) > 1 {
		ed, err := time.Parse("2006-01-02 15:04:05", t.EndDate)
		if err != nil {
			return errors.Wrap(err, "parse EndDate failed")
		}
		d.EndDate = &ed
	}

	return nil
}

// DailyMenu gets daily menu using Zomato restaurant ID.
func (c Client) DailyMenu(ctx context.Context, restaurantID int64) (resp DailyMenuResp, err error) {
	if c.Auth == nil {
		return resp, ErrNoAuth
	}

	err = c.Do(c.Auth(WithCtx(ctx, DailyMenuReq{RestaurantID: restaurantID})), &resp)
	return resp, errors.Wrap(err, "Client.Do failed")
}

// RestaurantReq parameters
type RestaurantReq struct {
	// ID of restaurant whose details are requested
	RestaurantID int64 `url:"res_id" validate:"required"`
}

// Request encodes RestaurantReq parameters returning a new http.Request
func (r RestaurantReq) Request() (*http.Request, error) {
	err := validate.Struct(r)
	if err != nil {
		return nil, errors.Wrap(err, "invalid request")
	}

	urlStr := DefaultBaseURL + "/v2.1/restaurant"

	values, err := query.Values(r)
	if err != nil {
		return nil, errors.Wrap(err, "encoding query params failed")
	}
	if params := values.Encode(); params != "" {
		urlStr += "?" + params
	}

	return http.NewRequest(http.MethodGet, urlStr, nil)
}

// Restaurant gets restaurant details.
//
// Get detailed restaurant information using Zomato restaurant ID.
// Partner Access is required to access photos and reviews.
func (c Client) Restaurant(ctx context.Context, restaurantID int64) (resp Restaurant, err error) {
	if c.Auth == nil {
		return resp, ErrNoAuth
	}

	err = c.Do(c.Auth(WithCtx(ctx, RestaurantReq{RestaurantID: restaurantID})), &resp)
	return resp, errors.Wrap(err, "Client.Do failed")
}

// Restaurant holds a restaurant details
type Restaurant struct {
	ID       *int64              `json:"id,string,omitempty"` // ID of the restaurant
	Name     *string             `json:"name,omitempty"`      // Name of the restaurant
	URL      *string             `json:"url,omitempty"`       // URL of the restaurant page
	Location *RestaurantLocation `json:"location,omitempty"`  // Restaurant location details

	// List of cuisines served at the restaurant in csv format
	Cuisines []string // `json:"cuisines,omitempty"`
	// Average price of a meal for two people
	AverageCostForTwo *int64 `json:"average_cost_for_two,omitempty"`
	// Price bracket of the restaurant (1 being pocket friendly and 4 being the costliest)
	PriceRange *uint8 `json:"price_range,omitempty"`
	// Local currency symbol; to be used with price
	Currency *string `json:"currency,omitempty"`
	// Restaurant rating details
	UserRating *UserRating `json:"user_rating,omitempty"`

	// URL of the low resolution header image of restaurant
	ThumbnailURL *string `json:"thumb,omitempty"`
	// URL of the restaurant's photos page
	PhotosURL *string `json:"photos_url,omitempty"`
	// URL of the restaurant's menu page
	MenuURL *string `json:"menu_url,omitempty"`
	// URL of the high resolution header image of restaurant
	FeaturedImageURL *string `json:"featured_image,omitempty"`
	// URL of the restaurant's events page
	EventsURL *string `json:"events_url,omitempty"`
	// Short URL of the restaurant page; for use in apps or social shares
	DeeplinkURL *string `json:"deeplink,omitempty"`

	OrderURL         *string `json:"order_url,omitempty"`
	OrderDeeplinkURL *string `json:"order_deeplink,omitempty"`
	BookURL          *string `json:"book_url,omitempty"`

	// Whether the restaurant has online delivery enabled or not
	HasOnlineDelivery *bool // `json:"has_online_delivery,omitempty"`
	// Valid only if has_online_delivery = 1;
	// whether the restaurant is accepting online orders right now
	IsDeliveringNow   *bool // `json:"is_delivering_now,omitempty"`
	HasTableBooking   *bool // `json:"has_table_booking,omitempty"`
	SwitchToOrderMenu *bool // `json:"switch_to_order_menu,omitempty"`

	// TODO find their structure
	Offers             []interface{} `json:"offers,omitempty"`
	EstablishmentTypes []interface{} `json:"establishment_types,omitempty"`

	// ZomatoEvents are the events available in this restaurant
	ZomatoEvents []struct {
		Event *Event `json:"event,omitempty"`
	} `json:"zomato_events,omitempty"`

	// APIKey used to make the request
	APIkey *string `json:"apikey,omitempty"`
	// R holds restaurant ID's
	R struct {
		RestaurantID *int64 `json:"res_id,omitempty"`
	} `json:"R,omitempty"`

	// Partner Access fields
	ReviewsCount *int64   `json:"all_reviews_count,omitempty"` // [Partner access] Number of reviews for the restaurant
	PhotoCount   *int64   `json:"photo_count,omitempty"`       // [Partner access] Total number of photos for the restaurant, at max 10 photos for partner access
	PhoneNumbers *string  `json:"phone_numbers,omitempty"`     // [Partner access] Restaurant's contact numbers in csv format
	Photos       []Photo  `json:"photos,omitempty"`            // [Partner access] List of restaurant photos
	Reviews      []Review `json:"all_reviews,omitempty"`       // [Partner access] List of restaurant reviews
}

// UnmarshalJSON convert JSON data to struct
func (r *Restaurant) UnmarshalJSON(data []byte) error {
	type Alias Restaurant
	t := struct {
		Alias
		Cuisines          string `json:"cuisines,omitempty"`
		HasOnlineDelivery uint8  `json:"has_online_delivery,omitempty"`
		IsDeliveringNow   uint8  `json:"is_delivering_now,omitempty"`
		HasTableBooking   uint8  `json:"has_table_booking,omitempty"`
		SwitchToOrderMenu uint8  `json:"switch_to_order_menu,omitempty"`
	}{}
	if err := json.Unmarshal(data, &t); err != nil {
		return errors.Wrap(err, "UnmarshalJSON failed")
	}

	*r = Restaurant(t.Alias)
	r.HasOnlineDelivery = newBool(t.HasOnlineDelivery == 1)
	r.IsDeliveringNow = newBool(t.IsDeliveringNow == 1)
	r.HasTableBooking = newBool(t.HasTableBooking == 1)
	r.SwitchToOrderMenu = newBool(t.SwitchToOrderMenu == 1)
	r.Cuisines = strings.Split(t.Cuisines, ",")
	return nil
}

// RestaurantLocation holds restaurant location details
type RestaurantLocation struct {
	Address         *string  `json:"address,omitempty"`  // Complete address of the restaurant
	Locality        *string  `json:"locality,omitempty"` // Name of the locality
	City            *string  `json:"city,omitempty"`     // Name of the city
	CityID          *int64   `json:"city_id,omitempty"`
	Latitude        *float64 `json:"latitude,string,omitempty"`  // Coordinates of the restaurant
	Longitude       *float64 `json:"longitude,string,omitempty"` // Coordinates of the restaurant
	Zipcode         *int64   // `json:"zipcode,omitempty"`          // Zipcode
	CountryID       *int64   `json:"country_id,omitempty"` // ID of the country
	LocalityVerbose *string  `json:"locality_verbose,omitempty"`
}

// UnmarshalJSON convert JSON data to struct
func (r *RestaurantLocation) UnmarshalJSON(data []byte) error {
	type Alias RestaurantLocation
	t := struct {
		Alias
		Zipcode string `json:"zipcode"`
	}{}
	if err := json.Unmarshal(data, &t); err != nil {
		return errors.Wrap(err, "UnmarshalJSON failed")
	}

	*r = RestaurantLocation(t.Alias)

	if t.Zipcode != "" {
		zc, err := strconv.Atoi(t.Zipcode)
		if err != nil {
			return errors.Wrap(err, "parse Zipcode failed")
		}
		z := int64(zc)
		r.Zipcode = &z
	}

	return nil
}

// UserRating stores user rating details
type UserRating struct {
	// Restaurant rating on a scale of 0.0 to 5.0 in increments of 0.1
	AggregateRating *float64 `json:"aggregate_rating,string,omitempty"`
	// Short description of the rating
	RatingText *string `json:"rating_text,omitempty"`
	// Color hex code used with the rating on Zomato
	RatingColor *string `json:"rating_color,omitempty"`
	// Number of ratings received
	Votes *int64 `json:"votes,string,omitempty"`
}

// Photo holds photo details
type Photo struct {
	URL          *string `json:"url,omitempty"`       // URL of the image file
	ThumbnailURL *string `json:"thumb_url,omitempty"` // URL for 200 X 200 thumb image file
	Order        *int64  `json:"order,omitempty"`
	MD5Sum       *string `json:"md5sum,omitempty"`
	PhotoID      *int64  `json:"photo_id,omitempty"`
	UUID         *int64  `json:"uuid,omitempty"`
	Type         *string `json:"type,omitempty"`

	ID           *string `json:"id,omitempty"`            // ID of the photo
	User         *User   `json:"user,omitempty"`          // User who uploaded the photo
	RestaurantID *int64  `json:"res_id,string,omitempty"` // ID of restaurant for which the image was uploaded
	Caption      *string `json:"caption,omitempty"`       // Caption of the photo
	// Unix timestamp when the photo was uploaded
	Timestamp *time.Time // `json:"timestamp,string"`
	// User friendly time string; denotes when the photo was uploaded
	FriendlyTime  *string `json:"friendly_time,omitempty"`
	Width         *int64  `json:"width,string,omitempty"`          // Image width in pixel; usually 640
	Height        *int64  `json:"height,string,omitempty"`         // Image height in pixel; usually 640
	CommentsCount *int64  `json:"comments_count,string,omitempty"` // Number of comments on photo
	LikesCount    *int64  `json:"likes_count,string,omitempty"`    // Number of likes on photo
}

// UnmarshalJSON convert JSON data to struct
func (p *Photo) UnmarshalJSON(data []byte) error {
	type Alias Photo
	t := struct {
		Alias
		Timestamp int64 `json:"timestamp,string"`
	}{}
	if err := json.Unmarshal(data, &t); err != nil {
		return errors.Wrap(err, "UnmarshalJSON failed")
	}

	*p = Photo(t.Alias)

	ts := time.Unix(t.Timestamp, 0)
	if !ts.IsZero() {
		p.Timestamp = &ts
	}
	return nil
}

// User holds user details
type User struct {
	// User's name
	Name *string `json:"name,omitempty"`
	// User's @handle; uniquely identifies a user on Zomato
	ZomatoHandle *string `json:"zomato_handle,omitempty"`
	// Text for user's foodie level
	FoodieLevel *string `json:"foodie_level,omitempty"`
	// Number to identify user's foodie level; ranges from 0 to 10
	FoodieLevelNumber *uint8 `json:"foodie_level_num,omitempty"`
	// Color hex code used with foodie level on Zomato
	FoodieColor *string `json:"foodie_color,omitempty"`
	// URL for user's profile on Zomato
	ProfileURL *string `json:"profile_url,omitempty"`
	// short URL for user's profile on Zomato; for use in apps or social sharing
	ProfileDeeplinkURL *string `json:"profile_deeplink,omitempty"`
	// URL for user's profile image
	ProfileImageURL *string `json:"profile_image,omitempty"`
}

// Event holds zomato event details
type Event struct {
	ID *int64 `json:"event_id,omitempty"`

	StartDate *time.Time // `json:"start_date"`
	EndDate   *time.Time // `json:"end_date"`
	EndTime   *time.Time // `json:"end_time"`
	StartTime *time.Time // `json:"start_time"`
	DateAdded *time.Time // `json:"date_added"`

	IsActive     *bool // `json:"is_active"`
	IsValid      *bool // `json:"is_valid"`
	ShowShareURL *bool // `json:"show_share_url"`
	IsEndTimeSet *bool // `json:"is_end_time_set"`

	Photos []struct {
		Photo *Photo `json:"photo,omitempty"`
	} `json:"photos,omitempty"`

	Restaurants       []Restaurant `json:"restaurants,omitempty"`
	ShareURL          *string      `json:"share_url,omitempty"`
	Title             *string      `json:"title,omitempty"`
	Description       *string      `json:"description,omitempty"`
	DisplayTime       *string      `json:"display_time,omitempty"`
	DisplayDate       *string      `json:"display_date,omitempty"`
	Disclaimer        *string      `json:"disclaimer,omitempty"`
	EventCategory     *int64       `json:"event_category,omitempty"`
	EventCategoryName *string      `json:"event_category_name,omitempty"`
	BookLinkURL       *string      `json:"book_link,omitempty"`

	FriendlyStartDate *string `json:"friendly_start_date,omitempty"`
	FriendlyEndDate   *string `json:"friendly_end_date,omitempty"`
	FriendlyTiming    *string `json:"friendly_timing_str,omitempty"`
}

// UnmarshalJSON convert JSON data to struct
func (e *Event) UnmarshalJSON(data []byte) error {
	type Alias Event
	t := struct {
		Alias

		StartDate string `json:"start_date,omitempty"`
		EndDate   string `json:"end_date,omitempty"`
		StartTime string `json:"start_time,omitempty"`
		EndTime   string `json:"end_time,omitempty"`
		DateAdded string `json:"date_added,omitempty"`

		IsActive     uint8 `json:"is_active,omitempty"`
		IsValid      uint8 `json:"is_valid,omitempty"`
		ShowShareURL uint8 `json:"show_share_url,omitempty"`
		IsEndTimeSet uint8 `json:"is_end_time_set,omitempty"`
	}{}
	if err := json.Unmarshal(data, &t); err != nil {
		return errors.Wrap(err, "UnmarshalJSON failed")
	}

	*e = Event(t.Alias)

	e.IsActive = newBool(t.IsActive == 1)
	e.IsValid = newBool(t.IsValid == 1)
	e.ShowShareURL = newBool(t.ShowShareURL == 1)
	e.IsEndTimeSet = newBool(t.IsEndTimeSet == 1)

	if len(t.StartDate) > 1 {
		sd, err := time.Parse("2006-01-02", t.StartDate)
		if err != nil {
			return errors.Wrap(err, "parse StartDate failed")
		}
		e.StartDate = &sd
	}

	if len(t.EndDate) > 1 {
		ed, err := time.Parse("2006-01-02", t.EndDate)
		if err != nil {
			return errors.Wrap(err, "parse EndDate failed")
		}
		e.EndDate = &ed
	}

	if len(t.StartTime) > 1 {
		st, err := time.Parse("15:04:05", t.StartTime)
		if err != nil {
			return errors.Wrap(err, "parse StartTime failed")
		}
		e.StartTime = &st
	}

	if len(t.EndTime) > 1 {
		et, err := time.Parse("15:04:05", t.EndTime)
		if err != nil {
			return errors.Wrap(err, "parse EndTime failed")
		}
		e.EndTime = &et
	}

	if len(t.DateAdded) > 1 {
		da, err := time.Parse("2006-01-02 15:04:05", t.DateAdded)
		if err != nil {
			return errors.Wrap(err, "parse DateAdded failed")
		}
		e.DateAdded = &da
	}

	return nil
}

// ReviewsReq parameters
type ReviewsReq struct {
	// ID of restaurant whose details are requested
	RestaurantID int64 `url:"res_id" validate:"required"`
	// Fetch results after this offset
	Start uint64 `url:"start,omitempty"`
	// Max number of results to retrieve
	Count uint64 `url:"count,omitempty"`
}

// Request encodes ReviewsReq parameters returning a new http.Request
func (r ReviewsReq) Request() (*http.Request, error) {
	err := validate.Struct(r)
	if err != nil {
		return nil, errors.Wrap(err, "invalid request")
	}

	urlStr := DefaultBaseURL + "/v2.1/reviews"

	values, err := query.Values(r)
	if err != nil {
		return nil, errors.Wrap(err, "encoding query params failed")
	}
	if params := values.Encode(); params != "" {
		urlStr += "?" + params
	}

	return http.NewRequest(http.MethodGet, urlStr, nil)
}

// ReviewsResp holds reviews for a restaurant
type ReviewsResp struct {
	ReviewsCount *int64 `json:"reviews_count,omitempty"`
	ReviewsStart *int64 `json:"reviews_start,omitempty"`
	ReviewsShown *int64 `json:"reviews_shown,omitempty"`

	UserReviews []struct {
		Review *Review `json:"review,omitempty"`
	} `json:"user_reviews,omitempty"`

	RespondToReviewsViaZomatoDashboardURL *string `json:"Respond to reviews via Zomato Dashboard,omitempty"`
}

// Review holds review details
type Review struct {
	// ID of the review
	ID *int64 // `json:"id,string,omitempty"`
	// Rating on scale of 0 to 5 in increments of 0.5
	Rating *float64 // `json:"rating,string,omitempty"`
	// Review text
	ReviewText *string `json:"review_text,omitempty"`
	// Color hex code used with the rating on Zomato
	RatingColor *string `json:"rating_color,omitempty"`
	// User friendly time string corresponding to time of review posting
	ReviewTimeFriendly *string `json:"review_time_friendly,omitempty"`
	// Short description of the rating
	RatingText *string `json:"rating_text,omitempty"`
	// Unix timestamp for review_time_friendly
	Timestamp *time.Time // `json:"timestamp,string,omitempty"`
	// No of likes received for review
	Likes *int64 // `json:"likes,string,omitempty"`
	// User details of author of review
	User *User `json:"user"`
	// No of comments on review
	CommentsCount *int64 // `json:"comments_count,string,omitempty"`
}

// UnmarshalJSON convert JSON data to struct
func (r *Review) UnmarshalJSON(data []byte) error {
	type Alias Review
	t := struct {
		Alias
		Timestamp     json.Number `json:"timestamp"`
		Rating        json.Number `json:"rating"`
		Likes         json.Number `json:"likes"`
		ID            json.Number `json:"id"`
		CommentsCount json.Number `json:"comments_count"`
	}{}
	if err := json.Unmarshal(data, &t); err != nil {
		return errors.Wrap(err, "UnmarshalJSON failed")
	}

	*r = Review(t.Alias)

	if t.Timestamp.String() != "" {
		tt, err := t.Timestamp.Int64()
		if err != nil {
			return errors.Wrap(err, "parse Timestamp failed")
		}

		ts := time.Unix(tt, 0)
		if !ts.IsZero() {
			r.Timestamp = &ts
		}
	}

	if t.Rating.String() != "" {
		rating, err := t.Rating.Float64()
		if err != nil {
			return errors.Wrap(err, "parse Rating failed")
		}
		r.Rating = &rating
	}

	if t.Likes.String() != "" {
		likes, err := t.Likes.Int64()
		if err != nil {
			return errors.Wrap(err, "parse Likes failed")
		}
		r.Likes = &likes
	}

	if t.ID.String() != "" {
		id, err := t.ID.Int64()
		if err != nil {
			return errors.Wrap(err, "parse ID failed")
		}
		r.ID = &id
	}

	if t.CommentsCount.String() != "" {
		cc, err := t.CommentsCount.Int64()
		if err != nil {
			return errors.Wrap(err, "parse CommentsCount failed")
		}
		r.CommentsCount = &cc
	}
	return nil
}

// Reviews gets restaurant reviews.
//
// Get restaurant reviews using the Zomato restaurant ID.
// Only 5 latest reviews are available under the Basic API plan.
func (c Client) Reviews(ctx context.Context, req ReviewsReq) (resp ReviewsResp, err error) {
	if c.Auth == nil {
		return resp, ErrNoAuth
	}

	err = c.Do(c.Auth(WithCtx(ctx, req)), &resp)
	return resp, errors.Wrap(err, "Client.Do failed")
}

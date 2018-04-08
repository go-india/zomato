package zomato

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// LocationDetailsReq parameters
type LocationDetailsReq struct {
	EntityID   int64      `url:"entity_id" validate:"required"`   // location id obtained from locations api
	EntityType EntityType `url:"entity_type" validate:"required"` // location type obtained from locations api
}

// Request encodes LocationDetailsReq parameters returning a new http.Request
func (r LocationDetailsReq) Request() (*http.Request, error) {
	err := validate.Struct(r)
	if err != nil {
		return nil, errors.Wrap(err, "invalid request")
	}

	urlStr := DefaultBaseURL + "/v2.1/location_details"

	values, err := query.Values(r)
	if err != nil {
		return nil, errors.Wrap(err, "encoding query params failed")
	}
	if params := values.Encode(); params != "" {
		urlStr += "?" + params
	}

	return http.NewRequest(http.MethodGet, urlStr, nil)
}

// LocationDetailsResp holds location details
type LocationDetailsResp struct {
	Location           *Location `json:"location,omitempty"`
	NumberOfRestaurant int64     `json:"num_restaurant,omitempty"`

	BestRatedRestaurant []struct {
		Restaurant *Restaurant `json:"restaurant,omitempty"`
	} `json:"best_rated_restaurant,omitempty"`

	Experts []struct {
		User *User `json:"user,omitempty"`
	} `json:"experts,omitempty"`
}

// Location holds details of location
type Location struct {
	// Type of location: one of [city, zone, subzone, landmark, group, metro, street]
	EntityType *string `json:"entity_type,omitempty"`
	// ID of location: (entity_id, entity_type) tuple uniquely identifies a location
	EntityID *int64 `json:"entity_id,omitempty"`
	// Name of the location
	Title *string `json:"title,omitempty"`
	// Coordinates of the (centre of) location
	Latitude *float64 // `json:"latitude,string,omitempty"`
	// Coordinates of the (centre of) location
	Longitude *float64 // `json:"longitude,string,omitempty"`
	// ID of city
	CityID *int64 `json:"city_id,omitempty"`
	// Name of the city
	CityName *string `json:"city_name,omitempty"`
	// ID of country
	CountryID *int64 `json:"country_id,omitempty"`
	// Name of the country
	CountryName *string `json:"country_name,omitempty"`
}

// UnmarshalJSON convert JSON data to struct
func (l *Location) UnmarshalJSON(data []byte) error {
	type Alias Location
	t := struct {
		Alias

		Latitude  json.Number `json:"latitude,omitempty"`
		Longitude json.Number `json:"longitude,omitempty"`
	}{}
	if err := json.Unmarshal(data, &t); err != nil {
		return errors.Wrap(err, "UnmarshalJSON failed")
	}

	*l = Location(t.Alias)

	if t.Latitude.String() != "" {
		lat, err := t.Latitude.Float64()
		if err != nil {
			return errors.Wrap(err, "parse Latitude failed")
		}
		l.Latitude = &lat
	}

	if t.Longitude.String() != "" {
		lon, err := t.Longitude.Float64()
		if err != nil {
			return errors.Wrap(err, "parse Longitude failed")
		}
		l.Longitude = &lon
	}

	return nil
}

// LocationDetails gets Zomato location details.
//
// Get Foodie Index, Nightlife Index, Top Cuisines and Best rated restaurants in a given location.
func (c Client) LocationDetails(ctx context.Context,
	entityID int64, entityType EntityType) (resp LocationDetailsResp, err error) {
	if c.Auth == nil {
		return resp, ErrNoAuth
	}

	err = c.Do(c.Auth(WithCtx(ctx, LocationDetailsReq{
		EntityID:   entityID,
		EntityType: entityType,
	})), &resp)
	return resp, errors.Wrap(err, "Client.Do failed")
}

// LocationsReq parameters
type LocationsReq struct {
	Query     string  `url:"query" validate:"required"` // suggestion for location name
	Latitude  float64 `url:"lat,omitempty"`             // latitude
	Longitude float64 `url:"lon,omitempty"`             // longitude
	Count     uint64  `url:"count,omitempty"`           // number of max results to display
}

// Request encodes LocationsReq parameters returning a new http.Request
func (r LocationsReq) Request() (*http.Request, error) {
	err := validate.Struct(r)
	if err != nil {
		return nil, errors.Wrap(err, "invalid request")
	}

	urlStr := DefaultBaseURL + "/v2.1/locations"

	values, err := query.Values(r)
	if err != nil {
		return nil, errors.Wrap(err, "encoding query params failed")
	}
	if params := values.Encode(); params != "" {
		urlStr += "?" + params
	}

	return http.NewRequest(http.MethodGet, urlStr, nil)
}

// LocationsResp holds locations from LocationsReq query
type LocationsResp struct {
	LocationSuggestions []Location `json:"location_suggestions,omitempty"`
	Status              *string    `json:"status,omitempty"`
	HasMore             *bool      // json:"has_more,omitempty"`
	HasTotal            *bool      // `json:"has_total,omitempty"`
}

// UnmarshalJSON convert JSON data to struct
func (l *LocationsResp) UnmarshalJSON(data []byte) error {
	type Alias LocationsResp
	t := struct {
		Alias
		HasMore  uint8 `json:"has_more,omitempty"`
		HasTotal uint8 `json:"has_total,omitempty"`
	}{}
	if err := json.Unmarshal(data, &t); err != nil {
		return errors.Wrap(err, "UnmarshalJSON failed")
	}

	*l = LocationsResp(t.Alias)
	l.HasMore = newBool(t.HasMore == 1)
	l.HasTotal = newBool(t.HasTotal == 1)
	return nil
}

// Locations searchs for locations.
//
// Search for Zomato locations by keyword.
// Provide coordinates to get better search results.
func (c Client) Locations(ctx context.Context, req LocationsReq) (resp LocationsResp, err error) {
	if c.Auth == nil {
		return resp, ErrNoAuth
	}

	err = c.Do(c.Auth(WithCtx(ctx, req)), &resp)
	return resp, errors.Wrap(err, "Client.Do failed")
}

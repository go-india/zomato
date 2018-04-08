package zomato

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// SearchReq parameters
type SearchReq struct {
	Query      string     `url:"q,omitempty"`           // search keyword
	EntityID   int64      `url:"entity_id,omitempty"`   // location id
	EntityType EntityType `url:"entity_type,omitempty"` // location type
	Latitude   float64    `url:"lat,omitempty"`         // latitude
	Longitude  float64    `url:"lon,omitempty"`         // longitude
	// Fetch results after this offset
	Start uint64 `url:"start,omitempty"`
	// Max number of results to retrieve
	Count uint64 `url:"count,omitempty"`
	// radius around (lat,lon); to define search area, defined in meters(M)
	Radius float64 `url:"radius,omitempty"`
	// establishment id obtained from establishments call
	Establishment string `url:"establishment_type,omitempty"`
	// list of cuisine id's separated by comma
	Cuisines []string `url:"cuisines,omitempty"`
	// collection id obtained from collections call
	Collection string `url:"collection_id,omitempty"`
	// category ids obtained from categories call
	Category string `url:"category,omitempty"`
	// sort restaurants by ...
	Sort Sort `url:"sort,omitempty"`
	// used with 'sort' parameter to define ascending or descending
	Order Order `url:"order,omitempty"`
}

// Request encodes SearchReq parameters returning a new http.Request
func (r SearchReq) Request() (*http.Request, error) {
	urlStr := DefaultBaseURL + "/v2.1/search"

	values, err := query.Values(r)
	if err != nil {
		return nil, errors.Wrap(err, "encoding query params failed")
	}
	if params := values.Encode(); params != "" {
		urlStr += "?" + params
	}

	return http.NewRequest(http.MethodGet, urlStr, nil)
}

// SearchResp holds search response from the search query
type SearchResp struct {
	// Number of results found
	ResultsFound int64 `json:"results_found,omitempty"`
	// The starting location within results from which the results were fetched
	// (used for pagination)
	ResultsStart int64 `json:"results_start,omitempty"`
	// The number of results fetched (used for pagination)
	ResultsShown int64 `json:"results_shown,omitempty"`

	Restaurants []struct {
		Restaurant *Restaurant `json:"restaurant,omitempty"`
	} `json:"restaurants,omitempty"`
}

// Search provides search for restaurants.
//
// The location input can be specified using Zomato location ID or coordinates.
// Cuisine/Establishment/Collection IDs can be obtained from respective API calls.
//
// Get up to 100 restaurants by changing the 'start' and 'count' parameters with
// the maximum value of count being 20.
//
// Examples:
//
// 		To search for 'Italian' restaurants in 'Manhattan, New York City', set cuisines = 55, entity_id = 94741 and entity_type = zone
// 		To search for 'cafes' in 'Manhattan, New York City', set establishment_type = 1, entity_type = zone and entity_id = 94741
// 		Get list of all restaurants in 'Trending this Week' collection in 'New York City' by using entity_id = 280, entity_type = city and collection_id = 1
//
// Partner Access is required to access photos and reviews.
func (c Client) Search(ctx context.Context, req SearchReq) (resp SearchResp, err error) {
	if c.Auth == nil {
		return resp, ErrNoAuth
	}

	err = c.Do(c.Auth(WithCtx(ctx, req)), &resp)
	return resp, errors.Wrap(err, "Client.Do failed")
}

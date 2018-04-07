package zomato

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// CategoriesReq parameters
type CategoriesReq struct{}

// Request encodes CategoriesReq parameters returning a new http.Request
func (r CategoriesReq) Request() (*http.Request, error) {
	urlStr := DefaultBaseURL + "/v2.1/categories"
	return http.NewRequest(http.MethodGet, urlStr, nil)
}

// Categorie has categorie details.
type Categorie struct {
	ID   int64  `json:"id"`   // ID of the category type
	Name string `json:"name"` // Name of the category type
}

// CategoriesResp holds categories response
type CategoriesResp struct {
	Categories []struct {
		Categorie *Categorie `json:"categories,omitempty"`
	} `json:"categories,omitempty"`
}

// CitiesReq parameters
type CitiesReq struct {
	Query     string  `url:"q,omitempty"`        // query by city name
	Latitude  float64 `url:"lat,omitempty"`      // latitude
	Longitude float64 `url:"lon,omitempty"`      // longitude
	CityIDs   []int64 `url:"city_ids,omitempty"` // comma separated city_id values
	Count     uint64  `url:"count,omitempty"`    // number of max results to display
}

// Request encodes CitiesReq parameters returning a new http.Request
func (r CitiesReq) Request() (*http.Request, error) {
	urlStr := DefaultBaseURL + "/v2.1/cities"

	values, err := query.Values(r)
	if err != nil {
		return nil, errors.Wrap(err, "encoding query params failed")
	}
	if params := values.Encode(); params != "" {
		urlStr += "?" + params
	}

	return http.NewRequest(http.MethodGet, urlStr, nil)
}

// City holds city details
type City struct {
	ID   int64  `json:"id"`   // ID of the city
	Name string `json:"name"` // City name

	CountryID      *int64  `json:"country_id,omitempty"`       // ID of the country
	CountryName    *string `json:"country_name,omitempty"`     // Name of the country
	CountryFlagURL *string `json:"country_flag_url,omitempty"` // Country Flag picture URL

	ShouldExperimentWith *bool // `json:"should_experiment_with,omitempty"`
	DiscoveryEnabled     *bool // `json:"discovery_enabled,omitempty"`
	HasNewAdFormat       *bool // `json:"has_new_ad_format,omitempty"`

	IsState   *bool   // `json:"is_state,omitempty"`   // Whether this location is a state
	StateID   *int64  `json:"state_id,omitempty"`   // ID of the state
	StateName *string `json:"state_name,omitempty"` // Name of the state
	StateCode *string `json:"state_code,omitempty"` // Short code for the state
}

// UnmarshalJSON convert JSON data to struct
func (c *City) UnmarshalJSON(data []byte) error {
	type Alias City
	t := struct {
		Alias
		ShouldExperimentWith uint8 `json:"should_experiment_with,omitempty"`
		DiscoveryEnabled     uint8 `json:"discovery_enabled,omitempty"`
		HasNewAdFormat       uint8 `json:"has_new_ad_format,omitempty"`
		IsState              uint8 `json:"is_state,omitempty"`
	}{}
	if err := json.Unmarshal(data, &t); err != nil {
		return errors.Wrap(err, "UnmarshalJSON failed")
	}

	*c = City(t.Alias)
	c.ShouldExperimentWith = newBool(t.ShouldExperimentWith == 1)
	c.DiscoveryEnabled = newBool(t.DiscoveryEnabled == 1)
	c.HasNewAdFormat = newBool(t.HasNewAdFormat == 1)
	c.IsState = newBool(t.IsState == 1)
	return nil
}

// CitiesResp has cities returned by CitiesReq query
type CitiesResp struct {
	LocationSuggestions []City  `json:"location_suggestions,omitempty"`
	Status              *string `json:"status,omitempty"`
	HasMore             *bool   // `json:"has_more,omitempty"`
	HasTotal            *bool   // `json:"has_total,omitempty"`
}

// UnmarshalJSON convert JSON data to struct
func (c *CitiesResp) UnmarshalJSON(data []byte) error {
	type Alias CitiesResp
	t := struct {
		Alias
		HasMore  uint8 `json:"has_more,omitempty"`
		HasTotal uint8 `json:"has_total,omitempty"`
	}{}
	if err := json.Unmarshal(data, &t); err != nil {
		return errors.Wrap(err, "UnmarshalJSON failed")
	}

	*c = CitiesResp(t.Alias)
	c.HasMore = newBool(t.HasMore == 1)
	c.HasTotal = newBool(t.HasTotal == 1)
	return nil
}

// CollectionsReq parameters
type CollectionsReq struct {
	CityID    int64   `url:"city_id,omitempty"` // id of the city for which collections are needed
	Latitude  float64 `url:"lat,omitempty"`     // latitude of any point within a city
	Longitude float64 `url:"lon,omitempty"`     // longitude  of any point within a city
	Count     uint64  `url:"count,omitempty"`   // max number of results needed
}

// Request encodes CollectionsReq parameters returning a new http.Request
func (r CollectionsReq) Request() (*http.Request, error) {
	urlStr := DefaultBaseURL + "/v2.1/collections"

	values, err := query.Values(r)
	if err != nil {
		return nil, errors.Wrap(err, "encoding query params failed")
	}
	if params := values.Encode(); params != "" {
		urlStr += "?" + params
	}

	return http.NewRequest(http.MethodGet, urlStr, nil)
}

// Collection holds collection details
type Collection struct {
	ID               *int64  `json:"collection_id,omitempty"` // ID of the collection of restaurants
	URL              *string `json:"url,omitempty"`           // Collection name
	Title            *string `json:"title,omitempty"`         // URL of the collection page
	Description      *string `json:"description,omitempty"`   // Short description of the collection
	RestaurantsCount *int64  `json:"res_count,omitempty"`     // URL for header image of the collection
	ImageURL         *string `json:"image_url,omitempty"`     // Number of restaurants in the collection
	ShareURL         *string `json:"share_url,omitempty"`     // Short URL for apps and social sharing
}

// CollectionsResp holds collections returned from CollectionsReq query
type CollectionsResp struct {
	Collections []struct {
		Collection *Collection `json:"collection,omitempty"`
	} `json:"collections,omitempty"`
	ShareURL    *string `json:"share_url,omitempty"`
	DisplayText *string `json:"display_text,omitempty"`
	HasMore     *bool   `json:"has_more,omitempty"`
	HasTotal    *bool   `json:"has_total,omitempty"`
}

// UnmarshalJSON convert JSON data to struct
func (c *CollectionsResp) UnmarshalJSON(data []byte) error {
	type Alias CollectionsResp
	t := struct {
		Alias
		HasMore  uint8 `json:"has_more,omitempty"`
		HasTotal uint8 `json:"has_total,omitempty"`
	}{}
	if err := json.Unmarshal(data, &t); err != nil {
		return errors.Wrap(err, "UnmarshalJSON failed")
	}

	*c = CollectionsResp(t.Alias)
	c.HasMore = newBool(t.HasMore == 1)
	c.HasTotal = newBool(t.HasTotal == 1)
	return nil
}

// CuisinesReq parameters
type CuisinesReq struct {
	CityID    int64   `url:"city_id,omitempty"` // id of the city for which cuisines are needed
	Latitude  float64 `url:"lat,omitempty"`     // latitude of any point within a city
	Longitude float64 `url:"lon,omitempty"`     // longitude  of any point within a city
}

// Request encodes CuisinesReq parameters returning a new http.Request
func (r CuisinesReq) Request() (*http.Request, error) {
	urlStr := DefaultBaseURL + "/v2.1/cuisines"

	values, err := query.Values(r)
	if err != nil {
		return nil, errors.Wrap(err, "encoding query params failed")
	}
	if params := values.Encode(); params != "" {
		urlStr += "?" + params
	}

	return http.NewRequest(http.MethodGet, urlStr, nil)
}

// Cuisine holds cuisine details
type Cuisine struct {
	ID   int64  `json:"cuisine_id"`   // ID of the cuisine
	Name string `json:"cuisine_name"` // Name of the cuisine
}

// CuisinesResp holds cuisines from CuisinesReq query
type CuisinesResp struct {
	Cuisines []struct {
		Cuisine *Cuisine `json:"cuisine,omitempty"`
	} `json:"cuisines,omitempty"`
}

// EstablishmentsReq parameters
type EstablishmentsReq struct {
	CityID    int64   `url:"city_id,omitempty"` // id of the city
	Latitude  float64 `url:"lat,omitempty"`     // latitude of any point within a city
	Longitude float64 `url:"lon,omitempty"`     // longitude  of any point within a city
}

// Request encodes EstablishmentsReq parameters returning a new http.Request
func (r EstablishmentsReq) Request() (*http.Request, error) {
	urlStr := DefaultBaseURL + "/v2.1/establishments"

	values, err := query.Values(r)
	if err != nil {
		return nil, errors.Wrap(err, "encoding query params failed")
	}
	if params := values.Encode(); params != "" {
		urlStr += "?" + params
	}

	return http.NewRequest(http.MethodGet, urlStr, nil)
}

// Establishment holds establishment details
type Establishment struct {
	ID   int64  `json:"id"`   // ID of the establishment type
	Name string `json:"name"` // Name of the establishment type
}

// EstablishmentsResp holds establishments from CuisinesReq query
type EstablishmentsResp struct {
	Establishments []struct {
		Establishment *Establishment `json:"establishment,omitempty"`
	} `json:"establishments,omitempty"`
}

// GeoCodeReq parameters
type GeoCodeReq struct {
	Latitude  float64 `url:"lat" validate:"required"` // latitude of any point within a city
	Longitude float64 `url:"lon" validate:"required"` // longitude  of any point within a city
}

// Request encodes GeoCodeReq parameters returning a new http.Request
func (r GeoCodeReq) Request() (*http.Request, error) {
	err := validate.Struct(r)
	if err != nil {
		return nil, errors.Wrap(err, "invalid request")
	}

	urlStr := DefaultBaseURL + "/v2.1/geocode"

	values, err := query.Values(r)
	if err != nil {
		return nil, errors.Wrap(err, "encoding query params failed")
	}
	if params := values.Encode(); params != "" {
		urlStr += "?" + params
	}

	return http.NewRequest(http.MethodGet, urlStr, nil)
}

// GeoCodeResp holds foodie and Nightlife Index,
// list of popular cuisines and nearby restaurants around the given coordinates
type GeoCodeResp struct {
	Location          *Location   `json:"location,omitempty"`
	Popularity        *Popularity `json:"popularity,omitempty"`
	LinkURL           *string     `json:"link,omitempty"`
	NearbyRestaurants []struct {
		Restaurant *Restaurant `json:"restaurant,omitempty"`
	} `json:"nearby_restaurants,omitempty"`
}

// Popularity has popularity details
type Popularity struct {
	Popularity           *float64 `json:"popularity,string,omitempty"`      // Foodie index of a location out of 5.00
	NightlifeIndex       *float64 `json:"nightlife_index,string,omitempty"` // Nightlife index of a location out of 5.00
	NearbyRestaurantIDs  []int64  // `json:"nearby_res,omitempty"`
	TopCuisines          []string `json:"top_cuisines,omitempty"` // Most popular cuisines in the locality
	PopularityRestaurant *int64   `json:"popularity_res,string,omitempty"`
	NightlifeRestaurant  *int64   `json:"nightlife_res,string,omitempty"`
	Subzone              *string  `json:"subzone,omitempty"`
	SubzoneID            *int64   `json:"subzone_id,omitempty"`
	City                 *string  `json:"city,omitempty"`
}

// UnmarshalJSON convert JSON data to struct
func (p *Popularity) UnmarshalJSON(data []byte) error {
	type Alias Popularity
	t := struct {
		Alias
		NearbyRestaurantIDs []string `json:"nearby_res,omitempty"`
	}{}
	if err := json.Unmarshal(data, &t); err != nil {
		return errors.Wrap(err, "UnmarshalJSON failed")
	}

	*p = Popularity(t.Alias)

	if len(t.NearbyRestaurantIDs) > 0 {
		for _, id := range t.NearbyRestaurantIDs {
			i, err := strconv.Atoi(id)
			if err != nil {
				continue
			}
			p.NearbyRestaurantIDs = append(p.NearbyRestaurantIDs, int64(i))
		}
	}
	return nil
}

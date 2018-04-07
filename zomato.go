package zomato

import "gopkg.in/go-playground/validator.v9"

// Sort defines sort types used for sorting in searching
type Sort string

// Sort types
const (
	Cost         Sort = "cost"
	Rating       Sort = "rating"
	RealDistance Sort = "real_distance"
)

// Order defines order types used for ordering in searching
type Order string

// Order types
const (
	Ascending  Order = "asc"
	Descending Order = "desc"
)

// EntityType defines entity types used for location type
type EntityType string

// Entity Types
const (
	CityEntity EntityType = "city"
	SubZone    EntityType = "subzone"
	Zone       EntityType = "zone"
	Landmark   EntityType = "landmark"
	Metro      EntityType = "metro"
	Group      EntityType = "group"
)

// use a single instance of Validate, it caches struct info
var validate = validator.New()

// newBool initialises a new bool and returns its address
func newBool(b bool) *bool { return &b }

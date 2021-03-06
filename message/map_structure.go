package message

// point position in latitude and longitude projection wgs84
// swagger:response point
type Point struct {
	// description of the Point
	Description string `json:"description,omitempty"`
	// latitude in decimal value
	Latitude float64 `json:"lat"`
	// longitude in decimal value
	Longitude float64 `json:"lng"`
}

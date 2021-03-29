package fptf

// A Route represents a single set of stations, of a single Line.
//
// For a very consistent subway service, there may be one route
// for each direction. Planned detours, trains stopping early and
// additional directions would have their own route.
type Route struct {
	Type    string  `json:"type"`
	Id      string  `json:"id"`
	Line    *Line   `json:"line"`
	Mode    Mode    `json:"mode,omitempty"`
	SubMode string  `json:"subMode,omitempty"`
	Stops   []*Stop `json:"stops"`
}

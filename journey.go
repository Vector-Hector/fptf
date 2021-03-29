package fptf

// A journey is a computed set of directions to get from A to B at a
// specific time. It would typically be the result of a route planning
// algorithm.
type Journey struct {
	Type  string `json:"type"`
	Id    string `json:"id"`
	Trips []Trip `json:"legs"`
	Price Price  `json:"price"`
}

// A formalized, inferred version of a journey leg
type Trip struct {
	Origin      *StopStation `json:"origin"`
	Destination *StopStation `json:"destination"`

	Departure         TimeNullable `json:"departure"`
	DepartureDelay    int          `json:"departureDelay,omitempty"`
	DeparturePlatform string       `json:"departurePlatform,omitempty"`

	Arrival         TimeNullable `json:"arrival"`
	ArrivalDelay    int          `json:"arrivalDelay,omitempty"`
	ArrivalPlatform string       `json:"arrivalPlatform,omitempty"`

	Schedule *Schedule `json:"schedule,omitempty"`

	Stopovers []*Stopover `json:"stopovers,omitempty"`

	Mode    Mode `json:"mode,omitempty"`
	SubMode Mode `json:"subMode,omitempty"`

	Public bool `json:"public,omitempty"`

	Operator *Operator `json:"operator,omitempty"`

	Price *Price `json:"price,omitempty"`
}

type Price struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

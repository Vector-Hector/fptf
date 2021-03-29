package fptf

// A Stopover represents a vehicle stopping at a stop/station at a specific time.
type Stopover struct {
	Type              string       `json:"type"`
	StopStation       *StopStation `json:"stop"`
	Arrival           TimeNullable `json:"arrival"`
	ArrivalDelay      int          `json:"arrivalDelay,omitempty"`
	ArrivalPlatform   string       `json:"arrivalPlatform,omitempty"`
	Departure         TimeNullable `json:"departure"`
	DepartureDelay    int          `json:"departureDelay,omitempty"`
	DeparturePlatform string       `json:"departurePlatform,omitempty"`
}

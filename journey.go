package fptf

import "encoding/json"

// A journey is a computed set of directions to get from A to B at a
// specific time. It would typically be the result of a route planning
// algorithm.
type Journey struct {
	Id    string
	Trips *[]Trip
	Price *Price
	Meta  interface{} // any additional data
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

	Mode    Mode   `json:"mode,omitempty"`
	SubMode string `json:"subMode,omitempty"`

	Public *bool `json:"public,omitempty"`

	Operator *Operator `json:"operator,omitempty"`

	Price *Price `json:"price,omitempty"`

	// Some additional arguments, inspired by https://github.com/public-transport/hafas-client
	Line      *Line  `json:"line,omitempty"`      // The line on which this trip is going
	Direction string `json:"direction,omitempty"` // The direction string on the train

	Meta interface{} `json:"meta,omitempty"` // any additional data
}

type Price struct {
	Amount   float64     `json:"amount"`
	Currency string      `json:"currency"`
	Meta     interface{} `json:"meta,omitempty"` // any additional data
}

type mJourney struct {
	typed
	Id    string      `json:"id"`
	Trips *[]Trip     `json:"legs"`
	Price *Price      `json:"price"`
	Meta  interface{} `json:"meta,omitempty"`
}

func (w *Journey) toM() *mJourney {
	return &mJourney{
		typed: typedJourney,
		Id:    w.Id,
		Trips: w.Trips,
		Price: w.Price,
		Meta:  w.Meta,
	}
}

func (w *Journey) fromM(m *mJourney) {
	w.Id = m.Id
	w.Trips = m.Trips
	w.Price = m.Price
	w.Meta = m.Meta
}

// as it is optional to give either line id or Line object,
// we have to unmarshal|marshal it ourselves.
func (w *Journey) UnmarshalJSON(data []byte) error {
	var m mJourney
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	w.fromM(&m)
	return nil
}

func (w *Journey) MarshalJSON() ([]byte, error) {
	return json.Marshal(w.toM())
}

package fptf

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"time"
)

// Journey is a computed set of directions to get from A to B at a
// specific time. It would typically be the result of a route planning
// algorithm.
type Journey struct {
	Id    string
	Trips []*Trip
	Price *Price
	Meta  interface{} // any additional data
}

func (j *Journey) GetFirstTrip() *Trip {
	if len(j.Trips) == 0 {
		return nil
	}
	return j.Trips[0]
}

func (j *Journey) GetOrigin() *StopStation {
	first := j.GetFirstTrip()
	if first == nil {
		return nil
	}
	return first.Origin
}

func (j *Journey) GetDeparture() time.Time {
	first := j.GetFirstTrip()
	if first == nil {
		return time.Time{}
	}
	return first.Departure.Time
}

func (j *Journey) GetDeparturePlatform() string {
	first := j.GetFirstTrip()
	if first == nil {
		return ""
	}
	return first.DeparturePlatform
}

func (j *Journey) GetDepartureDelay() *int {
	first := j.GetFirstTrip()
	if first == nil {
		return nil
	}
	return first.DepartureDelay
}

func (j *Journey) GetLastTrip() *Trip {
	if len(j.Trips) == 0 {
		return nil
	}
	return j.Trips[len(j.Trips)-1]
}

func (j *Journey) GetDestination() *StopStation {
	last := j.GetLastTrip()
	if last == nil {
		return nil
	}
	return last.Origin
}

func (j *Journey) GetArrival() time.Time {
	last := j.GetLastTrip()
	if last == nil {
		return time.Time{}
	}
	return last.Arrival.Time
}

func (j *Journey) GetArrivalPlatform() string {
	last := j.GetLastTrip()
	if last == nil {
		return ""
	}
	return last.ArrivalPlatform
}

func (j *Journey) GetArrivalDelay() *int {
	last := j.GetLastTrip()
	if last == nil {
		return nil
	}
	return last.ArrivalDelay
}

// Trip is a formalized, inferred version of a journey leg
type Trip struct {
	Origin      *StopStation `json:"origin,omitempty" bson:"origin,omitempty"`
	Destination *StopStation `json:"destination,omitempty" bson:"destination,omitempty"`

	Departure         TimeNullable `json:"departure,omitempty" bson:"departure,omitempty"`
	DepartureDelay    *int         `json:"departureDelay,omitempty" bson:"departureDelay,omitempty"`
	DeparturePlatform string       `json:"departurePlatform,omitempty" bson:"departurePlatform,omitempty"`

	Arrival         TimeNullable `json:"arrival,omitempty" bson:"arrival,omitempty"`
	ArrivalDelay    *int         `json:"arrivalDelay,omitempty" bson:"arrivalDelay,omitempty"`
	ArrivalPlatform string       `json:"arrivalPlatform,omitempty" bson:"arrivalPlatform,omitempty"`

	Schedule *Schedule `json:"schedule,omitempty" bson:"schedule,omitempty"`

	Stopovers []*Stopover `json:"stopovers,omitempty" bson:"stopovers,omitempty"`

	Mode    Mode   `json:"mode,omitempty" bson:"mode,omitempty"`
	SubMode string `json:"subMode,omitempty" bson:"subMode,omitempty"`

	Public *bool `json:"public,omitempty" bson:"public,omitempty"`

	Operator *Operator `json:"operator,omitempty" bson:"operator,omitempty"`

	Price *Price `json:"price,omitempty" bson:"price,omitempty"`

	// Some additional arguments, inspired by https://github.com/public-transport/hafas-client
	Line      *Line  `json:"line,omitempty" bson:"line,omitempty"`           // The line on which this trip is going
	Direction string `json:"direction,omitempty" bson:"direction,omitempty"` // The direction string on the train

	Meta interface{} `json:"meta,omitempty" bson:"meta,omitempty"` // any additional data
}

type Price struct {
	Amount   float64     `json:"amount,omitempty" bson:"amount,omitempty"`
	Currency string      `json:"currency,omitempty" bson:"currency,omitempty"`
	Meta     interface{} `json:"meta,omitempty" bson:"meta,omitempty"` // any additional data
}

// GetMode The mode of a trip can be defined in many places.
// This method finds the mode of a given trip.
func (trip *Trip) GetMode() *Mode {
	if trip.Mode != "" {
		return &trip.Mode
	}
	if trip.Line != nil && trip.Line.Mode != "" {
		return &trip.Line.Mode
	}
	if trip.Schedule != nil {
		if trip.Schedule.Mode != "" {
			return &trip.Schedule.Mode
		}
		if trip.Schedule.Route != nil && trip.Schedule.Route.Mode != "" {
			return &trip.Schedule.Route.Mode
		}
		if trip.Schedule.Route != nil && trip.Schedule.Route.Line != nil && trip.Schedule.Route.Line.Mode != "" {
			return &trip.Schedule.Route.Line.Mode
		}
	}
	return nil
}

// GetLine The line of a trip can be defined in multiple places
// This method finds it
func (trip *Trip) GetLine() *Line {
	if trip.Line != nil {
		return trip.Line
	}
	if trip.Schedule != nil && trip.Schedule.Route != nil && trip.Schedule.Route.Line != nil {
		return trip.Schedule.Route.Line
	}
	return nil
}

func (trip *Trip) SubTrip(startInclusive int, endExclusive int) *Trip {
	stopovers := trip.Stopovers[startInclusive:endExclusive]
	origin := stopovers[0]
	dest := stopovers[len(stopovers)-1]

	return &Trip{
		Origin:            origin.StopStation,
		Destination:       dest.StopStation,
		Departure:         origin.Departure,
		DepartureDelay:    origin.DepartureDelay,
		DeparturePlatform: origin.DeparturePlatform,
		Arrival:           dest.Arrival,
		ArrivalDelay:      dest.ArrivalDelay,
		ArrivalPlatform:   dest.ArrivalPlatform,
		Schedule:          trip.Schedule,
		Stopovers:         stopovers,
		Mode:              trip.Mode,
		SubMode:           trip.SubMode,
		Public:            trip.Public,
		Operator:          trip.Operator,
		Price:             nil,
		Line:              trip.Line,
		Direction:         trip.Direction,
		Meta:              trip.Meta,
	}
}

type mJourney struct {
	Typed `bson:"inline"`
	Id    string      `json:"id,omitempty" bson:"id,omitempty"`
	Trips []*Trip     `json:"legs,omitempty" bson:"trips,omitempty"`
	Price *Price      `json:"price,omitempty" bson:"price,omitempty"`
	Meta  interface{} `json:"meta,omitempty" bson:"meta,omitempty"`
}

func (j *Journey) toM() *mJourney {
	return &mJourney{
		Typed: typedJourney,
		Id:    j.Id,
		Trips: j.Trips,
		Price: j.Price,
		Meta:  j.Meta,
	}
}

func (j *Journey) fromM(m *mJourney) {
	j.Id = m.Id
	j.Trips = m.Trips
	j.Price = m.Price
	j.Meta = m.Meta
}

// as it is optional to give either line id or Line object,
// we have to unmarshal|marshal it ourselves.

func (j *Journey) UnmarshalJSON(data []byte) error {
	var m mJourney
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	j.fromM(&m)
	return nil
}

func (j *Journey) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.toM())
}

func (j *Journey) UnmarshalBSON(data []byte) error {
	var m mJourney
	err := bson.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	j.fromM(&m)
	return nil
}

func (j *Journey) UnmarshalBSONValue(typ bsontype.Type, data []byte) error {
	var m mJourney
	err := bson.UnmarshalValue(typ, data, &m)
	if err != nil {
		return err
	}
	j.fromM(&m)
	return nil
}

func (j Journey) MarshalBSON() ([]byte, error) {
	return bson.Marshal(j.toM())
}

func (j Journey) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(j.toM())
}

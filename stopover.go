package fptf

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

// A Stopover represents a vehicle stopping at a stop/station at a specific time.
type Stopover struct {
	StopStation       *StopStation
	Arrival           TimeNullable
	ArrivalDelay      *int
	ArrivalPlatform   string
	Departure         TimeNullable
	DepartureDelay    *int
	DeparturePlatform string
	Meta              interface{}
}

// intermediate, Typed format used by marshal
type mStopover struct {
	Typed             `bson:"inline"`
	StopStation       *StopStation `json:"stop,omitempty" bson:"stop,omitempty"`
	Arrival           TimeNullable `json:"arrival,omitempty" bson:"arrival,omitempty"`
	ArrivalDelay      *int         `json:"arrivalDelay,omitempty" bson:"arrivalDelay,omitempty"`
	ArrivalPlatform   string       `json:"arrivalPlatform,omitempty" bson:"arrivalPlatform,omitempty"`
	Departure         TimeNullable `json:"departure,omitempty" bson:"departure,omitempty"`
	DepartureDelay    *int         `json:"departureDelay,omitempty" bson:"departureDelay,omitempty"`
	DeparturePlatform string       `json:"departurePlatform,omitempty" bson:"departurePlatform,omitempty"`
	Meta              interface{}  `json:"meta,omitempty" bson:"meta,omitempty"`
}

func (s *Stopover) toM() *mStopover {
	return &mStopover{
		Typed:             typedStopover,
		StopStation:       s.StopStation,
		Arrival:           s.Arrival,
		ArrivalDelay:      s.ArrivalDelay,
		ArrivalPlatform:   s.ArrivalPlatform,
		Departure:         s.Departure,
		DepartureDelay:    s.DepartureDelay,
		DeparturePlatform: s.DeparturePlatform,
		Meta:              s.Meta,
	}
}

func (s *Stopover) fromM(m *mStopover) {
	s.StopStation = m.StopStation
	s.Arrival = m.Arrival
	s.ArrivalDelay = m.ArrivalDelay
	s.ArrivalPlatform = m.ArrivalPlatform
	s.Departure = m.Departure
	s.DepartureDelay = m.DepartureDelay
	s.DeparturePlatform = m.DeparturePlatform
	s.Meta = m.Meta
}

// as it is optional to give either stop|station id or Stop or Station object,
// we have to unmarshal|marshal it ourselves.

func (s *Stopover) UnmarshalJSON(data []byte) error {
	var m mStopover
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	s.fromM(&m)
	return nil
}

func (s *Stopover) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.toM())
}

func (s *Stopover) UnmarshalBSONValue(typ bsontype.Type, data []byte) error {
	var m mStopover
	err := bson.UnmarshalValue(typ, data, &m)
	if err != nil {
		return err
	}
	s.fromM(&m)
	return nil
}

func (s Stopover) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(s.toM())
}

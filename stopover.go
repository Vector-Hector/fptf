package fptf

import (
	"encoding/json"
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

// intermediate, typed format used by marshal
type mStopover struct {
	typed
	StopStation       *StopStation `json:"stop"`
	Arrival           TimeNullable `json:"arrival"`
	ArrivalDelay      *int         `json:"arrivalDelay,omitempty"`
	ArrivalPlatform   string       `json:"arrivalPlatform,omitempty"`
	Departure         TimeNullable `json:"departure"`
	DepartureDelay    *int         `json:"departureDelay,omitempty"`
	DeparturePlatform string       `json:"departurePlatform,omitempty"`
	Meta              interface{}  `json:"meta,omitempty"`
}

func (s *Stopover) toM() *mStopover {
	return &mStopover{
		typed:             typedStopover,
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

package fptf

import (
	"encoding/json"
)

// A Stop is a single small point or structure at which vehicles stop.
// A Stop always belongs to a Station. It may for example be a sign,
// a basic shelter or a railway platform.
//
// If the underlying data source does not allow such a fine-grained
// distinction, use stations instead.
type Stop struct {
	Id       string
	Name     string
	Station  *Station
	Location *Location
	Meta     interface{}

	Partial bool // only show the id in the json response?
}

// used by marshal
type mStop struct {
	typed
	Id       string      `json:"id"`
	Name     string      `json:"name"`
	Station  *Station    `json:"station"`
	Location *Location   `json:"location,omitempty"`
	Meta     interface{} `json:"meta,omitempty"`
}

func (s *Stop) toM() *mStop {
	return &mStop{
		typed:    typedStop,
		Id:       s.Id,
		Name:     s.Name,
		Station:  s.Station,
		Location: s.Location,
		Meta:     s.Meta,
	}
}

func (s *Stop) fromM(m *mStop) {
	s.Id = m.Id
	s.Name = m.Name
	s.Station = m.Station
	s.Location = m.Location
	s.Meta = m.Meta
}

// as it is optional to give either line id or Line object,
// we have to unmarshal|marshal it ourselves.
func (s *Stop) UnmarshalJSON(data []byte) error {
	var id string
	if err := json.Unmarshal(data, &id); err == nil {
		s.Id = id
		s.Partial = true
		return nil
	}
	s.Partial = false
	var m mStop
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	s.fromM(&m)
	return nil
}

func (s *Stop) MarshalJSON() ([]byte, error) {
	if s.Partial {
		return json.Marshal(s.Id)
	}
	return json.Marshal(s.toM())
}

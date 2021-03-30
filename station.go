package fptf

import (
	"encoding/json"
)

// A Station is a larger building or area that can be identified by a name.
// It is usually represented by a single node on a public transport map.
// Whereas a Stop usually specifies a location, a Station often is a
// broader area that may span across multiple levels or buildings.
type Station struct {
	Id       string
	Name     string
	Location *Location
	Regions  *[]Region
	Meta     interface{} // any additional data

	Partial bool // only show the id in the json response?
}

// used by marshal
type mStation struct {
	typed
	Id       string      `json:"id"`
	Name     string      `json:"name"`
	Location *Location   `json:"location,omitempty"`
	Regions  *[]Region   `json:"regions,omitempty"`
	Meta     interface{} `json:"meta,omitempty"`
}

func (s *Station) toM() *mStation {
	return &mStation{
		typed:    typedStation,
		Id:       s.Id,
		Name:     s.Name,
		Location: s.Location,
		Regions:  s.Regions,
		Meta: s.Meta,
	}
}

func (s *Station) fromM(m *mStation) {
	s.Id = m.Id
	s.Name = m.Name
	s.Location = m.Location
	s.Regions = m.Regions
	s.Meta = m.Meta
}

// as it is optional to give either station id or Station object,
// we have to unmarshal|marshal it ourselves.
func (s *Station) UnmarshalJSON(data []byte) error {
	var id string
	if err := json.Unmarshal(data, &id); err == nil {
		s.Id = id
		s.Partial = true
		return nil
	}
	s.Partial = false
	var m mStation
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	s.fromM(&m)
	return nil
}

func (s *Station) MarshalJSON() ([]byte, error) {
	if s.Partial {
		return json.Marshal(s.Id)
	}
	return json.Marshal(s.toM())
}

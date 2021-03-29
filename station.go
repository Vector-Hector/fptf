package fptf

import (
	"encoding/json"
)

// A Station is a larger building or area that can be identified by a name.
// It is usually represented by a single node on a public transport map.
// Whereas a Stop usually specifies a location, a Station often is a
// broader area that may span across multiple levels or buildings.
type Station struct {
	_Station

	Partial bool `json:"-"` // only show the id in the json response?
}

type _Station struct {
	Type     string    `json:"type"`
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Location *Location `json:"location,omitempty"`
	Regions  []*Region `json:"regions,omitempty"`
}

func NewStation(id string, name string, location *Location, regions []*Region) Station {
	return Station{
		_Station: _Station{
			Type:     "station",
			Id:       id,
			Name:     name,
			Location: location,
			Regions:  regions,
		},
		Partial:  false,
	}
}

// as it is optional to give either station id or Station object,
// we have to unmarshal|marshal it ourselves.
func (w *Station) UnmarshalJSON(data []byte) error {
	var id string
	if err := json.Unmarshal(data, &id); err == nil {
		w.Id = id
		w.Partial = true
		return nil
	}
	w.Partial = false
	return json.Unmarshal(data, &w._Station)
}

func (w *Station) MarshalJSON() ([]byte, error) {
	if w.Partial {
		return json.Marshal(w.Id)
	}
	return json.Marshal(w._Station)
}

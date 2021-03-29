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
	_Stop
	Partial bool `json:"-"` // only show the id in the json response?
}

type _Stop struct {
	Type     string    `json:"type"`
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Station  *Station  `json:"station"`
	Location *Location `json:"location,omitempty"`
}

// as it is optional to give either line id or Line object,
// we have to unmarshal|marshal it ourselves.
func (w *Stop) UnmarshalJSON(data []byte) error {
	var id string
	if err := json.Unmarshal(data, &id); err == nil {
		w.Id = id
		w.Partial = true
		return nil
	}
	w.Partial = false
	return json.Unmarshal(data, &w._Stop)
}

func (w *Stop) MarshalJSON() ([]byte, error) {
	if w.Partial {
		return json.Marshal(w.Id)
	}
	return json.Marshal(w._Stop)
}

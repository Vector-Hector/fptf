package fptf

import "encoding/json"

// Note: There are many ways to format schedules of public transport
// routes. This one tries to balance the amount of data and
// consumability. It is specifically geared towards urban public
// transport, with frequent trains and homogenous travels.
type Schedule struct {
	_Schedule

	Partial bool `json:"-"` // only show the id in the json response?
}

type _Schedule struct {
	Type     string             `json:"type"`
	Id       string             `json:"id"`
	Route    *Route             `json:"route"`
	Mode     Mode               `json:"mode,omitempty"`
	SubMode  string             `json:"subMode,omitempty"`
	Sequence []*SequenceElement `json:"sequence"`
	Starts   []TimeUnix         `json:"starts"`
}

type SequenceElement struct {
	Arrival   int64 `json:"arrival"`
	Departure int64 `json:"departure"`
}

// as it is optional to give either schedule id or Schedule object,
// we have to unmarshal|marshal it ourselves.
func (w *Schedule) UnmarshalJSON(data []byte) error {
	var id string
	if err := json.Unmarshal(data, &id); err == nil {
		w.Id = id
		w.Partial = true
		return nil
	}
	w.Partial = false
	return json.Unmarshal(data, &w._Schedule)
}

func (w *Schedule) MarshalJSON() ([]byte, error) {
	if w.Partial {
		return json.Marshal(w.Id)
	}
	return json.Marshal(w._Schedule)
}

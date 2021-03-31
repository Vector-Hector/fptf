package fptf

import (
	"encoding/json"
)

// Note: There are many ways to format schedules of public transport
// routes. This one tries to balance the amount of data and
// consumability. It is specifically geared towards urban public
// transport, with frequent trains and homogenous travels.
type Schedule struct {
	Id       string
	Route    *Route
	Mode     Mode
	SubMode  string
	Sequence []*SequenceElement
	Starts   []TimeUnix
	Meta     interface{} // any additional data

	Partial bool // only show the id in the json response?
}

type SequenceElement struct {
	Arrival   int64 `json:"arrival"`
	Departure int64 `json:"departure"`
}

// used by marshal
type mSchedule struct {
	typed
	Id       string             `json:"id"`
	Route    *Route             `json:"route"`
	Mode     Mode               `json:"mode,omitempty"`
	SubMode  string             `json:"subMode,omitempty"`
	Sequence []*SequenceElement `json:"sequence"`
	Starts   []TimeUnix         `json:"starts"`
	Meta     interface{}        `json:"meta,omitempty"`
}

func (w *Schedule) toM() *mSchedule {
	return &mSchedule{
		typed:    typedSchedule,
		Id:       w.Id,
		Route:    w.Route,
		Mode:     w.Mode,
		SubMode:  w.SubMode,
		Sequence: w.Sequence,
		Starts:   w.Starts,
		Meta:     w.Meta,
	}
}

func (w *Schedule) fromM(m *mSchedule) {
	w.Id = m.Id
	w.Route = m.Route
	w.Mode = m.Mode
	w.SubMode = m.SubMode
	w.Sequence = m.Sequence
	w.Starts = m.Starts
	w.Meta = m.Meta
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
	var m mSchedule
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	w.fromM(&m)
	return nil
}

func (w *Schedule) MarshalJSON() ([]byte, error) {
	if w.Partial {
		return json.Marshal(w.Id)
	}
	return json.Marshal(w.toM())
}

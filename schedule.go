package fptf

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

// Schedule There are many ways to format schedules of public transport
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
	Arrival   *int64 `json:"arrival,omitempty" bson:"arrival,omitempty"`
	Departure *int64 `json:"departure,omitempty" bson:"departure,omitempty"`
}

// used by marshal
type mSchedule struct {
	Typed    `bson:"inline"`
	Id       string             `json:"id,omitempty" bson:"id,omitempty"`
	Route    *Route             `json:"route,omitempty" bson:"route,omitempty"`
	Mode     Mode               `json:"mode,omitempty" bson:"mode,omitempty"`
	SubMode  string             `json:"subMode,omitempty" bson:"subMode,omitempty"`
	Sequence []*SequenceElement `json:"sequence,omitempty" bson:"sequence,omitempty"`
	Starts   []TimeUnix         `json:"starts,omitempty" bson:"starts,omitempty"`
	Meta     interface{}        `json:"meta,omitempty" bson:"meta,omitempty"`
}

func (w *Schedule) toM() *mSchedule {
	return &mSchedule{
		Typed:    typedSchedule,
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

func (w *Schedule) UnmarshalBSONValue(typ bsontype.Type, data []byte) error {
	if typ == bson.TypeString {
		var id string
		err := bson.UnmarshalValue(bson.TypeString, data, &id)
		if err != nil {
			return err
		}
		w.Id = id
		w.Partial = true
		return nil
	}

	w.Partial = false
	var m mSchedule
	err := bson.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	w.fromM(&m)
	return nil
}

func (w Schedule) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if w.Partial {
		return bson.MarshalValue(w.Id)
	}
	return bson.MarshalValue(w.toM())
}

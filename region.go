package fptf

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

// A Region is a group of Stations, for example a metropolitan area
// or a geographical or cultural region.
//
// In many urban areas, there are several
// long-distance train & bus stations, all distinct but well-connected
// through local public transport. It makes sense to keep
// them as Stations, because they may still have individual Stop s,
// but clustering them enables more advanced routing information.
//
// A Station can be part of multiple Region's.
type Region struct {
	Id       string
	Name     string
	Stations []*Station
	Meta     interface{} // any additional data

	Partial bool // only show the id in the json response?
}

type mRegion struct {
	Typed    `bson:"inline"`
	Id       string      `json:"id,omitempty" bson:"id,omitempty"`
	Name     string      `json:"name,omitempty" bson:"name,omitempty"`
	Stations []*Station  `json:"stations,omitempty" bson:"stations,omitempty"`
	Meta     interface{} `json:"meta,omitempty" bson:"meta,omitempty"`
}

func (w *Region) toM() *mRegion {
	return &mRegion{
		Typed:    typedRegion,
		Id:       w.Id,
		Name:     w.Name,
		Stations: w.Stations,
		Meta:     w.Meta,
	}
}

func (w *Region) fromM(m *mRegion) {
	w.Id = m.Id
	w.Name = m.Name
	w.Stations = m.Stations
	w.Meta = m.Meta
}

// as it is optional to give either region id or Region object,
// we have to unmarshal|marshal it ourselves.
func (w *Region) UnmarshalJSON(data []byte) error {
	var id string
	if err := json.Unmarshal(data, &id); err == nil {
		w.Id = id
		w.Partial = true
		return nil
	}
	w.Partial = false
	var m mRegion
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	w.fromM(&m)
	return nil
}

func (w *Region) MarshalJSON() ([]byte, error) {
	if w.Partial {
		return json.Marshal(w.Id)
	}
	return json.Marshal(w.toM())
}

func (w *Region) UnmarshalBSONValue(typ bsontype.Type, data []byte) error {
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
	var m mRegion
	err := bson.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	w.fromM(&m)
	return nil
}

func (w Region) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if w.Partial {
		return bson.MarshalValue(w.Id)
	}
	return bson.MarshalValue(w.toM())
}

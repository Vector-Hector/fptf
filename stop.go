package fptf

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

// Stop is a single small point or structure at which vehicles stop.
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
	Typed    `bson:"inline"`
	Id       string      `json:"id,omitempty" bson:"id,omitempty"`
	Name     string      `json:"name,omitempty" bson:"name,omitempty"`
	Station  *Station    `json:"station,omitempty" bson:"station,omitempty"`
	Location *Location   `json:"location,omitempty" bson:"location,omitempty"`
	Meta     interface{} `json:"meta,omitempty" bson:"meta,omitempty"`
}

func (s *Stop) toM() *mStop {
	return &mStop{
		Typed:    typedStop,
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

func (s *Stop) UnmarshalBSONValue(typ bsontype.Type, data []byte) error {
	if typ == bson.TypeString {
		var id string
		err := bson.UnmarshalValue(bson.TypeString, data, &id)
		if err != nil {
			return err
		}
		s.Id = id
		s.Partial = true
		return nil
	}
	s.Partial = false
	var m mStop
	err := bson.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	s.fromM(&m)
	return nil
}

func (s *Stop) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if s.Partial {
		return bson.MarshalValue(s.Id)
	}
	return bson.MarshalValue(s.toM())
}

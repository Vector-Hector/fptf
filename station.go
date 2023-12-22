package fptf

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

// Station is a larger building or area that can be identified by a name.
// It is usually represented by a single node on a public transport map.
// Whereas a Stop usually specifies a location, a Station often is a
// broader area that may span across multiple levels or buildings.
type Station struct {
	Id       string
	Name     string
	Location *Location
	Regions  []*Region
	Meta     interface{} // any additional data

	Partial bool // only show the id in the json response?
}

// used by marshal
type mStation struct {
	Typed    `bson:"inline"`
	Id       string      `json:"id,omitempty" bson:"id,omitempty"`
	Name     string      `json:"name,omitempty" bson:"name,omitempty"`
	Location *Location   `json:"location,omitempty" bson:"location,omitempty"`
	Regions  []*Region   `json:"regions,omitempty" bson:"regions,omitempty"`
	Meta     interface{} `json:"meta,omitempty" bson:"meta,omitempty"`
}

func (s *Station) toM() *mStation {
	return &mStation{
		Typed:    typedStation,
		Id:       s.Id,
		Name:     s.Name,
		Location: s.Location,
		Regions:  s.Regions,
		Meta:     s.Meta,
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

func (s *Station) UnmarshalBSONValue(typ bsontype.Type, data []byte) error {
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
	var m mStation
	err := bson.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	s.fromM(&m)
	return nil
}

func (s Station) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if s.Partial {
		return bson.MarshalValue(s.Id)
	}
	return bson.MarshalValue(s.toM())
}

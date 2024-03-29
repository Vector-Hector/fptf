package fptf

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type Operator struct {
	Id   string
	Name string
	Meta interface{} // any additional data

	Partial bool // only show the id in the json response?
}

// intermediate format used by marshal. this is to work out the partial and type part.
type mOperator struct {
	Id   string      `json:"id,omitempty" bson:"id,omitempty"`
	Name string      `json:"name,omitempty" bson:"name,omitempty"`
	Meta interface{} `json:"meta,omitempty" bson:"meta,omitempty"`

	Typed `bson:"inline"`
}

func (o *Operator) toM() *mOperator {
	return &mOperator{
		Typed: typedOperator,
		Id:    o.Id,
		Name:  o.Name,
		Meta:  o.Meta,
	}
}

func (o *Operator) fromM(m *mOperator) {
	o.Id = m.Id
	o.Name = m.Name
	o.Meta = m.Meta
}

// as it is optional to give either operator id or Operator object,
// we have to unmarshal|marshal it ourselves.
func (o *Operator) UnmarshalJSON(data []byte) error {
	var id string
	if err := json.Unmarshal(data, &id); err == nil {
		o.Id = id
		o.Partial = true
		return nil
	}
	o.Partial = false
	var m mOperator
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	o.fromM(&m)
	return nil

}

func (o *Operator) MarshalJSON() ([]byte, error) {
	if o.Partial {
		return json.Marshal(o.Id)
	}
	return json.Marshal(o.toM())
}

func (o *Operator) UnmarshalBSONValue(typ bsontype.Type, data []byte) error {
	if typ == bson.TypeString {
		var id string
		err := bson.UnmarshalValue(bson.TypeString, data, &id)
		if err != nil {
			return err
		}
		o.Id = id
		o.Partial = true
		return nil
	}
	o.Partial = false
	var m mOperator
	err := bson.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	o.fromM(&m)
	return nil
}

func (o Operator) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if o.Partial {
		return bson.MarshalValue(o.Id)
	}
	return bson.MarshalValue(o.toM())
}

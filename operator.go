package fptf

import "encoding/json"

type Operator struct {
	_Operator
	Partial bool `json:"-"` // only show the id in the json response?
}

type _Operator struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewOperator(id string, name string) Operator {
	return Operator{
		_Operator: _Operator{
			Type: "operator",
			Id:   id,
			Name: name,
		},
		Partial:   false,
	}
}

// as it is optional to give either operator id or Operator object,
// we have to unmarshal|marshal it ourselves.
func (w *Operator) UnmarshalJSON(data []byte) error {
	var id string
	if err := json.Unmarshal(data, &id); err == nil {
		w.Id = id
		w.Partial = true
		return nil
	}
	w.Partial = false
	return json.Unmarshal(data, &w._Operator)
}

func (w *Operator) MarshalJSON() ([]byte, error) {
	if w.Partial {
		return json.Marshal(w.Id)
	}
	return json.Marshal(w._Operator)
}

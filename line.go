package fptf

import "encoding/json"

type Line struct {
	_Line
	Partial bool `json:"-"` // only show the id in the json response?
}

type _Line struct {
	Type     string    `json:"type"`
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Mode     Mode      `json:"mode"`
	SubMode  string    `json:"subMode,omitempty"`
	Routes   []*Route  `json:"routes,omitempty"`
	Operator *Operator `json:"operator"`
}

// as it is optional to give either line id or Line object,
// we have to unmarshal|marshal it ourselves.
func (w *Line) UnmarshalJSON(data []byte) error {
	var id string
	if err := json.Unmarshal(data, &id); err == nil {
		w.Id = id
		w.Partial = true
		return nil
	}
	w.Partial = false
	return json.Unmarshal(data, &w._Line)
}

func (w *Line) MarshalJSON() ([]byte, error) {
	if w.Partial {
		return json.Marshal(w.Id)
	}
	return json.Marshal(w._Line)
}

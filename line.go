package fptf

import "encoding/json"

type Line struct {
	Id       string
	Name     string
	Mode     Mode
	SubMode  string
	Routes   []*Route
	Operator *Operator
	Meta     interface{} // any additional data

	Partial bool // only show the id in the json response?
}

type mLine struct {
	typed
	Id       string      `json:"id"`
	Name     string      `json:"name"`
	Mode     Mode        `json:"mode"`
	SubMode  string      `json:"subMode,omitempty"`
	Routes   []*Route    `json:"routes,omitempty"`
	Operator *Operator   `json:"operator"`
	Meta     interface{} `json:"meta,omitempty"`
}

func (w *Line) toM() *mLine {
	return &mLine{
		typed:    typedLine,
		Id:       w.Id,
		Name:     w.Name,
		Mode:     w.Mode,
		SubMode:  w.SubMode,
		Routes:   w.Routes,
		Operator: w.Operator,
		Meta:     w.Meta,
	}
}

func (w *Line) fromM(m *mLine) {
	w.Id = m.Id
	w.Name = m.Name
	w.Mode = m.Mode
	w.SubMode = m.SubMode
	w.Routes = m.Routes
	w.Operator = m.Operator
	w.Meta = m.Meta
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
	var m mLine
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	w.fromM(&m)
	return nil
}

func (w *Line) MarshalJSON() ([]byte, error) {
	if w.Partial {
		return json.Marshal(w.Id)
	}
	return json.Marshal(w.toM())
}

package fptf

import "encoding/json"

// A Route represents a single set of stations, of a single Line.
//
// For a very consistent subway service, there may be one route
// for each direction. Planned detours, trains stopping early and
// additional directions would have their own route.
type Route struct {
	Id      string
	Line    *Line
	Mode    Mode
	SubMode string
	Stops   []*Stop
	Meta    interface{} // any additional data
}

// used by marshal
type mRoute struct {
	typed
	Id      string      `json:"id"`
	Line    *Line       `json:"line"`
	Mode    Mode        `json:"mode,omitempty"`
	SubMode string      `json:"subMode,omitempty"`
	Stops   []*Stop     `json:"stops"`
	Meta    interface{} `json:"meta,omitempty"`
}

func (w *Route) toM() *mRoute {
	return &mRoute{
		typed:   typedRoute,
		Id:      w.Id,
		Line:    w.Line,
		Mode:    w.Mode,
		SubMode: w.SubMode,
		Stops:   w.Stops,
		Meta:    w.Meta,
	}
}

func (w *Route) fromM(m *mRoute) {
	w.Id = m.Id
	w.Line = m.Line
	w.Mode = m.Mode
	w.SubMode = m.SubMode
	w.Stops = m.Stops
	w.Meta = m.Meta
}

func (w *Route) UnmarshalJSON(data []byte) error {
	var m mRoute
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	w.fromM(&m)
	return nil
}

func (w *Route) MarshalJSON() ([]byte, error) {
	return json.Marshal(w.toM())
}

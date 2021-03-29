package fptf

import "encoding/json"

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
	_Region
	Partial bool `json:"-"` // only show the id in the json response?
}

type _Region struct {
	Type     string     `json:"type"`
	Id       string     `json:"id"`
	Name     string     `json:"name"`
	Stations []*Station `json:"stations"`
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
	return json.Unmarshal(data, &w._Region)
}

func (w *Region) MarshalJSON() ([]byte, error) {
	if w.Partial {
		return json.Marshal(w.Id)
	}
	return json.Marshal(w._Region)
}

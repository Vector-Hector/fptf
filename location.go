package fptf

import "encoding/json"

type Location struct {
	Name      string
	Address   string
	Longitude float64
	Latitude  float64
	Altitude  float64
	Meta      interface{}
}

type mLocation struct {
	typed
	Name      string      `json:"name,omitempty"`
	Address   string      `json:"address,omitempty"`
	Longitude float64     `json:"longitude,omitempty"`
	Latitude  float64     `json:"latitude,omitempty"`
	Altitude  float64     `json:"altitude,omitempty"`
	Meta      interface{} `json:"meta,omitempty"`
}

func (w *Location) toM() *mLocation {
	return &mLocation{
		typed:     typedLocation,
		Name:      w.Name,
		Address:   w.Address,
		Longitude: w.Longitude,
		Latitude:  w.Latitude,
		Altitude:  w.Altitude,
		Meta:      w.Meta,
	}
}

func (w *Location) fromM(m *mLocation) {
	w.Name = m.Name
	w.Address = m.Address
	w.Longitude = m.Longitude
	w.Latitude = m.Latitude
	w.Altitude = m.Altitude
	w.Meta = m.Meta
}

func (w *Location) UnmarshalJSON(data []byte) error {
	var m mLocation
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	w.fromM(&m)
	return nil
}

func (w *Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(w.toM())
}

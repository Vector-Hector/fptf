package fptf

type Location struct {
	Type      string  `json:"type"`
	Name      string  `json:"name,omitempty"`
	Address   string  `json:"address,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Altitude  float64 `json:"altitude,omitempty"`
}

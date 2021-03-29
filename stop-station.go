package fptf

type StopStation struct {
	Stop
	Station

	IsStation bool `json:"-"`
}

// as it is optional to give either stop|station id or Stop or Station object,
// we have to unmarshal|marshal it ourselves.
func (w *StopStation) UnmarshalJSON(data []byte) error {
	var station Station
	if err := station.UnmarshalJSON(data); err == nil && (station.Type == "station" || station.Partial) {
		w.Station = station
		w.IsStation = true
		return nil
	}

	w.IsStation = false
	return w.Stop.UnmarshalJSON(data)
}

func (w *StopStation) MarshalJSON() ([]byte, error) {
	if w.IsStation {
		return w.Station.MarshalJSON()
	}
	return w.Stop.MarshalJSON()
}

package fptf

import (
	"encoding/json"
	"errors"
)

type StopStation struct {
	Stop    *Stop    // if it is a stop, this will be not Stop{}
	Station *Station // if it is a station, this will be not Station{}
	Id      *string  // if it is just an id, this will be not ""
}

func (s *Station) ToStopStation() *StopStation {
	return &StopStation{
		Station:   s,
	}
}

func (s *Stop) ToStopStation() *StopStation {
	return &StopStation{
		Stop:      s,
	}
}

// as it is optional to give either stop|station id or Stop or Station object,
// we have to unmarshal|marshal it ourselves.
func (w *StopStation) UnmarshalJSON(data []byte) error {
	var id string
	if err := json.Unmarshal(data, &id); err == nil {
		w.Id = &id
		return nil
	}

	var station mStation
	if err := json.Unmarshal(data, &station); err == nil && station.Type == objectTypeStation {
		w.Station = new(Station)
		w.Station.fromM(&station)
		return nil
	}

	var stop mStop
	if err := json.Unmarshal(data, &stop); err == nil && stop.Type == objectTypeStop {
		w.Stop = new(Stop)
		w.Stop.fromM(&stop)
		return nil
	}

	return errors.New("could not unmarshall to any of type string, station or stop")
}

func (w *StopStation) MarshalJSON() ([]byte, error) {
	if w.Id != nil {
		return json.Marshal(w.Id)
	}
	if w.Station != nil {
		return json.Marshal(w.Station)
	}
	if w.Stop != nil {
		return json.Marshal(w.Stop)
	}
	return []byte("null"), nil
}

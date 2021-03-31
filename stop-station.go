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

func (s *StopStation) GetLocation() *Location {
	if s.Stop != nil {
		return s.Stop.Location
	}
	if s.Station != nil {
		return s.Station.Location
	}
	return nil
}

// as it is optional to give either stop|station id or Stop or Station object,
// we have to unmarshal|marshal it ourselves.
func (s *StopStation) UnmarshalJSON(data []byte) error {
	var id string
	if err := json.Unmarshal(data, &id); err == nil {
		s.Id = &id
		return nil
	}

	var station mStation
	if err := json.Unmarshal(data, &station); err == nil && station.Type == objectTypeStation {
		s.Station = new(Station)
		s.Station.fromM(&station)
		return nil
	}

	var stop mStop
	if err := json.Unmarshal(data, &stop); err == nil && stop.Type == objectTypeStop {
		s.Stop = new(Stop)
		s.Stop.fromM(&stop)
		return nil
	}

	return errors.New("could not unmarshall to any of type string, station or stop")
}

func (s *StopStation) MarshalJSON() ([]byte, error) {
	if s.Id != nil {
		return json.Marshal(s.Id)
	}
	if s.Station != nil {
		return json.Marshal(s.Station)
	}
	if s.Stop != nil {
		return json.Marshal(s.Stop)
	}
	return []byte("null"), nil
}

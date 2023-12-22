package fptf

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type StopStation struct {
	Stop    *Stop    // if it is a stop, this will be not Stop{}
	Station *Station // if it is a station, this will be not Station{}
	Id      *string  // if it is just an id, this will be not ""
}

func (s *Station) ToStopStation() *StopStation {
	return &StopStation{
		Station: s,
	}
}

func (s *Stop) ToStopStation() *StopStation {
	return &StopStation{
		Stop: s,
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

func (s *StopStation) SetLocation(loc *Location) {
	if s.Stop != nil {
		s.Stop.Location = loc
		return
	}
	if s.Station != nil {
		s.Station.Location = loc
		return
	}
	if s.Id != nil {
		s.Station = &Station{
			Id:       *s.Id,
			Location: loc,
		}
		s.Id = nil
	}
}

func (s *StopStation) GetName() string {
	if s.Stop != nil {
		if s.Stop.Name != "" {
			return s.Stop.Name
		}

		if s.Stop.Station != nil {
			if stationName := s.Stop.Station.ToStopStation().GetName(); stationName != "" {
				return stationName
			}
		}

		if s.Stop.Location != nil && s.Stop.Location.Name != "" {
			return s.Stop.Location.Name
		}
	}
	if s.Station != nil {
		if s.Station.Name != "" {
			return s.Station.Name
		}

		if s.Station.Location != nil && s.Station.Location.Name != "" {
			return s.Station.Location.Name
		}
	}
	if s.Id != nil {
		return *s.Id
	}
	return ""
}

func (s *StopStation) SetName(name string) {
	if s.Stop != nil {
		s.Stop.Name = name
		return
	}
	if s.Station != nil {
		s.Station.Name = name
		return
	}
	if s.Id != nil {
		s.Station = &Station{
			Id:   *s.Id,
			Name: name,
		}
		s.Id = nil
	}
}

func (s *StopStation) GetId() string {
	if s.Stop != nil {
		return s.Stop.Id
	}

	if s.Station != nil {
		return s.Station.Id
	}

	if s.Id != nil {
		return *s.Id
	}
	return ""
}

func (s *StopStation) SetId(id string) {
	if s.Stop != nil {
		s.Stop.Id = id
		return
	}

	if s.Station != nil {
		s.Station.Id = id
		return
	}

	if s.Id != nil {
		s.Id = &id
		return
	}
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

func (s *StopStation) UnmarshalBSONValue(typ bsontype.Type, data []byte) error {
	if typ == bson.TypeString {
		var id string
		err := bson.UnmarshalValue(bson.TypeString, data, &id)
		if err != nil {
			return err
		}
		s.Id = &id
		return nil
	}

	var station mStation
	stationErr := bson.Unmarshal(data, &station)
	if stationErr == nil && station.Type == objectTypeStation {
		s.Station = new(Station)
		s.Station.fromM(&station)
		return nil
	}

	var stop mStop
	stopErr := bson.Unmarshal(data, &stop)
	if stopErr == nil && stop.Type == objectTypeStop {
		s.Stop = new(Stop)
		s.Stop.fromM(&stop)
		return nil
	}

	var m bson.M
	err := bson.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	return fmt.Errorf("could not unmarshall to any of type string, station or stop. stationErr: %v, stopErr: %v, data type: %v, fptf type: %s, object as map: %v", stationErr, stopErr, typ, string(stop.Type), m)
}

func (s StopStation) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if s.Id != nil {
		return bson.MarshalValue(s.Id)
	}
	if s.Station != nil {
		return s.Station.MarshalBSONValue()
	}
	if s.Stop != nil {
		return s.Stop.MarshalBSONValue()
	}
	return bson.MarshalValue(nil)
}

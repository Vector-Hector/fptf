package fptf

// describes, that type of json object it is. set by marshal
type objectType string

const (
	objectTypeLocation objectType = "location"
	objectTypeStation  objectType = "station"
	objectTypeStop     objectType = "stop"
	objectTypeRegion   objectType = "region"
	objectTypeLine     objectType = "line"
	objectTypeRoute    objectType = "route"
	objectTypeSchedule objectType = "schedule"
	objectTypeOperator objectType = "operator"
	objectTypeStopover objectType = "stopover"
	objectTypeJourney  objectType = "journey"
)

type Typed struct {
	Type objectType `json:"type,omitempty" bson:"type,omitempty"`
}

var (
	typedLocation = Typed{Type: objectTypeLocation}
	typedStation  = Typed{Type: objectTypeStation}
	typedStop     = Typed{Type: objectTypeStop}
	typedRegion   = Typed{Type: objectTypeRegion}
	typedLine     = Typed{Type: objectTypeLine}
	typedRoute    = Typed{Type: objectTypeRoute}
	typedSchedule = Typed{Type: objectTypeSchedule}
	typedOperator = Typed{Type: objectTypeOperator}
	typedStopover = Typed{Type: objectTypeStopover}
	typedJourney  = Typed{Type: objectTypeJourney}
)

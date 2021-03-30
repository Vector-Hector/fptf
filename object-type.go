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

type typed struct {
	Type objectType `json:"type"`
}

var (
	typedLocation = typed{Type: objectTypeLocation}
	typedStation  = typed{Type: objectTypeStation}
	typedStop     = typed{Type: objectTypeStop}
	typedRegion   = typed{Type: objectTypeRegion}
	typedLine     = typed{Type: objectTypeLine}
	typedRoute    = typed{Type: objectTypeRoute}
	typedSchedule = typed{Type: objectTypeSchedule}
	typedOperator = typed{Type: objectTypeOperator}
	typedStopover = typed{Type: objectTypeStopover}
	typedJourney  = typed{Type: objectTypeJourney}
)

package fptf

type Mode string

// As discussed in #4, we decided to have two fields mode and subMode.
// The following list shows all possible values for a mode property.
// For consumers to be able to use mode meaningfully, we will keep
// this list very short.
//
// In order to convey more details, we will add the subMode field in
// the future. It will differentiate means of transport in a more
// fine-grained way, in order to enable consumers to provide more
// context and a better service.
const (
	ModeTrain      = Mode("train")
	ModeBus        = Mode("bus")
	ModeWatercraft = Mode("watercraft")
	ModeTaxi       = Mode("taxi")
	ModeGondola    = Mode("gondola")
	ModeAircraft   = Mode("aircraft")
	ModeCar        = Mode("car")
	ModeBicycle    = Mode("bicycle")
	ModeWalking    = Mode("walking")
)

package test

import (
	"encoding/json"
	"fmt"
	"github.com/Vector-Hector/friendly-public-transport-format"
	"io/ioutil"
	"testing"
)

func TestParseInvalidJourney(t *testing.T) {
	_, err := getJourneyFromFile("invalid-journey.json")
	if err != nil {
		return
	}
	t.Error("Invalid journey parsed as valid")
}

func TestParseValidJourney(t *testing.T) {
	_, err := getJourneyFromFile("valid-journey.json")
	if err != nil {
		t.Error(err)
	}
}

func TestParseValidSimpleJourney(t *testing.T) {
	_, err := getJourneyFromFile("valid-simple-journey.json")
	if err != nil {
		t.Error(err)
	}
}

func TestWriteValidJourney(t *testing.T) {
	testRewriteJourney(t, "valid-journey.json")
}

func TestWriteSimpleValidJourney(t *testing.T) {
	testRewriteJourney(t, "valid-simple-journey.json")
}

func testRewriteJourney(t *testing.T, file string) {
	rawDat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	journey, err := getJourneyFromBytes(rawDat)
	if err != nil {t.Error(err)}

	remarshalledDat, err := json.Marshal(journey)
	if err != nil {
		t.Error(err)
	}

	var journeyRawObj interface{}
	err = json.Unmarshal(rawDat, &journeyRawObj)
	if err != nil {t.Error(err)}

	var remarshalledRawObj interface{}
	err = json.Unmarshal(remarshalledDat, &remarshalledRawObj)
	if err != nil {t.Error(err)}

		fmt.Println("Remarshalled", file, ":")
		fmt.Println(string(remarshalledDat))
	if !deepEqual(journeyRawObj, remarshalledRawObj) {
		t.Error("Marshalling the parsed data did not give the original data")
	}
}

func getJourneyFromBytes(dat []byte) (*fptf.Journey, error) {
	var journey fptf.Journey
	err := json.Unmarshal(dat, &journey)
	if err != nil {
		return nil, err
	}
	return &journey, nil
}

func getJourneyFromFile(file string) (*fptf.Journey, error) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return getJourneyFromBytes(dat)
}

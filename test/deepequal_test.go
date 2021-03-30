package test

import (
	"testing"
)

func TestDeepEqual(t *testing.T) {
	m1 := map[string]string{
		"a": "",
		"b": "hello",
	}
	m2 := map[string]string{
		"b": "hello",
	}

	if !deepEqual(m1, m2) {
		t.Error("Failed equality of", m1, "and", m2)
	}
}

func TestDeepEqual2(t *testing.T) {
	m1 := map[string]string{
		"b": "hello2",
	}
	m2 := map[string]string{
		"b": "hello",
	}

	if deepEqual(m1, m2) {
		t.Error("Failed inequality of", m1, "and", m2)
	}
}

func TestDeepEqual3(t *testing.T) {
	m1 := map[string]interface{}{
		"b": "hello2",
		"a": map[string]string{

		},
	}
	m2 := map[string]interface{}{
		"b": "hello2",
		"a": nil,
	}

	if !deepEqual(m1, m2) {
		t.Error("Failed equality of", m1, "and", m2)
	}
}

func TestDeepEqual4(t *testing.T) {
	m1 := map[string]interface{}{
		"b": "hello2",
		"a": 0,
	}
	m2 := map[string]interface{}{
		"b": "hello2",
		"a": nil,
	}

	if !deepEqual(m1, m2) {
		t.Error("Failed equality of", m1, "and", m2)
	}
}

package test

import (
	"fmt"
	"reflect"
	"time"
)

func deepEqual(v1 interface{}, v2 interface{}) bool {
	return deepValueEqual(reflect.ValueOf(v1), reflect.ValueOf(v2))
}

// tests, if a is equivalent to b. That is it's reflect.DeepEqual but we accept, if there are zero values in between
// as we only look at json unmarshalled values, we can ignore the recursive type problem
func deepValueEqual(v1 reflect.Value, v2 reflect.Value) bool {
	if isZero(v1) {
		bothZero := isZero(v2)
		if !bothZero {
			fmt.Println("Deep equal failed at value comparing values\n", v1, "\n", v2, "(v1 is zero, v2 is not)")
		}
		return bothZero
	}
	if isZero(v2) {
		fmt.Println("Deep equal failed at value comparing values\n", v1, "\n", v2, "(v2 is zero, v1 is not)")
		return false
	}

	switch v1.Kind() {
	case reflect.Array:
		for i := 0; i < v1.Len(); i++ {
			if !deepValueEqual(v1.Index(i), v2.Index(i)) {
				//fmt.Println("Deep equal failed at array comparing values\n", v1.Index(i), "\n", v2.Index(i))
				return false
			}
		}
		return true
	case reflect.Slice:
		if v1.IsNil() != v2.IsNil() {
			fmt.Println("Deep equal failed at slice comparing values\n", v1, "\n", v2, "(one is nil, the other is not)")
			return false
		}
		if v1.Len() != v2.Len() {
			fmt.Println("Deep equal failed at slice comparing values\n", v1, "\n", v2, "(lengths differ)")
			return false
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		for i := 0; i < v1.Len(); i++ {
			if !deepValueEqual(v1.Index(i), v2.Index(i)) {
				//fmt.Println("Deep equal failed at slice comparing values\n", v1.Index(i), "\n", v2.Index(i))
				return false
			}
		}
		return true
	case reflect.Interface:
		if v1.IsNil() || v2.IsNil() {
			bothNil := v1.IsNil() == v2.IsNil()
			if !bothNil {
				fmt.Println("Deep equal failed at interface comparing values\n", v1, "\n", v2, "(one is nil, the other is not)")
			}
			return bothNil
		}
		return deepValueEqual(v1.Elem(), v2.Elem())
	case reflect.Ptr:
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		return deepValueEqual(v1.Elem(), v2.Elem())
	case reflect.Struct:
		for i, n := 0, v1.NumField(); i < n; i++ {
			if !deepValueEqual(v1.Field(i), v2.Field(i)) {
				return false
			}
		}
		return true
	case reflect.Map:
		if v1.IsNil() != v2.IsNil() {
			return false
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		for _, k := range v1.MapKeys() {
			val1 := v1.MapIndex(k)
			val2 := v2.MapIndex(k)
			if !deepValueEqual(val1, val2) {
				//fmt.Println("Deep equal failed at map comparing values\n", val1, "\n", val2)
				return false
			}
		}
		return true
	default:
		t1, v1IsTime := getTime(v1)
		t2, v2IsTime := getTime(v2)

		if v1IsTime && v2IsTime {
			eq := t1.Equal(t2)
			if !eq {
				fmt.Println("Deep equal failed at time comparing values\n", v1, "\n", v2)
			}
			return eq
		}

		interfacesEqual := v1.Interface() == v2.Interface()
		if !interfacesEqual {
			fmt.Println("Deep equal failed at default comparing values\n", v1, "\n", v2, "(interfaces not equal)")
		}
		return interfacesEqual
	}
}

func getTime(val reflect.Value) (time.Time, bool) {
	if val.Type().String() == "time.Time" {
		return val.Interface().(time.Time), true
	}

	if val.Type().String() == "string" {
		t, err := time.Parse(time.RFC3339, val.Interface().(string))
		if err != nil {
			return time.Time{}, false
		}
		return t, true
	}

	return time.Time{}, false
}

func isZero(val reflect.Value) bool {
	if val.Kind() == reflect.Ptr {
		return isZero(val.Elem())
	}

	if !val.IsValid() || val.IsZero() {
		return true
	}

	if val.CanInterface() {
		val = reflect.ValueOf(val.Interface())

		if !val.IsValid() || val.IsZero() {
			return true
		}
	}

	switch val.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return val.Len() == 0
	}
	return false
}

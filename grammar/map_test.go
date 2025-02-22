package grammar

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	mapV := make(map[string]string)
	mapV["Key1"] = "1"
	mapV["Key2"] = "2"

	for s := range mapV {
		println(s + ": " + mapV[s])
	}

	delete(mapV, "Key2")

	for s, s2 := range mapV {
		println(s + ": " + s2)
	}

	s, ok := mapV["Key1"]
	println(s, ok)

	typeS := reflect.TypeOf(s)
	println(typeS.Name())
}

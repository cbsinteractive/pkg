package timecode

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestJSON(t *testing.T) {
	s0 := Splice{{1, 5}}
	v, err := json.Marshal(s0)
	if err != nil {
		t.Fatal(err)
	}
	s1 := Splice{}
	if err = json.Unmarshal(v, &s1); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s0, s1) {
		t.Fatalf("bad codec roundtrip: have %q want %q", s1, s0)
	}
}

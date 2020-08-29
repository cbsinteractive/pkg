package timecode

import "testing"

func TestRangeTimecode(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  Range
	}{
		{"0s", "00:00:00:00", Range{0, 0}},
		{"1s", "00:00:01:00", Range{0, 1}},
		{"59s", "00:00:59:00", Range{0, 59}},
		// TODO(as): These don't work right, but they aren't a priority to fix
		{"frame59", "00:00:00:59", Range{0, 1}},
		{"frame60", "00:00:00:60", Range{0, 1}},
		{"frame600", "00:00:00:600", Range{0, 10}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			have, err := Parse(tc.input, 60)
			if err != nil {
				t.Fatalf("parse error: %v", err)
			}
			if have.String() != tc.want.String() {
				t.Fatalf("have %v want %v", have, tc.want)
			}
			if have.Timecode(60) != tc.input {
				have := have.Timecode(60)
				t.Fatalf("timecode and input dont match: have %v want %v", have, tc.want)
			}
		})
	}
}

// This API doesn't know about AWS InputClippings directly, this is just a due-dilligence test
// to ensure the conversion works as intended with previously encountered values
func TestRangeToInputClippings(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input Range
		want  [2]string
	}{
		{"0", Range{0, 0}, [2]string{"00:00:00:00", "00:00:00:00"}},
		{"5s-10s", Range{5, 10}, [2]string{"00:00:05:00", "00:00:10:00"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s, e := tc.input.Timecodes(60)
			if s != tc.want[0] || e != tc.want[1] {
				t.Fatalf("timecodes and input dont match: have %v want %v", s+"-"+e, tc.want[0]+"-"+tc.want[1])
			}
		})
	}
}

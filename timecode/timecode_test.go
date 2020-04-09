package timecode

import (
	"fmt"
	"testing"
)

func TestBasicExample(t *testing.T) {
	// TODO(as): turn this into a real test
	video, _ := ParseTimecode("01:00:00:00", 60)
	splice := Splice{{5, 10}, {25, 30}}
	checkSplice(video, splice)

	video, _ = ParseTimecode("00:00:10:00", 60)
	splice = Splice{{25, 30}, {5, 10}}
	checkSplice(video, splice)

	video, _ = ParseTimecode("00:01:11:05", 60)
	splice = Splice{{25, 30}, {5, 10}}
	checkSplice(video, splice)
}

func TestUnmarshal(t *testing.T) {
	want := Splice{{5, 10}}
	have := Splice{}
	have.UnmarshalText([]byte("[[5.0, 10.0]]"))
	if len(have) != len(want) || have[0] != want[0] {
		t.Fatalf("have %v, want %v", have, want)
	}
}

func TestUnmarshal3(t *testing.T) {
	// TODO(as): test table
	want := Splice{{5, 10}}
	have := Splice{}
	have.UnmarshalText([]byte("[[5.0, 10, 5]]"))
	if len(have) != len(want) || have[0] != want[0] {
		t.Fatalf("have %v, want %v", have, want)
	}
}

func checkSplice(video Range, s Splice) {
	return
	fmt.Println()
	fmt.Println("in bounds?", s.In(video))
	fmt.Println("minimum source input range for processing:", s.Union())
	fmt.Println("video range", video)
	fmt.Println("splice", s, "sorted?", s.Sorted())
	fmt.Println("output duration:", s.Size())
	fmt.Println("video timecode", video.Timecode(60))
}

func ExampleUsage() {
	r, _ := ParseTimecode("00:00:10:00", 60)
	s := Splice{{1, 5}}
	fmt.Println("video timecode", r.Timecode(60))
	fmt.Println(s, "in", r, "?", s.In(r))
	fmt.Println("splice sorted?", s.Sorted())
	fmt.Println("input: minimum span", s.Union())
	fmt.Println("output length", s.Size())

	// Output: video timecode 00:00:10:00
	// [(1s-5s)] in (0s-11s) ? true
	// splice sorted? true
	// input: minimum span (1s-5s)
	// output length 4s
}

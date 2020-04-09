package timecode

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

const defaultFps = 23.997

// Splice is a list of Ranges
type Splice []Range

func (s Splice) Len() int      { return len(s) }
func (s Splice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Splice) Less(i, j int) bool {
	if s[i][0] < s[j][0] {
		return true
	}
	return s[i][0] == s[j][0] && s[i].Size() < s[j].Size()
}

// Size returns the cummulative duration of the splice
func (s Splice) Size() (dt time.Duration) {
	for _, r := range s {
		dt += r.Size()
	}
	return dt

}

// Union returns the smallest Range that contains s
func (s Splice) Union() Range {
	if s.Len() == 0 {
		return Range{}
	}
	u := s[0]
	for _, r := range s[1:] {
		if r[0] < u[0] {
			u[0] = r[0]
		}
	}
	for _, r := range s[1:] {
		if r[1] > u[1] {
			u[1] = r[1]
		}
	}
	return u
}

// In returns true if the splice is contained by r
func (c Splice) In(r Range) bool {
	for _, c := range c {
		if c[0] < r[0] || c[1] > r[1] {
			return false
		}
	}
	return true
}

// Sort returns true if the splice is sorted
func (s Splice) Sorted() bool {
	return sort.IsSorted(s)
}

// UnmarshalText unmarshals the splice into s in the format of
// a two-dimensional JSON array of tuples: [[%f,%f], [%f,%f], ... [%f, %f]
func (s *Splice) UnmarshalText(p []byte) error {
	if len(p) == 0 {
		return nil
	}
	// NOTE(as): Technically, this will accept strange input like ranges
	// with three numbers by filling in [2]float64 with the first and second
	// value for the range; do we want to be more strict about this? What
	// does videorobot do?
	return json.Unmarshal(p, (*[]Range)(s))
}

// ParseTimecode parses an input string in HH:MM:SS:FF, HH:MM:SS;FF, or
// HH:MM:SS format, defined by the following convention
// HH = hour, MM = minute, SS = second, and FF is the frame number, the
// frameRate argument is either 0, or a fractional frame rate upon which
// to calculate the precise Range value based on the FF argument, if present.
func ParseTimecode(tc string, fps float64) (Range, error) {
	if fps == 0 {
		fps = defaultFps
	}
	var h, m, s, f float64
	n, err := fmt.Sscanf(tc, "%f:%f:%f:%f", &h, &m, &s, &f)
	if n < 3 {
		n, _ = fmt.Sscanf(tc, "%f:%f:%f;%f", &h, &m, &s, &f)
	}
	if n < 3 {
		return Range{}, err
	}
	if f == 0 {
		f = fps // avoid NaN condition
	}
	return Range{0, h*3600 + m*60 + s + (fps / f)}, err

}

package timecode

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

const defaultFps = 23.997

// Range is a pair of decimal seconds defining a time interval
// starting at Range[0] and ending at Range[1]
type Range [2]float64

// Canon returns the range in proper order, where r[0] <= r[1]
func (r Range) Canon() Range {
	if r[0] > r[1] {
		r[0], r[1] = r[1], r[0]
	}
	return r
}

// Size returns the duration of the Range
func (r Range) Size() time.Duration {
	dx := r[1] - r[0]
	if dx < 0 {
		dx = -dx
	}
	return time.Duration(dx * float64(time.Second))
}

func (r Range) String() string {
	return fmt.Sprintf("(%s-%s)", time.Duration(r[0])*time.Second, time.Duration(r[1])*time.Second)
}

// Timecode outputs the timecode in HH:MM:SS:FF format
func (r Range) Timecode(fps float64) string {
	if fps == 0 {
		fps = defaultFps
	}
	d := int64(r[1])
	h := d / 3600
	d %= 3600
	m := d / 60
	m %= 60
	s := d % 60
	f := 0 // TODO(as): frame number

	return fmt.Sprintf("%02d:%02d:%02d:%02d", h, m, s, f)
}

func (r Range) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("[%f,%f]", r[0], r[1])), nil
}

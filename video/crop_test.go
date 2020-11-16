package video

import (
	"image"
	"testing"
)

var (
	xvga  = R(0, 0, 640, 480)
	x720  = R(0, 0, 1280, 720)
	x1080 = R(0, 0, 1920, 1080)
)

var R = image.Rect

func TestCrop(t *testing.T) {
	for i, tt := range []struct {
		src, crop image.Rectangle
		want      Crop
	}{
		{R(0, 0, 100, 100), R(0, 0, 0, 0), Crop{0, 0, 100, 100}},
		{R(0, 0, 100, 100), R(0, 0, 100, 100), Crop{0, 0, 0, 0}},
		{R(0, 0, 100, 100), R(10, 10, 90, 90), Crop{10, 10, 10, 10}},
		{R(0, 0, 1920, 1080), R(0, 0, 1280, 720), Crop{0, 0, 640, 360}},
	} {
		c := Crop{}
		c.From(tt.src, tt.crop)
		if have := c; have != tt.want {
			t.Fatalf("%d: src %s crop %s: bad crop:\n\t\thave: %+v\n\t\twant: %+v", i, tt.src, tt.crop, have, tt.want)
		}
		if have := c.Rect(tt.src); have != tt.crop {
			t.Fatalf("%d: src %s crop %s: bad bijection:\n\t\thave: %v\n\t\twant: %v", i, tt.src, tt.crop, have, tt.crop)
		}
	}
}

func TestScale(t *testing.T) {
	for i, tt := range []struct {
		src, ffmpeg, crop image.Rectangle
	}{
		{R(0, 0, 10, 10), R(0, 0, 5, 5), R(0, 0, 5, 5)},
		{R(0, 0, 10, 10), R(0, 0, 5, 10), R(0, 3, 5, 8)},
		{x720, x720, x720},
		{x1080, x1080, x1080},
		{x1080, x720, x720},
		{x1080, x720.Add(image.Pt(10, 10)), R(10, 10, 1290, 730)},
		{x720, R(0, 40, 1280, 680), R(72, 41, 1208, 680)},
		{x1080, R(244, 4, 1684, 1076), R(244, 135, 1684, 945)},
	} {
		have := Scale(tt.src, tt.ffmpeg)
		want := tt.crop
		if have, src := aspect(have), aspect(tt.src); have != src {
			t.Fatalf("%d: bad scale: differs from source aspect\nhave:\t\t%s\nffmpeg:\t\t%s", i, have, src)
		}
		if !have.In(tt.ffmpeg) {
			t.Fatalf("%d: bad scale: outside ffmpeg params\nhave:\t\t%s\nffmpeg:\t\t%s", i, have, tt.ffmpeg)
		}
		if have != want {
			t.Fatalf("%d: bad scale:\nhave:\t\t%s\nwant:\t\t%s", i, have, want)
		}

	}
}

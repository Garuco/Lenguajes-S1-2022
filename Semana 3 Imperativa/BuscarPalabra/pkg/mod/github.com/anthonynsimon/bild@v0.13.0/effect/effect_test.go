package effect

import (
	"image"
	"testing"

	"github.com/anthonynsimon/bild/util"
)

// benchResult is used to avoid having the compiler optimize the benchmark code calls
var benchResult interface{}

func TestSepia(t *testing.T) {
	cases := []struct {
		value    image.Image
		expected *image.RGBA
	}{
		{
			value:    &image.RGBA{},
			expected: &image.RGBA{},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x7f, 0x7f, 0x7f, 0x80, 0x7f, 0x7f, 0x7f, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0xAC, 0x99, 0x77, 0x80, 0xAC, 0x99, 0x77, 0xFF,
					0xFF, 0xFF, 0xEF, 0xFF, 0x0, 0x0, 0x0, 0xFF,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0xAA, 0xCA, 0x7B, 0x80, 0x7f, 0x7f, 0x7f, 0xFF,
					0x31, 0xFF, 0x0, 0xFF, 0x0, 0xFF, 0x0, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0xF5, 0xDB, 0xAA, 0x80, 0xAC, 0x99, 0x77, 0xFF,
					0xD7, 0xC0, 0x95, 0xFF, 0xC4, 0xAF, 0x88, 0xFF,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Sepia(c.value)
		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s:\nexpected:\n%v\nactual:\n%v", "Sepia", util.RGBAToString(c.expected), util.RGBAToString(actual))
		}
	}
}

func TestInvert(t *testing.T) {
	cases := []struct {
		value    image.Image
		expected *image.RGBA
	}{
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x7f, 0x7f, 0x7f, 0x80, 0x7f, 0x7f, 0x7f, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0xFF,
					0x0, 0x0, 0x0, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0xAA, 0xCA, 0x7B, 0x80, 0x7f, 0x7f, 0x7f, 0xFF,
					0x31, 0xFF, 0x0, 0xFF, 0x0, 0xFF, 0x0, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x55, 0x35, 0x84, 0x80, 0x80, 0x80, 0x80, 0xFF,
					0xCE, 0x0, 0xFF, 0xFF, 0xFF, 0x0, 0xFF, 0xFF,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Invert(c.value)
		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s: expected: %#v, actual: %#v", "Invert", c.expected, actual)
		}
	}
}

func TestGrayscale(t *testing.T) {
	cases := []struct {
		value    image.Image
		expected *image.RGBA
	}{
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0xAA, 0xCA, 0x7B, 0x80, 0x7f, 0x7f, 0x7f, 0xFF,
					0x31, 0xFF, 0x0, 0xFF, 0x0, 0xFF, 0x0, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xb9, 0xb9, 0xb9, 0x80, 0x7f, 0x7f, 0x7f, 0xff,
					0xa8, 0xa8, 0xa8, 0xff, 0x99, 0x99, 0x99, 0xff,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0xCA, 0xAA, 0x7B, 0x80, 0x7f, 0x7f, 0x7A, 0xFF,
					0x71, 0xFF, 0x0, 0xFF, 0x0, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xaf, 0xaf, 0xaf, 0x80, 0x7f, 0x7f, 0x7f, 0xff,
					0xbb, 0xbb, 0xbb, 0xff, 0xb3, 0xb3, 0xb3, 0xff,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0xFF, 0x88, 0x88, 0x88, 0xFF,
					0x88, 0x88, 0x88, 0x00, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0xFF, 0x88, 0x88, 0x88, 0xFF,
					0x88, 0x88, 0x88, 0x00, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Grayscale(c.value)
		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s: expected: %#v, actual: %#v", "Grayscale", c.expected, actual)
		}
	}
}

func TestGrayscaleWithWeigths(t *testing.T) {
	cases := []struct {
		value    image.Image
		expected *image.RGBA
	}{
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0xAA, 0x00, 0x7B, 0x80,
					0x7f, 0x00, 0x7f, 0xFF,
					0x31, 0x00, 0x0, 0xFF,
					0x0, 0x00, 0x0, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0x93, 0x93, 0x93, 0x80,
					0x7f, 0x7f, 0x7f, 0xff,
					0x19, 0x19, 0x19, 0xff,
					0x0, 0x0, 0x0, 0xff,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 8,
				Pix: []uint8{
					0xAA, 0xCA, 0x7B, 0x80,
					0x7f, 0x7f, 0x7f, 0xFF,
					0x31, 0xFF, 0x0, 0xFF,
					0x0, 0xFF, 0x0, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0x93, 0x93, 0x93, 0x80,
					0x7f, 0x7f, 0x7f, 0xff,
					0x19, 0x19, 0x19, 0xff,
					0x0, 0x0, 0x0, 0xff,
				},
			},
		},
	}

	for _, c := range cases {
		actual := GrayscaleWithWeights(c.value, 0.5, 0.0, 0.5)
		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s: expected: %#v, actual: %#v", "GrayscaleWithWeigths", c.expected, actual)
		}
	}
}

func TestEdgeDetection(t *testing.T) {
	cases := []struct {
		radius   float64
		value    image.Image
		expected *image.RGBA
	}{
		{
			radius: 0.0,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
			},
		},
		{
			radius: 1.0,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x80, 0x80, 0xFF,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x80, 0x80, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
	}

	for _, c := range cases {
		actual := EdgeDetection(c.value, c.radius)
		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s:\nexpected:%v\nactual:%v\n", "EdgeDetection", util.RGBAToString(c.expected), util.RGBAToString(actual))
		}
	}
}

func TestSobel(t *testing.T) {
	cases := []struct {
		desc     string
		value    image.Image
		expected *image.RGBA
	}{
		{
			desc: "top right corner",
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
		{
			desc: "bottom right corner",
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x40, 0x40, 0x40, 0xff, 0x40, 0x40, 0x40, 0xff,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Sobel(c.value)
		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s: expected: %#v, actual: %#v", "Sobel "+c.desc, c.expected, actual)
		}
	}
}

func TestEmboss(t *testing.T) {
	cases := []struct {
		value    image.Image
		expected *image.RGBA
	}{
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x2, 0x2, 0x2, 0xff,
					0x80, 0x80, 0x80, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Emboss(c.value)
		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s: expected: %#v, actual: %#v", "Emboss", c.expected, actual)
		}
	}
}

func TestMedian(t *testing.T) {
	cases := []struct {
		radius   float64
		value    image.Image
		expected *image.RGBA
	}{
		{
			radius: 0,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
		},
		{
			radius: 1,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
				},
			},
		},
		{
			radius: 2,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0x00, 0xFF, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x0, 0x0, 0xFF, 0xFF, 0x0, 0x0, 0xFF, 0xFF, 0x0, 0x0, 0xFF,
					0xFF, 0x0, 0x0, 0xFF, 0xFF, 0x0, 0x0, 0xFF, 0xFF, 0x0, 0x0, 0xFF,
					0xFF, 0x0, 0x0, 0xFF, 0xFF, 0x0, 0x0, 0xFF, 0xFF, 0x0, 0x0, 0xFF,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Median(c.value, c.radius)
		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s:\nexpected:%v\nactual:%v\n", "Median", util.RGBAToString(c.expected), util.RGBAToString(actual))
		}
	}
}

func TestDilate(t *testing.T) {
	cases := []struct {
		radius   float64
		value    image.Image
		expected *image.RGBA
	}{
		{
			radius: 0,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
		},
		{
			radius: 1,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF, 0x40, 0x00, 0x00, 0xFF,
					0x00, 0xFF, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0x00, 0xFF, 0x00, 0xFF, 0x00, 0xFF, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0x00, 0xFF, 0x00, 0xFF, 0x00, 0xFF, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
				},
			},
		},
		{
			radius: 2,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF, 0x40, 0x00, 0x00, 0xFF,
					0x00, 0xFF, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0xFF, 0x00, 0xFF, 0x00, 0xFF, 0x0, 0xFF, 0x00, 0xFF, 0x0, 0xFF,
					0x00, 0xFF, 0x00, 0xFF, 0x00, 0xFF, 0x0, 0xFF, 0x00, 0xFF, 0x0, 0xFF,
					0x00, 0xFF, 0x00, 0xFF, 0x00, 0xFF, 0x0, 0xFF, 0x00, 0xFF, 0x0, 0xFF,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Dilate(c.value, c.radius)
		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s:\nexpected:%v\nactual:%v\n", "Dilate", util.RGBAToString(c.expected), util.RGBAToString(actual))
		}
	}
}

func TestErode(t *testing.T) {
	cases := []struct {
		radius   float64
		value    image.Image
		expected *image.RGBA
	}{
		{
			radius: 0,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
		},
		{
			radius: 1,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0x00, 0xFF, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0x00, 0xFF, 0x00, 0xFF, 0x20, 0x00, 0x00, 0xFF,
					0x00, 0xFF, 0x00, 0xFF, 0x40, 0x00, 0x00, 0xFF, 0x40, 0x40, 0x00, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF,
					0x40, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF,
					0x40, 0x00, 0x00, 0xFF, 0x20, 0x00, 0x00, 0xFF, 0x20, 0x00, 0x00, 0xFF,
				},
			},
		},
		{
			radius: 2,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0x80, 0x80, 0x80, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0x80, 0xFF, 0xFF, 0xFF, 0x40, 0xFF, 0xFF, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0x60, 0x60, 0x60, 0xFF, 0x60, 0x60, 0x60, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF,
					0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Erode(c.value, c.radius)
		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s:\nexpected:%v\nactual:%v\n", "Erode", util.RGBAToString(c.expected), util.RGBAToString(actual))
		}
	}
}

func TestSharpen(t *testing.T) {
	cases := []struct {
		value    image.Image
		expected *image.RGBA
	}{
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
				},
			},
		},
		{
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0x00, 0xFF, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Sharpen(c.value)
		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s: expected: %#v, actual: %#v", "Sharpen", c.expected, actual)
		}
	}
}

func TestUnsharpMask(t *testing.T) {
	cases := []struct {
		value    image.Image
		radius   float64
		amount   float64
		expected *image.RGBA
	}{
		{
			radius: 0.0,
			amount: 0.0,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0x00, 0xFF, 0xFF, 0x40, 0x40, 0xFF,
					0x40, 0x80, 0x80, 0xFF, 0xB2, 0x00, 0xCA, 0xFF, 0x80, 0x40, 0x80, 0xFF,
					0x00, 0x00, 0xAA, 0xFF, 0xFF, 0xCC, 0xCC, 0xFF, 0xFF, 0x00, 0xAA, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0x00, 0xFF, 0xFF, 0x40, 0x40, 0xFF,
					0x40, 0x80, 0x80, 0xFF, 0xB2, 0x00, 0xCA, 0xFF, 0x80, 0x40, 0x80, 0xFF,
					0x00, 0x00, 0xAA, 0xFF, 0xFF, 0xCC, 0xCC, 0xFF, 0xFF, 0x00, 0xAA, 0xFF,
				},
			},
		},
		{
			radius: 10.0,
			amount: 0.0,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0x00, 0xFF, 0xFF, 0x40, 0x40, 0xFF,
					0x40, 0x80, 0x80, 0xFF, 0xB2, 0x00, 0xCA, 0xFF, 0x80, 0x40, 0x80, 0xFF,
					0x00, 0x00, 0xAA, 0xFF, 0xFF, 0xCC, 0xCC, 0xFF, 0xFF, 0x00, 0xAA, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0x00, 0xFF, 0xFF, 0x40, 0x40, 0xFF,
					0x40, 0x80, 0x80, 0xFF, 0xB2, 0x00, 0xCA, 0xFF, 0x80, 0x40, 0x80, 0xFF,
					0x00, 0x00, 0xAA, 0xFF, 0xFF, 0xCC, 0xCC, 0xFF, 0xFF, 0x00, 0xAA, 0xFF,
				},
			},
		},
		{
			radius: 0.0,
			amount: 10.0,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0x00, 0xFF, 0xFF, 0x40, 0x40, 0xFF,
					0x40, 0x80, 0x80, 0xFF, 0xB2, 0x00, 0xCA, 0xFF, 0x80, 0x40, 0x80, 0xFF,
					0x00, 0x00, 0xAA, 0xFF, 0xFF, 0xCC, 0xCC, 0xFF, 0xFF, 0x00, 0xAA, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0x00, 0xFF, 0xFF, 0x40, 0x40, 0xFF,
					0x40, 0x80, 0x80, 0xFF, 0xB2, 0x00, 0xCA, 0xFF, 0x80, 0x40, 0x80, 0xFF,
					0x00, 0x00, 0xAA, 0xFF, 0xFF, 0xCC, 0xCC, 0xFF, 0xFF, 0x00, 0xAA, 0xFF,
				},
			},
		},
		{
			radius: 1.0,
			amount: 1.0,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0x00, 0xFF, 0xFF, 0x40, 0x40, 0xFF,
					0x40, 0x80, 0x80, 0xFF, 0xB2, 0x00, 0xCA, 0xFF, 0x80, 0x40, 0x80, 0xFF,
					0x00, 0x00, 0xAA, 0xFF, 0xFF, 0xCC, 0xCC, 0xFF, 0xFF, 0x00, 0xAA, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x0, 0x0, 0xFF, 0xFF, 0xFF, 0x0, 0xFF, 0xFF, 0x48, 0x26, 0xFF,
					0x0, 0xD2, 0x9A, 0xFF, 0xAA, 0x0, 0xFF, 0xFF, 0x36, 0x4D, 0x93, 0xFF,
					0x0, 0x0, 0xDA, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0xD5, 0xFF,
				},
			},
		},
		{
			radius: 1.0,
			amount: 0.5,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0x00, 0xFF, 0xFF, 0x40, 0x40, 0xFF,
					0x40, 0x80, 0x80, 0xFF, 0xB2, 0x00, 0xCA, 0xFF, 0x80, 0x40, 0x80, 0xFF,
					0x00, 0x00, 0xAA, 0xFF, 0xFF, 0xCC, 0xCC, 0xFF, 0xFF, 0x00, 0xAA, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x0, 0x0, 0xFF, 0xFF, 0xFF, 0x0, 0xFF, 0xFF, 0x44, 0x33, 0xFF,
					0xB, 0xA9, 0x8D, 0xFF, 0xAE, 0x0, 0xFA, 0xFF, 0x5B, 0x46, 0x89, 0xFF,
					0x0, 0x0, 0xC2, 0xFF, 0xFF, 0xFF, 0xF3, 0xFF, 0xFF, 0x0, 0xBF, 0xFF,
				},
			},
		},
	}

	for _, c := range cases {
		actual := UnsharpMask(c.value, c.radius, c.amount)
		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s:\nexpected: %v\nactual: %v\n", "UnsharpMask", util.RGBAToString(c.expected), util.RGBAToString(actual))
		}
	}
}

func BenchmarkMedian1(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, 256, 256))
	for n := 0; n < b.N; n++ {
		benchResult = Median(img, 1)
	}
}

func BenchmarkMedian4(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, 256, 256))
	for n := 0; n < b.N; n++ {
		benchResult = Median(img, 4)
	}
}

func BenchmarkMedian8(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, 256, 256))
	for n := 0; n < b.N; n++ {
		benchResult = Median(img, 8)
	}
}

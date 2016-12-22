package thumbnail

import (
	"bytes"
	"io"
)

type Size struct {
	Width  uint
	Height uint
}

type Point struct {
	X uint
	Y uint
}

type Rectangle struct {
	Point Point
	Size  Size
}

func Rect(x, y, w, h uint) Rectangle {
	return Rectangle{
		Point: Point{x, y},
		Size:  Size{w, h},
	}
}

func Thumbnail(w io.Writer, reader io.Reader, size Size) error {
	return thumbnail(reader, size, w)
}

func ThumbnailBytes(w io.Writer, bs []byte, size Size) error {
	return Thumbnail(w, bytes.NewReader(bs), size)
}

package thumbnail

import (
	"errors"
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

type EncodingType int

const (
	JPEG EncodingType = iota + 1
	PNG
)

type (
	Thumbnailer interface {
		Thumbnail(writer io.Writer, reader io.Reader, o Options) error
		ThumbnailBytes(writer io.Writer, bs []byte, o Options) error
	}

	Options struct {
		Type EncodingType
		Size Size
	}
)

var (
	generators map[string]Thumbnailer

	DefaultSize     = Size{100, 0}
	DefaultEncoding = PNG
)

func Register(mime string, fn Thumbnailer) {
	generators[mime] = fn
}

func init() {
	generators = make(map[string]Thumbnailer)
}

func Can(mime string) bool {
	return generators[mime] != nil
}

func Thumbnail(mime string, writer io.Writer, reader io.Reader, options ...Options) error {

	opts := Options{DefaultEncoding, DefaultSize}

	if len(options) > 0 {
		opts = options[0]
	}

	if !Can(mime) {
		return errors.New("cannot thumnail: " + mime)
	}
	return generators[mime].Thumbnail(writer, reader, opts)
}

func ThumbnailBytes(mime string, writer io.Writer, bs []byte, options ...Options) error {
	if !Can(mime) {
		return errors.New("cannot thumnail: " + mime)
	}

	opts := Options{DefaultEncoding, DefaultSize}

	if len(options) > 0 {
		opts = options[0]
	}

	return generators[mime].ThumbnailBytes(writer, bs, opts)

}

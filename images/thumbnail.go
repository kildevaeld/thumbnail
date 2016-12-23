package images

import (
	"bytes"
	"io"

	"github.com/kildevaeld/thumbnail"
)

type thumbnailer struct {
}

func (self *thumbnailer) Thumbnail(w io.Writer, reader io.Reader, o thumbnail.Options) error {
	return thumb(w, reader, o)
}

func (self *thumbnailer) ThumbnailBytes(w io.Writer, bs []byte, o thumbnail.Options) error {
	return thumb(w, bytes.NewReader(bs), o)
}

// +build !imagemagick

package images

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/kildevaeld/thumbnail"
	"github.com/nfnt/resize"
)

func thumb(w io.Writer, reader io.Reader, o thumbnail.Options) error {

	size := o.Size

	image, _, err := image.Decode(reader)
	if err != nil {
		return err
	}

	m := resize.Resize(size.Width, size.Height, image, resize.Lanczos3)

	/*switch t {
	case "jpeg", "jpg":
		return jpeg.Encode(w, m, nil)
	case "png":
		return png.Encode(w, m)
	case "gif":
		return gif.Encode(w, m, nil)

	}*/
	switch o.Type {
	case thumbnail.PNG:
		return png.Encode(w, m)
	case thumbnail.JPEG:
		return jpeg.Encode(w, m, nil)
	}

	return errors.New("invalid image")

}

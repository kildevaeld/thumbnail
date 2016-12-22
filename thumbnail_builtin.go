// +build !imagemagick

package thumbnail

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/nfnt/resize"
)

func thumbnail(reader io.Reader, size Size, w io.Writer) error {

	image, t, err := image.Decode(reader)
	if err != nil {
		return err
	}

	m := resize.Resize(size.Width, size.Height, image, resize.Lanczos3)

	switch t {
	case "jpeg", "jpg":
		return jpeg.Encode(w, m, nil)
	case "png":
		return png.Encode(w, m)
	case "gif":
		return gif.Encode(w, m, nil)

	}

	return errors.New("invalid image")

}

/*

func crop(reader io.Reader, rect Rectangle, w io.Writer) error {

	img, _, err := image.Decode(reader)
	if err != nil {
		return err
	}

	dst := image.NewRGBA(image.Rect(0, 0, int(rect.Size.Width), int(rect.Size.Height)))

	r := image.Rect(int(rect.Point.X), int(rect.Point.Y), int(rect.Size.Width), int(rect.Size.Height))

	draw.Draw(dst, r, img, draw.Src)

	switch t {
	case "jpeg", "jpg":
		return jpeg.Encode(w, m, nil)
	case "png":
		return png.Encode(w, m)
	case "gif":
		return gif.Encode(w, m, nil)

	}

	return errors.New("invalid image")

}
*/

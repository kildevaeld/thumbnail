// +build imagemagick

package images

import (
	"io"
	"io/ioutil"

	"github.com/kildevaeld/thumbnail"

	"gopkg.in/gographics/imagick.v2/imagick"
)

func thumb(w io.Writer, reader io.Reader, o thumbnail.Options) error {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()

	bs, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	if err := mw.ReadImageBlob(bs); err != nil {
		return err
	}

	wi := mw.GetImageWidth()
	he := mw.GetImageHeight()

	size := o.Size

	if size.Height == 0 {
		size.Height = uint(round(float64(he) / float64(wi) * float64(size.Width)))
	} else if size.Width == 0 {
		size.Width = uint(round(float64(wi) / float64(he) * float64(size.Height)))
	}

	if err := mw.ThumbnailImage(size.Width, size.Height); err != nil {
		return err
	}

	switch o.Type {
	case thumbnail.PNG:
		if err := mw.SetImageFormat("PNG"); err != nil {
			return err
		}
	case thumbnail.JPEG:
		if err := mw.SetImageFormat("JPEG"); err != nil {

		}
	}

	bs = mw.GetImageBlob()

	if _, err = w.Write(bs); err != nil {
		return err
	}

	return nil
}

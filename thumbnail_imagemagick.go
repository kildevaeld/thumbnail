// +build imagemagick

package thumbnail

import (
	"io"
	"io/ioutil"

	"gopkg.in/gographics/imagick.v2/imagick"
)

func thumbnail(reader io.Reader, size Size, w io.Writer) error {
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

	if size.Height == 0 {
		size.Height = uint(round(float64(he) / float64(wi) * float64(size.Width)))
	} else if size.Width == 0 {
		size.Width = uint(round(float64(wi) / float64(he) * float64(size.Height)))
	}

	if err := mw.ThumbnailImage(size.Width, size.Height); err != nil {
		return err
	}

	bs = mw.GetImageBlob()

	if _, err = w.Write(bs); err != nil {
		return err
	}

	return nil
}

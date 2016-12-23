package images

import "github.com/kildevaeld/thumbnail"

func init() {

	th := &thumbnailer{}

	thumbnail.Register("image/png", th)
	thumbnail.Register("image/jpg", th)
	thumbnail.Register("image/jpeg", th)
	thumbnail.Register("image/gif", th)
}

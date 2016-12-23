package images

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/kildevaeld/thumbnail"
)

func TestThumbnail(t *testing.T) {

	r, e := os.Open("cat.jpg")
	if e != nil {
		t.Fatal(e)
	}

	writer, werr := os.Create("cat2.jpg")
	if werr != nil {
		t.Fatal(werr)
	}

	defer writer.Close()

	th := &thumbnailer{}

	if err := th.Thumbnail(writer, r, thumbnail.Options{
		Type: thumbnail.JPEG,
		Size: thumbnail.Size{100, 0},
	}); err != nil {
		t.Fatal(err)
	}

}

/*func TestCrop(t *testing.T) {

	r, e := os.Open("cat.jpg")
	if e != nil {
		t.Fatal(e)
	}

	writer, werr := os.Create("cat_crop.jpg")
	if werr != nil {
		t.Fatal(werr)
	}

	defer writer.Close()

	if err := crop(r, Rect(0, 0, 100, 100), writer); err != nil {
		t.Fatal(err)
	}

}*/

func BenchmarkThumbnail(t *testing.B) {
	reader, e := os.Open("cat.jpg")
	if e != nil {
		t.Fatal(e)
	}
	th := &thumbnailer{}
	r, _ := ioutil.ReadAll(reader)
	w := bytes.NewBuffer(nil)
	t.ResetTimer()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		if err := th.ThumbnailBytes(w, r, thumbnail.Options{
			Type: thumbnail.JPEG,
			Size: thumbnail.Size{100, 0},
		}); err != nil {
			t.Fatal(err)
		}
		w.Reset()
	}

}

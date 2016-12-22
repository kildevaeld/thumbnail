package thumbnail

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
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

	if err := Thumbnail(writer, r, Size{100, 0}); err != nil {
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

	r, _ := ioutil.ReadAll(reader)
	w := bytes.NewBuffer(nil)
	t.ResetTimer()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		if err := ThumbnailBytes(w, r, Size{100, 0}); err != nil {
			t.Fatal(err)
		}
		w.Reset()
	}

}

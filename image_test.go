package mergi_test

import (
	"github.com/noelyahan/mergi"
	"github.com/noelyahan/impexp"
	"github.com/pkg/errors"
	"testing"
)

func TestNewWithFiles(t *testing.T) {
	tests := []struct {
		in  string
		out error
	}{
		{"./testdata/nothing.png", errors.New("")},
		{"./testdata/evraiki-2514543_240_180.jpg", nil},
		{"./testdata/nothing.jpg", errors.New("")},
		{"./testdata/mergi_logo_watermark.png", nil},
	}

	for _, test := range tests {
		_, err := mergi.Import(impexp.NewFileImporter(test.in))
		if test.out == nil {
			if err != test.out {
				t.Errorf("Want [%v] got [%v]", test.out, err)
			}
		}
	}
}

func TestWithURLs(t *testing.T) {
	tests := []struct {
		in  string
		out error
	}{
		{"https://via.placeholder.com/xxx", errors.New("")},
		//{"https://via.placeholder.com/350x150", nil},
		//{"https://via.placeholder.com/yyy", errors.Import("")},
		//{"https://via.placeholder.com/500x500.jpg", nil},
	}

	for _, test := range tests {
		img, err := mergi.Import(impexp.NewFileImporter(test.in))

		if err == nil {
			t.Error("Error", err)
		} else {
			mergi.Export(impexp.NewFileExporter(img, "out.png"))
		}

		//if test.out == nil {
		//	if err != test.out {
		//		t.Errorf("Want [%v] got [%v]", test.out, err)
		//	}
		//}
	}
}

func TestSaveFiles(t *testing.T) {
	file, _ := mergi.Import(impexp.NewFileImporter("./testdata/evraiki-2514543_240_180.jpg"))
	err := mergi.Export(impexp.NewFileExporter(file, "out.png"))
	if err != nil {
		t.Errorf("Want [%v] got [%v]", nil, err)
	}

	err = mergi.Export(impexp.NewFileExporter(nil, "out.png"))
	if err == nil {
		t.Errorf("Want [%v] got [%v]", nil, err)
	}
}

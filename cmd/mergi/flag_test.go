package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestFlagOrder(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"mergi -t TBTB -i flowers_small.png -r 500 500 -c 100 100 400 400", "0#resize 0#crop"},
		{"mergi -t TBTB -i flowers_small.png -c 500 500 -r 100 100 400 400", "0#crop 0#resize"},
		{"mergi -t TBTB -i flowers_small.png -w XXX -r 500 500 -c 100 100 400 400", "0#watermark 0#resize 0#crop"},
		{"mergi -t TBTB -i flowers_small.png -r XXX -c 500 500 -w 100 100 400 400", "0#resize 0#crop 0#watermark"},
		{"mergi -t TBTB -i flowers_small.png -r XXX -c XXX -w XXX -i XXX -c XXX -i XXX -i XXX -c XXX -r XXX -f true -c XXX -r", "0#resize 0#crop 0#watermark 1#crop 3#crop 3#resize 4#final 5#crop 5#resize"},
		{`mergi -t TBTBTB -i ./testdata/friends_group_1.jpg -c "85 97 133 170" -i ./testdata/friends_group_1.jpg -c "756 25 133 170" -i ./testdata/friends_group_1.jpg -c "408 113 133 170" -i ./testdata/friends_group_1.jpg -c "507 41 133 170" -i ./testdata/friends_group_1.jpg -c "638 106 133 170" -i ./testdata/friends_group_1.jpg -c "204 30 133 170" -f true -r "800 680" -w "./testdata/mergi_logo_watermark.png -r 180 80 620 600"
`, "0#crop 1#crop 2#crop 3#crop 4#crop 5#crop 6#final 7#resize 7#watermark 7#resize"},
	}

	for _, test := range tests {
		arr := strings.Split(test.in, " ")
		out := strings.Split(test.out, " ")

		res := getFlagOrder(arr)
		if len(res) != len(out) {
			t.Errorf("Result length mismatch, Want [%v] got %v", test.out, res)
		} else {
			for i, v := range res {
				if v != out[i] {
					t.Errorf("Want [%v] got %v", test.out, res)
					return
				}
			}
		}
	}
}

func TestGetWatermarkImageXY(t *testing.T) {
	s := "./testdata/mergi_logo_watermark.png -r 180 80 10 10"
	_, x, y := getWatermarkImageXY(s)
	fmt.Println(x, y)
}

func TestGetAnimationParams(t *testing.T) {

}

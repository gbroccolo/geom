package wkb_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/go-spatial/geom/encoding/ewkb"
	"github.com/go-spatial/geom/encoding/ewkb/internal/tcase"
)

func TestEWKBEncode(t *testing.T) {
	fnames, err := tcase.GetFiles("testdata")
	if err != nil {
		t.Fatalf("error getting files: %v", err)
	}
	var fname string

	fn := func(t *testing.T, tc tcase.C) {

		if tc.Skip.Is(tcase.TypeEncode) {
			t.Skip("instructed to skip.")
		}

		bs, err := ewkb.EncodeBytes(tc.Expected)
		if err != nil {
			log.Println("TestCase:", tc)
			t.Errorf("error, expected nil got %v", err)
			return
		}
		if !reflect.DeepEqual(bs, tc.Bytes) {
			t.Errorf(" encoded geometry, expected %v got %v", tcase.SprintBinary(tc.Bytes, "\t"), tcase.SprintBinary(bs, "\t"))
		}
	}

	for _, fname = range fnames {
		t.Run(fname, func(t *testing.T) {
			cases, err := tcase.ParseFile(fname)
			if err != nil {
				t.Fatalf("error parsing file: %v : %v ", fname, err)
			}
			for _, tc := range cases {
				tc := tc
				t.Run(tc.Desc, func(t *testing.T) { fn(t, tc) })
			}

		})
	}
}

package base64

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	for _, tc := range []struct {
		str  []byte
		want []byte
	}{
		{[]byte{77, 97, 110, 121, 32, 104, 97, 110, 100, 115, 32, 109, 97, 107, 101, 32, 108, 105, 103, 104, 116, 32, 119, 111, 114, 107, 46}, []byte{84, 87, 70, 117, 101, 83, 66, 111, 89, 87, 53, 107, 99, 121, 66, 116, 89, 87, 116, 108, 73, 71, 120, 112, 90, 50, 104, 48, 73, 72, 100, 118, 99, 109, 115, 117}},
	} {
		got := Encode(tc.str)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("ERR: Encode(%v): got = %v, want = %v", tc.str, got, tc.want)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, tc := range []struct {
		str  []byte
		want []byte
	}{
		{[]byte{84, 87, 70, 117, 101, 83, 66, 111, 89, 87, 53, 107, 99, 121, 66, 116, 89, 87, 116, 108, 73, 71, 120, 112, 90, 50, 104, 48, 73, 72, 100, 118, 99, 109, 115, 117}, []byte{77, 97, 110, 121, 32, 104, 97, 110, 100, 115, 32, 109, 97, 107, 101, 32, 108, 105, 103, 104, 116, 32, 119, 111, 114, 107, 46}},
	} {
		got := Decode(tc.str)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("ERR: Decode(%v): got = %v, want = %v", tc.str, got, tc.want)
		}
	}
}

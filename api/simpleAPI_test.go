package api

import (
	"testing"
)

type testCase1 struct {
	in   Book
	want []byte
}

type testCase2 struct {
	in   []byte
	want Book
}

func TestToJSON(t *testing.T) {
	cases := []testCase1{
		{
			in:   Book{"My Life in Norway", "John Doe", "1423118"},
			want: []byte(`{"Title":"My Life in Norway","Author":"John Doe","ISBN":"1423118"}`),
		},
	}

	// range over test cases
	for _, c := range cases {
		got := c.in.ToJSON()
		if len(got) != len(c.want) {
			t.Errorf("len got: %v, len want: %v\n", len(got), len(c.want))
		}

		// assert identiy of want and got
		for idx := range c.want {
			if got[idx] != c.want[idx] {
				t.Logf("idx: %v, got: %v, want: %v\n", idx, got[idx], c.want[idx])
				t.Errorf("got: %v, want: %v, input: %v\n", string(got), string(c.want), c.in)
			}
		}
	}
}

func TestFromJSON(t *testing.T) {
	// TODO: implement
}

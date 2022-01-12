package clipboard

import (
	"reflect"
	"testing"
)

func init() {
	write([]byte{})
}

func TestRead(t *testing.T) {
	want := []byte("foo bar 1337")
	err := write(want)

	if err != nil {
		t.Fatal(err)
	}

	got, err := read()

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %q; want %q", got, want)
	}
}

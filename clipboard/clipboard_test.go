package clipboard

import (
	"bytes"
	"testing"
)

func TestWriteAndRead(t *testing.T) {
	want := []byte("foo bar 1337")

	if err := write(want); err != nil {
		t.Fatal(err)
	}

	got, err := read()

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(want, got) {
		t.Errorf("got %q; want %q", got, want)
	}
}

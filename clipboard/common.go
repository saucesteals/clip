package clipboard

import "errors"

var (
	errFailedToWrite = errors.New("clipboard: failed to write")
	errFailedToRead  = errors.New("clipboard: failed to read")
)

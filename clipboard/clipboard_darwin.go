//go:build darwin

package clipboard

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework Cocoa
#include <stdlib.h>
unsigned long clipboard_read(void **out);
int clipboard_write(const void *bytes, long n);
*/
import "C"
import (
	"errors"
	"unsafe"
)

var (
	errFailedToWrite = errors.New("clipboard: failed to write")
	errFailedToRead  = errors.New("clipboard: failed to read")
)

func read() ([]byte, error) {
	var out unsafe.Pointer

	size := C.clipboard_read(&out)

	if out == nil {
		return nil, errFailedToRead
	}

	defer C.free(out)

	if size == 0 {
		return []byte{}, nil
	}

	return C.GoBytes(out, C.int(size)), nil
}

func write(buffer []byte) error {
	if len(buffer) == 0 {
		ok := C.clipboard_write(unsafe.Pointer(nil), 0)
		if ok != 1 {
			return errFailedToWrite
		}
		return nil
	}
	ok := C.clipboard_write(unsafe.Pointer(&buffer[0]), C.long(len(buffer)))

	if ok != 1 {
		return errFailedToWrite
	}

	return nil

}

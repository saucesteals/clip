package clipboard

/*
#include <stdlib.h>
#include <stdbool.h>
#include <windows.h>
size_t clipboard_read(void **out);
bool clipboard_write(const void *bytes, long n);
*/
import "C"
import (
	"unsafe"
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
		if !ok {
			return errFailedToWrite
		}
		return nil
	}
	ok := C.clipboard_write(unsafe.Pointer(&buffer[0]), C.long(len(buffer)+1))

	if !ok {
		return errFailedToWrite
	}

	return nil
}

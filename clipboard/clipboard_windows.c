#include <windows.h>
#include <stdbool.h>

bool clipboard_write(const void *bytes, size_t size) {
    OpenClipboard(0);
    EmptyClipboard();
    HGLOBAL h = GlobalAlloc(GMEM_MOVEABLE, size);
    memcpy(GlobalLock(h), bytes, size);
    GlobalUnlock(h);
    SetClipboardData(CF_TEXT, h);
    CloseClipboard();
    return true;
}

size_t clipboard_read(void **bytes) {
    OpenClipboard(0);
    if (!IsClipboardFormatAvailable(CF_TEXT)) {
        return -1;
    }
    HANDLE h = GetClipboardData(CF_TEXT);
    int size = strlen(h);
    *bytes = malloc(size);
    memcpy(*bytes, h, size);
    CloseClipboard();
    return size;
}
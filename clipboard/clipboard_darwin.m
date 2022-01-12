#import <Foundation/Foundation.h>
#import <Cocoa/Cocoa.h>

unsigned long clipboard_read(void **out) {
	NSPasteboard * pasteboard = [NSPasteboard generalPasteboard];
	NSData *data = [pasteboard dataForType:NSPasteboardTypeString];
	if (data == nil) {
		return true;
	}
	unsigned long size = [data length];
	*out = malloc(size);
	[data getBytes: *out length: size];
	return size;
}

bool clipboard_write(const void *bytes, long size) {
	NSPasteboard *pasteboard = [NSPasteboard generalPasteboard];
	NSData *data = [NSData dataWithBytes: bytes length: size];
	[pasteboard clearContents];
	return [pasteboard setData: data forType:NSPasteboardTypeString];
}

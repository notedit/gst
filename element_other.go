// +build !arm

package gst

/*
#cgo pkg-config: gstreamer-1.0 gstreamer-app-1.0
#include "gst.h"
*/
import "C"
import "time"

func (e *Element) Seek(duration time.Duration) bool {
	result := C.gst_element_seek_simple(e.GstElement, C.GST_FORMAT_TIME, C.GST_SEEK_FLAG_FLUSH, C.long(duration.Nanoseconds()))
	return result == C.TRUE
}

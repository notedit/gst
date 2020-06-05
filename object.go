package gst

/*
#cgo pkg-config: gstreamer-1.0
#include "gst.h"
*/
import "C"

type Object struct {
	C *C.GstObject
}

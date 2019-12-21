package gst

/*
#cgo pkg-config: gstreamer-1.0
#include "gst.h"
*/
import "C"

func init() {
	C.X_gst_shim_init()
}

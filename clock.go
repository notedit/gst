package gst

/*
#cgo pkg-config: gstreamer-1.0
#include "gst.h"
*/
import "C"

type Clock struct {
	C *C.GstClock
}

func (c *Clock) GetClockTime() uint64 {

	clocktime := C.gst_clock_get_time(c.C)
	return uint64(clocktime)
}

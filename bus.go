package gst



/*
#cgo pkg-config: gstreamer-1.0 gstreamer-base-1.0 gstreamer-app-1.0 gstreamer-plugins-base-1.0 gstreamer-video-1.0 gstreamer-audio-1.0 gstreamer-plugins-bad-1.0
#include "gst.h"
*/
import "C"


type GstBus struct {
	C           *C.GstBus
}


func BusNew() (bus *GstBus) {
	CGstBus := C.gst_bus_new()

	bus = &GstBus{
		C: CGstBus,
	}

	return
}
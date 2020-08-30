package gst

/*
#cgo pkg-config: gstreamer-1.0
#include "gst.h"
*/
import "C"

type Event struct {
	GstEvent *C.GstEvent
}

func NewEosEvent() (event *Event) {
	CGstEvent := C.gst_event_new_eos()

	event = &Event{
		GstEvent: CGstEvent,
	}

	return
}

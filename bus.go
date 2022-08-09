package gst

/*
#cgo pkg-config: gstreamer-1.0
#include "gst.h"
*/
import "C"

import (
	"fmt"
	"runtime"
	"time"
)

type Bus struct {
	C *C.GstBus
}

func (b *Bus) Pop() (message *Message) {
	CGstMessage := C.gst_bus_pop(b.C)
	message = &Message{
		C: CGstMessage,
	}

	runtime.SetFinalizer(message, func(message *Message) {
		C.gst_message_unref(message.C)
	})

	return
}

func (b *Bus) PopTimed() (message *Message, err error) {
	var timeNone int64 = C.GST_CLOCK_TIME_NONE

	CGstMessage := C.gst_bus_timed_pop(b.C, C.ulonglong(timeNone))
	if CGstMessage == nil {
		// Timeout hit, no message
		err = fmt.Errorf("no message in bus")
		return
	}

	message = &Message{
		C: CGstMessage,
	}

	runtime.SetFinalizer(message, func(message *Message) {
		C.gst_message_unref(message.C)
	})

	return
}

func (b *Bus) Pull(messageType MessageType) (message *Message) {

	CGstMessage := C.gst_bus_poll(b.C, C.GstMessageType(messageType), 18446744073709551615)
	if CGstMessage == nil {
		return nil
	}

	message = &Message{
		C: CGstMessage,
	}

	runtime.SetFinalizer(message, func(message *Message) {
		C.gst_message_unref(message.C)
	})

	return
}

func (b *Bus) TimedPopFiltered(timeout time.Duration, messageType MessageType) (message *Message) {

	CGstMessage := C.gst_bus_timed_pop_filtered(b.C, C.GstClockTime(uint64(timeout)), C.GstMessageType(messageType))
	if CGstMessage == nil {
		return nil
	}

	message = &Message{
		C: CGstMessage,
	}

	runtime.SetFinalizer(message, func(message *Message) {
		C.gst_message_unref(message.C)
	})

	return
}

func (b *Bus) HavePending() bool {
	return C.gst_bus_have_pending(b.C) != 0
}

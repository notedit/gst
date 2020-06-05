package gst

/*
#cgo pkg-config: gstreamer-1.0 gstreamer-app-1.0
#include "gst.h"
*/
import "C"

//export go_callback_new_pad
func go_callback_new_pad(CgstElement *C.GstElement, CgstPad *C.GstPad, callbackID C.guint64) {

	mutex.Lock()
	element := callbackStore[uint64(callbackID)]
	mutex.Unlock()

	if element == nil {
		return
	}

	callback := element.onPadAdded

	pad := &Pad{
		pad: CgstPad,
	}

	callback(element, pad)
}

//export go_callback_event_function
func go_callback_event_function(CgstPad *C.GstPad, CgstObject *C.GstObject, CgstEvent *C.GstEvent) (ret C.gboolean) {
	wPad := &Pad{
		pad: CgstPad,
	}

	name := wPad.Name()
	padCbMutex.Lock()
	pad := padCbStore[name]
	padCbMutex.Unlock()

	if pad == nil {
		ret = C.gboolean(0)
		return
	}

	callback := pad.eventFunction

	wObject := &Object{
		C: CgstObject,
	}
	wEvent := &Event{
		event: CgstEvent,
	}

	r := callback(wPad, wObject, wEvent)
	if r {
		ret = C.gboolean(1)
	} else {
		ret = C.gboolean(0)
	}
	return
}

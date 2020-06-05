package gst

/*
#cgo pkg-config: gstreamer-1.0 gstreamer-video-1.0
#include "gst.h"
#include <glib.h>
#include <gst/video/videooverlay.h>
*/
import (
	"C"
)
import "unsafe"

// VideoOverlaySetWindowHandle will call the video overlay's set_window_handle method.
// You should use this method to tell to an overlay to display video output to a specific window
// (e.g. an XWindow on X11). Passing 0 as the handle will tell the overlay to stop using that
// window and create an internal one. registers the windowID for video output of the element.
func (e *Element) VideoOverlaySetWindowHandle(windowID uintptr) {
	C.gst_video_overlay_set_window_handle((*C.GstVideoOverlay)(unsafe.Pointer(e.GstElement)), (C.guintptr)(windowID))
}

// VideoOverlayExpose tells an overlay that it has been exposed. This will redraw the current frame in the drawable even if the pipeline is PAUSED.
func (e *Element) VideoOverlayExpose() {
	C.gst_video_overlay_expose((*C.GstVideoOverlay)(unsafe.Pointer(e.GstElement)))
}

// VideoOverlayHandleEvents tells an overlay that it should handle events from the window system.
// These events are forwarded upstream as navigation events. In some window system, events are not
// propagated in the window hierarchy if a client is listening for them. This method allows you to
// disable events handling completely from the GstVideoOverlay.
func (e *Element) VideoOverlayHandleEvents(handleEvents bool) {
	bInt := int(0)
	if handleEvents {
		bInt = 1
	}
	C.gst_video_overlay_handle_events((*C.GstVideoOverlay)(unsafe.Pointer(e.GstElement)), (C.gboolean)(bInt))
}

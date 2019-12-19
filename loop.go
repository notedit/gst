package gst

/*
#cgo pkg-config: gstreamer-1.0 gstreamer-base-1.0 gstreamer-app-1.0 gstreamer-plugins-base-1.0 gstreamer-video-1.0 gstreamer-audio-1.0 gstreamer-plugins-bad-1.0
#include "gst.h"
*/
import "C"

import (
	"runtime"
)

type GMainLoop struct {
	C *C.GMainLoop
}

func MainLoopNew() (loop *GMainLoop) {
	CLoop := C.g_main_loop_new(nil, C.gboolean(0))
	loop = &GMainLoop{C: CLoop}
	runtime.SetFinalizer(loop, func(loop *GMainLoop) {
		C.g_main_loop_unref(loop.C)
	})

	return
}

func (loop *GMainLoop) Run() {
	C.g_main_loop_run(loop.C)
}

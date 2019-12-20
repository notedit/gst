package gst



/*
#cgo pkg-config: gstreamer-1.0 gstreamer-base-1.0 gstreamer-app-1.0 gstreamer-plugins-base-1.0 gstreamer-video-1.0 gstreamer-audio-1.0 gstreamer-plugins-bad-1.0
#include "gst.h"
*/
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
	"errors"
)


type Pipeline struct {
	Bin
}


func ParseLaunch(pipelineStr string,) (p *Pipeline, err error) {
	var gError *C.GError

	pDesc := (*C.gchar)(unsafe.Pointer(C.CString(pipelineStr)))
	defer C.g_free(C.gpointer(unsafe.Pointer(pDesc)))

	gstElt := C.gst_parse_launch(pDesc, &gError)

	if gError != nil {
		err = errors.New("create pipeline error")
		return
	}

	p = &Pipeline{}
	p.GstElement = gstElt
	return
}


func PipelineNew(name string) (e *Pipeline, err error) {
	var pName *C.gchar

	if name == "" {
		pName = nil
	} else {
		pName := (*C.gchar)(unsafe.Pointer(C.CString(name)))
		defer C.g_free(C.gpointer(unsafe.Pointer(pName)))
	}

	gstElt := C.gst_pipeline_new(pName)
	if gstElt == nil {
		err = errors.New(fmt.Sprintf("could not create a Gstreamer pipeline name %s", name))
		return
	}

	e = &Pipeline{}

	e.GstElement = gstElt

	runtime.SetFinalizer(e, func(e *Pipeline) {
		fmt.Printf("CLEANING PIPELINE")
		C.gst_object_unref(C.gpointer(unsafe.Pointer(e.GstElement)))
	})

	return
}


func (p *Pipeline) SetState(state StateOptions) {
	C.gst_element_set_state(p.GstElement, C.GstState(state))
}



func (p *Pipeline) GetBus() (bus *Bus) {

	CBus := C.X_gst_pipeline_get_bus(p.GstElement)

	bus = &Bus{
		C: CBus,
	}

	runtime.SetFinalizer(bus, func(bus *Bus) {
		C.gst_object_unref(C.gpointer(unsafe.Pointer(bus.C)))
	})

	return
}

// func (p *GstPipeline) GetClock() (gstClock *GstClock) {

// 	CElementClock := C.gst_pipeline_get_clock(p.gstElement)

// 	gstClock = &GstClock{
// 		C: CElementClock,
// 	}

// 	runtime.SetFinalizer(gstClock, func(gstClock *GstClock) {
// 		C.gst_object_unref(C.gpointer(unsafe.Pointer(gstClock.C)))
// 	})

// 	return
// }


// TODO  
// get/set delay 
// get/set latency


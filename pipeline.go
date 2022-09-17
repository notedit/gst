package gst

/*
#cgo pkg-config: gstreamer-1.0
#include "gst.h"


const char* ErrorMessage(GError *err) {
    return err->message;
}
*/
import "C"

import (
	"errors"
	"fmt"
	"runtime"
	"unsafe"
)

type Pipeline struct {
	Bin
}

func ParseLaunch(pipelineStr string) (p *Pipeline, err error) {
	var gError *C.GError

	pDesc := (*C.gchar)(unsafe.Pointer(C.CString(pipelineStr)))
	defer C.g_free(C.gpointer(unsafe.Pointer(pDesc)))

	gstElt := C.gst_parse_launch(pDesc, &gError)

	if gError != nil {
		err = errors.New(C.GoString(C.ErrorMessage(gError)))
		return
	}

	p = &Pipeline{}
	p.GstElement = gstElt

	C.X_gst_pipeline_use_clock_real(p.GstElement)

	runtime.SetFinalizer(p, func(p *Pipeline) {
		C.gst_object_unref(C.gpointer(unsafe.Pointer(p.GstElement)))
	})

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

	C.X_gst_pipeline_use_clock_real(e.GstElement)

	runtime.SetFinalizer(e, func(e *Pipeline) {
		C.gst_object_unref(C.gpointer(unsafe.Pointer(e.GstElement)))
	})

	return
}

func (p *Pipeline) SetState(state StateOptions) StateChangeReturn {
	Cint := C.gst_element_set_state(p.GstElement, C.GstState(state))
	return StateChangeReturn(Cint)
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

func (p *Pipeline) GetClock() (clock *Clock) {

	CElementClock := C.X_gst_pipeline_get_clock(p.GstElement)

	clock = &Clock{
		C: CElementClock,
	}

	runtime.SetFinalizer(clock, func(clock *Clock) {
		C.gst_object_unref(C.gpointer(unsafe.Pointer(clock.C)))
	})

	return
}

func (p *Pipeline) GetDelay() uint64 {

	CClockTime := C.X_gst_pipeline_get_delay(p.GstElement)
	return uint64(CClockTime)
}

func (p *Pipeline) GetLatency() uint64 {

	CClockTime := C.X_gst_pipeline_get_latency(p.GstElement)
	return uint64(CClockTime)
}

func (p *Pipeline) SetLatency(latency uint64) {

	C.X_gst_pipeline_set_latency(p.GstElement, C.GstClockTime(latency))
}

package gst


/*
#cgo pkg-config: gstreamer-1.0
#include "gst.h"
*/
import "C"

import (
	"runtime"
	"unsafe"
)


type GstBin struct {
	gstElement *C.GstElement
	GstElement
}

func BinNew() (bin *GstBin) {
	Celement := C.gst_bin_new(nil)
	bin = &GstBin{
		gstElement: Celement,
	}

	runtime.SetFinalizer(bin, func(bin *GstBin) {
		C.gst_object_unref(C.gpointer(unsafe.Pointer(bin.gstElement)))
	})

	return
}

func (b *GstBin) Add(child *GstElement) {

	C.X_gst_bin_add(b.gstElement, child.gstElement)
	return
}


func (b *GstBin) Remove(child *GstElement) {
	
	C.X_gst_bin_remove(b.gstElement, child.gstElement)
	return
}


func (b *GstBin) AddMany(elements ...*GstElement) {
	for _, e := range elements {
		if e != nil {
			C.X_gst_bin_add(b.gstElement, e.gstElement)
		}
	}

	return
}


func (b *GstBin) GetByName(name string) (element *GstElement) {

	n := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.g_free(C.gpointer(unsafe.Pointer(n)))
	e := C.X_gst_bin_get_by_name(b.gstElement, n)
	element = &GstElement{
		gstElement: e,
	}

	return
}
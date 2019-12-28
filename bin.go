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

type Bin struct {
	Element
}

func BinNew() (bin *Bin) {
	Celement := C.gst_bin_new(nil)
	bin = &Bin{}

	bin.GstElement = Celement

	runtime.SetFinalizer(bin, func(bin *Bin) {
		C.gst_object_unref(C.gpointer(unsafe.Pointer(bin.GstElement)))
	})

	return
}

func (b *Bin) Add(child *Element) {

	C.X_gst_bin_add(b.GstElement, child.GstElement)
	return
}

func (b *Bin) Remove(child *Element) {

	C.X_gst_bin_remove(b.GstElement, child.GstElement)
	return
}

func (b *Bin) AddMany(elements ...*Element) {
	for _, e := range elements {
		if e != nil {
			C.X_gst_bin_add(b.GstElement, e.GstElement)
		}
	}

	return
}

func (b *Bin) GetByName(name string) (element *Element) {

	n := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.g_free(C.gpointer(unsafe.Pointer(n)))
	CElement := C.X_gst_bin_get_by_name(b.GstElement, n)

	if CElement == nil {
		return
	}

	element = &Element{
		GstElement: CElement,
	}

	return
}

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

type PadDirection C.GstPadDirection

const (
	PAD_UNKNOWN = PadDirection(C.GST_PAD_UNKNOWN)
	PAD_SRC     = PadDirection(C.GST_PAD_SRC)
	PAD_SINK    = PadDirection(C.GST_PAD_SINK)
)

func (p PadDirection) String() string {
	switch p {
	case PAD_UNKNOWN:
		return "PAD_UNKNOWN"
	case PAD_SRC:
		return "PAD_SRC"
	case PAD_SINK:
		return "PAD_SINK"
	}
	panic("Wrong value of PadDirection variable")
}

type PadLinkReturn int

const (
	PadLinkOk             PadLinkReturn = C.GST_PAD_LINK_OK
	PadLinkWrongHierarchy               = C.GST_PAD_LINK_WRONG_HIERARCHY
	PadLinkWasLinked                    = C.GST_PAD_LINK_WAS_LINKED
	PadLinkWrongDirection               = C.GST_PAD_LINK_WRONG_DIRECTION
	PadLinkNoFormat                     = C.GST_PAD_LINK_NOFORMAT
	PadLinkNoSched                      = C.GST_PAD_LINK_NOSCHED
	PadLinkRefused                      = C.GST_PAD_LINK_REFUSED
)

type PadTemplate struct {
	C *C.GstPadTemplate
}

type Pad struct {
	pad *C.GstPad
}

func (p *Pad) Link(sink *Pad) (padLinkReturn PadLinkReturn) {
	padLinkReturn = PadLinkReturn(C.gst_pad_link(p.pad, sink.pad))
	return
}

func (p *Pad) Unlink(sink *Pad) (padLinkReturn PadLinkReturn) {
	padLinkReturn = PadLinkReturn(C.gst_pad_unlink(p.pad, sink.pad))
	return
}

func (p *Pad) GetCurrentCaps() (gstCaps *Caps) {
	Ccaps := C.gst_pad_get_current_caps(p.pad)

	gstCaps = &Caps{
		caps: Ccaps,
	}

	runtime.SetFinalizer(gstCaps, func(gstCaps *Caps) {
		C.gst_caps_unref(gstCaps.caps)
	})
	return
}

func (p *Pad) Name() string {

	CStr := C.X_gst_pad_get_name(p.pad)
	defer C.g_free(C.gpointer(unsafe.Pointer(CStr)))
	str := C.GoString((*C.char)(unsafe.Pointer(CStr)))

	return str
}

func (p *Pad) IsEOS() bool {
	// todo
	return false
}

func (p *Pad) IsLinked() bool {
	// todo
	return false
}

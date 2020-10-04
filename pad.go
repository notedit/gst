package gst

/*
#cgo pkg-config: gstreamer-1.0
#include "gst.h"
*/
import "C"

import (
	"fmt"
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

func (e *Pad) SetObject(name string, value interface{}) {

	cname := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.g_free(C.gpointer(unsafe.Pointer(cname)))
	switch value.(type) {
	case string:
		str := (*C.gchar)(unsafe.Pointer(C.CString(value.(string))))
		defer C.g_free(C.gpointer(unsafe.Pointer(str)))
		C.X_gst_g_pad_set_string(e.pad, cname, str)
	case int:
		C.X_gst_g_pad_set_int(e.pad, cname, C.gint(value.(int)))
	case uint32:
		C.X_gst_g_pad_set_uint(e.pad, cname, C.guint(value.(uint32)))
	case bool:
		var cvalue int
		if value.(bool) == true {
			cvalue = 1
		} else {
			cvalue = 0
		}
		C.X_gst_g_pad_set_bool(e.pad, cname, C.gboolean(cvalue))
	case float64:
		C.X_gst_g_pad_set_gdouble(e.pad, cname, C.gdouble(value.(float64)))
	case *Caps:
		caps := value.(*Caps)
		C.X_gst_g_pad_set_caps(e.pad, cname, caps.caps)
	case *Structure:
		structure := value.(*Structure)
		C.X_gst_g_pad_set_structure(e.pad, cname, structure.C)
	default:
		panic(fmt.Errorf("SetObject: don't know how to set value for type %T", value))
	}
}

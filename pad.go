
package gst



/*
#cgo pkg-config: gstreamer-1.0 gstreamer-base-1.0 gstreamer-app-1.0 gstreamer-plugins-base-1.0 gstreamer-video-1.0 gstreamer-audio-1.0 gstreamer-plugins-bad-1.0
#include "gst.h"
*/
import "C"

import (
	"runtime"
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



type GstPadTemplate struct {
	C *C.GstPadTemplate
}



type GstPad struct {
	pad *C.GstPad
}


func (p *GstPad) Link(sink *GstPad) (padLinkReturn PadLinkReturn) {
	padLinkReturn = PadLinkReturn(C.gst_pad_link(p.pad, sink.pad))
	return
}


func (p *GstPad) Unlink(sink *GstPad) (padLinkReturn PadLinkReturn) {
	padLinkReturn = PadLinkReturn(C.gst_pad_unlink(p.pad, sink.pad))
	return
}

func (p *GstPad) GetCurrentCaps() (gstCaps *GstCaps) {
	Ccaps := C.gst_pad_get_current_caps(p.pad)

	gstCaps = &GstCaps{
		caps: Ccaps,
	}

	runtime.SetFinalizer(gstCaps, func(gstCaps *GstCaps) {
		C.gst_caps_unref(gstCaps.caps)
	})

	return
}


func (p *GstPad) IsEOS() bool {
	// todo 
	return false
}
package gst

/*
#cgo pkg-config: gstreamer-1.0
#include "gst.h"
*/
import "C"
import (
	"errors"
	"runtime"
	"time"
	"unsafe"
)

type FormatOptions int

const (
	FormatUndefined FormatOptions = C.GST_FORMAT_UNDEFINED
	FormatDefault   FormatOptions = C.GST_FORMAT_DEFAULT
	FormatBytes     FormatOptions = C.GST_FORMAT_BYTES
	FormatTime      FormatOptions = C.GST_FORMAT_TIME
	FormatBuffers   FormatOptions = C.GST_FORMAT_BUFFERS
	FormatPercent   FormatOptions = C.GST_FORMAT_PERCENT
)

type Query struct {
	C *C.GstQuery
}

func QueryNewSeeking(format FormatOptions) (q *Query, err error) {

	gstQuery := C.gst_query_new_seeking(C.GstFormat(format))
	if gstQuery == nil {
		err = errors.New("could not create a Gstreamer query")
		return
	}

	q = &Query{}

	q.C = gstQuery

	runtime.SetFinalizer(q, func(q *Query) {
		C.gst_object_unref(C.gpointer(unsafe.Pointer(q.C)))
	})

	return
}

func (q *Query) ParseSeeking(format *FormatOptions) (seekable bool, segmentStart, segmentEnd time.Duration) {

	var Cformat C.GstFormat
	var Cseekable C.gboolean
	var CsegmentStart, CsegmentEnd C.gint64

	if format != nil {
		Cformat = C.GstFormat(*format)
	}

	C.gst_query_parse_seeking(q.C, &Cformat, &Cseekable, &CsegmentStart, &CsegmentEnd)

	seekable = Cseekable == 1
	segmentStart = time.Duration(CsegmentStart)
	segmentEnd = time.Duration(CsegmentEnd)

	return
}

package gst

/*
#cgo pkg-config: gstreamer-1.0
#include "gst.h"
*/
import "C"

type Event struct {
	event *C.GstEvent
}

const (
	EventUnknown                MessageType = C.GST_EVENT_UNKNOWN
	EventFlushStart             MessageType = C.GST_EVENT_FLUSH_START
	EventFlushStop              MessageType = C.GST_EVENT_FLUSH_STOP
	EventStreamStart            MessageType = C.GST_EVENT_STREAM_START
	EventCaps                   MessageType = C.GST_EVENT_CAPS
	EventSegment                MessageType = C.GST_EVENT_SEGMENT
	EventTag                    MessageType = C.GST_EVENT_TAG
	EventBuffersize             MessageType = C.GST_EVENT_BUFFERSIZE
	EventSinkMessage            MessageType = C.GST_EVENT_SINK_MESSAGE
	EventEos                    MessageType = C.GST_EVENT_EOS
	EventToc                    MessageType = C.GST_EVENT_TOC
	EventSegmentDone            MessageType = C.GST_EVENT_SEGMENT_DONE
	EventGap                    MessageType = C.GST_EVENT_GAP
	EventQos                    MessageType = C.GST_EVENT_QOS
	EventSeek                   MessageType = C.GST_EVENT_SEEK
	EventNavigation             MessageType = C.GST_EVENT_NAVIGATION
	EventLatency                MessageType = C.GST_EVENT_LATENCY
	EventStep                   MessageType = C.GST_EVENT_STEP
	EventReconfigure            MessageType = C.GST_EVENT_RECONFIGURE
	EventTocSelect              MessageType = C.GST_EVENT_TOC_SELECT
	EventCustomUpstream         MessageType = C.GST_EVENT_CUSTOM_UPSTREAM
	EventCustomDownstream       MessageType = C.GST_EVENT_CUSTOM_DOWNSTREAM
	EventCustomDownstreamOob    MessageType = C.GST_EVENT_CUSTOM_DOWNSTREAM_OOB
	EventCustomDownstreamSticky MessageType = C.GST_EVENT_CUSTOM_DOWNSTREAM_STICKY
	EventCustomBoth             MessageType = C.GST_EVENT_CUSTOM_BOTH
	EventCustomBothOob          MessageType = C.GST_EVENT_CUSTOM_BOTH_OOB
)

type EventType C.GstEventType

func (e *Event) GetType() EventType {
	ctype := C.X_GST_EVENT_TYPE(e.event)
	t := EventType(ctype)

	return t
}

package gst


/*
#cgo pkg-config: gstreamer-1.0 gstreamer-base-1.0 gstreamer-app-1.0 gstreamer-plugins-base-1.0 gstreamer-video-1.0 gstreamer-audio-1.0 gstreamer-plugins-bad-1.0
#include "gst.h"
*/
import "C"


import "unsafe"

type GstMessageTypeOption C.GstMessageType

const (
	MessageUnknown          GstMessageTypeOption = C.GST_MESSAGE_UNKNOWN
	MessageEos              GstMessageTypeOption = C.GST_MESSAGE_EOS
	MessageError            GstMessageTypeOption = C.GST_MESSAGE_ERROR
	MessageWarning          GstMessageTypeOption = C.GST_MESSAGE_WARNING
	MessageInfo             GstMessageTypeOption = C.GST_MESSAGE_INFO
	MessageTag              GstMessageTypeOption = C.GST_MESSAGE_TAG
	MessageBuffering        GstMessageTypeOption = C.GST_MESSAGE_BUFFERING
	MessageStateChanged     GstMessageTypeOption = C.GST_MESSAGE_STATE_CHANGED
	MessageStateDirty       GstMessageTypeOption = C.GST_MESSAGE_STATE_DIRTY
	MessageStepDone         GstMessageTypeOption = C.GST_MESSAGE_STEP_DONE
	MessageClockProvide     GstMessageTypeOption = C.GST_MESSAGE_CLOCK_PROVIDE
	MessageClockLost        GstMessageTypeOption = C.GST_MESSAGE_CLOCK_LOST
	MessageStructureChange  GstMessageTypeOption = C.GST_MESSAGE_STREAM_STATUS
	MessageApplication      GstMessageTypeOption = C.GST_MESSAGE_APPLICATION
	MessageElement          GstMessageTypeOption = C.GST_MESSAGE_ELEMENT
	MessageSegmentStart     GstMessageTypeOption = C.GST_MESSAGE_SEGMENT_START
	MessageSegmentDone      GstMessageTypeOption = C.GST_MESSAGE_SEGMENT_DONE
	MessageDurationChanged  GstMessageTypeOption = C.GST_MESSAGE_DURATION_CHANGED
	MessageLatency          GstMessageTypeOption = C.GST_MESSAGE_LATENCY
	MessageAsyncStart       GstMessageTypeOption = C.GST_MESSAGE_ASYNC_START
	MessageAsyncDone        GstMessageTypeOption = C.GST_MESSAGE_ASYNC_DONE
	MessageRequestState     GstMessageTypeOption = C.GST_MESSAGE_REQUEST_STATE
	MessageStepStart        GstMessageTypeOption = C.GST_MESSAGE_STEP_START
	MessageQos              GstMessageTypeOption = C.GST_MESSAGE_QOS
	MessageProgress         GstMessageTypeOption = C.GST_MESSAGE_PROGRESS
	MessageToc              GstMessageTypeOption = C.GST_MESSAGE_TOC
	MessageResetTime        GstMessageTypeOption = C.GST_MESSAGE_RESET_TIME
	MessageStreamStart      GstMessageTypeOption = C.GST_MESSAGE_STREAM_START
	MessageNeedContext      GstMessageTypeOption = C.GST_MESSAGE_NEED_CONTEXT
	MessageHaveContext      GstMessageTypeOption = C.GST_MESSAGE_HAVE_CONTEXT
	MessageExtended         GstMessageTypeOption = C.GST_MESSAGE_EXTENDED
	MessageDeviceAdded      GstMessageTypeOption = C.GST_MESSAGE_DEVICE_ADDED
	MessageDeviceRemoved    GstMessageTypeOption = C.GST_MESSAGE_DEVICE_REMOVED
	MessagePropertyNotify   GstMessageTypeOption = C.GST_MESSAGE_PROPERTY_NOTIFY
	MessageStreamCollection GstMessageTypeOption = C.GST_MESSAGE_STREAM_COLLECTION
	MessageStreamsSelected  GstMessageTypeOption = C.GST_MESSAGE_STREAMS_SELECTED
	MessageRedirect         GstMessageTypeOption = C.GST_MESSAGE_REDIRECT
	MessageAny              GstMessageTypeOption = C.GST_MESSAGE_ANY
)



type GstMessage struct {
	C *C.GstMessage
}

func (message *GstMessage) GetType() (messageType GstMessageTypeOption) {
	CMessageType := C.X_GST_MESSAGE_TYPE(message.C)
	messageType = GstMessageTypeOption(CMessageType)

	return
}

func (message *GstMessage) GetName() (name string) {
	messageType := message.GetType()
	Cname := C.gst_message_type_get_name(C.GstMessageType(messageType))
	name = C.GoString((*C.char)(unsafe.Pointer(Cname)))

	return
}

func (message *GstMessage) GetStructure() (structure *GstStructure) {
	Cstructure := C.gst_message_get_structure(message.C)
	structure = &GstStructure{
		C: Cstructure,
	}

	return
}


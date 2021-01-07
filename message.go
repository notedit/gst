package gst

/*
#cgo pkg-config: gstreamer-1.0
#include "gst.h"
*/
import "C"

import "unsafe"

type MessageType C.GstMessageType

const (
	MessageUnknown         MessageType = C.GST_MESSAGE_UNKNOWN
	MessageEos             MessageType = C.GST_MESSAGE_EOS
	MessageError           MessageType = C.GST_MESSAGE_ERROR
	MessageWarning         MessageType = C.GST_MESSAGE_WARNING
	MessageInfo            MessageType = C.GST_MESSAGE_INFO
	MessageTag             MessageType = C.GST_MESSAGE_TAG
	MessageBuffering       MessageType = C.GST_MESSAGE_BUFFERING
	MessageStateChanged    MessageType = C.GST_MESSAGE_STATE_CHANGED
	MessageStateDirty      MessageType = C.GST_MESSAGE_STATE_DIRTY
	MessageStepDone        MessageType = C.GST_MESSAGE_STEP_DONE
	MessageClockProvide    MessageType = C.GST_MESSAGE_CLOCK_PROVIDE
	MessageClockLost       MessageType = C.GST_MESSAGE_CLOCK_LOST
	MessageStructureChange MessageType = C.GST_MESSAGE_STREAM_STATUS
	MessageApplication     MessageType = C.GST_MESSAGE_APPLICATION
	MessageElement         MessageType = C.GST_MESSAGE_ELEMENT
	MessageSegmentStart    MessageType = C.GST_MESSAGE_SEGMENT_START
	MessageSegmentDone     MessageType = C.GST_MESSAGE_SEGMENT_DONE
	MessageDurationChanged MessageType = C.GST_MESSAGE_DURATION_CHANGED
	MessageLatency         MessageType = C.GST_MESSAGE_LATENCY
	MessageAsyncStart      MessageType = C.GST_MESSAGE_ASYNC_START
	MessageAsyncDone       MessageType = C.GST_MESSAGE_ASYNC_DONE
	MessageRequestState    MessageType = C.GST_MESSAGE_REQUEST_STATE
	MessageStepStart       MessageType = C.GST_MESSAGE_STEP_START
	MessageQos             MessageType = C.GST_MESSAGE_QOS
	MessageProgress        MessageType = C.GST_MESSAGE_PROGRESS
	MessageToc             MessageType = C.GST_MESSAGE_TOC
	MessageResetTime       MessageType = C.GST_MESSAGE_RESET_TIME
	MessageStreamStart     MessageType = C.GST_MESSAGE_STREAM_START
	MessageNeedContext     MessageType = C.GST_MESSAGE_NEED_CONTEXT
	MessageHaveContext     MessageType = C.GST_MESSAGE_HAVE_CONTEXT
	MessageExtended        MessageType = C.GST_MESSAGE_EXTENDED
	MessageDeviceAdded     MessageType = C.GST_MESSAGE_DEVICE_ADDED
	MessageDeviceRemoved   MessageType = C.GST_MESSAGE_DEVICE_REMOVED
	//MessagePropertyNotify   MessageType = C.GST_MESSAGE_PROPERTY_NOTIFY
	//MessageStreamCollection MessageType = C.GST_MESSAGE_STREAM_COLLECTION
	//MessageStreamsSelected  MessageType = C.GST_MESSAGE_STREAMS_SELECTED
	//MessageRedirect         MessageType = C.GST_MESSAGE_REDIRECT
	MessageAny MessageType = C.GST_MESSAGE_ANY
)

type Message struct {
	C *C.GstMessage
}

func (message *Message) GetType() (messageType MessageType) {
	CMessageType := C.X_GST_MESSAGE_TYPE(message.C)
	messageType = MessageType(CMessageType)

	return
}

func (message *Message) GetName() (name string) {
	messageType := message.GetType()
	Cname := C.gst_message_type_get_name(C.GstMessageType(messageType))
	name = C.GoString((*C.char)(unsafe.Pointer(Cname)))

	return
}

func (message *Message) GetStructure() (structure *Structure) {
	Cstructure := C.gst_message_get_structure(message.C)
	structure = &Structure{
		C: Cstructure,
	}

	return
}

func (message *Message) ParseStateChanged() (oldState, newState, pending StateOptions) {
	var Coldstate, Cnewstate, Cpending C.GstState
	C.gst_message_parse_state_changed(message.C, &Coldstate, &Cnewstate, &Cpending)
	oldState = StateOptions(Coldstate)
	newState = StateOptions(Cnewstate)
	pending = StateOptions(Cpending)
	return
}

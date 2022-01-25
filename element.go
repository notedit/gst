package gst

/*
#cgo pkg-config: gstreamer-1.0 gstreamer-app-1.0
#include "gst.h"
*/
import "C"

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"sync"
	"time"
	"unsafe"
)

var (
	mutex         sync.Mutex
	callbackStore = map[uint64]*Element{}
)

type PadAddedCallback func(element *Element, pad *Pad)

type StateOptions int

const (
	StateVoidPending StateOptions = C.GST_STATE_VOID_PENDING
	StateNull        StateOptions = C.GST_STATE_NULL
	StateReady       StateOptions = C.GST_STATE_READY
	StatePaused      StateOptions = C.GST_STATE_PAUSED
	StatePlaying     StateOptions = C.GST_STATE_PLAYING
)

type StateChangeReturn int

const (
	StateChangeFailure StateChangeReturn = C.GST_STATE_CHANGE_FAILURE
	StateChangeSuccess StateChangeReturn = C.GST_STATE_CHANGE_SUCCESS
	StateChangeAsync   StateChangeReturn = C.GST_STATE_CHANGE_ASYNC
	StateChangePreroll StateChangeReturn = C.GST_STATE_CHANGE_NO_PREROLL
)

type SeekFlags int

const (
	SeekFlagNone                      SeekFlags = C.GST_SEEK_FLAG_NONE
	SeekFlagFlush                     SeekFlags = C.GST_SEEK_FLAG_FLUSH
	SeekFlagAccurate                  SeekFlags = C.GST_SEEK_FLAG_ACCURATE
	SeekFlagKeyUnit                   SeekFlags = C.GST_SEEK_FLAG_KEY_UNIT
	SeekFlagSegment                   SeekFlags = C.GST_SEEK_FLAG_SEGMENT
	SeekFlagTrickmode                 SeekFlags = C.GST_SEEK_FLAG_TRICKMODE
	SeekFlagSkip                      SeekFlags = C.GST_SEEK_FLAG_SKIP
	SeekFlagSnapBefore                SeekFlags = C.GST_SEEK_FLAG_SNAP_BEFORE
	SeekFlagSnapAfter                 SeekFlags = C.GST_SEEK_FLAG_SNAP_AFTER
	SeekFlagSnapNearest               SeekFlags = C.GST_SEEK_FLAG_SNAP_NEAREST
	SeekFlagTrickmodeKeyUnits         SeekFlags = C.GST_SEEK_FLAG_TRICKMODE_KEY_UNITS
	SeekFlagTrickmodeNoAudio          SeekFlags = C.GST_SEEK_FLAG_TRICKMODE_NO_AUDIO
	SeekFlagTrickmodeForwardPredicted SeekFlags = C.GST_SEEK_FLAG_TRICKMODE_FORWARD_PREDICTED
	SeekFlagInstantRateChange         SeekFlags = C.GST_SEEK_FLAG_INSTANT_RATE_CHANGE
)

type Element struct {
	GstElement *C.GstElement
	onPadAdded PadAddedCallback
	callbackID uint64
}

func (e *Element) Name() (name string) {
	n := (*C.char)(unsafe.Pointer(C.gst_object_get_name((*C.GstObject)(unsafe.Pointer(e.GstElement)))))
	if n != nil {
		name = string(nonCopyCString(n, C.int(C.strlen(n))))
	}

	return
}

func (e *Element) DumpDot(filePath string) {
	n := (*C.gchar)(unsafe.Pointer(C.CString(filePath)))
	defer C.g_free(C.gpointer(unsafe.Pointer(n)))
	C.X_GST_DEBUG_BIN_TO_DOT_FILE(e.GstElement, n)
}

func (e *Element) Link(dst *Element) bool {

	result := C.gst_element_link(e.GstElement, dst.GstElement)
	if result == C.TRUE {
		return true
	}
	return false
}

func (e *Element) GetPadTemplate(name string) (padTemplate *PadTemplate) {

	n := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.g_free(C.gpointer(unsafe.Pointer(n)))
	CPadTemplate := C.gst_element_class_get_pad_template(C.X_GST_ELEMENT_GET_CLASS(e.GstElement), n)
	padTemplate = &PadTemplate{
		C: CPadTemplate,
	}

	return
}

func (e *Element) GetRequestPad(padTemplate *PadTemplate, name string, caps *Caps) (pad *Pad) {

	var n *C.gchar
	var c *C.GstCaps

	if name == "" {
		n = nil
	} else {
		n = (*C.gchar)(unsafe.Pointer(C.CString(name)))
		defer C.g_free(C.gpointer(unsafe.Pointer(n)))
	}
	if caps == nil {
		c = nil
	} else {
		c = caps.caps
	}
	CPad := C.gst_element_request_pad(e.GstElement, padTemplate.C, n, c)
	pad = &Pad{
		pad: CPad,
	}

	return
}

func (e *Element) GetStaticPad(name string) (pad *Pad) {

	n := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.g_free(C.gpointer(unsafe.Pointer(n)))
	CPad := C.gst_element_get_static_pad(e.GstElement, n)
	pad = &Pad{
		pad: CPad,
	}

	return
}

func (e *Element) Query(q *Query) bool {
	Cboolean := C.gst_element_query(e.GstElement, q.C)
	return Cboolean == 1
}

func (e *Element) QueryPosition() (time.Duration, error) {
	var position C.gint64

	if int(C.gst_element_query_position(e.GstElement, C.GST_FORMAT_TIME, &position)) == 0 {
		return 0, fmt.Errorf("position query failed")
	}

	return time.Duration(position/C.GST_SECOND) * time.Second, nil
}

func (e *Element) QueryDuration() (time.Duration, error) {
	var duration C.gint64

	if int(C.gst_element_query_duration(e.GstElement, C.GST_FORMAT_TIME, &duration)) == 0 {
		return 0, fmt.Errorf("duration query failed")
	}

	return time.Duration(duration/C.GST_SECOND) * time.Second, nil
}

func (e *Element) AddPad(pad *Pad) bool {

	Cret := C.gst_element_add_pad(e.GstElement, pad.pad)
	if Cret == 1 {
		return true
	}

	return false
}

func (e *Element) GetClockBaseTime() uint64 {

	CClockTime := C.gst_element_get_base_time(e.GstElement)

	return uint64(CClockTime)
}

func (e *Element) GetClock() (gstClock *Clock) {

	CElementClock := C.gst_element_get_clock(e.GstElement)

	gstClock = &Clock{
		C: CElementClock,
	}

	runtime.SetFinalizer(gstClock, func(gstClock *Clock) {
		C.gst_object_unref(C.gpointer(unsafe.Pointer(gstClock.C)))
	})

	return
}

func (e *Element) EndOfStream() (err error) {
	// EndOfStream signals that the appsrc will not receive any further
	// input via PushBuffer and permits the pipeline to finish properly.

	var gstReturn C.GstFlowReturn
	gstReturn = C.gst_app_src_end_of_stream((*C.GstAppSrc)(unsafe.Pointer(e.GstElement)))
	if gstReturn != C.GST_FLOW_OK {
		err = errors.New("could not send end_of_stream")
	}
	return
}

func (e *Element) PushBuffer(data []byte) (err error) {

	b := C.CBytes(data)
	defer C.free(b)

	var gstReturn C.GstFlowReturn

	gstReturn = C.X_gst_app_src_push_buffer(e.GstElement, b, C.int(len(data)))

	if gstReturn != C.GST_FLOW_OK {
		err = errors.New("could not push buffer on appsrc element")
		return
	}

	return
}

func (e *Element) PullSample() (sample *Sample, err error) {

	CGstSample := C.gst_app_sink_pull_sample((*C.GstAppSink)(unsafe.Pointer(e.GstElement)))
	if CGstSample == nil {
		err = errors.New("could not pull a sample from appsink")
		return
	}

	gstBuffer := C.gst_sample_get_buffer(CGstSample)

	if gstBuffer == nil {
		err = errors.New("could not pull a sample from appsink")
		return
	}

	mapInfo := (*C.GstMapInfo)(unsafe.Pointer(C.malloc(C.sizeof_GstMapInfo)))
	defer C.free(unsafe.Pointer(mapInfo))

	if int(C.X_gst_buffer_map(gstBuffer, mapInfo)) == 0 {
		err = errors.New(fmt.Sprintf("could not map gstBuffer %#v", gstBuffer))
		return
	}

	CData := (*[1 << 30]byte)(unsafe.Pointer(mapInfo.data))
	data := make([]byte, int(mapInfo.size))
	copy(data, CData[:])

	duration := uint64(C.X_gst_buffer_get_duration(gstBuffer))
	pts := uint64(C.X_gst_buffer_get_pts(gstBuffer))
	dts := uint64(C.X_gst_buffer_get_dts(gstBuffer))
	offset := uint64(C.X_gst_buffer_get_offset(gstBuffer))

	sample = &Sample{
		Data:     data,
		Duration: duration,
		Pts:      pts,
		Dts:      dts,
		Offset:   offset,
	}

	C.gst_buffer_unmap(gstBuffer, mapInfo)
	C.gst_sample_unref(CGstSample)

	return
}

// appsink
func (e *Element) IsEOS() bool {

	Cbool := C.gst_app_sink_is_eos((*C.GstAppSink)(unsafe.Pointer(e.GstElement)))
	if Cbool == 1 {
		return true
	}

	return false
}

func (e *Element) SetObject(name string, value interface{}) {

	cname := (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.g_free(C.gpointer(unsafe.Pointer(cname)))
	switch value.(type) {
	case string:
		str := (*C.gchar)(unsafe.Pointer(C.CString(value.(string))))
		defer C.g_free(C.gpointer(unsafe.Pointer(str)))
		C.X_gst_g_object_set_string(e.GstElement, cname, str)
	case int:
		C.X_gst_g_object_set_int(e.GstElement, cname, C.gint(value.(int)))
	case uint32:
		C.X_gst_g_object_set_uint(e.GstElement, cname, C.guint(value.(uint32)))
	case bool:
		var cvalue int
		if value.(bool) == true {
			cvalue = 1
		} else {
			cvalue = 0
		}
		C.X_gst_g_object_set_bool(e.GstElement, cname, C.gboolean(cvalue))
	case float64:
		C.X_gst_g_object_set_gdouble(e.GstElement, cname, C.gdouble(value.(float64)))
	case *Caps:
		caps := value.(*Caps)
		C.X_gst_g_object_set_caps(e.GstElement, cname, caps.caps)
	case *Structure:
		structure := value.(*Structure)
		C.X_gst_g_object_set_structure(e.GstElement, cname, structure.C)
	case *Element:
		element := value.(*Element)
		C.X_gst_g_object_set_element(e.GstElement, cname, element.GstElement)

	default:
		panic(fmt.Errorf("SetObject: don't know how to set value for type %T", value))
	}
}

func (e *Element) SendEvent(event *Event) bool {

	Cbool := C.gst_element_send_event(e.GstElement, event.GstEvent)
	if Cbool == 1 {
		return true
	}

	return false
}

func (e *Element) SetState(state StateOptions) StateChangeReturn {
	Cint := C.gst_element_set_state(e.GstElement, C.GstState(state))
	return StateChangeReturn(Cint)
}

func (e *Element) GetBus() (bus *Bus) {

	CBus := C.X_gst_element_get_bus(e.GstElement)

	bus = &Bus{
		C: CBus,
	}

	runtime.SetFinalizer(bus, func(bus *Bus) {
		C.gst_object_unref(C.gpointer(unsafe.Pointer(bus.C)))
	})

	return
}

func (e *Element) SeekSimple(format FormatOptions, flags SeekFlags, seekPos time.Duration) bool {

	Cbool := C.gst_element_seek_simple(e.GstElement, C.GstFormat(format), C.GstSeekFlags(flags), C.gint64(seekPos))

	return Cbool == 1
}

func (e *Element) cleanCallback() {

	if e.onPadAdded == nil {
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	delete(callbackStore, e.callbackID)
}

//export go_callback_new_pad
func go_callback_new_pad(CgstElement *C.GstElement, CgstPad *C.GstPad, callbackID C.guint64) {

	mutex.Lock()
	element := callbackStore[uint64(callbackID)]
	mutex.Unlock()

	if element == nil {
		return
	}

	callback := element.onPadAdded

	pad := &Pad{
		pad: CgstPad,
	}

	callback(element, pad)
}

func (e *Element) SetPadAddedCallback(callback PadAddedCallback) {
	e.onPadAdded = callback

	var callbackID uint64
	mutex.Lock()
	for {
		callbackID = rand.Uint64()
		if callbackStore[callbackID] != nil {
			continue
		}
		callbackStore[callbackID] = e
		break
	}
	mutex.Unlock()

	e.callbackID = callbackID

	detailedSignal := (*C.gchar)(C.CString("pad-added"))
	defer C.free(unsafe.Pointer(detailedSignal))

	runtime.SetFinalizer(e, func(e *Element) {
		e.cleanCallback()
	})

	C.X_g_signal_connect(e.GstElement, detailedSignal, C.guint64(callbackID))
}

func ElementFactoryMake(factoryName string, name string) (e *Element, err error) {
	var pName *C.gchar

	pFactoryName := (*C.gchar)(unsafe.Pointer(C.CString(factoryName)))
	defer C.g_free(C.gpointer(unsafe.Pointer(pFactoryName)))
	if name == "" {
		pName = nil
	} else {
		pName = (*C.gchar)(unsafe.Pointer(C.CString(name)))
		defer C.g_free(C.gpointer(unsafe.Pointer(pName)))
	}
	gstElt := C.gst_element_factory_make(pFactoryName, pName)

	if gstElt == nil {
		err = errors.New(fmt.Sprintf("could not create a GStreamer element factoryName %s, name %s", factoryName, name))
		return
	}

	e = &Element{
		GstElement: gstElt,
	}

	return
}

func nonCopyGoBytes(ptr uintptr, length int) []byte {
	var slice []byte
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	header.Cap = length
	header.Len = length
	header.Data = ptr
	return slice
}

func nonCopyCString(data *C.char, size C.int) []byte {
	return nonCopyGoBytes(uintptr(unsafe.Pointer(data)), int(size))
}

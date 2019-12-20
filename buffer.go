package gst



/*
#cgo pkg-config: gstreamer-1.0 gstreamer-base-1.0 gstreamer-app-1.0 gstreamer-plugins-base-1.0 gstreamer-video-1.0 gstreamer-audio-1.0 gstreamer-plugins-bad-1.0
#include "gst.h"
*/
import "C"


import (
	"fmt"
	"errors"
	"runtime"
	"unsafe"
)


type Buffer struct {
	C *C.GstBuffer
}



type Sample struct {
	C      *C.GstSample
	Width  uint32
	Height uint32
}


func BufferNewAndAlloc(size uint) (gstBuffer *Buffer, err error) {
	CGstBuffer := C.gst_buffer_new_allocate(nil, C.gsize(size), nil)

	if CGstBuffer == nil {
		err = errors.New("could not allocate a new GstBuffer")
		return
	}

	gstBuffer = &Buffer{C: CGstBuffer}

	runtime.SetFinalizer(gstBuffer, func(gstBuffer *Buffer) {
		C.gst_buffer_unref(gstBuffer.C)
	})

	return
}


func BufferNewWrapped(data []byte) (gstBuffer *Buffer, err error) {
	Cdata := (*C.gchar)(unsafe.Pointer(C.malloc(C.size_t(len(data)))))
	C.bcopy(unsafe.Pointer(&data[0]), unsafe.Pointer(Cdata), C.size_t(len(data)))
	CGstBuffer := C.X_gst_buffer_new_wrapped(Cdata, C.gsize(len(data)))
	if CGstBuffer == nil {
		err = errors.New("could not allocate and wrap a new GstBuffer")
		return
	}
	gstBuffer = &Buffer{C: CGstBuffer}

	return
}


func BufferGetData(gstBuffer *Buffer) (data []byte, err error) {
	mapInfo := (*C.GstMapInfo)(unsafe.Pointer(C.malloc(C.sizeof_GstMapInfo)))
	defer C.free(unsafe.Pointer(mapInfo))

	if int(C.X_gst_buffer_map(gstBuffer.C, mapInfo)) == 0 {
		err = errors.New(fmt.Sprintf("could not map gstBuffer %#v", gstBuffer))
		return
	}
	CData := (*[1 << 30]byte)(unsafe.Pointer(mapInfo.data))
	data = make([]byte, int(mapInfo.size))
	copy(data, CData[:])
	C.gst_buffer_unmap(gstBuffer.C, mapInfo)
	
	return
}




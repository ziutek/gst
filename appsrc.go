package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
#include <gst/app/gstappsrc.h>

#cgo LDFLAGS: -lgstapp-1.0
*/
import "C"

import (
	//	"github.com/ziutek/glib"
	"unsafe"
)

// =================================
type AppSrc struct {
	BaseSrc
}

// =================================
func (b *AppSrc) g() *C.GstAppSrc {
	return (*C.GstAppSrc)(b.GetPtr())
}

// =================================
func (b *AppSrc) AsAppSrc() *AppSrc {
	return b
}

// =================================
func NewAppSrc(name string) *AppSrc {
	return (*AppSrc)(unsafe.Pointer(ElementFactoryMake("appsrc", name)))
}

// =================================
func (b *AppSrc) PushBuffer(buffer *Buffer) int {
	return (int)(C.gst_app_src_push_buffer((*C.GstAppSrc)(b.g()), (*C.GstBuffer)(buffer.GstBuffer)))
}

// =================================
//GstFlowReturn    gst_app_src_end_of_stream           (GstAppSrc *appsrc);
func (b *AppSrc) EndOfStream() int {
	return (int)(C.gst_app_src_end_of_stream((*C.GstAppSrc)(b.g())))
}

// =================================
// void             gst_app_src_set_caps                (GstAppSrc *appsrc, const GstCaps *caps);
func (b *AppSrc) SetCaps(caps *Caps) {
	C.gst_app_src_set_caps((*C.GstAppSrc)(b.g()), (*C.GstCaps)(caps))
}

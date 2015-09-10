package gst

/*
#include <stdlib.h>
#include <gst/gst.h>

int capsRefCount(GstCaps *c) {
	return GST_CAPS_REFCOUNT(c);
}
*/
import "C"

import (
	"github.com/ziutek/glib"
	"unsafe"
)

type Caps C.GstCaps

func (c *Caps) g() *C.GstCaps {
	return (*C.GstCaps)(c)
}

func (c *Caps) Ref() *Caps {
	return (*Caps)(C.gst_caps_ref(c.g()))
}

func (c *Caps) Unref() {
	C.gst_caps_unref(c.g())
}

func (c *Caps) RefCount() int {
	return int(C.capsRefCount(c.g()))
}

func (c *Caps) AppendStructure(media_type string, fields glib.Params) {
	C.gst_caps_append_structure(c.g(), makeGstStructure(media_type, fields))
}

func (c *Caps) GetSize() int {
	return int(C.gst_caps_get_size(c.g()))
}

func (c *Caps) String() string {
	s := (*C.char)(C.gst_caps_to_string(c.g()))
	defer C.free(unsafe.Pointer(s))
	return C.GoString(s)
}

func NewCapsAny() *Caps {
	return (*Caps)(C.gst_caps_new_any())
}

func NewCapsEmpty() *Caps {
	return (*Caps)(C.gst_caps_new_empty())
}

func NewCapsSimple(media_type string, fields glib.Params) *Caps {
	c := NewCapsEmpty()
	c.AppendStructure(media_type, fields)
	return c
}

func CapsFromString(s string) *Caps {
	cs := (*C.gchar)(C.CString(s))
	defer C.free(unsafe.Pointer(cs))
	return (*Caps)(C.gst_caps_from_string(cs))
}

package gst

/*
#include <gst/gst.h>
*/
import "C"

import (
	"unsafe"
	"github.com/ziutek/glib"
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
	return int(c.refcount)
}

func (c *Caps) AppendStructure(media_type string, fields glib.Params) {
	mt := (*C.gchar)(C.CString(media_type))
	s := C.gst_structure_empty_new(mt)
	C.free(unsafe.Pointer(mt))
	for k, v := range fields {
		n := (*C.gchar)(C.CString(k))
		C.gst_structure_take_value(s, n, v2g(glib.ValueOf(v)))
		C.free(unsafe.Pointer(n))
	}
	C.gst_caps_append_structure(c.g(), s)
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


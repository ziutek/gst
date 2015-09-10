package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

import (
	"github.com/ziutek/glib"
	"unsafe"
)

type GstObj struct {
	glib.Object
}

func (o *GstObj) g() *C.GstObject {
	return (*C.GstObject)(o.GetPtr())
}

func (o *GstObj) AsGstObj() *GstObj {
	return o
}

// Sets the name of object.
// Returns true if the name could be set. MT safe.
func (o *GstObj) SetName(name string) bool {
	s := C.CString(name)
	defer C.free(unsafe.Pointer(s))
	return C.gst_object_set_name(o.g(), (*C.gchar)(s)) != 0
}

// MT safe.
func (o *GstObj) GetName() string {
	s := C.gst_object_get_name(o.g())
	if s == nil {
		return ""
	}
	defer C.g_free(C.gpointer(s))
	return C.GoString((*C.char)(s))
}

// Sets the parent of o to p. This function causes the parent-set signal to be
// emitted when the parent was successfully set.
func (o *GstObj) SetParent(p *GstObj) bool {
	return C.gst_object_set_parent(o.g(), p.g()) != 0
}

// Returns the parent of o. Increases the refcount of the parent object so you
// should Unref it after usage.
func (o *GstObj) GetParent() *GstObj {
	p := new(GstObj)
	p.SetPtr(glib.Pointer(C.gst_object_get_parent(o.g())))
	return p
}

// Clear the parent of object, removing the associated reference. This function
// decreases the refcount of o. MT safe. Grabs and releases object's lock.
func (o *GstObj) Unparent() {
	C.gst_object_unparent(o.g())
}

// Generates a string describing the path of object in the object hierarchy.
// Only useful (or used) for debugging.
func (o *GstObj) GetPathString() string {
	s := C.gst_object_get_path_string(o.g())
	defer C.g_free(C.gpointer(s))
	return C.GoString((*C.char)(s))
}

/*func (o *GstObj) ImplementsInterfaceCheck(typ glib.Type) bool {
	return C.gst_implements_interface_check(C.gpointer(o.GetPtr()),
		C.GType(typ)) != 0
}

func (o *GstObj) ImplementsInterfaceCast(typ glib.Type) glib.Pointer {
	return glib.Pointer(C.gst_implements_interface_cast(C.gpointer(o.GetPtr()),
		C.GType(typ)))
}*/

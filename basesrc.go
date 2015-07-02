package gst

/*
#include <stdlib.h>
#include <gst/base/gstbasesrc.h>
*/
import "C"

import (
//	"github.com/ziutek/glib"
//	"unsafe"
)

type BaseSrc struct {
	Element
}

func (b *BaseSrc) g() *C.GstBaseSrc {
	return (*C.GstBaseSrc)(b.GetPtr())
}

func (b *BaseSrc) AsBaseSrc() *BaseSrc {
	return b
}

/*
//func NewBaseSrc(name string) *BaseSrc {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	b := new(BaseSrc)
	//b.SetPtr(glib.Pointer(C.gst_basesrc_new(s)))
	return b
}
*/

/*
func (b *BaseSrc) Add(els ...*Element) bool {
	for _, e := range els {
		if C.gst_BaseSrc_add(b.g(), e.g()) == 0 {
			return false
		}
	}
	return true
}

func (b *BaseSrc) Remove(els ...*Element) bool {
	for _, e := range els {
		if C.gst_BaseSrc_remove(b.g(), e.g()) == 0 {
			return false
		}
	}
	return true
}

// GetByName returns the element with the given name from a BaseSrc. Returns nil
// if no element with the given name is found in the BaseSrc.
func (b *BaseSrc) GetByName(name string) *Element {
	en := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(en))
	p := glib.Pointer(C.gst_BaseSrc_get_by_name(b.g(), en))
	if p == nil {
		return nil
	}
	e := new(Element)
	e.SetPtr(p)
	return e
}
*/

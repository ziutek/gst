package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

import (
	"unsafe"
	"github.com/ziutek/glib"
)

type Bin struct {
	Element
}

func (b *Bin) g() *C.GstBin {
	return (*C.GstBin)(b.Pointer())
}

func (b *Bin) AsBin() *Bin {
	return b
}

func NewBin(name string) *Bin {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	b := new(Bin)
	b.Set(glib.Pointer(C.gst_bin_new(s)))
	return b
}

func (b *Bin) Add(e *Element) bool {
	return C.gst_bin_add(b.g(), e.g()) != 0
}

func (b *Bin) Remove(e *Element) bool {
	return C.gst_bin_remove(b.g(), e.g()) != 0
}

func (b *Bin) AddMany(els ...*Element) bool {
	for _, e := range els {
		if !b.Add(e) {
			return false
		}
	}
	return true
}

func (b *Bin) RemoveMany(els ...*Element) bool {
	for _, e := range els {
		if !b.Remove(e) {
			return false
		}
	}
	return true
}

package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

import (
	"errors"
	"github.com/ziutek/glib"
	"unsafe"
)

type Bin struct {
	Element
}

func (b *Bin) g() *C.GstBin {
	return (*C.GstBin)(b.GetPtr())
}

func (b *Bin) AsBin() *Bin {
	return b
}

func NewBin(name string) *Bin {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	b := new(Bin)
	b.SetPtr(glib.Pointer(C.gst_bin_new(s)))
	return b
}

func (b *Bin) Add(els ...*Element) bool {
	for _, e := range els {
		if C.gst_bin_add(b.g(), e.g()) == 0 {
			return false
		}
	}
	return true
}

func (b *Bin) Remove(els ...*Element) bool {
	for _, e := range els {
		if C.gst_bin_remove(b.g(), e.g()) == 0 {
			return false
		}
	}
	return true
}

// GetByName returns the element with the given name from a bin. Returns nil
// if no element with the given name is found in the bin.
func (b *Bin) GetByName(name string) *Element {
	en := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(en))
	p := glib.Pointer(C.gst_bin_get_by_name(b.g(), en))
	if p == nil {
		return nil
	}
	e := new(Element)
	e.SetPtr(p)
	return e
}

// Similar to gst_parse_launch, but makes a bin instead of a pipeline
func ParseBinFromDescription(desc string) (*Bin, error) {
	en := (*C.gchar)(C.CString(desc))
	defer C.free(unsafe.Pointer(en))

	ghost_unlinked_pads := C.gboolean(1) // probably should be true? http://gstreamer.freedesktop.org/data/doc/gstreamer/head/gstreamer/html/gstreamer-GstParse.html#gst-parse-bin-from-description
	var Cerr *C.GError
	p := glib.Pointer(C.gst_parse_bin_from_description(en, ghost_unlinked_pads, &Cerr))
	if Cerr != nil {
		errStr := (*glib.Error)(unsafe.Pointer(Cerr)).Error()
		C.g_error_free(Cerr)
		return nil, errors.New(errStr)
	}
	if p == nil {
		return nil, nil
	}
	b := new(Bin)
	b.SetPtr(p)

	return b, nil
}

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

type Pipeline struct {
	Bin
}

func (p *Pipeline) g() *C.GstPipeline {
	return (*C.GstPipeline)(p.Pointer())
}

func (p *Pipeline) AsPipeline() *Pipeline {
	return p
}

func NewPipeline(name string) *Pipeline {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	p := new(Pipeline)
	p.Set(glib.Pointer(C.gst_pipeline_new(s)))
	return p
}

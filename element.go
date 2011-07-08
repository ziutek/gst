package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

type State C.GstState

const (
	STATE_VOID_PENDING = State(C.GST_STATE_VOID_PENDING)
	STATE_NULL         = State(C.GST_STATE_NULL)
	STATE_READY        = State(C.GST_STATE_READY)
	STATE_PAUSED       = State(C.GST_STATE_PAUSED)
	STATE_PLAYING      = State(C.GST_STATE_PLAYING)
)

type Element struct {
	GstObj
}

func (o *Element) g() *C.GstElement {
	return (*C.GstElement)(o.Pointer())
}

func (o *Element) AsElement() *Element {
	return o
}

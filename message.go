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

type MessageType C.GstMessageType

const (
	MESSAGE_UNKNOWN          = MessageType(C.GST_MESSAGE_UNKNOWN)
	MESSAGE_EOS              = MessageType(C.GST_MESSAGE_EOS)
	MESSAGE_ERROR            = MessageType(C.GST_MESSAGE_ERROR)
	MESSAGE_WARNING          = MessageType(C.GST_MESSAGE_WARNING)
	MESSAGE_INFO             = MessageType(C.GST_MESSAGE_INFO)
	MESSAGE_TAG              = MessageType(C.GST_MESSAGE_TAG)
	MESSAGE_BUFFERING        = MessageType(C.GST_MESSAGE_BUFFERING)
	MESSAGE_STATE_CHANGED    = MessageType(C.GST_MESSAGE_STATE_CHANGED)
	MESSAGE_STATE_DIRTY      = MessageType(C.GST_MESSAGE_STATE_DIRTY)
	MESSAGE_STEP_DONE        = MessageType(C.GST_MESSAGE_STEP_DONE)
	MESSAGE_CLOCK_PROVIDE    = MessageType(C.GST_MESSAGE_CLOCK_PROVIDE)
	MESSAGE_CLOCK_LOST       = MessageType(C.GST_MESSAGE_CLOCK_LOST)
	MESSAGE_NEW_CLOCK        = MessageType(C.GST_MESSAGE_NEW_CLOCK)
	MESSAGE_STRUCTURE_CHANGE = MessageType(C.GST_MESSAGE_STRUCTURE_CHANGE)
	MESSAGE_STREAM_STATUS    = MessageType(C.GST_MESSAGE_STREAM_STATUS)
	MESSAGE_APPLICATION      = MessageType(C.GST_MESSAGE_APPLICATION)
	MESSAGE_ELEMENT          = MessageType(C.GST_MESSAGE_ELEMENT)
	MESSAGE_SEGMENT_START    = MessageType(C.GST_MESSAGE_SEGMENT_START)
	MESSAGE_SEGMENT_DONE     = MessageType(C.GST_MESSAGE_SEGMENT_DONE)
	MESSAGE_DURATION         = MessageType(C.GST_MESSAGE_DURATION)
	MESSAGE_LATENCY          = MessageType(C.GST_MESSAGE_LATENCY)
	MESSAGE_ASYNC_START      = MessageType(C.GST_MESSAGE_ASYNC_START)
	MESSAGE_ASYNC_DONE       = MessageType(C.GST_MESSAGE_ASYNC_DONE)
	MESSAGE_REQUEST_STATE    = MessageType(C.GST_MESSAGE_REQUEST_STATE)
	MESSAGE_STEP_START       = MessageType(C.GST_MESSAGE_STEP_START)
	MESSAGE_QOS              = MessageType(C.GST_MESSAGE_QOS)
	//MESSAGE_PROGRESS         = MessageType(C.GST_MESSAGE_PROGRESS)
	MESSAGE_ANY              = MessageType(C.GST_MESSAGE_ANY)
)

func (t MessageType) String() string {
	switch t {
	case MESSAGE_UNKNOWN:
		return "MESSAGE_UNKNOWN"
	case MESSAGE_EOS:
		return "MESSAGE_EOS"
	case MESSAGE_ERROR:
		return "MESSAGE_ERROR"
	case MESSAGE_WARNING:
		return "MESSAGE_WARNING"
	case MESSAGE_INFO:
		return "MESSAGE_INFO"
	case MESSAGE_TAG:
		return "MESSAGE_TAG"
	case MESSAGE_BUFFERING:
		return "MESSAGE_BUFFERING"
	case MESSAGE_STATE_CHANGED:
		return "MESSAGE_STATE_CHANGED"
	case MESSAGE_STATE_DIRTY:
		return "MESSAGE_STATE_DIRTY"
	case MESSAGE_STEP_DONE:
		return "MESSAGE_STEP_DONE"
	case MESSAGE_CLOCK_PROVIDE:
		return "MESSAGE_CLOCK_PROVIDE"
	case MESSAGE_CLOCK_LOST:
		return "MESSAGE_CLOCK_LOST"
	case MESSAGE_NEW_CLOCK:
		return "MESSAGE_NEW_CLOCK"
	case MESSAGE_STRUCTURE_CHANGE:
		return "MESSAGE_STRUCTURE_CHANGE"
	case MESSAGE_STREAM_STATUS:
		return "MESSAGE_STREAM_STATUS"
	case MESSAGE_APPLICATION:
		return "MESSAGE_APPLICATION"
	case MESSAGE_ELEMENT:
		return "MESSAGE_ELEMENT"
	case MESSAGE_SEGMENT_START:
		return "MESSAGE_SEGMENT_START"
	case MESSAGE_SEGMENT_DONE:
		return "MESSAGE_SEGMENT_DONE"
	case MESSAGE_DURATION:
		return "MESSAGE_DURATION"
	case MESSAGE_LATENCY:
		return "MESSAGE_LATENCY"
	case MESSAGE_ASYNC_START:
		return "MESSAGE_ASYNC_START"
	case MESSAGE_ASYNC_DONE:
		return "MESSAGE_ASYNC_DONE"
	case MESSAGE_REQUEST_STATE:
		return "MESSAGE_REQUEST_STATE"
	case MESSAGE_STEP_START:
		return "MESSAGE_STEP_START"
	case MESSAGE_QOS:
		return "MESSAGE_QOS"
	//case MESSAGE_PROGRESS:
	//	return "MESSAGE_PROGRESS"
	case MESSAGE_ANY:
		return "MESSAGE_ANY"
	}
	panic("Unknown value of gst.MessageType")
}

type Message C.GstMessage

func (m *Message) g() *C.GstMessage {
	return (*C.GstMessage)(m)
}

func (m *Message) Type() glib.Type {
	return glib.TypeFromName("GstMessage")
}

func (m *Message) Ref() *Message {
	return (*Message)(C.gst_message_ref(m.g()))
}

func (m *Message) Unref() {
	C.gst_message_unref(m.g())
}

func (m *Message) GetType() MessageType {
	return MessageType(m._type)
}

func (m *Message) GetStructure() (string, glib.Params) {
	s := C.gst_message_get_structure(m.g())
	if s == nil {
		return "", nil
	}
	return parseGstStructure(s)
}

func (m *Message) GetSrc() *GstObj {
	src := new(GstObj)
	src.SetPtr(glib.Pointer(m.src))
	return src
}

func (m *Message) ParseError() (err *glib.Error, debug string) {
	var d *C.gchar
	var	e, ret_e *C.GError

	C.gst_message_parse_error(m.g(), &e, &d)
	defer C.free(unsafe.Pointer(e))
	defer C.free(unsafe.Pointer(d))

	debug = C.GoString((*C.char)(d))
	ret_e = new(C.GError)
	*ret_e = *e
	err = (*glib.Error)(unsafe.Pointer(ret_e))
	return
}

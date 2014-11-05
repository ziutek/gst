package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

import (
	"github.com/ziutek/glib"
)

type Bus struct {
	GstObj
}

func (b *Bus) g() *C.GstBus {
	return (*C.GstBus)(b.GetPtr())
}

func (b *Bus) AsBus() *Bus {
	return b
}

func (b *Bus) Post(msg *Message) bool {
	return C.gst_bus_post(b.g(), msg.g()) != 0
}

func (b *Bus) HavePending() bool {
	return C.gst_bus_have_pending(b.g()) != 0
}

func (b *Bus) Peek() *Message {
	return (*Message)(C.gst_bus_peek(b.g()))
}

func (b *Bus) Pop() *Message {
	return (*Message)(C.gst_bus_pop(b.g()))
}

func (b *Bus) PopFiltered(types MessageType) *Message {
	return (*Message)(C.gst_bus_pop_filtered(b.g(), C.GstMessageType(types)))
}

func (b *Bus) TimedPop(timeout uint64) *Message {
	return (*Message)(C.gst_bus_timed_pop(b.g(), C.GstClockTime(timeout)))
}

func (b *Bus) TimedPopFiltered(timeout uint64, types MessageType) *Message {
	return (*Message)(C.gst_bus_timed_pop_filtered(b.g(),
		C.GstClockTime(timeout), C.GstMessageType(types)))
}

func (b *Bus) SetFlushing(flushing bool) {
	var f C.gboolean
	if flushing {
		f = 1
	}
	C.gst_bus_set_flushing(b.g(), f)
}

func (b *Bus) DisableSyncMessageEmission() {
	C.gst_bus_disable_sync_message_emission(b.g())
}

func (b *Bus) EnableSyncMessageEmission() {
	C.gst_bus_enable_sync_message_emission(b.g())
}

func (b *Bus) AddSignalWatch() {
	C.gst_bus_add_signal_watch(b.g())
}

func (b *Bus) AddSignalWatchFull(priority int) {
	C.gst_bus_add_signal_watch_full(b.g(), C.gint(priority))
}

func (b *Bus) RemoveSignalWatch() {
	C.gst_bus_remove_signal_watch(b.g())
}

func (b *Bus) Poll(events MessageType, timeout int64) *Message {
	return (*Message)(C.gst_bus_poll(b.g(), C.GstMessageType(events),
		C.GstClockTime(timeout)))
}

func NewBus() *Bus {
	b := new(Bus)
	b.SetPtr(glib.Pointer(C.gst_bus_new()))
	return b
}

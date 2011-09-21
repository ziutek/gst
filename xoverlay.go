package gst

/*
#include <gst/interfaces/xoverlay.h>

GstXOverlay* _gst_x_overlay_cast(GstObject* o) {
	return GST_X_OVERLAY(o);
}
*/
import "C"

import (
	"github.com/ziutek/glib"
)

type XOverlay C.GstXOverlay

func (x *XOverlay) g() *C.GstXOverlay {
	return (*C.GstXOverlay)(x)
}

func (x *XOverlay) Type() glib.Type {
	return glib.TypeFromName("GstXOverlay")
}


func (x *XOverlay) SetXwindowId(id uint) {
	C.gst_x_overlay_set_xwindow_id(x.g(), C.gulong(id))
}

func (x *XOverlay) GotXwindowId(id uint) {
	C.gst_x_overlay_got_xwindow_id(x.g(), C.gulong(id))
}

func (x *XOverlay) PrepareWwindowId() {
	C.gst_x_overlay_prepare_xwindow_id(x.g())
}

func (x *XOverlay) Expose() {
	C.gst_x_overlay_expose(x.g())
}

func (x *XOverlay) HandleEvents(handle_events bool) {
	var he C.gboolean
	if handle_events {
		he = 1
	}
	C.gst_x_overlay_handle_events(x.g(), he)
}

func (o *XOverlay) SetRenderRectangle(x, y, width, height int) bool {
	return C.gst_x_overlay_set_render_rectangle(o.g(), C.gint(x), C.gint(y),
		C.gint(width), C.gint(height)) != 0
}

func XOverlayCast(o *GstObj) *XOverlay {
	return (*XOverlay)(C._gst_x_overlay_cast(o.g()))
}

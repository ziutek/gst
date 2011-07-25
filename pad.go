package gst

/*
#include <gst/gst.h>
*/
import "C"

type PadLinkReturn C.GstPadLinkReturn

const (
  PAD_LINK_OK = PadLinkReturn(C.GST_PAD_LINK_OK)
  PAD_LINK_WRONG_HIERARCHY = PadLinkReturn(C.GST_PAD_LINK_WRONG_HIERARCHY)
  PAD_LINK_WAS_LINKED = PadLinkReturn(C.GST_PAD_LINK_WAS_LINKED)
  PAD_LINK_WRONG_DIRECTION = PadLinkReturn(C.GST_PAD_LINK_WRONG_DIRECTION)
  PAD_LINK_NOFORMAT = PadLinkReturn(C.GST_PAD_LINK_NOFORMAT)
  PAD_LINK_NOSCHED = PadLinkReturn(C.GST_PAD_LINK_NOSCHED)
  PAD_LINK_REFUSED = PadLinkReturn(C.GST_PAD_LINK_REFUSED)
)

type PadDirection C.GstPadDirection

const (
  PAD_UNKNOWN = PadDirection(C.GST_PAD_UNKNOWN)
  PAD_SRC = PadDirection(C.GST_PAD_SRC)
  PAD_SINK = PadDirection(C.GST_PAD_SINK)
)

type Pad struct {
	GstObj
}

func (p *Pad) g() *C.GstPad {
	return (*C.GstPad)(p.GetPtr())
}

func (p *Pad) AsPad() *Pad {
	return p
}

func (p *Pad) CanLink(sink_pad *Pad) bool {
	return C.gst_pad_can_link(p.g(), sink_pad.g()) != 0
}

func (p *Pad) Link(sink_pad *Pad) PadLinkReturn {
	return PadLinkReturn(C.gst_pad_link(p.g(), sink_pad.g()))
}


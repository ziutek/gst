package gst

/*
#include <stdlib.h>
#include <gst/gstconfig.h>
#include <gst/gstbufferpool.h>
*/
import "C"

import (
//	"github.com/ziutek/glib"
//	"unsafe"
)

const GST_PADDING = C.GST_PADDING

type gint C.gint
type gpointer C.gpointer
type GstBufferPoolPrivate C.GstBufferPoolPrivate

/*
struct _GstBufferPool {
  GstObject            object;

  //< protected >
  gint                 flushing;

  //< private >/
  GstBufferPoolPrivate *priv;

  gpointer _gst_reserved[GST_PADDING];
};

*/
type BufferPool struct {
	Object GstObj

	Flushing gint

	GstBufferPoolPrivate *GstBufferPoolPrivate

	_gst_reserved [GST_PADDING]gpointer
}

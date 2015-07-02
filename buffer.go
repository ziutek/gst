package gst

/*
#include <stdlib.h>
#include <gst/gstbuffer.h>
*/
import "C"

import (
	//	"github.com/ziutek/glib"
	"unsafe"
)

/*
struct _GstBuffer {
  GstMiniObject          mini_object; // +++

  //< public > with COW
  GstBufferPool         *pool;

  // timestamp
  GstClockTime           pts;
  GstClockTime           dts;
  GstClockTime           duration;

  // media specific offset
  guint64                offset;
  guint64                offset_end;
};
*/

// =================================
type GstBufferStruct C.GstBuffer

type Buffer struct {
	GstBuffer *GstBufferStruct
}

func NewBuffer() *Buffer {
	buffer := new(Buffer)
	buffer.GstBuffer = (*GstBufferStruct)(C.gst_buffer_new())
	return buffer
}

// =================================
//GstBuffer * gst_buffer_new_allocate        (GstAllocator * allocator, gsize size,  GstAllocationParams * params);
func NewBufferAllocate(size uint) *Buffer {
	buffer := new(Buffer)
	buffer.GstBuffer = (*GstBufferStruct)(C.gst_buffer_new_allocate(nil, C.gsize(size), nil))
	return buffer
}

// =================================
//gsize       gst_buffer_get_size            (GstBuffer *buffer);
func (this *Buffer) GetSize() uint {
	return (uint)(C.gst_buffer_get_size((*C.GstBuffer)(this.GstBuffer)))
}

// =================================
//void        gst_buffer_append_memory        (GstBuffer *buffer, GstMemory *mem);
func (this *Buffer) AppendMemory(memory *Memory) {
	C.gst_buffer_append_memory((*C.GstBuffer)(this.GstBuffer), (*C.GstMemory)(memory))
}

// =================================
//gsize       gst_buffer_memset              (GstBuffer *buffer, gsize offset, guint8 val, gsize size);
func (this *Buffer) MemSet(offset uint, val byte, size uint) int {
	return (int)(C.gst_buffer_memset((*C.GstBuffer)(this.GstBuffer), C.gsize(offset), C.guint8(val), C.gsize(size)))
}

// =================================
//  gsize       gst_buffer_fill                (GstBuffer *buffer, gsize offset, gconstpointer src, gsize size);
func (this *Buffer) Fill(offset uint, src unsafe.Pointer, size uint) int {
	return (int)(C.gst_buffer_fill((*C.GstBuffer)(this.GstBuffer), C.gsize(offset), C.gconstpointer(src), C.gsize(size)))
}

// =================================
func (this *Buffer) Unref() {
	C.gst_buffer_unref((*C.GstBuffer)(this.GstBuffer))
}

/*type Buffer struct {
  Object    GstMiniObject
  Pool      *BufferPool
}*/

/*func (b *AppSrc) g() *C.GstAppSrc {
	return (*C.GstAppSrc)(b.GetPtr())
}

func (b *AppSrc) AsAppSrc() *AppSrc {
	return b
}

func NewAppSrc(name string) *AppSrc {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	b := new(AppSrc)
	//b.SetPtr(glib.Pointer(C.gst_appsrc_new(s)))
	return b
}

/*
func (b *AppSrc) Add(els ...*Element) bool {
	for _, e := range els {
		if C.gst_AppSrc_add(b.g(), e.g()) == 0 {
			return false
		}
	}
	return true
}

func (b *AppSrc) Remove(els ...*Element) bool {
	for _, e := range els {
		if C.gst_AppSrc_remove(b.g(), e.g()) == 0 {
			return false
		}
	}
	return true
}

// GetByName returns the element with the given name from a AppSrc. Returns nil
// if no element with the given name is found in the AppSrc.
func (b *AppSrc) GetByName(name string) *Element {
	en := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(en))
	p := glib.Pointer(C.gst_AppSrc_get_by_name(b.g(), en))
	if p == nil {
		return nil
	}
	e := new(Element)
	e.SetPtr(p)
	return e
}
*/

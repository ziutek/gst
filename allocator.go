package gst

/*
#include <stdlib.h>
#include <gst/gstobject.h>
#include <gst/gstminiobject.h>
#include <gst/gstmemory.h>
#include <gst/gstallocator.h>
*/
import "C"

import (
//	"github.com/ziutek/glib"
//	"unsafe"
)

/*
GstMemory *    gst_allocator_alloc           (GstAllocator * allocator, gsize size,
                                              GstAllocationParams *params);
void           gst_allocator_free            (GstAllocator * allocator, GstMemory *memory);
*/

type Memory C.GstMemory

func Allocate(size uint32) *Memory {
	return (*Memory)(C.gst_allocator_alloc(nil, C.gsize(size), nil))
}

func Free(ptr *Memory) {
	C.gst_allocator_free(nil, (*C.GstMemory)(ptr))
}

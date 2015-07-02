// Bindings for GStreamer API
package gst

/*
#include <stdlib.h>
#include <gst/gst.h>

char** _gst_init(int* argc, char** argv) {
	gst_init(argc, &argv);
	return argv;
}

typedef struct {
	const char *name;
	const GValue *val;
} Field;

typedef struct {
	Field* tab;
	int    n;
} Fields;

gboolean _parse_field(GQuark id, const GValue* val, gpointer data) {
	Fields *f = (Fields*)(data);
	f->tab[f->n].name = g_quark_to_string(id);
	f->tab[f->n].val = val;
	++f->n;
	return TRUE;
}

Fields _parse_struct(GstStructure *s) {
	int n = gst_structure_n_fields(s);
	Fields f = { malloc(n * sizeof(Field)), 0 };
	gst_structure_foreach(s, _parse_field, (gpointer)(&f));
	return f;
}

#cgo pkg-config: gstreamer-1.0
*/
import "C"

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/ziutek/glib"
)

func v2g(v *glib.Value) *C.GValue {
	return (*C.GValue)(unsafe.Pointer(v))
}

func g2v(v *C.GValue) *glib.Value {
	return (*glib.Value)(unsafe.Pointer(v))
}

type Fourcc uint32

func (f Fourcc) Type() glib.Type {
	return glib.TYPE_UINT
}

func (f Fourcc) Value() *glib.Value {
	return glib.ValueOf(f)
}

func (f Fourcc) String() string {
	buf := make([]byte, 4)
	buf[0] = byte(f)
	buf[1] = byte(f >> 8)
	buf[2] = byte(f >> 16)
	buf[3] = byte(f >> 32)
	return string(buf)
}

func MakeFourcc(a, b, c, d byte) Fourcc {
	return Fourcc(uint32(a) | uint32(b)<<8 | uint32(c)<<16 | uint32(d)<<24)
}

func StrFourcc(s string) Fourcc {
	if len(s) != 4 {
		panic("Fourcc string length != 4")
	}
	return MakeFourcc(s[0], s[1], s[2], s[3])
}

func ValueFourcc(v *glib.Value) Fourcc {
	return Fourcc(v.GetUint32())
}

type IntRange struct {
	Start, End int
}

func (r *IntRange) Type() glib.Type {
	return TYPE_INT_RANGE
}

func (r *IntRange) Value() *glib.Value {
	v := glib.NewValue(r.Type())
	C.gst_value_set_int_range(v2g(v), C.gint(r.Start), C.gint(r.End))
	return v
}

func (r *IntRange) String() string {
	return fmt.Sprintf("[%d,%d]", r.Start, r.End)
}

func ValueRange(v *glib.Value) *IntRange {
	return &IntRange{
		int(C.gst_value_get_int_range_min(v2g(v))),
		int(C.gst_value_get_int_range_max(v2g(v))),
	}
}

type Fraction struct {
	Numer, Denom int
}

func (f *Fraction) Type() glib.Type {
	return TYPE_FRACTION
}

func (f *Fraction) Value() *glib.Value {
	v := glib.NewValue(f.Type())
	C.gst_value_set_fraction(v2g(v), C.gint(f.Numer), C.gint(f.Denom))
	return v
}

func (r *Fraction) String() string {
	return fmt.Sprintf("%d/%d", r.Numer, r.Denom)
}

func ValueFraction(v *glib.Value) *Fraction {
	return &Fraction{
		int(C.gst_value_get_fraction_numerator(v2g(v))),
		int(C.gst_value_get_fraction_denominator(v2g(v))),
	}
}

var TYPE_FOURCC, TYPE_INT_RANGE, TYPE_FRACTION glib.Type

func init() {
	alen := C.int(len(os.Args))
	argv := make([]*C.char, alen)
	for i, s := range os.Args {
		argv[i] = C.CString(s)
	}
	ret := C._gst_init(&alen, &argv[0])
	argv = (*[1 << 16]*C.char)(unsafe.Pointer(ret))[:alen]
	os.Args = make([]string, alen)
	for i, s := range argv {
		os.Args[i] = C.GoString(s)
	}

	TYPE_INT_RANGE = glib.Type(C.gst_int_range_get_type())
	TYPE_FRACTION = glib.Type(C.gst_fraction_get_type())
}

func makeGstStructure(name string, fields glib.Params) *C.GstStructure {
	nm := (*C.gchar)(C.CString(name))
	s := C.gst_structure_new_empty(nm)
	C.free(unsafe.Pointer(nm))
	for k, v := range fields {
		n := (*C.gchar)(C.CString(k))
		C.gst_structure_take_value(s, n, v2g(glib.ValueOf(v)))
		C.free(unsafe.Pointer(n))
	}
	return s
}

func parseGstStructure(s *C.GstStructure) (name string, fields glib.Params) {
	name = C.GoString((*C.char)(C.gst_structure_get_name(s)))
	ps := C._parse_struct(s)
	n := (int)(ps.n)
	tab := (*[1 << 16]C.Field)(unsafe.Pointer(ps.tab))[:n]
	fields = make(glib.Params)
	for _, f := range tab {
		fields[C.GoString(f.name)] = g2v(f.val).Get()
	}
	return
}

var CLOCK_TIME_NONE = int64(C.GST_CLOCK_TIME_NONE)

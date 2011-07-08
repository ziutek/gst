include $(GOROOT)/src/Make.inc

TARG = github.com/ziutek/gst
CGOFILES = functions.go object.go

include $(GOROOT)/src/Make.pkg

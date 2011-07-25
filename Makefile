include $(GOROOT)/src/Make.inc

TARG = github.com/ziutek/gst
CGOFILES = functions.go object.go element.go bin.go pipeline.go clock.go pad.go

include $(GOROOT)/src/Make.pkg

include $(GOROOT)/src/Make.inc

TARG = github.com/ziutek/gst
CGOFILES = common.go object.go element.go bin.go pipeline.go clock.go pad.go caps.go bus.go message.go xoverlay.go

include $(GOROOT)/src/Make.pkg

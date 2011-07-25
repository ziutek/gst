package main

import (
	"github.com/ziutek/glib"
	"github.com/ziutek/gst"
)

func main() {
	src := gst.ElementFactoryMake("videotestsrc", "VideoSrc")
	sink := gst.ElementFactoryMake("autovideosink", "VideoSink")
	pl := gst.NewPipeline("MyPipeline")

	pl.Add(src, sink)
	src.Link(sink)
	pl.SetState(gst.STATE_PLAYING)

	glib.NewMainLoop(nil).Run()
}

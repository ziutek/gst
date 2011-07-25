package main

import (
	"github.com/ziutek/glib"
	"github.com/ziutek/gst"
	"fmt"
)

func main() {
	src := gst.ElementFactoryMake("videotestsrc", "VideoSrc")
	sink := gst.ElementFactoryMake("autovideosink", "VideoSink")
	pl := gst.NewPipeline("MyPipeline")

	pl.Add(src, sink)
	filter := gst.NewCapsSimple(
		"video/x-raw-yuv",
		glib.Params{"width": 192, "height": 108},
	)
	fmt.Println(filter)
	src.LinkFiltered(sink, filter)
	pl.SetState(gst.STATE_PLAYING)

	glib.NewMainLoop(nil).Run()
}

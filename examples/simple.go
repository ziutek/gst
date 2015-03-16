package main

import (
	"fmt"
	"os"

	"github.com/ziutek/glib"
	"github.com/ziutek/gst"
)

func checkElem(e *gst.Element, name string) {
	if e == nil {
		fmt.Fprintln(os.Stderr, "can't make element: ", name)
		os.Exit(1)
	}
}

func main() {
	src := gst.ElementFactoryMake("videotestsrc", "VideoSrc")
	checkElem(src, "videotestsrc")
	//vsink := "autovideosink"
	vsink := "xvimagesink"
	sink := gst.ElementFactoryMake(vsink, "VideoSink")
	checkElem(sink, vsink)

	pl := gst.NewPipeline("MyPipeline")

	pl.Add(src, sink)
	filter := gst.NewCapsSimple(
		"video/x-raw,format=yuv",
		glib.Params{
			"width":     int32(192),
			"height":    int32(108),
			"framerate": &gst.Fraction{25, 1},
		},
	)
	fmt.Println(filter)
	//src.LinkFiltered(sink, filter)
	src.Link(sink)
	pl.SetState(gst.STATE_PLAYING)

	glib.NewMainLoop(nil).Run()
}

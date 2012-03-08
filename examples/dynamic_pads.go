// This simple test application create live H264 (or WebM - see commented lines)
// content from test source, decode it and display.
package main

import (
	"fmt"
	"github.com/ziutek/glib"
	"github.com/ziutek/gst"
)

func main() {
	src := gst.ElementFactoryMake("videotestsrc", "Test source")
	src.SetProperty("do-timestamp", true)
	src.SetProperty("pattern", 18) // ball

	//enc := gst.ElementFactoryMake("vp8enc", "VP8 encoder")
	enc := gst.ElementFactoryMake("x264enc", "H.264 encoder")

	//mux := gst.ElementFactoryMake("webmmux", "WebM muxer")
	mux := gst.ElementFactoryMake("matroskamux", "matroskamux muxer")
	mux.SetProperty("streamable", true)

	demux := gst.ElementFactoryMake("matroskademux", "Matroska demuxer")

	//dec := gst.ElementFactoryMake("vp8dec", "VP8 dcoder")
	dec := gst.ElementFactoryMake("ffdec_h264", "H.264 dcoder")

	sink := gst.ElementFactoryMake("autovideosink", "Video display element")

	pl := gst.NewPipeline("MyPipeline")

	pl.Add(src, enc, mux, demux, dec, sink)

	src.Link(enc, mux, demux)
	demux.ConnectNoi("pad-added", cbPadAdded, dec.GetStaticPad("sink"))
	dec.Link(sink)
	pl.SetState(gst.STATE_PLAYING)

	glib.NewMainLoop(nil).Run()
}

// Callback function for "pad-added" event
func cbPadAdded(dec_sink_pad, demux_new_pad *gst.Pad) {
	fmt.Println("New pad:", demux_new_pad.GetName())
	if demux_new_pad.CanLink(dec_sink_pad) {
		if demux_new_pad.Link(dec_sink_pad) != gst.PAD_LINK_OK {
			fmt.Println("Link error")
		}
	} else {
		fmt.Println("Can't link it with:", dec_sink_pad.GetName())
	}
}

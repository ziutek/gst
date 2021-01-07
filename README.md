Retired. I don't use/develop this package anymore.

### Go bindings for GStreamer at a very early stage of maturity.

This package is based on [GLib bindings](https://github.com/ziutek/glib). It
should be goinstalable. Try

    $ go get github.com/ziutek/gst

#### Documentation

See *examples* directory and http://gopkgdoc.appspot.com/pkg/github.com/ziutek/gst

To run examples use `go run` command:

	$ cd examples
	$ go run simple.go

To run live WebM example use `go run live_webm.go` and open
http://127.0.0.1:8080 with your browser. You probably need to wait a long time
for video because of small bitrate of this stream and big buffer in you browser.

// This simple test application serve live generated WebM content on webpage
// using HTML5 <video> element.
// The bitrate is low so you need to wait long for video if you browser has
// big input buffer.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"syscall"

	"github.com/ziutek/glib"
	"github.com/ziutek/gst"
)

type Index struct {
	width, height int
}

func (ix *Index) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	html := `<!doctype html>
	<html>
		<head>
			<meta charset='utf-8'>
			<title>Live WebM video</title>
		</head>
		<body>
			<img src='/images/logo.png' alt=logo1><br>
			<video src='/video' width=%d height=%d autoplay></video><br>
			<img src='/images/logo-153x55.png' alt=logo2>
		</body>
	</html>`

	fmt.Fprintf(wr, html, ix.width, ix.height)
}

type WebM struct {
	pl    *gst.Pipeline
	sink  *gst.Element
	conns map[int]net.Conn
}

func (wm *WebM) Play() {
	wm.pl.SetState(gst.STATE_PLAYING)
}

func (wm *WebM) Stop() {
	wm.pl.SetState(gst.STATE_READY)
}

func (wm *WebM) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	/*wr.Header().Set("Content-Type", "video/webm")
	wr.Header().Set("Transfer-Encoding", "identity")
	wr.WriteHeader(http.StatusOK)
	wr.(http.Flusher).Flush()*/

	// Obtain fd
	conn, _, err := wr.(http.Hijacker).Hijack()
	if err != nil {
		log.Println("http.Hijacker.Hijack:", err)
		return
	}
	file, err := conn.(*net.TCPConn).File()
	if err != nil {
		log.Println("net.TCPConn.File:", err)
		return
	}
	fd, err := syscall.Dup(int(file.Fd()))
	if err != nil {
		log.Println("syscall.Dup:", err)
		return
	}
	// Send HTTP header
	_, err = io.WriteString(
		file,
		"HTTP/1.1 200 OK\r\n"+
			"Content-Type: video/webm\r\n\r\n",
	)
	if err != nil {
		log.Println("io.WriteString:", err)
		return
	}
	file.Close()

	// Save connection in map (workaround)
	wm.conns[fd] = conn

	// Pass fd to the multifdsink
	wm.sink.Emit("add", fd)
}

// Handler for connection closing
func (wm *WebM) cbClientFdRemoved(fd int32) {
	wm.conns[int(fd)].Close()
	syscall.Close(int(fd))
	delete(wm.conns, int(fd))
}

func NewWebM(width, height, fps int) *WebM {
	wm := new(WebM)
	wm.conns = make(map[int]net.Conn)

	src := gst.ElementFactoryMake("videotestsrc", "Test source")
	src.SetProperty("do-timestamp", true)
	src.SetProperty("pattern", 18) // ball

	enc := gst.ElementFactoryMake("vp8enc", "VP8 encoder")

	mux := gst.ElementFactoryMake("webmmux", "WebM muxer")
	mux.SetProperty("streamable", true)

	wm.sink = gst.ElementFactoryMake("multifdsink", "Multifd sink")
	wm.sink.SetProperty("sync", true)
	wm.sink.SetProperty("recover-policy", 3) // keyframe
	wm.sink.SetProperty("sync-method", 2)    // latest-keyframe

	wm.pl = gst.NewPipeline("WebM generator")
	wm.pl.Add(src, enc, mux, wm.sink)

	filter := gst.NewCapsSimple(
		"video/x-raw,format=yuv",
		glib.Params{
			"width":     int32(width),
			"height":    int32(height),
			"framerate": &gst.Fraction{fps, 1},
		},
	)
	src.LinkFiltered(enc, filter)
	enc.Link(mux, wm.sink)

	wm.sink.ConnectNoi("client-fd-removed", (*WebM).cbClientFdRemoved, wm)

	return wm
}

func staticHandler(wr http.ResponseWriter, req *http.Request) {
	http.ServeFile(wr, req, req.URL.Path[1:])
}

func main() {
	index := &Index{384, 216}
	wm := NewWebM(index.width, index.height, 25)
	wm.Play()

	http.Handle("/", index)
	http.Handle("/video", wm)
	http.HandleFunc("/images/", staticHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("http.ListenAndServe:", err)
	}
}

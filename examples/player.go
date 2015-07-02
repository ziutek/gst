package main

import (
	"fmt"
	"github.com/ziutek/gst"
	"github.com/ziutek/gtk"
)

type Player struct {
	pipe       *gst.Element
	bus        *gst.Bus
	window     *gtk.Window
	movie_area *gtk.DrawingArea
	file_path  string
	xid        uint
}

func (p *Player) onPlayClicked(b *gtk.Button) {
	if p.file_path != "" {
		p.pipe.SetProperty("uri", "file://"+p.file_path)
		p.pipe.SetState(gst.STATE_PLAYING)
	}
}

func (p *Player) onPauseClicked(b *gtk.Button) {
	state, _, _ := p.pipe.GetState(gst.CLOCK_TIME_NONE)
	if state == gst.STATE_PLAYING {
		p.pipe.SetState(gst.STATE_PAUSED)
	}
}

func (p *Player) onStopClicked(b *gtk.Button) {
	p.pipe.SetState(gst.STATE_NULL)
}

func (p *Player) onFileSelected(chooser *gtk.FileChooserButton) {
	p.file_path = gtk.FileChooserCast(chooser.AsWidget()).GetFilename()
}

func (p *Player) onMessage(bus *gst.Bus, msg *gst.Message) {
	switch msg.GetType() {
	case gst.MESSAGE_EOS:
		p.pipe.SetState(gst.STATE_NULL)
	case gst.MESSAGE_ERROR:
		p.pipe.SetState(gst.STATE_NULL)
		err, debug := msg.ParseError()
		fmt.Printf("Error: %s (debug: %s)\n", err, debug)
	}
}

func (p *Player) onSyncMessage(bus *gst.Bus, msg *gst.Message) {
	name, _ := msg.GetStructure()
	if name != "prepare-xwindow-id" {
		return
	}
	img_sink := msg.GetSrc()
	xov := gst.XOverlayCast(img_sink)
	if p.xid != 0 && xov != nil {
		img_sink.SetProperty("force-aspect-ratio", true)
		xov.SetXwindowId(p.xid)
	} else {
		fmt.Println("Error: xid =", p.xid, "xov =", xov)
	}
}

func (p *Player) onVideoWidgetRealize(w *gtk.Widget) {
	p.xid = p.movie_area.GetWindow().GetXid()
}

func NewPlayer() *Player {
	p := new(Player)

	p.window = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	p.window.SetTitle("Player")
	p.window.ConnectNoi("destroy", gtk.MainQuit, nil)

	vbox := gtk.NewVBox(false, 0)
	p.window.Add(vbox.AsWidget())
	hbox := gtk.NewHBox(false, 0)
	vbox.PackStart(hbox.AsWidget(), false, false, 0)

	fcb := gtk.NewFileChooserButton(
		"Choose media file",
		gtk.FILE_CHOOSER_ACTION_OPEN,
	)
	fcb.Connect("selection-changed", (*Player).onFileSelected, p)
	hbox.Add(fcb.AsWidget())

	button := gtk.NewButtonFromStock("gtk-media-play")
	button.Connect("clicked", (*Player).onPlayClicked, p)
	hbox.PackStart(button.AsWidget(), false, false, 0)

	button = gtk.NewButtonFromStock("gtk-media-pause")
	button.Connect("clicked", (*Player).onPauseClicked, p)
	hbox.PackStart(button.AsWidget(), false, false, 0)

	button = gtk.NewButtonFromStock("gtk-media-stop")
	button.Connect("clicked", (*Player).onStopClicked, p)
	hbox.PackStart(button.AsWidget(), false, false, 0)

	p.movie_area = gtk.NewDrawingArea()
	p.movie_area.Connect("realize", (*Player).onVideoWidgetRealize, p)
	p.movie_area.SetDoubleBuffered(false)
	p.movie_area.SetSizeRequest(640, 360)
	vbox.Add(p.movie_area.AsWidget())

	p.window.ShowAll()
	p.window.Realize()

	p.pipe = gst.ElementFactoryMake("playbin2", "autoplay")
	p.bus = p.pipe.GetBus()
	p.bus.AddSignalWatch()
	p.bus.Connect("message", (*Player).onMessage, p)
	p.bus.EnableSyncMessageEmission()
	p.bus.Connect("sync-message::element", (*Player).onSyncMessage, p)

	return p
}

func (p *Player) Run() {
	gtk.Main()
}

func main() {
	NewPlayer().Run()
}

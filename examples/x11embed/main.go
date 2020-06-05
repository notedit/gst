/*
This example embeds a gstream pipeline output into an X11 window.
Left-clicking the mouse in the window will start/pause the output video.
*/
package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/notedit/gst"
)

func main() {
	pipeline, err := gst.ParseLaunch("videotestsrc ! xvimagesink sync=false name=output")
	if err != nil {
		log.Fatalf("ParseLaunch: %s", err)
	}
	e := pipeline.GetByName("output")
	if e == nil {
		log.Fatal("Output element not found")
	}

	X, err := xgb.NewConn()
	if err != nil {
		log.Fatalf("NewConn: %s", err)
	}

	setup := xproto.Setup(X)
	screen := setup.DefaultScreen(X)
	wid, _ := xproto.NewWindowId(X)
	xproto.CreateWindow(X, screen.RootDepth, wid, screen.Root,
		0, 0, 500, 500, 0,
		xproto.WindowClassInputOutput, screen.RootVisual, 0, []uint32{})
	xproto.ChangeWindowAttributes(X, wid,
		xproto.CwBackPixel|xproto.CwEventMask,
		[]uint32{
			0xffffffff,
			xproto.EventMaskStructureNotify | xproto.EventMaskButtonPress})
	err = xproto.MapWindowChecked(X, wid).Check()
	if err != nil {
		log.Fatalf("Checked Error for mapping window %d: %s\n", wid, err)
	}

	isPlaying := false
	for {
		ev, xerr := X.WaitForEvent()
		if ev == nil && xerr == nil {
			fmt.Println("Both event and error are nil. Exiting...")
			return
		}
		if _, ok := ev.(xproto.ButtonPressEvent); ok {
			if isPlaying {
				isPlaying = false
				pipeline.SetState(gst.StatePaused)
			} else {
				isPlaying = true
				e.VideoOverlaySetWindowHandle(uintptr(wid))
				pipeline.SetState(gst.StatePlaying)
			}
		}
	}
}

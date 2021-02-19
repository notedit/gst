package gst

import "testing"

func TestCaps_GetStructure(t *testing.T) {
	pipeline, err := ParseLaunch("videotestsrc name=src ! video/x-raw,width=640,height=480 ! fakesink")
	if err != nil {
		t.Fatal(err)
	}
	src := pipeline.GetByName("src")
	if src == nil {
		t.Fatal("element 'src' not found")
	}
	pipeline.SetState(StatePlaying)
	bus := pipeline.GetBus()
	for {
		msg := bus.Pull(MessageStateChanged)
		_, newState, _ := msg.ParseStateChanged()
		if newState == StatePlaying {
			structure := src.GetStaticPad("src").GetCurrentCaps().GetStructure(0)
			width, err := structure.GetInt("width")
			if err != nil {
				t.Fatal(err)
			}
			height, err := structure.GetInt("height")
			if err != nil {
				t.Fatal(err)
			}
			if width != 640 || height != 480 {
				t.Fatal(err)
			}
			break
		}
	}
}

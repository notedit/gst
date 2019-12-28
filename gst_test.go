package gst

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
	"time"
)

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func TestPipeline(t *testing.T) {

	PrintMemUsage()

	pipeline, err := ParseLaunch("videotestsrc  ! capsfilter name=filter ! autovideosink")

	if err != nil {
		t.Error("pipeline create error", err)
		t.FailNow()
	}

	fmt.Println(pipeline.Name())

	element := pipeline.GetByName("filter")

	if element == nil {
		t.Error("pipe find element error")
	}

	PrintMemUsage()

	time.Sleep(1000000)

}

func TestAppsrc(t *testing.T) {

	PrintMemUsage()

	pipeline, err := ParseLaunch("appsrc name=mysource format=time is-live=true do-timestamp=true ! videoconvert ! autovideosink")

	if err != nil {
		t.Error("pipeline create error", err)
		t.FailNow()
	}

	videoCap := CapsFromString("video/x-raw,format=RGB,width=320,height=240,bpp=24,depth=24")

	element := pipeline.GetByName("mysource")

	element.SetObject("caps", videoCap)

	pipeline.SetState(StatePlaying)

	time.Sleep(100000000)

	i := 0
	for {

		if i > 10 {
			break
		}

		data := make([]byte, 320*240*3)

		err := element.PushBuffer(data)

		if err != nil {
			t.Error("push buffer error")
			t.FailNow()
			break
		}

		fmt.Println("push one")

		i += 1

		time.Sleep(50000000)
	}

	pipeline.SetState(StateNull)

	pipeline = nil
	element = nil
	videoCap = nil

	PrintMemUsage()

}

func TestAppsink(t *testing.T) {

	PrintMemUsage()

	pipeline, err := ParseLaunch("videotestsrc  num-buffers=10 ! appsink name=sink")

	if err != nil {
		t.Error("pipeline create error", err)
		t.FailNow()
	}

	element := pipeline.GetByName("sink")

	pipeline.SetState(StatePlaying)

	time.Sleep(1000000)

	for {

		sample, err := element.PullSample()
		if err != nil {
			if element.IsEOS() == true {
				fmt.Println("eos")
				return
			} else {
				fmt.Println(err)
				continue
			}
		}
		fmt.Println("got sample", sample.Duration)

	}

	pipeline.SetState(StateNull)

	pipeline = nil
	element = nil

	PrintMemUsage()

	time.Sleep(1000000)
}

func TestDynamicPipeline(t *testing.T) {

	pipeline, err := PipelineNew("test-pipeline")

	if err != nil {
		panic(err)
	}

	source, _ := ElementFactoryMake("uridecodebin", "source")
	convert, _ := ElementFactoryMake("audioconvert", "convert")
	sink, _ := ElementFactoryMake("autoaudiosink", "sink")

	pipeline.Add(source)
	pipeline.Add(convert)
	pipeline.Add(sink)

	convert.Link(sink)

	source.SetObject("uri", "https://www.freedesktop.org/software/gstreamer-sdk/data/media/sintel_trailer-480p.webm")

	source.SetPadAddedCallback(func(element *Element, pad *Pad) {
		capstr := pad.GetCurrentCaps().ToString()

		if strings.HasPrefix(capstr, "audio") {
			sinkpad := convert.GetStaticPad("sink")
			pad.Link(sinkpad)
		}

	})

	pipeline.SetState(StatePlaying)

	bus := pipeline.GetBus()

	for {
		message := bus.Pull(MessageError | MessageEos)
		fmt.Println("message:", message.GetName())
		if message.GetType() == MessageEos {
			break
		}

	}
}

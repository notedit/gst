package gst

import (
	"fmt"
	"runtime"
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

	fmt.Println(element.Name())

}

func TestAppsink(t *testing.T) {

	pipeline, err := ParseLaunch("videotestsrc  num-buffers=15 ! appsink name=sink")

	if err != nil {
		t.Error("pipeline create error", err)
		t.FailNow()
	}

	fmt.Println(pipeline)

	element := pipeline.GetByName("sink")

	pipeline.SetState(StatePlaying)

	for {

		gstSample, err := element.PullSample()
		if err != nil {
			if element.IsEOS() == true {
				fmt.Println("eos")
				return
			} else {
				fmt.Println(err)
				continue
			}
		}
		fmt.Println("got sample", gstSample)

	}
}

func TestAppsrc(t *testing.T) {

	pipeline, err := ParseLaunch("appsrc name=mysource format=time is-live=true do-timestamp=true ! videoconvert ! autovideosink")

	fmt.Println("push one")

	if err != nil {
		t.Error("pipeline create error", err)
		t.FailNow()
	}

	videoCap := CapsFromString("video/x-raw,format=RGB,width=320,height=240,bpp=24,depth=24")

	element := pipeline.GetByName("mysource")

	element.SetObject("caps", videoCap)

	pipeline.SetState(StatePlaying)

	fmt.Println("push one")

	time.Sleep(1000)

	i := 0
	for {

		if i > 100 {
			break
		}

		data := make([]byte, 320*240*3)

		err = element.PushBuffer2(data)

		if err != nil {
			t.Error("push buffer error")
			t.FailNow()
			break
		}

		fmt.Println("push one")

		i += 1
	}

}

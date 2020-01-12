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

func TestCheckPlugins(t *testing.T) {

	error := CheckPlugins([]string{"flv", "rtmp"})

	if error != nil {
		t.Error("CheckPlugins", error)
		t.FailNow()
	}
}

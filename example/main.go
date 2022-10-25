package main

import (
	"fmt"
	"github.com/buYoung/gst"
	"time"
)

func testGST() {
	println(1)
	pipeline, err := gst.ParseLaunch("appsrc name=mysource format=time is-live=true do-timestamp=true ! videoconvert ! autovideosink")
	println(2)
	if err != nil {
		panic("pipeline error")
	}
	println(3)
	videoCap := gst.CapsFromString("video/x-raw,format=RGB,width=320,height=240,bpp=24,depth=24")
	println(4)
	element := pipeline.GetByName("mysource")
	println(5)
	element.SetObject("caps", videoCap)
	println(6)
	pipeline.SetState(gst.StatePlaying)
	println(7)
	i := 0
	for {

		if i > 100 {
			break
		}

		data := make([]byte, 320*240*3)

		err := element.PushBuffer(data)

		if err != nil {
			fmt.Println("push buffer error")
			break
		}

		fmt.Println("push one")
		i++
		time.Sleep(50000000)
	}

	pipeline.SetState(gst.StateNull)

	pipeline = nil
	element = nil
	videoCap = nil
}
func main() {

	testGST()
}

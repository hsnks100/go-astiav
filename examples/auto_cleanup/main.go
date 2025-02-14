package main

import (
	"runtime"
	"time"

	"github.com/asticode/go-astiav"
)

func main() {
	_ = astiav.NewAudioFifo(astiav.SampleFormatFltp, 2, 2000)
	runtime.GC()
	time.Sleep(1 * time.Second)
	return
}

package main

import (
	"machine"
	"time"

	"github.com/thegrumpylion/gopher-lua"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(`print("hello TinyGo from lua")`); err != nil {
		panic(err)
	}

	for {
		led.Low()
		time.Sleep(time.Millisecond * 500)

		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}

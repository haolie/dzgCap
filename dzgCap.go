package main

import (
	"fmt"
	"unsafe"

	"github.com/go-vgo/robotgo"

	"dzgCap/win"
)

func main() {
	fmt.Println(2345)
	robotgo.Move(2382, 1084)
	robotgo.Click("left", true)
	//ShowMessage2("","")
}

func ShowMessage2(title, text string) {
	minput:=win.MOUSEINPUT{Dx:2382,Dy:1084,MouseData:0,DwFlags:32769,Time:0,DwExtraInfo:0}
	m_input:=win.MOUSE_INPUT{
		0,
		minput,
	}

	list:=[]win.MOUSE_INPUT{m_input}

	r:= win.SendInput(1,unsafe.Pointer(&list),int32(unsafe.Sizeof(m_input)))
	fmt.Println(r)
}

package main

import (
	"os"

	"golang.org/x/sys/windows"
)

func ConsoleMode() {
	os.Mkdir("files", os.ModePerm)
	var mode uint32
	//handle := C.GetStdHandle(C.STD_OUTPUT_HANDLE)
	handle := windows.Handle(os.Stdout.Fd())
	windows.GetConsoleMode(handle, &mode)
	windows.SetConsoleMode(handle, mode|0x0004)
}

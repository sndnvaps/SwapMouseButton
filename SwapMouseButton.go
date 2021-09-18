//build +windows

package main

import (
	"syscall"
)

func IntPtr(n int) uintptr {
	return uintptr(n)
}

// SwapMouseButton(IsLeft int)
// IsLeft = 0 | 1 -> 0 表示左键，1表示右键
func SwapMouseButton(IsLeft int) {

	user32, err := syscall.LoadLibrary("user32.dll")
	if err != nil {
		abort("LoadLibrary", err.Error())
	}
	defer syscall.FreeLibrary(user32)
	swapMouseButton, err := syscall.GetProcAddress(user32, "SwapMouseButton")
	if err != nil {
		abort("GetProcAdress", err.Error())
	}
	syscall.Syscall(uintptr(swapMouseButton), 1, IntPtr(IsLeft), 0, 0)

}

func abort(funcname string, err string) {
	panic(funcname + "failed: " + err)
}

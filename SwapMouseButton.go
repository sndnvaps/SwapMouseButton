//build +windows

package main

import (
	"embed"
	"image"
	"image/draw"
	"syscall"
)

//go:embed res/mouse_left.png
//go:embed res/mouse_right.png
var res embed.FS
var (
	user32, _ = syscall.LoadLibrary("user32.dll")
	//int SwapMouseButton(int bSwap)
	swapMouseButton, _ = syscall.GetProcAddress(user32, "SwapMouseButton")
	//int GetSystemMetrics(int nIndes)
	getSystemMetrics, _ = syscall.GetProcAddress(user32, "GetSystemMetrics")
)

var (
	SM_SWAPBUTTON = 23 //如果左右鼠标已经交换，则为TRUE
)

func IntPtr(n int) uintptr {
	return uintptr(n)
}

// GetSystemMetrics
// 用于获取鼠标是否已经交换左右键，当交换左右键的时候为TRUE，否则为FALSE
func GetSystemMetrics() bool {
	var nargs uintptr = 1
	ret, _, callErr := syscall.Syscall(uintptr(getSystemMetrics), nargs, IntPtr(SM_SWAPBUTTON), 0, 0)
	if callErr != 0 {
		abort("Call GetSystemMetrics", callErr.Error())
	}

	return ret != 0
}

// SwapMouseButton(IsLeft int)
// IsLeft = 0 | 1 -> 0 表示左键，1表示右键
func SwapMouseButton(IsLeft int) bool {
	var nargs uintptr = 1
	ret, _, callErr := syscall.Syscall(uintptr(swapMouseButton), nargs, IntPtr(IsLeft), 0, 0)
	if callErr != 0 {
		abort("Call SwapMouseButton", callErr.Error())
	}

	return ret != 0
}

func abort(funcname string, err string) {
	panic(funcname + "failed: " + err)
}

func imageToRGBA(src image.Image) *image.RGBA {

	// No conversion needed if image is an *image.RGBA.
	if dst, ok := src.(*image.RGBA); ok {
		return dst
	}

	// Use the image/draw package to convert to *image.RGBA.
	b := src.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), src, b.Min, draw.Src)
	return dst
}

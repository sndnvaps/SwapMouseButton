//build +windows

package main

import (
	"syscall"

	"github.com/andlabs/ui"
)

func setupUI() {

	group := ui.NewGroup("鼠标键配置")
	group.SetMargined(true)

	hbox := ui.NewHorizontalBox()
	hbox.Append(group, true)
	/*
		picbox := ui.NewImage(64, 64)

		b, _ := res.ReadFile("res/mouse_left.png")
		img, _, _ := image.Decode(bytes.NewReader(b))
		mouse_left := imageToRGBA(img)

		picbox.Append(mouse_left)
	*/
	cb := ui.NewCheckbox("切换主要和次要的按钮")

	//获取系统中鼠标主次键的状态
	if GetSystemMetrics() {
		cb.SetChecked(true)
	}

	cb.OnToggled(func(*ui.Checkbox) {
		if cb.Checked() {
			SwapMouseButton(1)
		} else {
			SwapMouseButton(0)
		}
	})
	vbox := ui.NewVerticalBox()

	vbox.Append(cb, true)
	group.SetChild(vbox)

	mainLabel := ui.NewLabel("选择此复选框来将右键设成如选择和拖放等主要性能使用。")
	vbox.Append(mainLabel, true)

	version := ui.NewLabel("本程序由sndnvaps编写")

	//hbox.Append(picbox, true)
	// 生成：垂直容器
	vbox = ui.NewVerticalBox()

	// 往 垂直容器 中添加 控件
	vbox.Append(hbox, false)
	vbox.Append(version, false)

	// 生成：窗口（标题，宽度，高度，是否有 菜单 控件）
	window := ui.NewWindow(`设置鼠标主键`, 200, 100, false)

	// 窗口容器绑定
	window.SetChild(vbox)

	// 设置：窗口关闭时
	window.OnClosing(func(*ui.Window) bool {
		// 窗体关闭
		ui.Quit()
		return true
	})

	// 窗体显示
	window.Show()
}
func main() {
	defer syscall.FreeLibrary(user32)

	err := ui.Main(setupUI)
	if err != nil {
		panic(err)
	}
}

//build +windows

package main

import (
	"github.com/andlabs/ui"
)

func setupUI() {
	// 生成：标签
	isLeft := ui.NewCombobox()
	isLeft.Append("左键") //左键
	isLeft.Append("右键") //右键

	isLeft.SetSelected(1) //默认选择为右键
	//选择
	setIsLeftLabel := ui.NewLabel("请选择要设置的鼠标主键")

	hbox := ui.NewHorizontalBox()
	hbox.Append(isLeft, false)
	hbox.Append(setIsLeftLabel, false)

	//isLeft.Selected()
	// 生成：按钮
	button := ui.NewButton("确认操作")
	// 设置：按钮点击事件
	button.OnClicked(func(*ui.Button) {
		//当为左键的时候，设置参数为0
		if isLeft.Selected() == 0 {
			SwapMouseButton(0)
		}
		//当为右键的时候，设置参数为1
		if isLeft.Selected() == 1 {
			SwapMouseButton(1)
		}
	})
	version := ui.NewLabel("本程序由sndnvaps编写")
	// 生成：垂直容器
	box := ui.NewVerticalBox()

	// 往 垂直容器 中添加 控件
	box.Append(hbox, false)
	box.Append(button, false)
	box.Append(version, false)

	// 生成：窗口（标题，宽度，高度，是否有 菜单 控件）
	window := ui.NewWindow(`设置鼠标主键`, 200, 100, false)

	// 窗口容器绑定

	window.SetChild(box)

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
	err := ui.Main(setupUI)
	if err != nil {
		panic(err)
	}
}

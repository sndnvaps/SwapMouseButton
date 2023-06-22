# SwapMouseButton程序是用go 编写的切换鼠标主键的程序

# 实现方法
 调用系统中的user32.dll动态库中的SwapMouseButton函数来实现

# 编译环境
 1. 需要安装go(测试编译环境为 go 1.17),开启MODULE功能
 2. 需要安装Mingw32-gcc,用于连接动态库
 3. 需要windows系统（交叉编译没有测试过）

# 依赖库
 1. https://github.com/andlabs/ui
 2. http://github.com/josephspurrier/goversioninfo (主要用于生成文件信息)
 3. 国内的用户建议，设置GOPROXY,这样可以更快地下载项目
 
# 如何编译
 1. go get github.com/josephspurrier/goversioninfo
 2. go get github.com/sndnvaps/SwapMouseButton
 3. 如果在国内，可以使用 
     go get gitee.com/sndnvaps/SwapMouseButton

# 直接下载编译好的程序
  1. [Github下载](https://github.com/sndnvaps/SwapMouseButton/releases)
  2. [Gitee下载](https://gitee.com/sndnvaps/SwapMouseButton/releases)

# 运行效果图

![pic1](./pic/pic_01.png)

----------------------------
# 已经测试过的系统

 1. win7 x64
 2. Win vista (没有测试)
 3. Win10 x64
 4. win8 (没有测试)

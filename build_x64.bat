goversioninfo -64 -icon=mouse.ico -manifest=SwapMouseButton.exe.manifest -o=rsrc_x64.syso
rem rsrc -manifest SwapMouseButton.exe.manifest -ico mouse.ico -arch amd64 -o rsrc_x64.syso
go build -ldflags="-H windowsgui -w -s" -o SwapMouseButton.exe
del rsrc_x64.syso
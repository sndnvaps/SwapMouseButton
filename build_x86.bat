goversioninfo -icon=mouse.ico -manifest=SwapMouseButton.exe.manifest
rem rsrc -manifest SwapMouseButton.exe.manifest -ico mouse.ico -arch 386 -o rsrc_x86.syso
go build -ldflags="-H windowsgui -w -s" -o SwapMouseButton.exe
del resource.syso
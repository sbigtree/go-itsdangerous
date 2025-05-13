
### 安装 MinGW-w64
x86_64-w64-mingw32-gcc 在 MinGW-w64 里面
下载地址 https://winlibs.com/
找到 MinGW-w64 下载  附上最终下载地址[link:https://objects.githubusercontent.com/github-production-release-asset-2e65be/220996547/cda8e2e9-6de1-41c3-bd1f-b08b165563b8?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=releaseassetproduction%2F20250512%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20250512T121450Z&X-Amz-Expires=300&X-Amz-Signature=f013de7b55a864784d6379e87c9116e60349a141a6bd8201d25a3d1faf8fdf1b&X-Amz-SignedHeaders=host&response-content-disposition=attachment%3B%20filename%3Dwinlibs-x86_64-posix-seh-gcc-15.1.0-mingw-w64msvcrt-12.0.0-r1.zip&response-content-type=application%2Foctet-stream]

解压后 mingw64 复制到 C:\Program Files (x86)
环境变量 PATH 加上 C:\Program Files (x86)\mingw64\bin

### 下载zlib源码
官网 https://zlib.net/

最终下载地址
https://zlib.net/zlib-1.3.1.tar.gz

解压得到 zlib-1.3.1 

### 下载Cmake
https://cmake.org/download/
最终下载地址 https://objects.githubusercontent.com/github-production-release-asset-2e65be/537699/8218aa6e-d1cf-4ac6-953b-dcbb36c4836a?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=releaseassetproduction%2F20250513%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20250513T013344Z&X-Amz-Expires=300&X-Amz-Signature=ab7fac394fb62aabe19c6c186d005898dfd91aa71dfdf506bc31906b03478586&X-Amz-SignedHeaders=host&response-content-disposition=attachment%3B%20filename%3Dcmake-4.0.2-windows-x86_64.msi&response-content-type=application%2Foctet-stream
双击安装

### 编译
```shell
cd zlib-1.3.1

mkdir build && cd build
cmake .. -G "MinGW Makefiles" -DCMAKE_INSTALL_PREFIX=C:/zlib-install
mingw32-make
mingw32-make install
```

设置环境变量

变量名：C_INCLUDE_PATH
变量值：C:\zlib-install\include

变量名：LIBRARY_PATH
变量值：C:\zlib-install\lib

变量名：PATH（已存在，点击编辑）
新增值：C:\zlib-install\bin




### 安装 MinGW-w64
x86_64-w64-mingw32-gcc 在 MinGW-w64 里面
下载地址 https://winlibs.com/
找到 MinGW-w64 下载  附上最终下载地址[link:https://github.com/brechtsanders/winlibs_mingw/releases/download/15.1.0posix-12.0.0-msvcrt-r1/winlibs-x86_64-posix-seh-gcc-15.1.0-mingw-w64msvcrt-12.0.0-r1.zip]


解压后 mingw64 复制到 C:\Program Files (x86)
环境变量 PATH 加上 C:\Program Files (x86)\mingw64\bin

### 下载zlib源码
官网 https://zlib.net/

最终下载地址
https://zlib.net/zlib-1.3.1.tar.gz

解压得到 zlib-1.3.1 

### 下载Cmake
https://cmake.org/download/
最终下载地址 https://github.com/Kitware/CMake/releases/download/v4.0.2/cmake-4.0.2-windows-x86_64.msi
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



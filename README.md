[![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat-square)](https://github.com/iromli/go-itsdangerous/blob/master/LICENSE)
[![Test](https://github.com/iromli/go-itsdangerous/actions/workflows/test.yml/badge.svg)](https://github.com/iromli/go-itsdangerous/actions/workflows/test.yml)

go-itsdangerous
===============

Like [itsdangerous](https://pythonhosted.org/itsdangerous/) but for Go.

### 编译zlib
```shell
# macOS (darwin/arm64)
gcc -c zlib_wrapper.c -o zlib_wrapper_darwin.o
ar rcs libzlib_wrapper_darwin.a zlib_wrapper_darwin.o

# Linux
gcc -c zlib_wrapper.c -o zlib_wrapper_linux.o
ar rcs libzlib_wrapper_linux.a zlib_wrapper_linux.o

# Windows (可用 cross-compile 或 mingw)
x86_64-w64-mingw32-gcc -c zlib_wrapper.c -o zlib_wrapper_windows.o
ar rcs libzlib_wrapper_windows.a zlib_wrapper_windows.o

```



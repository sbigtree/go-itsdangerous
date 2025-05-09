package zlib

/*
#cgo CFLAGS: -g -Wall
#include "zlib_wrapper.h"  // 包含 C 头文件
#include <stdlib.h>
*/
import "C"
import (
	"encoding/base64"
	"errors"
	"unsafe"
)

// DeflateString 调用 zlib 压缩并返回 base64 编码后的数据
func DeflateString(s string) (string, error) {
	// 将 Go 字符串转换为 C 字符串
	input := C.CString(s)
	defer C.free(unsafe.Pointer(input))

	var outLen C.int
	// 调用 C 的 zlib_deflate 函数，压缩数据
	outPtr := C.zlib_deflate(input, C.int(len(s)), &outLen)
	if outPtr == nil {
		return "", errors.New("compression failed")
	}
	defer C.free(unsafe.Pointer(outPtr))

	// 将压缩后的 C 数据转为 Go 字节切片
	data := C.GoBytes(unsafe.Pointer(outPtr), outLen)
	//return data, nil
	// 对压缩后的字节数据进行 base64 编码
	//return base64.StdEncoding.EncodeToString(data), nil
	return base64.URLEncoding.EncodeToString(data), nil
}

//go:build windows
// +build windows

package zlib

/*
#cgo CFLAGS: -I${SRCDIR}/lib
#cgo LDFLAGS: -L${SRCDIR}/lib -lzlib_wrapper_windows -lzlibstatic
*/
import "C"

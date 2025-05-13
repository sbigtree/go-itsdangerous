//go:build windows
// +build windows

package zlib

/*
#cgo LDFLAGS: -L${SRCDIR}/lib -lzlib_wrapper_windows -lzlibstatic
*/
import "C"

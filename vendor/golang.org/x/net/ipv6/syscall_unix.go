// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin dragonfly freebsd linux,amd64 linux,arm linux,ppc64 linux,ppc64le netbsd openbsd

package ipv6

import (
	"syscall"
	"unsafe"
)

func getsockopt(fd, level, name int, v unsafe.Pointer, l *sysSockoptLen) error {
	if _, _, errno := syscall.Syscall6(syscall.SYS_GETSOCKOPT, uintptr(fd), uintptr(level), uintptr(name), uintptr(v), uintptr(unsafe.Pointer(l)), 0); errno != 0 {
		return error(errno)
	}
	return nil
}

func setsockopt(fd, level, name int, v unsafe.Pointer, l sysSockoptLen) error {
	if _, _, errno := syscall.Syscall6(syscall.SYS_SETSOCKOPT, uintptr(fd), uintptr(level), uintptr(name), uintptr(v), uintptr(l), 0); errno != 0 {
		return error(errno)
	}
	return nil
}

// Based on ssh/terminal:
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux,!appengine darwin freebsd

package isterminal

import (
	"syscall"
	"unsafe"
)

// IsTerminal returns true if the given file descriptor is a terminal.
func IsTerminal(fd int) bool {
	var termios Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}

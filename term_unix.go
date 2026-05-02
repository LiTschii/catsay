//go:build !windows

package main

import (
	"os"
	"strconv"
	"syscall"
	"unsafe"
)

func termWidth() int {
	if v := os.Getenv("COLUMNS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			return n
		}
	}
	type winsize struct {
		Row, Col, Xpixel, Ypixel uint16
	}
	var ws winsize
	if _, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		0x5413,
		uintptr(unsafe.Pointer(&ws)),
	); errno == 0 && ws.Col > 0 {
		return int(ws.Col)
	}
	return 80
}

// +build linux,!arm64 !linux,!windows

package panicwrap

import (
	// "syscall"
	"golang.org/x/sys/unix"
)

func dup2(oldfd, newfd int) error {
	// return syscall.Dup2(oldfd, newfd)
	return unix.Dup2(oldfd, newfd)
}

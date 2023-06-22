// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build linux && loong64
// +build linux,loong64

package netpoll

import (
	"syscall"
)

const EPOLLET = syscall.EPOLLET

// EpollCreate implements epoll_create1.
func EpollCreate(flag int) (fd int, err error) {
	var r0 uintptr
	r0, _, err = syscall.RawSyscall(syscall.SYS_EPOLL_CREATE1, uintptr(flag), 0, 0)
	if err == syscall.Errno(0) {
		err = nil
	}
	return int(r0), err
}

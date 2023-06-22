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

package netpoll

import (
	"net"
)

// CreateListener return a new Listener.
func CreateListener(network, addr string) (l Listener, err error) {
	if network == "udp" {
		// TODO: udp listener.
		return udpListener(network, addr)
	}
	// tcp, tcp4, tcp6, unix
	ln, err := net.Listen(network, addr)
	if err != nil {
		return nil, err
	}
	return ConvertListener(ln)
}

// ConvertListener converts net.Listener to Listener
func ConvertListener(l net.Listener) (nl Listener, err error) {
	return &listener{
		Listener: l,
	}, nil
}

// TODO: udpListener does not work now.
func udpListener(network, addr string) (l Listener, err error) {
	panic("not implement")
}

var _ net.Listener = &listener{}

type listener struct {
	Listener net.Listener
}

func (ln *listener) Accept() (net.Conn, error) {
	conn, err := ln.Listener.Accept()
	if err != nil {
		return nil, err
	}
	return newNetFD(conn), nil
}

func (ln *listener) Close() error {
	return ln.Listener.Close()
}

func (ln *listener) Addr() net.Addr {
	return ln.Listener.Addr()
}

// Fd implements Listener.
func (ln *listener) Fd() (fd int) {
	return getFd(ln.Listener)
}

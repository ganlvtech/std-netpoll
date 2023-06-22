// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// This file may have been modified by CloudWeGo authors. (“CloudWeGo Modifications”).
// All CloudWeGo Modifications are Copyright 2022 CloudWeGo authors.

package netpoll

import (
	"net"
)

// BUG(mikio): On JS, NaCl and Plan 9, methods and functions related
// to UnixConn and UnixListener are not implemented.

// BUG(mikio): On Windows, methods and functions related to UnixConn
// and UnixListener don't work for "unixgram" and "unixpacket".

// UnixAddr represents the address of a Unix domain socket end point.
type UnixAddr struct {
	net.UnixAddr
}

// ResolveUnixAddr returns an address of Unix domain socket end point.
//
// The network must be a Unix network name.
//
// See func Dial for a description of the network and address
// parameters.
func ResolveUnixAddr(network, address string) (*UnixAddr, error) {
	addr, err := net.ResolveUnixAddr(network, address)
	if err != nil {
		return nil, err
	}
	return &UnixAddr{*addr}, nil
}

// UnixConnection implements Connection.
type UnixConnection struct {
	connection
}

// newUnixConnection wraps UnixConnection.
func newUnixConnection(conn Conn) (connection *UnixConnection, err error) {
	connection = &UnixConnection{}
	err = connection.init(conn, nil)
	if err != nil {
		return nil, err
	}
	return connection, nil
}

// DialUnix acts like Dial for Unix networks.
//
// The network must be a Unix network name; see func Dial for details.
//
// If laddr is non-nil, it is used as the local address for the
// connection.
func DialUnix(network string, laddr, raddr *UnixAddr) (*UnixConnection, error) {
	switch network {
	case "unix", "unixgram", "unixpacket":
	default:
		return nil, &net.OpError{Op: "dial", Net: network, Source: &laddr.UnixAddr, Addr: &raddr.UnixAddr, Err: net.UnknownNetworkError(network)}
	}
	c, err := net.DialUnix(network, &laddr.UnixAddr, &raddr.UnixAddr)
	if err != nil {
		return nil, &net.OpError{Op: "dial", Net: network, Source: &laddr.UnixAddr, Addr: &raddr.UnixAddr, Err: err}
	}
	return newUnixConnection(newNetFD(c))
}

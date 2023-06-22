// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// This file may have been modified by CloudWeGo authors. (“CloudWeGo Modifications”).
// All CloudWeGo Modifications are Copyright 2022 CloudWeGo authors.

package netpoll

import (
	"context"
	"net"
)

// TCPAddr represents the address of a TCP end point.
type TCPAddr struct {
	net.TCPAddr
}

// ResolveTCPAddr returns an address of TCP end point.
//
// The network must be a TCP network name.
//
// If the host in the address parameter is not a literal IP address or
// the port is not a literal port number, ResolveTCPAddr resolves the
// address to an address of TCP end point.
// Otherwise, it parses the address as a pair of literal IP address
// and port number.
// The address parameter can use a host name, but this is not
// recommended, because it will return at most one of the host name's
// IP addresses.
//
// See func Dial for a description of the network and address
// parameters.
func ResolveTCPAddr(network, address string) (*TCPAddr, error) {
	addr, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		return nil, err
	}
	return &TCPAddr{*addr}, nil
}

// TCPConnection implements Connection.
type TCPConnection struct {
	connection
}

// newTCPConnection wraps *TCPConnection.
func newTCPConnection(conn Conn) (connection *TCPConnection, err error) {
	connection = &TCPConnection{}
	err = connection.init(conn, nil)
	if err != nil {
		return nil, err
	}
	return connection, nil
}

// DialTCP acts like Dial for TCP networks.
//
// The network must be a TCP network name; see func Dial for details.
//
// If laddr is nil, a local address is automatically chosen.
// If the IP field of raddr is nil or an unspecified IP address, the
// local system is assumed.
func DialTCP(ctx context.Context, network string, laddr, raddr *TCPAddr) (*TCPConnection, error) {
	switch network {
	case "tcp", "tcp4", "tcp6":
	default:
		return nil, &net.OpError{Op: "dial", Net: network, Source: &laddr.TCPAddr, Addr: &raddr.TCPAddr, Err: net.UnknownNetworkError(network)}
	}
	if raddr == nil {
		return nil, &net.OpError{Op: "dial", Net: network, Source: &laddr.TCPAddr, Addr: nil, Err: errMissingAddress}
	}
	if ctx == nil {
		ctx = context.Background()
	}
	c, err := net.DialTCP(network, &laddr.TCPAddr, &raddr.TCPAddr)
	if err != nil {
		return nil, &net.OpError{Op: "dial", Net: network, Source: &laddr.TCPAddr, Addr: &raddr.TCPAddr, Err: err}
	}
	return newTCPConnection(newNetFD(c))
}

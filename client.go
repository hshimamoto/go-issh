// MIT License Copyright (C) 2022 Hiroshi Shimamoto
package issh

import (
	"net"

	"golang.org/x/crypto/ssh"
)

type Client struct {
	*ssh.Client
}

// ssh handshake through conn
func Handshake(conn net.Conn, addr string, config *Config) (*Client, error) {
	cconn, cchans, creqs, err := ssh.NewClientConn(conn, addr, &config.config)
	if err != nil {
		return nil, err
	}
	client := ssh.NewClient(cconn, cchans, creqs)
	return &Client{client}, nil
}

// new ssh connection
func Dial(addr string, config *Config) (*Client, error) {
	client, err := ssh.Dial("tcp", addr, &config.config)
	return &Client{client}, err
}

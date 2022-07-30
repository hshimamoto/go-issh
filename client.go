// MIT License Copyright (C) 2022 Hiroshi Shimamoto
package issh

import (
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

type Client struct {
	*ssh.Client
	keepalivesec int
}

// ssh handshake through conn
func Handshake(conn net.Conn, addr string, config *Config) (*Client, error) {
	cconn, cchans, creqs, err := ssh.NewClientConn(conn, addr, &config.config)
	if err != nil {
		return nil, err
	}
	client := ssh.NewClient(cconn, cchans, creqs)
	return &Client{client, 0}, nil
}

// new ssh connection
func Dial(addr string, config *Config) (*Client, error) {
	client, err := ssh.Dial("tcp", addr, &config.config)
	return &Client{client, 0}, err
}

// start keepalive
func (cli *Client) StartKeepalive(sec int) {
	if cli.keepalivesec > 0 {
		return
	}
	go func() {
		for cli.keepalivesec > 0 {
			_, _, err := cli.SendRequest("keepalive@golang.org", true, nil)
			if err != nil {
				cli.Close()
				break
			}
			time.Sleep(time.Second * time.Duration(cli.keepalivesec))
		}
	}()
}

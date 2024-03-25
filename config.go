// MIT License Copyright (C) 2022, 2024 Hiroshi Shimamoto
package issh

import (
	"os"

	"golang.org/x/crypto/ssh"
)

type Config struct {
	config ssh.ClientConfig
}

// create inscure config with user and key
func NewConfigBytes(user string, key []byte) (*Config, error) {
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}
	config := ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return &Config{config: config}, nil
}

func NewConfig(user, keyfile string) (*Config, error) {
	key, err := os.ReadFile(keyfile)
	if err != nil {
		return nil, err
	}
	return NewConfigBytes(user, key)
}

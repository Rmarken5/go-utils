package ssh

import (
	cssh "golang.org/x/crypto/ssh"
)

type Options struct {
	username        string
	password        string
	address         string
	networkProtocol string
	hostKeyCallback cssh.HostKeyCallback
	bannerCallback  cssh.BannerCallback
}
type Option func(*Options)

func NewSSHClient(opts ...Option) (*cssh.Client, error) {
	options := &Options{
		username:        "",
		password:        "",
		address:         "",
		networkProtocol: "",
		hostKeyCallback: cssh.InsecureIgnoreHostKey(),
		bannerCallback:  cssh.BannerDisplayStderr(),
	}

	for _, o := range opts {
		o(options)
	}
	config := &cssh.ClientConfig{
		Config: cssh.Config{},
		User:   options.username,
		Auth: []cssh.AuthMethod{
			cssh.Password(options.password),
		},
		HostKeyCallback:   options.hostKeyCallback,
		BannerCallback:    options.bannerCallback,
		ClientVersion:     "",
		HostKeyAlgorithms: nil,
		Timeout:           0,
	}

	return cssh.Dial(options.networkProtocol, options.address, config)

}

func WithUsername(username string) func(options *Options) {
	return func(options *Options) {
		options.username = username
	}
}

func WithPassword(password string) func(*Options) {
	return func(options *Options) {
		options.password = password
	}
}

func WithAddress(address string) func(*Options) {
	return func(options *Options) {
		options.address = address
	}
}

func WithNetworkProtocol(protocol string) func(*Options) {
	return func(options *Options) {
		options.networkProtocol = protocol
	}
}

func WithHostKeyCallback(callback cssh.HostKeyCallback) func(*Options) {
	return func(options *Options) {
		options.hostKeyCallback = callback
	}
}

func WithBannerCallback(bannerCallback cssh.BannerCallback) func(*Options) {
	return func(options *Options) {
		options.bannerCallback = bannerCallback
	}
}

package ssh

import cssh "golang.org/x/crypto/ssh"

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

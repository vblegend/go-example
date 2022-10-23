package config

type Ssl struct {
	Key    string
	Pem    string
	Enable bool
	Domain string
}

var SslConfig = new(Ssl)

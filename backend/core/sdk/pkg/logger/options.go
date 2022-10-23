/*
 * @Author: lwnmengjing
 * @Date: 2021/6/10 10:26 上午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/10 10:26 上午
 */

package logger

type Option func(*options)

type options struct {
	path     string
	level    string
	stdout   string
	cap      uint
	location bool
}

func setDefault() options {
	return options{
		path:     "temp/logs",
		level:    "warn",
		stdout:   "default",
		location: false,
	}
}

func WithPath(s string) Option {
	return func(o *options) {
		o.path = s
	}
}

func WithLevel(s string) Option {
	return func(o *options) {
		o.level = s
	}
}

func WithStdout(s string) Option {
	return func(o *options) {
		o.stdout = s
	}
}

func WithCap(n uint) Option {
	return func(o *options) {
		o.cap = n
	}
}
func WithLocation(l bool) Option {
	return func(o *options) {
		o.location = l
	}
}

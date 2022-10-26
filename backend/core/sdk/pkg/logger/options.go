/*
 * @Author: lwnmengjing
 * @Date: 2021/6/10 10:26 上午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/10 10:26 上午
 */

package logger

type Option func(*options)

type options struct {
	path       string
	fileName   string
	fileSuffix string
	level      string
	enabled    bool
	cap        uint
	location   bool
}

func setDefault() options {
	return options{
		path:       "temp/logs",
		level:      "warn",
		enabled:    true,
		fileName:   "logfile",
		fileSuffix: "log",
		location:   false,
	}
}
func WithFileName(s string) Option {
	return func(o *options) {
		o.fileName = s
	}
}
func WithFileSuffix(s string) Option {
	return func(o *options) {
		o.fileSuffix = s
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

func WithEnabled(s bool) Option {
	return func(o *options) {
		o.enabled = s
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

/*
 * @Author: lwnmengjing
 * @Date: 2021/6/3 8:33 上午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/3 8:33 上午
 */

package writer

// Options 可配置参数
type Options struct {
	path     string
	fileName string
	suffix   string //文件扩展名
	cap      uint
}

func setDefault() Options {
	return Options{
		path:   "/temp/logs",
		suffix: "log",
		cap:    1024000,
	}
}

// Option set options
type Option func(*Options)

// WithPath set path
func WithPath(s string) Option {
	return func(o *Options) {
		o.path = s
	}
}
func WithFileName(s string) Option {
	return func(o *Options) {
		o.fileName = s
	}
}

// WithSuffix set suffix
func WithSuffix(s string) Option {
	return func(o *Options) {
		o.suffix = s
	}
}

// WithCap set cap
func WithCap(n uint) Option {
	return func(o *Options) {
		o.cap = n
	}
}

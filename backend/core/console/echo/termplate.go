package echo

import (
	"io"
	"regexp"
)

type IEchoTemplateWriter interface {
	AddRule(regex *regexp.Regexp, options ...Option)
	Write(data []byte) (n int, err error)
}

// 设置正则rule  直接write
func Template(w io.Writer) IEchoTemplateWriter {
	return &echoTemplate{
		regexs:  make([]*regexp.Regexp, 0),
		options: make([][]Option, 0),
		w:       w,
	}
}

type echoTemplate struct {
	regexs  []*regexp.Regexp
	options [][]Option
	w       io.Writer
}

func (t *echoTemplate) AddRule(regex *regexp.Regexp, options ...Option) {
	t.regexs = append(t.regexs, regex)
	t.options = append(t.options, options)
}

func (t *echoTemplate) Parse(data []byte, regex *regexp.Regexp, options []Option) []byte {
	if regex.Match(data) {
		val := regex.ReplaceAllStringFunc(string(data), func(element string) string {
			return Encode(element, options...)
		})
		return []byte(val)
	}
	return data
}

func (t *echoTemplate) Write(data []byte) (n int, err error) {
	datalen := len(data)
	for i := 0; i < len(t.regexs); i++ {
		regex := t.regexs[i]
		options := t.options[i]
		data = t.Parse(data, regex, options)
	}
	t.w.Write(data)

	return datalen, nil
}

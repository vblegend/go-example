package std

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type FontStyle uint8

type Color uint16

func (c Color) GetControl() uint16 {
	return uint16(c) >> 8
}

func (c Color) GetCode() uint16 {
	return (uint16(c) << 8) >> 8
}

const (
	Foreground = 3
	Background = 4
	// 前景色：黑色
	FBlack = Color(0x0300)
	// 前景色：红色
	FRed = Color(0x0301)
	// 前景色：绿色
	FGreen = Color(0x0302)
	// 前景色：黄色
	FYellow = Color(0x0303)
	// 前景色：蓝色
	FBlue = Color(0x0304)
	// 前景色：品红色
	FMagenta = Color(0x0305)
	// 前景色：青色
	FCyan = Color(0x0306)
	// 前景色：白色
	FWhite = Color(0x0307)
	// =========================

	// 背景色：黑色
	BBlack = Color(0x0400)
	// 背景色：红色
	BRed = Color(0x0401)
	// 背景色：绿色
	BGreen = Color(0x0402)
	// 背景色：黄色
	BYellow = Color(0x0403)
	// 背景色：蓝色
	BBlue = Color(0x0404)
	// 背景色：品红色
	BMagenta = Color(0x0405)
	// 背景色：青色
	BCyan = Color(0x0406)
	// 背景色：白色
	BWhite = Color(0x0407)
	// =========================
	// 字体风格：正常
	Normal = FontStyle(0)
	// 字体风格：粗体
	Bold = FontStyle(1)
	// 字体风格：透明
	Transparent = FontStyle(2)
	// 字体风格：斜体
	Italics = FontStyle(3)
	// 字体风格：下划线
	Underline = FontStyle(4)
)

type StdOutBuilder interface {
	// 设置输出流对象，使 Print系列函数输出至输出流（默认输出流为os.stdout）
	WithStdout(w io.Writer) StdOutBuilder
	// 打印内容到输出流（默认输出流为os.stdout）
	Print()
	// 打印换行内容到输出流（默认输出流为os.stdout）
	Println()
	// 获取构建的字符串内容
	String() string
	// 追加一个换行符
	Appendln() StdOutBuilder
	// 追加一个对象内容至结尾，可以设定内容参数
	Append(object interface{}, options ...interface{}) StdOutBuilder
}

// 编码一段标准输出流消息
func Encode(object interface{}, options ...interface{}) string {
	fontColor := 0
	backColor := 0
	fontStyle := Normal
	segment := ""
	for _, v := range options {
		switch t := v.(type) {
		case Color:
			{
				ctl := t.GetControl()
				if ctl == Foreground {
					fontColor = int(t.GetCode())
				} else if ctl == Background {
					backColor = int(t.GetCode())
				} else {
					panic("error：Invalid control character")
				}
			}
		case FontStyle:
			{
				fontStyle = t
			}
		}
	}
	if fontStyle > Normal {
		segment = fmt.Sprintf(";%d", fontStyle)
	}
	return fmt.Sprintf("\033[3%d;4%d%sm%v\033[0m", fontColor, backColor, segment, object)
}

func Builder() StdOutBuilder {
	return &stdBuilder{
		buffers: make([]string, 0),
		w:       os.Stdout,
	}
}

type stdBuilder struct {
	buffers []string
	w       io.Writer
}

func (std *stdBuilder) Appendln() StdOutBuilder {
	std.buffers = append(std.buffers, "\n")
	return std
}

func (std *stdBuilder) Append(obj interface{}, options ...interface{}) StdOutBuilder {
	std.buffers = append(std.buffers, Encode(obj, options...))
	return std
}

func (std *stdBuilder) String() string {
	return strings.Join(std.buffers, "")
}

func (std *stdBuilder) WithStdout(w io.Writer) StdOutBuilder {
	std.w = w
	return std
}

func (std *stdBuilder) Print() {
	fmt.Fprint(std.w, std.String())
}

func (std *stdBuilder) Println() {
	fmt.Fprintln(std.w, std.String())
}

package echo

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Option uint16

func (c Option) GetControl() uint16 {
	return uint16(c) >> 8
}

func (c Option) GetCode() uint16 {
	return (uint16(c) << 8) >> 8
}

const (
	FontStyleFlag  = 0
	ForegroundFlag = 3
	BackgroundFlag = 4
)

const (
	// 前景色：黑色
	FBlack = Option(0x0300)
	// 前景色：红色
	FRed = Option(0x0301)
	// 前景色：绿色
	FGreen = Option(0x0302)
	// 前景色：黄色
	FYellow = Option(0x0303)
	// 前景色：蓝色
	FBlue = Option(0x0304)
	// 前景色：品红色
	FMagenta = Option(0x0305)
	// 前景色：青色
	FCyan = Option(0x0306)
	// 前景色：白色
	FWhite = Option(0x0307)
	// =========================

	// 背景色：黑色
	BBlack = Option(0x0400)
	// 背景色：红色
	BRed = Option(0x0401)
	// 背景色：绿色
	BGreen = Option(0x0402)
	// 背景色：黄色
	BYellow = Option(0x0403)
	// 背景色：蓝色
	BBlue = Option(0x0404)
	// 背景色：品红色
	BMagenta = Option(0x0405)
	// 背景色：青色
	BCyan = Option(0x0406)
	// 背景色：白色
	BWhite = Option(0x0407)

	// =========================
	// 字体风格：正常
	Normal = Option(0x0000)
	// 字体风格：粗体
	Bold = Option(0x0001)
	// 字体风格：透明
	Transparent = Option(0x0002)
	// 字体风格：斜体
	Italics = Option(0x0003)
	// 字体风格：下划线
	Underline = Option(0x0004)
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
	Append(object interface{}, options ...Option) StdOutBuilder
}

// 编码一段标准输出流消息
func Encode(object interface{}, options ...Option) string {
	if len(options) == 0 {
		return Strval(object)
	}
	fontColor := -1
	backColor := -1
	fontStyle := Normal
	for _, v := range options {
		ctl := v.GetControl()
		if ctl == ForegroundFlag {
			fontColor = int(v.GetCode())
		} else if ctl == BackgroundFlag {
			backColor = int(v.GetCode())
		} else if ctl == FontStyleFlag {
			fontStyle = Option(v.GetCode())
		} else {
			panic("error：Invalid control character")
		}
	}

	result := make([]byte, 0)
	result = append(result, []byte("\033[")...)
	if fontColor >= 0 {
		result = append(result, '3')
		result = strconv.AppendInt(result, int64(fontColor), 8)
		result = append(result, ';')
	}
	if backColor >= 0 {
		result = append(result, '4')
		result = strconv.AppendInt(result, int64(backColor), 8)
		result = append(result, ';')
	}
	if fontStyle > Normal {
		result = strconv.AppendInt(result, int64(fontStyle), 8)
	}
	result = append(result, 'm')
	result = append(result, []byte(Strval(object))...)
	result = append(result, []byte("\033[0m")...)

	return string(result)
	// fmt.Println(string(result))
	// return fmt.Sprintf("\033[3%d;4%d%sm%v\033[0m", fontColor, backColor, segment, object)
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func Strval(value interface{}) string {
	// interface 转 string
	if value == nil {
		return ""
	}
	switch t := value.(type) {
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 64)
	case int:
		return strconv.Itoa(t)
	case uint:
		return strconv.Itoa(int(t))
	case int8:
		return strconv.Itoa(int(t))
	case uint8:
		return strconv.Itoa(int(t))
	case int16:
		return strconv.Itoa(int(t))
	case uint16:
		return strconv.Itoa(int(t))
	case int32:
		return strconv.Itoa(int(t))
	case uint32:
		return strconv.Itoa(int(t))
	case int64:
		return strconv.FormatInt(t, 10)
	case uint64:
		return strconv.FormatUint(t, 10)
	case string:
		return value.(string)
	case bool:
		if t {
			return "true"
		} else {
			return "false"
		}
	case []byte:
		return fmt.Sprintf("%v", value)
	default:
		return fmt.Sprintf("%v", value)
	}
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

func (std *stdBuilder) Append(obj interface{}, options ...Option) StdOutBuilder {
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

package echo

import (
	"fmt"
	"runtime"
)

// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
//
// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

// Colors
// Reset       = "\033[0m"
// Red         = "\033[31m"
// Green       = "\033[32m"
// Yellow      = "\033[33m"
// Blue        = "\033[34m"
// Magenta     = "\033[35m"
// Cyan        = "\033[36m"
// White       = "\033[37m"
// BlueBold    = "\033[34;1m"
// MagentaBold = "\033[35;1m"
// RedBold     = "\033[31;1m"
// YellowBold  = "\033[33;1m"

// 其中0x1B是标记，[开始定义颜色，依次为：模式，背景色，前景色，0代表恢复默认颜色。

// \033[34;45mHello\x1b[0m

func Black(msg string) string {
	return SetColor(msg, 0, 0, TextBlack)
}

func Red(msg string) string {
	return SetColor(msg, 0, 0, TextRed)
}

func Green(msg string) string {
	return SetColor(msg, 0, 0, TextGreen)
}

func Yellow(msg string) string {
	return SetColor(msg, 0, 0, TextYellow)
}

func Blue(msg string) string {
	return SetColor(msg, 0, 0, TextBlue)
}

func Magenta(msg string) string {
	return SetColor(msg, 0, 0, TextMagenta)
}

func Cyan(msg string) string {
	return SetColor(msg, 0, 0, TextCyan)
}

func White(msg string) string {
	return SetColor(msg, 0, 0, TextWhite)
}

func SetColor(msg string, conf, bg, text int) string {
	if runtime.GOOS == "linux" {
		// 其中0x1B是标记，[开始定义颜色，依次为：模式，背景色，前景色，0代表恢复默认颜色。
		return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, text, msg, 0x1B)
	}
	return msg
}

package types

type LoggerConfigure struct {
	Path       string // 日志路径
	Level      string // 日志等级  trace/debug/info/warn/error/fatal
	Enabled    bool   // 是否输出到文件 false 仅打印至console
	Cap        uint
	Location   bool // 是否显示代码位置
	FileName   string
	FileSuffix string
}

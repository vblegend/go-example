package env

type ApplicationRunMode string

const (
	//开发模式
	Develop ApplicationRunMode = ApplicationRunMode("dev")
	//生产模式
	Production ApplicationRunMode = ApplicationRunMode("prod")
	// 测试模式
	Test ApplicationRunMode = ApplicationRunMode("Test")
)

// 当前运行模式
var Mode ApplicationRunMode

func SetMode(_mode ApplicationRunMode) {
	Mode = _mode
}

func GetMode() ApplicationRunMode {
	return Mode
}

func ModeIs(_mode ApplicationRunMode) bool {
	return Mode == _mode
}

func ParseMode(str string) ApplicationRunMode {
	if str == "dev" {
		return Develop
	} else if str == "Test" {
		return Test
	} else {
		return Production
	}
}

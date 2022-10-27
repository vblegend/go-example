package jobs

import "fmt"

func init() {
	RegisterClass("测试任务", ExamplesOne{})
}

// 新job必须实现JobsExec接口
type ExamplesOne struct {
}

func (t ExamplesOne) Exec(arg string) error {
	fmt.Println(arg)
	return nil
}

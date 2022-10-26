package jobs

import "fmt"

func init() {
	RegisterClass(ExamplesOne{})
}

// 新job必须实现JobsExec接口
type ExamplesOne struct {
}

func (t ExamplesOne) Exec(arg string) error {
	fmt.Println(arg)
	return nil
}

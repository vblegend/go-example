package jobs

// 需要将定义的struct 添加到字典中；
// 字典 key 可以配置到 自动任务 调用目标 中；
func InitJob() {
	jobList = map[string]JobsExec{
		"ExamplesOne": ExamplesOne{},
		"TimeSync":    TimeSync{},
		// ...
	}
}

// 新添加的job 必须按照以下格式定义，并实现Exec函数
type ExamplesOne struct {
}

type TimeSync struct {
}

func (t ExamplesOne) Exec(arg interface{}) error {
	// TODO: 这里需要注意 Examples 传入参数是 string 所以 arg.(string)；请根据对应的类型进行转化；
	switch arg.(type) {
	case string:
		if arg.(string) != "" {
			// log.Infof("JobCore ExamplesOne[%s] exec success..", arg.(string))
		} else {
			// log.Info("JobCore ExamplesOne[] exec success..")
		}
		break
	}

	return nil
}

func (t TimeSync) Exec(arg interface{}) error {

	switch arg.(type) {
	case string:
		if arg.(string) != "" {

		} else {
			// log.Info("JobCore ExamplesOne[] exec success..")
		}
		break
	}

	return nil
}

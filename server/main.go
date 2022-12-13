package main

import (
	"server/cmd"
)

func main() {
	cmd.Execute()
	// p, err := plugin.Open("main.so")
	// if err != nil {
	// 	panic(err)
	// }
	// v, err := p.Lookup("V") //调用插件中的V字段
	// if err != nil {
	// 	panic(err)
	// }
	// f, err := p.Lookup("F") //调用插件中的F函数
	// if err != nil {
	// 	panic(err)
	// }
	// *v.(*int) = 100
	// f.(func())()
}

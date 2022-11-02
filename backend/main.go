package main

import (
	"backend/cmd"
	"plugin"
)

func main() {
	cmd.Execute()
	p, err := plugin.Open("main.so")
	if err != nil {
		panic(err)
	}
	v, err := p.Lookup("") //调用插件中的V字段
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("") //调用插件中的F函数
	if err != nil {
		panic(err)
	}
	*v.(*int) = 100
	f.(func())()
}

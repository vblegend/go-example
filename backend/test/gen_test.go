package test

import (
	//"backend/models/tools"
	//"os"

	"backend/core/sdk/console/std"
	"backend/core/sdk/pkg"
	"fmt"
	"os"
	"reflect"
	"testing"
	//"text/template"
)

func TestXxx(t *testing.T) {

	exist := pkg.ProcessExist("sleep")
	fmt.Println(exist)
}

func TestConsole(t *testing.T) {

	std.Builder().WithStdout(os.Stdout).
		Appendln().
		Append("Test", std.BRed, std.FYellow, std.Italics).
		Append("[", std.BGreen, std.FMagenta, std.Bold).
		Append(1234.65565677145, std.BRed, std.FBlue, std.Bold).
		Append("]", std.BGreen, std.FMagenta, std.Bold).
		Appendln().
		Append("[", std.BWhite, std.FBlack, std.Bold).
		Append(1234.65565677145, std.BWhite, std.FBlack, std.Bold).
		Append("]", std.BWhite, std.FBlack, std.Bold).
		Println()

}

func TestMap_d2(t *testing.T) {
	var m1 map[string]interface{}
	var m2 = &m1

	v2 := reflect.ValueOf(m2)
	// v2i := v2.IndirectValueRecursive()

	// t.Logf("v2i: %v     | v2: %v      | m2: %v", v2i.Type(), v2.Type(), m2)
	nmi := reflect.MakeMap(reflect.TypeOf(m1))
	nmi.SetMapIndex(reflect.ValueOf("today"), reflect.ValueOf("is monday"))
	// t.Logf("nmi: %v", nmi.Type())
	// t.Logf("     %v | %v | %v", v2.CanAddr(), v2i.CanAddr(), nmi.CanAddr())
	//*(v2.Interface().(*map[string]interface{})) = nmi.Interface().(map[string]interface{})
	v2.Elem().Set(nmi)
	t.Logf("m2 = %v", m2)
}

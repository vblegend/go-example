package test

import (
	//"backend/models/tools"
	//"os"

	"backend/core/console"
	"backend/core/sdk/pkg"
	"fmt"
	"reflect"
	"regexp"
	"testing"
	//"text/template"
)

// 0x1b[0;0;32m   0x1b[0m    \x1b[0;0;31mAA\x1b[0m
func TestXxx(t *testing.T) {
	raw := []byte(fmt.Sprintf("{ v1: %s ,v2: %s ,v3: %s }", console.Red("AA"), console.Green("BB"), console.Blue("CC")))
	re := regexp.MustCompile(`/\x1b\[*/\x1b\[0m`)
	if re.Match(raw) {
		fmt.Println("111")
	}
	fmt.Println("222")
	exist := pkg.ProcessExist("sleep")
	fmt.Println(exist)
}

func TestConsole(t *testing.T) {

	// std.Builder().WithStdout(os.Stdout).
	// 	Appendln().
	// 	Append("Test", std.BRed, std.FYellow, std.Italics).
	// 	Append("[", std.BGreen, std.FMagenta, std.Bold).
	// 	Append(1234.65565677145, std.BRed, std.FBlue, std.Bold).
	// 	Append("]", std.BGreen, std.FMagenta, std.Bold).
	// 	Appendln().
	// 	Append("[", std.BWhite, std.FBlack, std.Bold).
	// 	Append(1234.65565677145, std.BWhite, std.FBlack, std.Bold).
	// 	Append("]", std.BWhite, std.FBlack, std.Bold).
	// 	Println()

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

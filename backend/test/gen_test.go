package test

import (
	//"backend/models/tools"
	//"os"

	"backend/core/dataflow/flowtest"
	"backend/core/echo"
	"os"
	"reflect"
	"regexp"
	"testing"
	//"text/template"
)

func Testddddd(t *testing.T) {
	// poller.Start(conn, netpoll.EventRead, func() {
	// 	// We spawn goroutine here to prevent poller wait loop
	// 	// to become locked during receiving packet from ch.
	// 	go Receive(ch)
	// })

}

func TestXxxddd(t *testing.T) {

	a := flowtest.NewAddNode("", "")
	b := flowtest.NewAddNode("", "")
	c := flowtest.NewAddNode("", "")
	d := flowtest.NewAddNode("", "")
	e := flowtest.NewAddNode("", "")
	f := flowtest.NewPrintNode("", "")
	a.Outputs().Add(b)
	b.Outputs().Add(c)
	b.Outputs().Add(e)
	c.Outputs().Add(d)
	d.Outputs().Add(e)
	e.Outputs().Add(f)

	a.Next(1, 1)

	// exist := pkg.ProcessExist("sleep")
	// fmt.Println(exist)
}

// 0x1b[0;0;32m   0x1b[0m    \x1b[0;0;31mAA\x1b[0m
func TestXxx(t *testing.T) {

}

func TestConsole(t *testing.T) {

	tmp := echo.Template(os.Stdout)

	tmp.AddRule(regexp.MustCompile(`\[ERROR\]`), echo.FRed, echo.Bold)
	tmp.AddRule(regexp.MustCompile(`\[WARN\]`), echo.FYellow, echo.Bold)
	tmp.AddRule(regexp.MustCompile(`\[INFO\]`), echo.FCyan, echo.Bold)
	tmp.AddRule(regexp.MustCompile(`\[DEBUG\]`), echo.FGreen, echo.Bold)
	tmp.AddRule(regexp.MustCompile(`\[GORM\]`), echo.BYellow, echo.FRed, echo.Bold)
	tmp.AddRule(regexp.MustCompile(`([1-2][0-9][0-9][0-9]-[0-1]{0,1}[0-9]-[0-3]{0,1}[0-9])\s(20|21|22|23|[0-1]\d):[0-5]\d:[0-5]\d`), echo.FRed, echo.BCyan, echo.Italics)

	tmp.Write([]byte("2022-09-25 15:32:42 [DEBUG] ABCD.A_B_C_D.a*b*c*d\n"))
	tmp.Write([]byte("2022-09-25 15:32:42 [INFO] ABCD.A_B_C_D.a*b*c*d\n"))
	tmp.Write([]byte("2022-09-25 15:32:42 [WARN] ABCD.A_B_C_D.a*b*c*d\n"))
	tmp.Write([]byte("2022-09-25 15:32:42 [ERROR] ABCD.A_B_C_D.a*b*c*d\n"))
	tmp.Write([]byte("2022-09-25 15:32:42 [GORM] ABCD.A_B_C_D.a*b*c*d\n"))

	return

	echo.Builder().WithStdout(os.Stdout).
		Appendln().
		Append("Test", echo.BRed).
		Append("[", echo.BGreen, echo.Bold).
		Append(1234.65565677145, echo.FBlue, echo.Bold).
		Append("]", echo.FMagenta).
		Appendln().
		Append([]byte{1, 2, 3, 4}, echo.FMagenta).
		Appendln().
		Append("[", echo.BWhite, echo.FBlack, echo.Bold).
		Append(1234.65565677145, echo.BWhite, echo.FBlack).
		Append("]", echo.Bold).
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

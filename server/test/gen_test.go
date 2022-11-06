package test

import (
	//"server/models/tools"
	//"os"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"server/sugar/echo"
	"server/sugar/encoding/pkcs1"
	"server/sugar/mpool"
	"server/sugar/random"
	"server/sugar/types"
	"testing"
	"time"
	//"text/template"
)

type StructT struct {
	Id   int
	Name string
}

func TestSignData(t *testing.T) {
	// encoding.NewKey()
	data := []byte("Hello World")
	signer, err := pkcs1.NewSignature("./private.pem", "123456")
	if err != nil {
		fmt.Println(err)
		return
	}
	out, err := signer.Sign(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)
	validator, err := pkcs1.NewValidator("./public.pem")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = validator.Verify(data, out)

	fmt.Println(err)

}

func TestRandom(t *testing.T) {
	k := types.SIntMap{}
	s := types.Strings{}
	t2 := types.Ints{}
	for i := 0; i < 10; i++ {
		s.Append(fmt.Sprintf("%d", random.Between(0, 100)))
		k.Put(fmt.Sprintf("%d", i), random.Between(0, 100))
		t2.Append(random.Between(0, 100))
	}
	s.Push("-")
	s.Sort()
	t2.Sort()
	t2.Reverse()

	vss := k.SortKeys()
	t2.Remove(5)
	inx := t2.IndexOf(7)
	val := k.Get("5")
	k.Clear()
	res := s.Pop()
	fmt.Println(res, vss, val, inx)

}

func TestPool(t *testing.T) {
	options := &mpool.Options{
		Capacity: 100,
		MaxIdle:  90,
		New: func() interface{} {
			return &StructT{}
		},
		MinIdleTime: time.Hour,
	}
	pool := mpool.NewObjectPool(options)

	arry := make([]*StructT, 0)

	for i := 0; i < 1000; i++ {
		v, e := pool.Malloc()
		if v != nil {
			s := v.(*StructT)
			s.Id = i + 1
			arry = append(arry, s)
		}
		fmt.Println(i, v, e)
	}
	fmt.Println("==========================================")
	for _, v := range arry {
		pool.Free(v)
	}
	fmt.Println("==========================================")
	for i := 0; i < 100; i++ {
		v, e := pool.Malloc()
		fmt.Println(i, v, e)
	}

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

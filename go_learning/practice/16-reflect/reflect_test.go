package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f int64 = 12
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f), reflect.ValueOf(f).Type())
}

func checkType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Println("Integer.")
	case reflect.Float32, reflect.Float64:
		fmt.Println("float")
	default:
		fmt.Println("Unknow type.")
	}
}

func TestType(t *testing.T) {
	var f float64 = 12
	checkType(f)
}

type student struct {
	CookieId string
	Name     string `formate:"normal"`
	Age      int
}

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "name1", 2: "age"}
	b := map[int]string{1: "name1", 2: "age"}

	t.Log("a==b?", reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 3, 2}
	s3 := []int{2, 3, 1}

	t.Log("s1==s2?", reflect.DeepEqual(s1, s2))
	t.Log("s2==s3?", reflect.DeepEqual(s2, s3))

	stu1 := student{"stu0001", "wanghao1", 23}
	stu2 := student{"stu0001", "wanghao1", 23}
	t.Log("stu1==stu2", stu1 == stu2)

}

func (s *student) Update(name string) {
	s.Name = name
}

func TestInvokeByName(t *testing.T) {
	stu := student{"stu0001", "wanghao", 23}

	t.Log(reflect.ValueOf(stu).FieldByName("Name"))
	t.Log(reflect.ValueOf(stu).FieldByName("Name").Type())

	reflect.ValueOf(&stu).MethodByName("Update").Call([]reflect.Value{reflect.ValueOf("new_name")})
	t.Log(stu)
}

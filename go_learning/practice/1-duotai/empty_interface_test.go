package duotai

import (
	"fmt"
	"reflect"
	"testing"
)

//断言判断变量类型

func DoSomeThing(p interface{})  {
	//if v,ok := p.(int); ok {
	//	fmt.Println("Integer",v)
	//	return
	//}
	//
	//if s,ok := p.(string);ok {
	//	fmt.Println("String",s)
	//	return
	//}
	//fmt.Println("Unknow type")
}

func TestEmptyInterface(t *testing.T)  {
	//DoSomeThing(10)
	//DoSomeThing("10")

	age := map[int]string{1:"wanghao"}
	if v,ok := (interface{})(age).(int);ok {
		fmt.Println("Integer",v)
	} else if v,ok := (interface{})(age).(string);ok {
		fmt.Println("String",v)
	} else {
		fmt.Println("Unkonw type")
	}

	age_type := reflect.TypeOf(age)
	fmt.Println(age_type)

	switch age_type := (interface{})(age).(type) {
	case int:
		fmt.Println("Integer",age_type)
	case string:
		fmt.Println("String",age_type)
	default:
		fmt.Println("Unknow type",age_type)
	}
}

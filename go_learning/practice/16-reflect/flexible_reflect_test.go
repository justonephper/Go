package reflect

import (
	"errors"
	"reflect"
	"testing"
)

type student struct {
	CookieName string
	Name       string
	Age        int
}

func fillNameAndAge(stu *student, setting map[string]interface{}) error {
	if reflect.TypeOf(stu).Kind() != reflect.Ptr {
		return errors.New("stu is not a pointer.")
	}
	if reflect.TypeOf(stu).Elem().Kind() != reflect.Struct {
		return errors.New("stu's elem not a struct")
	}
	if setting == nil {
		return errors.New("setting is nil")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	for k, v := range setting {
		if field, ok = reflect.TypeOf(stu).Elem().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			reflect.ValueOf(stu).Elem().FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "wanghao", "Age": 23}
	stu := new(student)

	if err := fillNameAndAge(stu, settings); err != nil {
		t.Log("fill failed,errorinfo:", err)
	}
	t.Log(stu)
}

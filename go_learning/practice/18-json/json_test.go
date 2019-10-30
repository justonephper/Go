package jsontest

import (
	"encoding/json"
	"testing"
)

var jsonStr = `{
	"basic_info":{
		"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go","C"]
	}
}`

func TestEmbeddeJson(t *testing.T) {
	e := new(Employee)
	if err := json.Unmarshal([]byte(jsonStr), e); err != nil {
		t.Log(err)
	}
	t.Log(e.BasicInfo.Age)
	if v, err := json.Marshal(e); err != nil {
		t.Log(err)
	} else {
		t.Log(string(v))
	}
}

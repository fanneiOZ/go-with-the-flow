package handler

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type object struct {
	K *string `json:"key"`
	V int64   `json:"value,omitempty,string"`
	//c int    `arbitrary:"-"`
}

func TestJson(t *testing.T) {
	//jsonData := bytes.NewBuffer("asdfasdf")
	t.Run("case 1", func(t *testing.T) {
		b := "hello"
		hello := object{
			//K: "hello",
			K: &b,
			V: 2,
		}
		jsonString, err := json.Marshal(hello)
		t.Log(string(jsonString))
		var a object
		err = json.Unmarshal([]byte(`{"v":"20"}`), &a)
		t.Log(err)
		t.Log(a)
	})

	t.Run("case 2", func(t *testing.T) {
		a := "hello"
		hello := object{K: &a, V: 2}

		reflectedType := reflect.TypeOf(hello)
		//reflectedValue := reflect.ValueOf(hello)
		n := reflectedType.NumField()
		f1 := reflectedType.Field(n - 1)
		t.Log(f1.Tag)

		//t.Log(reflectedType)
		//t.Log(reflectedValue)

		//t.Log(f1)
	})

	t.Run("case 3", func(t *testing.T) {
		bolB, _ := json.Marshal(true)
		fmt.Println(string(bolB))
		intB, _ := json.Marshal(1)
		fmt.Println(string(intB))
		fltB, _ := json.Marshal(2.34)
		fmt.Println(string(fltB))
		strB, _ := json.Marshal("gopher")
		fmt.Println(string(strB))
		slcD := []string{"apple", "peach", "pear"}
		slcB, _ := json.Marshal(slcD)
		fmt.Println(string(slcB))
		mapD := map[string]int{"apple": 5, "lettuce": 7}
		mapB, _ := json.Marshal(mapD)
		fmt.Println(mapB)

	})

	t.Run("case 4", func(t *testing.T) {
		r := &Response{
			ResourceType: ResourceType{Id: "1234", Type: "response"},
			Key:          "abc",
			Value:        "bc",
		}

		result, _ := json.Marshal(r)

		fmt.Println(string(result))
	})

}

type ResourceType struct {
	Id   string `json:"object_id"`
	Type string `json:"type"`
}

type Response struct {
	ResourceType
	Key   string `json:"key"`
	Value string `json:"value"`
}

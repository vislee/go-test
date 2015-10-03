package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	var mm map[string]string = make(map[string]string)
	mm["name"] = "liwq"
	mm["age"] = "23"
	data, err := json.Marshal(mm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(data))

	data1 := `{"topics":[{"channels":[],"name":"hello","paused":false},{"channels":[],"name":"test","paused":false}],"version":"0.3.6-alpha"}`
	var res interface{}
	dec := json.NewDecoder(bytes.NewBuffer([]byte(data1)))
	dec.UseNumber()
	dec.Decode(&res)
	fmt.Println(res)
	m1, ok := res.(map[string]interface{})
	if ok {
		tty := m1["topics"].([]interface{})
		fmt.Println(tty)
		for _, val := range tty {
			mval, ok := val.(map[string]interface{})
			if !ok {
				fmt.Println("map error")
				return
			}
			fmt.Println(mval["name"])
		}
	}
}

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonstr := []byte(`{"foo": "fooooo", "bar": "barrrrrr\nbumbum", "blah": 12}`)
	blah := &Blah{}
	json.Unmarshal(jsonstr, blah)
	fmt.Printf("%+v\n", blah)

	blah.Bar = "Hello\nWorld"
	data, err:= json.Marshal(blah)
	if err == nil {
		fmt.Printf("Blah:=%s", string(data))
	}
}

type Blah struct {
	Foo string `json:"foo"`
	Bar string `json:"bar"`
}

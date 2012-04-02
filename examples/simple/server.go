package main

import (
	"github.com/zozor/webgui"
	"encoding/json"
)

func main() {
	webgui.SetHandler("test", Dostuff)
	webgui.SetHandler("reverse", ReverseData)
	webgui.StartServer("localhost:27000")
}

func Dostuff(js []byte) []byte {
	var data struct{Data string}
	err := json.Unmarshal(js, &data)
	if err != nil {
		return []byte(`{"Data": `+err.Error()+`}`)
	}
	
	//Do something with data
	//...
	
	out, err := json.Marshal(data)
	if err != nil {
		return []byte(`{"Data": `+err.Error()+`}`)
	}
	return out
}

func ReverseData(js []byte) []byte {
	out := make([]byte, len(js))
	for i, v := range js {
		out[len(out)-1-i] = v
	}
	return out
}

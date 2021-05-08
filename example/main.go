package main

import (
	"fmt"
	"github.com/gkuanine/httputils"
)

func main() {
	var a = make(map[string]string)
	a["a"] = "1"
	resp, err := httputils.PostJson("http://baidu.com", 3, a)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

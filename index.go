package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

	var std = Std{Name: "armando", Sex: "male"}
	b, _ := json.Marshal(std)
	res, err := http.Post("http://localhost:90", "application/json", bytes.NewReader(b))
	if err != nil {
		panic(err.Error())
	}
	byt, _ := io.ReadAll(res.Body)
	fmt.Println("request already sent")
	fmt.Println(string(byt))

	

}

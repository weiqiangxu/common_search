package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	query := getJSON()
	w := &Where{}
	z := []byte(query)
	e := json.Unmarshal(z, w)
	if e != nil {
		log.Fatal("error is", e)
	}
	zzz := Builder(w)
	fmt.Println(zzz)
}

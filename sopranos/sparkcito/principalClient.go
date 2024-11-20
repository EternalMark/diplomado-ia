package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"sparkcito/util"
)

func client(req util.SearchClienReq) {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(c).Encode(req)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	req := util.SearchClienReq{
		Palabra: "Spider",
	}

	client(req)

	// fmt.Println(res)
	// var input string
	// fmt.Scanln(&input)
}

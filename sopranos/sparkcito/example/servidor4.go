package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Persona struct {
	Nombre string
	Email  []string
}

func servidor() {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(c)
	}
}

// func handleClient(c net.Conn) {
// 	var persona Persona
// 	err := gob.NewDecoder(c).Decode(&persona)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	} else {
// 		fmt.Println("Mensaje: ", persona)
// 	}
// }

func handleClient(c net.Conn) {
	var persona Persona
	err := gob.NewDecoder(c).Decode(&persona)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Mensaje: ", persona)
	}
}


func main() {
	go servidor()

	var input string
	fmt.Scanln(&input)
}

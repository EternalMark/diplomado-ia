package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"sparkcito/util"
)

func client4workers(workerReq util.SearchWorkerReq,port string) {
	c, err := net.Dial("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(c).Encode(workerReq)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func handlePrincipalClient(c net.Conn) {

	var clientReq util.SearchClienReq
	err := gob.NewDecoder(c).Decode(&clientReq)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Mensaje: ", clientReq)

		workerReq:=util.SearchWorkerReq{
			Palabra: clientReq.Palabra,
			Archivo: "enwiki_1.xml",
		}
		go client4workers(workerReq,"9990")

		workerReq.Archivo="enwiki_2.xml"
		go client4workers(workerReq,"9991")

		workerReq.Archivo="enwiki_3.xml"
		go client4workers(workerReq,"9992")

		workerReq.Archivo="enwiki_4.xml"
		go client4workers(workerReq,"9993")
	}
}

func master() {
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
		go handlePrincipalClient(c)
	}
}

func main() {
	go master()
	fmt.Println("MasterServer Iniciado")
	var input string
	fmt.Scanln(&input)
}

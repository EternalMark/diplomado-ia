package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"sparkcito/util"
)

func handleMasterServer(c net.Conn) {

	var workerReq util.SearchWorkerReq
	err := gob.NewDecoder(c).Decode(&workerReq)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Iniciando búsqueda de: ", workerReq)

		var res util.RespuestaBuscaPalabra
		res = util.BuscaPalabra(workerReq.Archivo, workerReq.Palabra)
		// res = util.BuscaPalabra("./data/enwiki_1.xml", "Spider")

		fmt.Println("Respuesta: ", res)
	}
}

func worker(port string) {
	s, err := net.Listen("tcp", port)
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
		go handleMasterServer(c)
	}
}

func main() {
	port := ":9990"
	go worker(port)
	fmt.Println("Worker Server Iniciado en el puerto", port)
	var input string
	fmt.Scanln(&input)
}

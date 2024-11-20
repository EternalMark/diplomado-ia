package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func BuscaPalabra(archivo string, palabra string) RespuestaBuscaPalabra {
	lines := 0
	c := 0
	start := time.Now()
	file, err := os.Open(archivo)
	// file, err := os.Open("./data/enwiki_1.xml")

	if err != nil {
		log.Fatalf("Error al abrir el archivo: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines++
		// if strings.Contains(scanner.Text(), "Spider") {
		temp := strings.Count(scanner.Text(), palabra)
		if temp > 0 {
			// fmt.Println(scanner.Text())
			c += temp
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	duration := time.Since(start)
	fmt.Println("Duration = ", duration, " Lineas:", lines, "Coincid:", c)
	// var input string
	// fmt.Scanln(&input)
	resp := RespuestaBuscaPalabra{
		coincidencias: int64(c),
		duration:      duration,
		lineas:        int64(lines),
	}
	return resp
}

// func main() {
// 	resp := BuscaPalabra("./data/enwiki_1.xml", "Spider")
// 	fmt.Println("Duration = ", resp.duration, " Lineas:", resp.lineas, " Coincidencias", resp.coincidencias)
// }

package util

import (
	"time"
)

type SearchWorkerReq struct {
	Palabra string //Palabra a buscar
	Archivo string
}

type SearchClienReq struct {
	Palabra string //Palabra a buscar
}

type RespuestaBuscaPalabra struct {
	coincidencias int64
	duration      time.Duration
	lineas        int64
}
# Sparkcito v1  by Sopranos

Proyecto creado para el diplomado de inteligencia artificial.

Este proyecto tiene por objetivo leer toda la información de la wikipedia buscando una palabra dada por el usuario, en un corto periodo de tiempo

## Ejecución

* Es necesario que las particiones del xml se encuentren en una carpeta "data" en cada dispositivo al nivel del servicio worker. Ej. "./data/enwiki_1.xml"
* Ejecutar los servicios workers (workerServer.go) en varios dipositivos que esten conectados a una misma red. Insertando como parámetro el puerto. 
    go run workerServer.go 9990

* Ejecutar el servicio master (masterServer.go) en el dispositivo principal. Este debe conocer de antemano las ips y puertos de todos los workers
    go run masterServer.go

* Ejecutar el programa cliente (principalClient.go). En este programa colocar la ip y puerto del servicio master y enviar la palabra que se desea buscar.
    go run principalClient.go


## TO DO
* Al iniciar el master revisar el estatus de todos los workers
* Control de flujos cuando un worker no responda o no este disponible
* Hacer la suma de todas las coincidencias de todos los workers 
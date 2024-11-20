#Este programa lee y divide los articulos de wiki

using Pkg
Pkg.add("XML")
using XML
import Pkg
Pkg.add("EzXML")
using EzXML

function main()

    println("Here we go ...")

    #PARAMETROS
    # l=10000001
    l = 782745622 #Lineas totales del Archivo Original
    ns = 10 #Numero de slices requeridas

    #Abriendo Archivo original
    archivo = "./data/enwiki_2014_06.xml"
    io = open(archivo)

    #Creando primer slice
    sliceActual = string("./data/enwiki_1.xml")
    ioSlice = open(sliceActual, "w")

    
    i = 0 #Contador lineas archivo Original
    j = 1 #Contador de Slices
    
    #Obteniendo número de lineas por Slice
    nls = l ÷ ns
    println("nls ", nls)

    #lifecheck
    lc = nls ÷ 10

    while i < l
        #while !eof(io)

        i = i + 1

        #El corazón de este código
        linea = readline(io)
        write(ioSlice, string(linea, "\n"))

        #lifeCheck
        if mod(i, lc) == 0
            println("Leyendo linea ", i)
        end

        #Si el número de lineas es mayor al que debería tener el slice
        #y ya se encontró la palabra </page>
        #Entonces debemos crear un nuevo slice
        if i > nls*j && j < ns && contains(linea, "</page>" )

            j = j + 1
            println("CREANDO nuevo Slice ", j)
            close(ioSlice)

            sliceActual = string("./data/enwiki_", j, ".xml")
            ioSlice = open(sliceActual, "w")
        end

    end
    println("Proceso finalizado. Lineas recorridas ", i)
    close(io)
    close(ioSlice)

end

@time main()
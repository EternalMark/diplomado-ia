#Este programa lee y divide los articulos de wiki


using Pkg
Pkg.add("XML")
using XML
import Pkg
Pkg.add("EzXML")
using EzXML


function main()

    println("Leyendo archivo")
    archivo = "/home/cecilia/Documents/sopranos_ok/enwiki_2014_06.xml"
    
    i=0
    io = open(archivo)
    while  i<10 #!eof(io)
        linea = readline(io)
        println(linea)
        i = i +1
        # Operaciones con `linea`
    end
    close(io)

    # Cargar el archivo XML
   
    #open(archivo, "r") do file
    #    for line in eachline(file)
    #        # Procesa la línea actual
    #        println(line)  # Puedes reemplazar println por cualquier otra operación
    #    end
    #end

    #filename = joinpath(dirname(pathof(XML)), "..", "test", "data", "enwiki_2014_06.xml")
    #doc = read("/home/cecilia/Documents/sopranos_ok/enwiki_2014_06.xml", Node) 
   
    #hay que leer como texto 

end

@time main()
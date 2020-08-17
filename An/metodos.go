// Code generated by PABLO. DO NOT EDIT.

package An

import (
	"log"
	"os"
	"strconv"
)

// LAS FUNCIONES DE ESTE ARCHIVO ESTAN COMPARTIDAS PORQUE PERTENECEN AL MISMO PACKAGE , siempre deben de iniciar con MAYUSCULA el nombre del metodo para ser exportado

// QuitarComillas: igual para la declaracion de variables MAYUSCULA AL INICIO para usar ne otras partes

func init() {
	// es una funcion especial que permite iniciazar variables o estructuras
}

//quita \* de continuacion de linea
func QuitarSimboloNextLine(cadena string) string {
	salida := ""
	for i := 0; i < len(cadena); i++ {
		if (i != len(cadena)-2) && (i != len(cadena)-1) {
			salida += string(cadena[i])
		}
	}
	return salida
}

// QuitarComillas lo que hace es quitar comillas xd
func QuitarComillas(ruta string) string {

	salida := ""
	if ruta[0] == '"' {
		//fmt.Println("tiene comillas")
		for i := 0; i < len(ruta); i++ {
			if i != 0 && (i != len(ruta)-1) {
				salida += string(ruta[i])
			}

		}
		//fmt.Println("SALIDA: " + salida)
	} else {
		salida = ruta
	}
	return salida
}

func CrearDisco(numero string, ruta string, nombre string, K_o_M string) {
	ruta = QuitarComillas(ruta)
	tamanio, _ := strconv.ParseInt(numero, 10, 64)
	//	fmt.Printf("%v\n", tamanio)

	size := int64(0)
	if K_o_M == "K" || K_o_M == "k" {
		size = int64(tamanio * 1024)
	} else { // SINO SON MEGABYTES
		size = int64(tamanio * 1024 * 1024)
	}
	rutaCompleta := ruta + nombre
	fichero, err := os.Create(rutaCompleta)
	if err != nil {
		log.Fatal("fallo creando el archivo de salida")
	}

	_, err = fichero.Seek(size-1, 0)

	if err != nil {
		log.Fatal("FALLO EN SEEK")
	}
	_, err = fichero.Write([]byte{0})
	if err != nil {
		log.Fatal("write")
	}
	err = fichero.Close()
	if err != nil {
		log.Fatal("ERROR AL CEERAR EL PROGRAMA ")
	}

}

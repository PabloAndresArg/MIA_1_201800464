// Code generated by PABLO. DO NOT EDIT.

package An

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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

func verificarRuta(ruta string) {
	rutas := strings.Split(ruta, "/")

	var temporal string = ""
	if ruta[len(ruta)-1] != '/' {
		i := 0
		for i < len(rutas) {
			temporal += rutas[i] + "/"
			CrearDirectorio_si_no_exist(temporal)
			//mkdisk-SIze->1-path->/home/pablo/Escritorio/carpetae/hola/-name->prub.dsk-unit->K
			i++
		}
	} else {
		i := 0
		for i < len(rutas) {
			if i != (len(rutas) - 1) {
				temporal += rutas[i] + "/"
				CrearDirectorio_si_no_exist(temporal)
			}
			i++
		}
	}
}
func CrearDisco(numero string, ruta string, nombre string, K_o_M string) {
	ruta = QuitarComillas(ruta)
	tamanio, _ := strconv.ParseInt(numero, 10, 64)
	//	fmt.Printf("%v\n", tamanio)
	verificarRuta(ruta)
	size := int64(0)
	if K_o_M == "K" || K_o_M == "k" {
		size = int64(tamanio * 1024)
	} else { // SINO SON MEGABYTES
		size = int64(tamanio * 1024 * 1024)
	}
	rutaCompleta := ruta + nombre
	fichero, err := os.Create(rutaCompleta)
	defer fichero.Close()
	if err != nil {
		log.Fatal("fallo creando el archivo de salida")
	} else {
		fmt.Print("\n\nDisco creado Correctamente:")
	}
	var cero int8 = 0 // asignando el cero
	direccion_cero := &cero
	var binario_ bytes.Buffer
	binary.Write(&binario_, binary.BigEndian, direccion_cero) // SE ESCRIBE UN CERO AL INICIO DEL ARCHIVO
	escribirBinariamente(fichero, binario_.Bytes())
	fichero.Seek(size-1, 0) // posicionarse en la pos 0
	var bin2_ bytes.Buffer  // se escribe un cero al final del archivo
	binary.Write(&bin2_, binary.BigEndian, direccion_cero)
	escribirBinariamente(fichero, bin2_.Bytes())

	/*
		METIENDO EL STRUCT AL DISCO
	*/
	fichero.Seek(0, 0) // POS AL INICIO DEL ARCHIVO
	// SEREALIZACION DEL STRUCT , escribir al inicio del archivo el struct
	FechaFormatoTime := time.Now()
	mbr := TipoMbr{Tamanio: size}
	copy(mbr.Fecha[:], FechaFormatoTime.String())
	dirMemory_mbr := &mbr
	fmt.Printf("\nFECHA: %s\nTamanio: %v\n", mbr.Fecha, mbr.Tamanio)

	var bin3_ bytes.Buffer
	binary.Write(&bin3_, binary.BigEndian, dirMemory_mbr)
	escribirBinariamente(fichero, bin3_.Bytes())

	// limpiar variables
	Name_ = ""
	Path_ = ""
	Size_ = ""
	Unit_m_ = "M"

}

func escribirBinariamente(fichero *os.File, bytes []byte) {
	_, erro := fichero.Write(bytes)
	if erro != nil {
		log.Fatal(erro)
	}
}

// elimina un disco duro o archivo
func EliminarDisco(ruta_absoluta string) {
	ruta_absoluta = QuitarComillas(ruta_absoluta)
	fmt.Println("¿ESTA SEGURO DE QUERER ELIMINAR ESTE DISCO? ")
	fmt.Print("Presione 1 para confirmar, dsino presione 0")
	var decision int
	fmt.Scanln(&decision)
	if decision == 1 {
		erro := os.Remove(ruta_absoluta)
		if erro == nil {
			fmt.Println("Disco eliminado...")
		} else {
			fmt.Printf("ERROR , NO SE PUEDO ELIMINAR EL DISCO: %v\n", erro)
		}
	}
}

// crea un direcctorio si no encuentra la ruta
func CrearDirectorio_si_no_exist(dir__ string) {

	if _, err := os.Stat(dir__); os.IsNotExist(err) {

		err = os.Mkdir(dir__, 0755)
		fmt.Println("Crea la carpeta:  " + dir__)
		if err != nil {
			panic(err)
		}
	}

}

// LeerBinariamente es para leer archivos binarios
func LeerBinariamente(direccion_archivo_binario string) {
	// atributos obligatorios : NAME , PATH  SIZE
	// OPCIONALES  unit , type , fit , delete , add

	archivoDisco, err := os.Open(QuitarComillas(direccion_archivo_binario))
	defer archivoDisco.Close()
	if err != nil {
		log.Fatal(err)
	}
	mrbAuxiliar := TipoMbr{}
	tamanioMbr := binary.Size(mrbAuxiliar) // este mbrAuxiliar aunque este vacio ya tiene el tamanio por defecto por eso aca se usa
	//binary.Size(estructura1) // PROBAR CON ESTE PORQUE ES MAS EXACTO YA QUE MI STRUCT TIENE ATRIBUTOS int64
	datosEnBytes := leerBytePorByte(archivoDisco, tamanioMbr)
	buff := bytes.NewBuffer(datosEnBytes)                   // lo convierto a buffer porque eso pedia la funcion
	err = binary.Read(buff, binary.BigEndian, &mrbAuxiliar) //se decodifica y se guarda en el mbrAuxiliar , asi que despues de aca ya tengo el original
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}
	fmt.Println(mrbAuxiliar) // ACA YA TENGO EL MBR QUE ESTABA EN EL AUX
	//  v – formats the value in a default format
	//  d – formats decimal integers
	//  g – formats the floating-point numbers
	//  b – formats base 222 numbers
	//  o – formats base 888 numbers
	//  t – formats true or false values
	//  s – formats string values

}

func leerBytePorByte(archivoDisco *os.File, tamanio int) []byte {
	bytes := make([]byte, tamanio) // hago un arreglo dinamico de bytes
	_, err := archivoDisco.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}
func dameUnNumeroRandom() int64 { // para el signature del mbr
	return int64(rand.Intn(7*7*7-1) + (10 / 10))
}

package An

import (
	"fmt"

	"github.com/TwinProduction/go-color"
)

var RegistroSignatures []int64 //Slice sin tamaño definido
var DiscosMontados_ []disco    //Slice sin tamaño definido
// VARIABLES GLOBLALES

var Name_ string = ""
var Path_ string = ""
var Size_ string = ""
var Unit_m_ string = "M"

var Unit_k_ string = "k" // para fdisk el por defecto es K
var tipo_particion_ string = "p"
var FIT_ string = "wf" // por defecto
var OPCION_DELETE_ string = ""
var add_ string = ""

/*

	FUNCIONES EXTRA EXPORTADAS

*/
//
func addRegistroSignature(id int64) {
	RegistroSignatures = append(RegistroSignatures, id)
}

func esRepetido(id int64) bool {
	for t := 0; t < len(RegistroSignatures); t++ {
		if RegistroSignatures[t] == id {
			return true
		}
	}
	addRegistroSignature(id)
	return false
}

func addMonturaDisco(disk disco) {
	DiscosMontados_ = append(DiscosMontados_, disk)
}

func mostrarMounts() {
	fmt.Println("***************DISCOS MONTADOS****************")
	b := false
	for u_u := 0; u_u < len(DiscosMontados_); u_u++ {
		b = true
		//DiscosMontados_[u_u].imprimirMontura() // ACA HACER UN FOR CON LAS MONTURAS DEL DISCO MONTADAS E MOSTRARLAS
	}
	if !(b) {
		println(color.Yellow + "POR EL MOMENTO NO HAY PARTICIONES MONTADAS" + color.Reset)
	}
	fmt.Println("**********************************************")
}

func yaRegistreElPathEnElMount(path string) bool {

	return false
}

/*
	ALGORITMO PARA LAS PARTICIONES

*/

// SI EXISTE LA PARTICION , ENTONCES BUSCO EL PATH SI YA TENGO REGISTRADO EL PATH SOLO LO GUARDO EN MI disco particionesMontadas
// si es un nuevo disco mi contador de letrass incrementa y pasa a a ser el b ,
// si el path no esta registrado pero si existe entoneces LO AGREGO AL ARREGLO DE DISCOS DE DONDE SACO MIS PARTICIONES

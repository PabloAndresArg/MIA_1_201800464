package An

var RegistroSignatures []int64 //Slice sin tamaño definido
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

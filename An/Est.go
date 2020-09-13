package An

var RegistroSignatures []int64 //Slice sin tamaño definido
var DiscosMontados_ []disco    //Slice sin tamaño definido
// VARIABLES GLOBLALES

var Name_ string = ""
var Path_ string = ""
var Size_ string = ""
var Unit_m_ string = "M"

var Unit_k_ string = "k" // para fdisk el por defecto es K Y unit
var tipo_particion_ string = "p"
var FIT_ string = "wf" // por defecto
var OPCION_DELETE_ string = ""
var add_ string = ""

// variables para comando REP
// Name_ , Path_
var Id_vdlentraNumero_ = ""
var Commando_Ruta_ = ""

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

/*

	VARIABLES GLOBALES PARA CONTROLAR EL MONTAJE DE LA LETRA

*/

var CONT_lETRA int16 = 0

// PARA LA FASE 2 , unit_k , add_  Id_vdlentraNumero_
var type_ string = "full"

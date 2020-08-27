package An

import "fmt"

func generarReporte() {
	if len(Path_) != 0 && len(Name_) != 0 && len(Id_vdlentraNumero_) != 0 {
		switch Name_ {
		case "mbr":
			grahpMBR(Id_vdlentraNumero_)
		default:
			fmt.Println("ERRROR COMANDO INCORRECTO")
		}
	}
	limpiarVariablesRep()
}
func limpiarVariablesRep() {
	Name_ = ""
	Path_ = ""
	Id_vdlentraNumero_ = ""
	Commando_Ruta_ = ""
}

func grahpMBR(id string) {
	// NECESITO IR A ATRAER EL PATH , TENIENDO EN CUENTA QUE PUEDO BUSCAR EN MI LISTA id[2] me da la letra y ya tengo el disco que necesito
	var letraID = string(id[2])
	fmt.Println("DISCO: " + letraID)

}

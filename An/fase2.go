package An

import (
	"github.com/TwinProduction/go-color"
)

func metodoMKFS(id string, tipo string, add string, unit string) {
	if len(id) != 0 {
		if validacionQueEsteMontada(string(id[2]), id) {

		}
	} else {
		println(color.Red + "-----------------------------------" + color.Reset)
		println(color.Red + "NO PUSO UN PARAMETRO OBLIGATORIO :)" + color.Reset)
		println(color.Red + "-----------------------------------" + color.Reset)
	}
	limpiarVariablesMKFS()
}

func limpiarVariablesMKFS() {
	Id_vdlentraNumero_ = ""
	type_ = "full"
	add_ = ""
	Unit_k_ = "k"
}

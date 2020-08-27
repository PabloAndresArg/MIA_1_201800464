package An

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/TwinProduction/go-color"
)

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
	_disco_ := getDiscoMontadoPorLetraID(letraID)
	archivoDisco, err := os.OpenFile(QuitarComillas(_disco_.Path), os.O_RDWR, 0644)
	defer archivoDisco.Close()
	if err != nil {
		log.Fatal(err)
	}
	mrbAuxiliar := TipoMbr{}
	tamanioMbr := binary.Size(mrbAuxiliar)
	datosEnBytes := leerBytePorByte(archivoDisco, tamanioMbr)
	buff := bytes.NewBuffer(datosEnBytes)
	err = binary.Read(buff, binary.BigEndian, &mrbAuxiliar)
	if err != nil {
		log.Fatal("error al leer", err)
		println(color.Red + "NO SE PUDO ENCONTRAR EL MBR " + color.Reset)
		return
	}
	crearTxt(mrbAuxiliar)
	generarImg()
}
func generarImg() {
	consola := exec.Command("dot", "-Tjpg", "mbr.txt", "-o repMBR.jpg")
	if errOr := consola.Run(); errOr != nil {
		println(color.Red + "Error al ejecutar comando dot" + color.Reset)

		log.Fatal(errOr)
	}
}

func crearTxt(m TipoMbr) { // pasar tambien la ruta
	w, err := os.OpenFile("mbr.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
		println(color.Red + "Error al crear el archivo" + color.Reset)
		return
	}
	w.WriteString("Digraph tablaMbr{\n")
	w.WriteString(" charset = latin1;\n")
	w.WriteString("tbl[\n")
	w.WriteString("shape = plaintext\n")
	w.WriteString("label =<")
	w.WriteString("<table border = '4' cellborder = '3' color = 'black' cellspacing = '4' bgcolor = \"bisque4\">")
	w.WriteString("<tr>\n") // TITULO
	w.WriteString("<td color = \"grey\" colspan = '2'> TABLA MBR </td> ")
	w.WriteString("</tr>\n") // FIN TITULO
	w.WriteString("<tr>\n")  // FILA 1
	w.WriteString("<td color = \"grey\">Size</td>\n")
	w.WriteString("<td color = \"grey\">" + fmt.Sprint(m.Tamanio) + "</td>\n")
	w.WriteString("</tr>\n")
	w.WriteString("<tr>\n") // FILA 2
	w.WriteString("<td color = \"grey\">Signature</td>\n")
	w.WriteString("<td color = \"grey\">" + fmt.Sprint(m.DiskSignature) + "</td>\n")
	w.WriteString("</tr>\n")
	w.WriteString("<tr>\n") // FILA 3
	w.WriteString("<td color = \"grey\">Fecha</td>\n")
	w.WriteString("<td color = \"grey\">" + string(m.Fecha[:]) + "</td>\n")
	w.WriteString("</tr>\n")
	/*		las n particiones */

	w.WriteString("</table>\n")
	w.WriteString(">];\n")
	w.WriteString("}\n")
	if errOr := w.Close(); errOr != nil {
		log.Fatal(errOr)
		return
	}
}

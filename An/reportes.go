package An

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

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
		return
	}
	println(color.Blue + "REPORTE GENERADO" + color.Reset)

}

func crearTxt(m TipoMbr) { // pasar tambien la ruta
	w, err := os.Create("mbr.txt")
	if err != nil {
		println(color.Red + "Error al crear el archivo" + color.Reset)
		return
	}
	w.WriteString("Digraph tablaMbr{\n")
	w.WriteString("tbl[\n")
	w.WriteString("shape = plaintext\n")
	w.WriteString("label =<")
	w.WriteString("<table border = '4' cellborder = '3' color = 'black' cellspacing = '4' bgcolor = \"bisque4\">")
	w.WriteString("<tr>\n") // TITULO
	w.WriteString("<td color = \"grey\" colspan = '2'> REPORTE MBR </td> ")
	w.WriteString("</tr>\n") // FIN TITULO
	w.WriteString("<tr>\n")  // FILA 1
	w.WriteString("<td color = \"grey\">Size</td>\n")
	w.WriteString("<td color = \"grey\">\"" + strconv.Itoa(int(m.Tamanio)) + "\"</td>\n")
	w.WriteString("</tr>\n")
	w.WriteString("<tr>\n") // FILA 2
	w.WriteString("<td color = \"grey\">Signature</td>\n")
	w.WriteString("<td color = \"grey\">\"" + strconv.Itoa(int(m.DiskSignature)) + "\"</td>\n")
	w.WriteString("</tr>\n")
	w.WriteString("<tr>\n") // FILA 3
	w.WriteString("<td color = \"grey\">Fecha</td>\n")
	w.WriteString("<td color = \"grey\">" + string(m.Fecha[:]) + "</td>\n")
	w.WriteString("</tr>\n")
	/*		las n particiones */
	for x := 0; x < 4; x++ {
		status := (m.Particiones[x].Status)
		if status == 'y' {

			tipo := (m.Particiones[x].Tipo)
			//nombreParticion := m.Particiones[x].Nombre[:16]
			w.WriteString("<tr>\n") // TITULO
			w.WriteString("<td color = \"grey\" colspan = '2'> Particion" + strconv.Itoa(x) + "</td> ")
			w.WriteString("</tr>\n") // FIN TITULO
			//status
			w.WriteString("<tr>\n")
			w.WriteString("<td color = \"grey\">STATUS</td>\n")
			w.WriteString("<td color = \"grey\">" + string(status) + "</td>\n")
			w.WriteString("</tr>\n")

			/*// NOMBRE
			w.WriteString("<tr>\n")
			w.WriteString("<td color = \"grey\">NOMBRE</td>\n")
			w.WriteString("<td color = \"grey\">" + string(nombreParticion) + "</td>\n")
			w.WriteString("</tr>\n")*/

			// tipo
			w.WriteString("<tr>\n")
			w.WriteString("<td color = \"grey\">TIPO</td>\n")
			w.WriteString("<td color = \"grey\">" + strings.ToUpper(string(tipo)) + "</td>\n")
			w.WriteString("</tr>\n")
			// inicio
			w.WriteString("<tr>\n")
			w.WriteString("<td color = \"grey\">INICIO</td>\n")
			w.WriteString("<td color = \"grey\">\"" + strconv.Itoa(int(m.Particiones[x].Inicio)) + "\"</td>\n")
			w.WriteString("</tr>\n")
			// SIZE
			w.WriteString("<tr>\n")
			w.WriteString("<td color = \"grey\">SIZE</td>\n")
			w.WriteString("<td color = \"grey\">\"" + strconv.Itoa(int(m.Particiones[x].Size)) + "\"</td>\n")
			w.WriteString("</tr>\n")

			// FIT
			w.WriteString("<tr>\n")
			w.WriteString("<td color = \"grey\">FIT</td>\n")
			w.WriteString("<td color = \"grey\">" + (m.Particiones[x].getFitToString()) + "</td>\n")
			w.WriteString("</tr>\n")

		}
	}

	w.WriteString("</table>\n")
	w.WriteString(">];\n")
	w.WriteString("}\n")
	if errOr := w.Close(); errOr != nil {
		log.Fatal(errOr)
		return
	}
}

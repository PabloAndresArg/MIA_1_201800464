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
		Path_ = QuitarComillas(Path_)
		switch Name_ {
		case "mbr":
			grahpMBR(Id_vdlentraNumero_, Path_)
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

func separarRutaYnombreReporte(pathCompleto string) (string, string, string) { // ruta completa  , nombre , extension
	// siempre vendra /home/
	posFinRuta := 0
	ruta := ""
	nombre := ""
	extension := ""
	for x := len(pathCompleto) - 1; x >= 0; x-- {
		if pathCompleto[x] == '/' {
			posFinRuta = x
			break
		}
	}
	for i := posFinRuta + 1; i < len(pathCompleto); i++ {
		nombre += string(pathCompleto[i])
	}
	for k := 0; k <= posFinRuta; k++ { // debo incluir su /
		ruta += string(pathCompleto[k])
	}
	aux := ""
	for t := 0; t < len(nombre); t++ {
		if nombre[t] == '.' {
			aux = extension
			extension = ""
		} else {
			extension += string(nombre[t])
		}
	}
	nombre = aux
	return ruta, nombre, extension
}

func grahpMBR(id string, pathCompleto string) {
	rut, nom, ext := separarRutaYnombreReporte(pathCompleto)
	/*fmt.Println(rut)
	fmt.Println(nom)
	fmt.Println(ext)*/
	verificarRuta(rut) // la crea si no existe
	// NECESITO IR A ATRAER EL PATH , TENIENDO EN CUENTA QUE PUEDO BUSCAR EN MI LISTA id[2] me da la letra y ya tengo el disco que necesito
	var letraID = string(id[2])
	_disco_ := getDiscoMontadoPorLetraID(letraID)
	if _disco_.Letra == "NOENCONTRADO" { // EN TEORIA NUNCA ENTRARIA ACA
		println(color.Red + "ESE ID NO FUE ENCONTRADO DENTRO DEL DISCO" + color.Reset)
		return
	}

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

	crearTxt(mrbAuxiliar, rut+nom+".txt")
	generarImg(rut+nom, ext)

}
func generarImg(fuente string, extension string) {
	pos1 := "-T" + extension
	pos2 := fuente + ".txt"
	pos3 := fuente + "." + extension
	// dot -Tjpg /home/pablo/Escritorio/REP/ReporteMbr.txt -o /home/pablo/Escritorio/REP/ReporteMbr.jpg
	consola := exec.Command("dot", pos1, pos2, "-o", pos3)
	var out bytes.Buffer
	var stderr bytes.Buffer
	consola.Stdout = &out
	consola.Stderr = &stderr
	err := consola.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	println(color.Blue + "REPORTE GENERADO" + color.Reset)
}

func crearTxt(m TipoMbr, direccionDestino string) { // pasar tambien la ruta
	w, err := os.Create(direccionDestino)
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
			nombre := m.Particiones[x].getNameHowString()
			w.WriteString("<tr>\n") // TITULO
			w.WriteString("<td color = \"grey\" colspan = '2'> Particion" + strconv.Itoa(x) + "</td> ")
			w.WriteString("</tr>\n") // FIN TITULO
			//status
			w.WriteString("<tr>\n")
			w.WriteString("<td color = \"grey\">STATUS</td>\n")
			w.WriteString("<td color = \"grey\">" + string(status) + "</td>\n")
			w.WriteString("</tr>\n")

			// NOMBRE
			w.WriteString("<tr>\n")
			w.WriteString("<td color = \"grey\">NOMBRE</td>\n")
			w.WriteString("<td color = \"grey\">" + nombre + "</td>\n")
			w.WriteString("</tr>\n")

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

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
		case "disk":
			graphDisk(Id_vdlentraNumero_, Path_)
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
	verificarRuta(rut) // la crea si no existe
	// NECESITO IR A ATRAER EL PATH , TENIENDO EN CUENTA QUE PUEDO BUSCAR EN MI LISTA id[2] me da la letra y ya tengo el disco que necesito
	var letraID = string(id[2])
	_disco_ := getDiscoMontadoPorLetraID(letraID)
	if _disco_.Letra == "NOENCONTRADO" { // EN TEORIA NUNCA ENTRARIA ACA
		println(color.Red + "ESE ID NO FUE ENCONTRADO DENTRO DEL DISCO" + color.Reset)
		return
	}

	if _, err := os.Stat(_disco_.Path); !(os.IsNotExist(err)) {
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

		crearTxt(mrbAuxiliar, rut+nom+".txt", archivoDisco)
		generarImg(rut+nom, ext, rut)
	} else {
		fmt.Println("-----------------------")
		fmt.Println("EL DISCO YA NO EXISTE")
		fmt.Println("-----------------------")
	}
}
func generarImg(fuente string, extension string, direccionCarpeta string) {
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
	println(color.Blue + "-----------------")
	println(color.Blue + "REPORTE GENERADO")
	println(color.Blue + "En: " + direccionCarpeta)
	println(color.Blue + "-----------------" + color.Reset)

}

func crearTxt(m TipoMbr, direccionDestino string, archivoDisco *os.File) { // pasar tambien la ruta
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
			w.WriteString("<td color = \"grey\" colspan = '2'> Particion [" + strconv.Itoa(x) + "]</td> ")
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
		// ---------------------------------------------------------------------------- PARA LOS EBRS
		if m.Particiones[x].Status == 'y' && (m.Particiones[x].Tipo == 'E' || m.Particiones[x].Tipo == 'e') {
			w.WriteString("<tr>\n") // TITULO
			w.WriteString("<td bgcolor = \"#a1fc6a\" colspan = '2'> EBRS EN LA EXTENDIDA</td> ")
			w.WriteString("</tr>\n") // FIN TITULO
			archivoDisco.Seek(m.Particiones[x].Inicio, 0)
			ebrAux := Ebr{}
			tamanioEBR := binary.Size(ebrAux) //tamanio de lo que ire a traer
			ebr_en_bytes := leerBytePorByte(archivoDisco, tamanioEBR)
			buff := bytes.NewBuffer(ebr_en_bytes)               // lo convierto a buffer porque eso pedia la funcion
			err := binary.Read(buff, binary.BigEndian, &ebrAux) //ya tengo el original
			if err != nil {
				fmt.Println("error en lectura ebr ")
			}
			if ebrAux.Status == 'y' && ebrAux.Next == -1 {
				w.WriteString(ebrAux.getCadenaHTML())
			} else if ebrAux.Next == -1 {
				g := ("<tr>\n")
				g += ("<td bgcolor = \"#a1fc6a\">" + "SOLO EXISTE UN EBR VACIO APUNTANDO A -1" + "</td>\n")
				g += ("</tr>\n")
				w.WriteString(g)
			} else {
				w.WriteString(ebrAux.getCadenaHTML())
				for ebrAux.Next != -1 {
					archivoDisco.Seek(ebrAux.Next, 0)
					tamanioEBR := binary.Size(ebrAux)
					ebr_en_bytes := leerBytePorByte(archivoDisco, tamanioEBR)
					buff := bytes.NewBuffer(ebr_en_bytes)
					err = binary.Read(buff, binary.BigEndian, &ebrAux)
					w.WriteString(ebrAux.getCadenaHTML())
				}
			}
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

func graphDisk(id string, pathCompleto string) {
	rut, nom, ext := separarRutaYnombreReporte(pathCompleto)
	verificarRuta(rut) // la crea si no existe
	// NECESITO IR A ATRAER EL PATH , TENIENDO EN CUENTA QUE PUEDO BUSCAR EN MI LISTA id[2] me da la letra y ya tengo el disco que necesito
	var letraID = string(id[2])
	_disco_ := getDiscoMontadoPorLetraID(letraID)
	if _disco_.Letra == "NOENCONTRADO" { // EN TEORIA NUNCA ENTRARIA ACA
		println(color.Red + "ESE ID NO FUE ENCONTRADO DENTRO DEL DISCO" + color.Reset)
		return
	}
	if _, err := os.Stat(_disco_.Path); !(os.IsNotExist(err)) {
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

		crearTxtDisk(mrbAuxiliar, rut+nom+".txt", archivoDisco)
		generarImg(rut+nom, ext, rut)
	} else {
		fmt.Println("-----------------------")
		fmt.Println("EL DISCO YA NO EXISTE  ")
		fmt.Println("-----------------------")
	}

}

func crearTxtDisk(m TipoMbr, direccionDestino string, archivoDisco *os.File) {
	w, err := os.Create(direccionDestino)
	if err != nil {
		println(color.Red + "Error al crear el archivo" + color.Reset)
		return
	}
	w.WriteString("Digraph DiscoRep{\n")
	w.WriteString("tbl[\n")
	w.WriteString("shape = plaintext\n")
	w.WriteString("label =<\n")
	//<table border = '4' color = 'black' cellborder = '4' cellspacing = '4' bgcolor= "black">
	w.WriteString("<table border = '4' cellborder = '4' cellspacing = '4' bgcolor = \"black\">")
	//---------- MBR ----------------
	w.WriteString("<tr>\n\n\n\n")
	w.WriteString("<td height = \"100\" bgcolor = \"#11fc6a\">MBR</td>\n")

	//------------ PARTICIONES Y FRAGMENTACION--------------------------
	for x := 0; x < 4; x++ {
		status := (m.Particiones[x].Status)
		if status == 'y' && m.Particiones[x].Tipo == 'p' || m.Particiones[x].Tipo == 'P' {
			nombre := m.Particiones[x].getNameHowString()
			w.WriteString("<td height = \"100\" bgcolor = \"#11fc6a\">" + nombre + "</td>\n")
		} else if status == 'y' && m.Particiones[x].Tipo == 'e' || m.Particiones[x].Tipo == 'E' {
			nombre := m.Particiones[x].getNameHowString()
			// ACA CREO UNA TABLA , pero tengo que tener en cuenta la cantidad de ebrs para hacer un cols = cantidadEbrs * 2 +  bloques  espacio libre :'v
			encabezado := ""
			cuerpo := ""
			w.WriteString("<td bgcolor='black' height = '100'>") // INICIA LA COLUMNA  TD
			w.WriteString("\n\n\n")

			cuerpo += ("<tr>\n") //fila 2 ,  INICIAN LOS EBRS
			// con un for consigo todos los ebr *2 + ---- y tengo que tener un contador para poder generar el encabezado mas abajo con los fatos del colspan necesarios
			cuerpo += ("<td color = 'black' bgcolor='#01A9DB' height = '30'>EBR1</td>\n")
			cuerpo += ("<td color = 'black' bgcolor='#f2ff51' height = '30'>LOGICA 1</td>\n")
			cuerpo += ("</tr>\n")                                                      //FILA 2  FIN EBRS*/
			encabezado += ("<table color='blue' cellspacing='4' bgcolor = 'black'>\n") // NO MANDARLO A ESCRIBIR DE UNA SINO QUE GUARDAR TODO EN VARIABLES TEMPORALES Y LUEGO MANDARLAS A ESCRIBIR
			encabezado += ("<tr><td bgcolor='WHITE'  height = '50' colspan='2'>" + nombre + "</td></tr>\n")

			w.WriteString(encabezado)
			w.WriteString(cuerpo)
			w.WriteString("</table>\n")
			w.WriteString("\n")

			w.WriteString("\n\n\n</td>\n") // FIN DE LA COLUMNA TD*/

		} else if status == 'n' {
			w.WriteString("<td height = \"100\" bgcolor = \"##00FFFF\">" + "Espacio para Particion" + "</td>\n")
		}
		RangosPrincipales := m.getRangosParticiones("")
		if len(RangosPrincipales) != 0 && status == 'y' && x+1 != len(RangosPrincipales) {
			m.verFragmentacion(archivoDisco)
			fmt.Println(fmt.Sprint(RangosPrincipales[x].LimiteSuperior) + "-" + fmt.Sprint(RangosPrincipales[x+1].LimiteInferior))
			resulto := RangosPrincipales[x+1].LimiteInferior - RangosPrincipales[x].LimiteSuperior - 1
			fmt.Println(resulto)
			if resulto != 0 {
				w.WriteString("<td height = \"100\" bgcolor = \"#ff0f00\">" + "FREE " + fmt.Sprint(resulto) + " bytes" + "</td>\n")
			}

		}
	}
	w.WriteString("<td height = \"100\" bgcolor = \"#CEF6E3\">" + "FREE " + fmt.Sprint(m.getEspacioLibre()) + " bytes </td>\n")
	w.WriteString("\n\n\n\n</tr>\n")
	w.WriteString("</table>\n")
	w.WriteString(">];\n")
	w.WriteString("}\n")
	if errOr := w.Close(); errOr != nil {
		log.Fatal(errOr)
		return
	}

}

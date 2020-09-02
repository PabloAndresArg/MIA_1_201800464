package An

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strings"

	"github.com/TwinProduction/go-color"
)

// TipoMbr es un mbr: tabla de particiones y que tiene info del archivo , es lo primero que se guarda dentro de un disco(en este caso archivo por ser simulacion )
type TipoMbr struct {
	Tamanio       int64
	Fecha         [19]byte
	DiskSignature int64
	Particiones   [4]Particion // DE TIPO PRIMARIA
}

// Particion primaria, extendida o logica
type Particion struct {
	Status byte
	Fit    byte // son char de GOLANG
	Inicio int64
	Size   int64
	Nombre [16]byte
	Tipo   byte
}

//Ebr es un  EBR solo existen adentro de las particiones extendidas
type Ebr struct {
	Status byte
	Fit    byte // son char de GOLANG
	Inicio int64
	Size   int64
	Nombre [16]byte
	Next   int64 // APUNTA AL BYTE QUE SIGUE :v y es -1 cuando ya no le sigue nada
}

func (e Ebr) imprimirDatosEbr() {
	fmt.Printf("Status: %c\n", e.Status)
	fmt.Printf("Nombre: %s\n", e.Nombre)
	fmt.Printf("Fit: %c\n", e.Fit)
	fmt.Printf("inicio: %d\n", e.Inicio)
	fmt.Printf("size: %d\n", e.Size)
	fmt.Printf("ebr siguiente: %d\n", e.Next)
}
func (e *Ebr) deleteFastMenosElNext() {
	e.Status = 'n'
	for x := 0; x < len(e.Nombre); x++ {
		e.Nombre[x] = 0
	}
	e.Fit = ' '
	e.Size = 0
	e.Inicio = 0
}

// FILTRO 1
func (m TipoMbr) hayUnaParticionDisponible() bool { // retornar si si pudo agregar la particion o si no
	for x := 0; x < len(m.Particiones); x++ {
		if m.Particiones[x].Status == 'n' { // esta libre
			return true
		}
	}
	return false
}

// FILTRO 2
func (m TipoMbr) hayEspacioSuficiente(nuevoEspacio int64) bool { // retornar si si pudo agregar la particion o si no
	var tamanoOcupado int64 = int64(binary.Size(m)) // primero considero el tamanio del mbr
	for x := 0; x < len(m.Particiones); x++ {       // solo considero primarias
		tamanoOcupado += int64(m.Particiones[x].Size)
	}
	if m.Tamanio > (tamanoOcupado + nuevoEspacio) {
		fmt.Println("\n" + fmt.Sprint("tamanio del disco:", m.Tamanio) + fmt.Sprint(" tamanio Ocupado: ", tamanoOcupado) + fmt.Sprint(" Espacio de la nueva particion: ", nuevoEspacio))
		//TAMANIO DEL DISCO - (INICIO+SIZE) DE LA ULTIMA POSICION DE MI ARRAY DE PARTICIONES
		// disponibleDisco - disponible = FRAGMENTACION    y considerar que tengo un byte menos a la hora de escribir hasta el final
		if m.hayFragmentacion() { // TENER EN CUENTA QUE LA FRAGMENTACION SOLO APARECE ENTRE PARTICIONES , NO EN LOS EXTREMOS
			println(color.Yellow + "Cuidado este disco posee fragmentacion" + color.Reset)
		}

		return true
	}
	return false
}

// FILTRO 3
func (m TipoMbr) yaExisteUnaExtendida() bool { // retornar si si pudo agregar la particion o si no
	for x := 0; x < len(m.Particiones); x++ {

		if m.Particiones[x].Tipo == 'E' || m.Particiones[x].Tipo == 'e' { // esta libre

			return true
		}
	}
	return false
}

func (m TipoMbr) getExtendida() Particion { // retornar si si pudo agregar la particion o si no
	for x := 0; x < len(m.Particiones); x++ {

		if m.Particiones[x].Tipo == 'E' || m.Particiones[x].Tipo == 'e' { // esta libre

			return m.Particiones[x]
		}
	}
	erro := Particion{}
	return erro
}

func (m TipoMbr) crearParticion(fit string, size int64, nombre string, tipo byte) TipoMbr { // retornar si si pudo agregar la particion o si no
	for x := 0; x < len(m.Particiones); x++ { // NECESITO inicioParticion int64

		if m.Particiones[x].Status == 'n' { // ingresa en la libre y le cambia el status , PRIMER AJUSTE
			m.Particiones[x].Status = 'y'
			m.Particiones[x].Tipo = tipo
			m.Particiones[x].Fit = getFit(fit)
			m.Particiones[x].Size = size
			copy(m.Particiones[x].Nombre[:], nombre)
			m.Particiones[x].Inicio = m.getInicio(uint8(x))
			fmt.Println("---------------------------------")
			println(color.Yellow + "PARTICION PRIMARIA CREADA CON EXITO" + color.Reset)
			m.Particiones[x].imprimirDatosParticion()
			fmt.Println("---------------------------------")
			return m
			//break
		}
	}
	return m
}
func (m *TipoMbr) crearParticionExtendida(fit string, size int64, nombre string, tipo byte) uint8 { // retornar si si pudo agregar la particion o si no
	for x := 0; x < len(m.Particiones); x++ { // NECESITO inicioParticion int64

		if m.Particiones[x].Status == 'n' { // ingresa en la libre y le cambia el status , PRIMER AJUSTE
			m.Particiones[x].Status = 'y'
			m.Particiones[x].Tipo = tipo
			m.Particiones[x].Fit = getFit(fit)
			m.Particiones[x].Size = size
			copy(m.Particiones[x].Nombre[:], nombre)
			m.Particiones[x].Inicio = m.getInicio(uint8(x))
			fmt.Println("---------------------------------")
			println(color.Yellow + "PARTICION EXTENDIDA CREADA CON EXITO" + color.Reset)
			m.Particiones[x].imprimirDatosParticion()
			fmt.Println("---------------------------------")
			return uint8(x) // retorno la posicion para poder saber que pos del arreglo tiene y obtener el byte donde escribite un ebr
			//break
		}
	}
	return 0
}
func (p Particion) imprimirDatosParticion() {
	fmt.Printf("Nombre: %s\n", p.Nombre)
	fmt.Printf("Status: %c\n", p.Status)
	fmt.Printf("Tipo: %c\n", p.Tipo)
	fmt.Printf("Fit: %c\n", p.Fit)
	fmt.Printf("inicio: %d\n", p.Inicio)
	fmt.Printf("size: %d\n", p.Size)
}
func (p Particion) getFitToString() string {
	switch p.Fit {
	case 'w':
		return "W"
	case 'b':
		return "B"
	case 'f':
		return "F"
	default:
		return " "
	}
}
func (m TipoMbr) getInicio(indice uint8) int64 { // si donde esta disponible es la posicion 0 que pasa ?

	switch indice {
	case 0:
		return int64(binary.Size(m) + 1)
	default: // siempre abra una particion antes porque sino hubiera entrado en la libre
		return int64((m.Particiones[indice-1].Inicio + m.Particiones[indice-1].Size) + 1)
	}
}

func getTipoEnBytes(tipo string) byte {
	if len(tipo) == 0 { // solo por seguridad
		tipo = "p"
	} else {
		tipo = strings.ToLower(tipo) // E o L  , Precaucion con la l porque para ello se debe de tener una extendida
	}
	return byte(tipo[0]) // puedo retornarn p e l
}
func getFit(fit string) byte {
	if len(fit) == 0 { // por seguridad en realdad siempre deberia de traer un valor
		fit = "wf"
	}
	switch strings.ToLower(fit) {
	case "wf":
		return byte('w') // worst fit
	case "ff":
		return byte('f') // firts fit
	case "bf":
		return byte('b') // best fit
	default:
		return byte('w') // worst fit

	}
}
func (m TipoMbr) crearParticionLogica() {
	if m.yaExisteUnaExtendida() {
		for x := 0; x < len(m.Particiones); x++ {
			if m.Particiones[x].Tipo == 'E' || m.Particiones[x].Tipo == 'e' { // esta libre
				//
			}
		}
	}
}

func (m TipoMbr) hayFragmentacion() bool {
	// ARREGLODINAMICO
	aux := make([]Particion, 0, 1) // TENER EN CUENTA QUE CUANDO YA NO CABE ALGO LA CAPACIDAD SE DUPLCA
	for x := 0; x < len(m.Particiones); x++ {
		if m.Particiones[x].Status == 'y' { // en uso
			aux = append(aux, m.Particiones[x]) // INSERTA AL FINAL
		}
	}
	// ya tendres todas las particiones utilizadas en el aux de forma continua
	for b := 0; b < (len(aux) - 1); b++ {
		if (aux[b].Inicio + aux[b].Size + 1) != aux[b+1].Inicio {
			return true // si hay fragmentacion
		}
	}
	return false
}

func (m TipoMbr) imprimirDatosMBR() { // retornar si si pudo agregar la particion o si no
	println(color.Green + "***************** MBR ***********************" + color.Reset)
	fmt.Printf("\nFECHA: %s\nTamanio: %v\n", m.Fecha, m.Tamanio)
	fmt.Printf("Signature: %d\n", m.DiskSignature)
	for x := 0; x < 4; x++ {
		if m.Particiones[x].Status == 'y' { // activas
			println(color.Green + fmt.Sprint("---------------- Particion[", x) + "]----------------" + color.Reset)
			m.Particiones[x].imprimirDatosParticion()
			println(color.Green + "---------------------------------------------" + color.Reset)
		}
	}
	println(color.Green + "*********************************************" + color.Reset)
}

func (mont Montura) getImpresionMontura() {
	print(color.Green + "ID: " + mont.Id + " Path: " + mont.PathDisco + color.Reset)
	fmt.Printf(" Nombre: %s\n", mont.Nombre)
}

func (m TipoMbr) buscarExistenciaEnParticiones(nombreBuscar string) bool { // retornar si si pudo agregar la particion o si no
	var aux [16]byte
	copy(aux[:], nombreBuscar)
	for x := 0; x < len(m.Particiones); x++ {
		if m.Particiones[x].Status == 'y' && (m.Particiones[x].Tipo == 'p' || m.Particiones[x].Tipo == 'P') { // esta libre
			if string(m.Particiones[x].Nombre[:]) == string(aux[:]) {
				//fmt.Println(string(m.Particiones[x].Nombre[:]) + " == " + nombreBuscar)
				return true
			}
		} else if m.Particiones[x].Status == 'y' && (m.Particiones[x].Tipo == 'e' || m.Particiones[x].Tipo == 'E') {

			if string(m.Particiones[x].Nombre[:]) == string(aux[:]) {
				return true
			} else {

				// LESE HACER UN METODO PARA TRAER LOGICAS
			}
		}
	}
	// si adentro de las primarias no lo encontro busco en la extendida en las logicas , eso suena mas complejo
	return false
}

func (m *TipoMbr) eliminarFast(nombreBuscar string, archivoDisco *os.File) bool { // retorna si fue posible la ELIMINACION
	var aux [16]byte
	copy(aux[:], nombreBuscar)
	for x := 0; x < len(m.Particiones); x++ {
		if m.Particiones[x].Status == 'y' && (m.Particiones[x].Tipo == 'p' || m.Particiones[x].Tipo == 'P') { // esta libre
			if string(m.Particiones[x].Nombre[:]) == string(aux[:]) {
				m.Particiones[x].Status = 'n'
				for r := 0; r < len(m.Particiones[x].Nombre); r++ {
					m.Particiones[x].Nombre[r] = 0
				}
				m.Particiones[x].Inicio = 0
				m.Particiones[x].Size = 0
				return true
			}
		} else if m.Particiones[x].Status == 'y' && (m.Particiones[x].Tipo == 'E' || m.Particiones[x].Tipo == 'e') {

			if string(m.Particiones[x].Nombre[:]) == string(aux[:]) {
				m.Particiones[x].Status = 'n'
				for r := 0; r < len(m.Particiones[x].Nombre); r++ {
					m.Particiones[x].Nombre[r] = 0
				}
				m.Particiones[x].Inicio = 0
				m.Particiones[x].Size = 0
				return true
			} else {
				archivoDisco.Seek(m.Particiones[x].Inicio, 0)
				ebrAux := Ebr{}
				tamanioEBR := binary.Size(ebrAux) //tamanio de lo que ire a traer
				ebr_en_bytes := leerBytePorByte(archivoDisco, tamanioEBR)
				buff := bytes.NewBuffer(ebr_en_bytes)               // lo convierto a buffer porque eso pedia la funcion
				err := binary.Read(buff, binary.BigEndian, &ebrAux) //ya tengo el original
				if err != nil {
					fmt.Println("error en lectura ebr ")
				}
				if string(ebrAux.Nombre[:]) == string(aux[:]) { // el primer EBR NO PUEDE SER ELIMINADO
					println(color.Red + "EL PRIMER EBR NO SE PUEDE ELIMINAR , por tanto se marcara como que no la encontro" + color.Reset)
				} else if ebrAux.Next == -1 {
					continue
				} else {
					// toca recorrer en busca de la logica
					auxAnterior := Ebr{}
					for string(ebrAux.Nombre[:]) != string(aux[:]) && ebrAux.Next != -1 {
						auxAnterior = ebrAux
						// LEER EBR POR EBR

						archivoDisco.Seek(ebrAux.Next, 0)
						tamanioEBR := binary.Size(ebrAux) //tamanio de lo que ire a traer
						ebr_en_bytes := leerBytePorByte(archivoDisco, tamanioEBR)
						buff := bytes.NewBuffer(ebr_en_bytes)              // lo convierto a buffer porque eso pedia la funcion
						err = binary.Read(buff, binary.BigEndian, &ebrAux) //ya tengo el original
						fmt.Printf(color.Cyan+"EBR actual: %s\n", ebrAux.Nombre)
					}
					// LA PREGUNTA ES ¿EL QUE SALIO ES EL QUE ?
					println("SALIO EL EBR" + color.Reset)
					if ebrAux.Next == -1 && string(ebrAux.Nombre[:]) == string(aux[:]) { // si es el ultimo solo hago esto :v
						auxAnterior.Next = -1
						escribirUnEBR(archivoDisco, (auxAnterior.Inicio - int64(binary.Size(auxAnterior))), auxAnterior) // REFRESCTO SU NEXT , CON SOLO ESTO BASTA
						ebrAux.deleteFastMenosElNext()
						escribirUnEBR(archivoDisco, (ebrAux.Inicio - int64(binary.Size(ebrAux))), ebrAux) // LA PONGO EN INHABILITADA
						return true
					} else if string(ebrAux.Nombre[:]) == string(aux[:]) {
						auxAnterior.Next = ebrAux.Next
						escribirUnEBR(archivoDisco, (auxAnterior.Inicio - int64(binary.Size(auxAnterior))), auxAnterior) // REFRESCTO SU NEXT
						ebrAux.deleteFastMenosElNext()
						escribirUnEBR(archivoDisco, (ebrAux.Inicio - int64(binary.Size(ebrAux))), ebrAux) // LA PONGO EN INHABILITADA
						fmt.Println("ENLAZO A LAS DE EN MEDIO ")
						return true
					}

				}

			}

		}
	}
	// si adentro de las primarias no lo encontro busco en la extendida en las logicas , eso suena mas complejo
	return false
}

func (p Particion) getNameHowString() string {
	auxSalida := ""
	for i := 0; i < 16; i++ {
		if p.Nombre[i] != 0 {
			auxSalida += string(p.Nombre[i])
		}
	}
	return auxSalida
}

// GetParticionYposicion me da la particion y la posicion del array a la que pertenece
func (m TipoMbr) GetParticionYposicion(nombreBuscar string) (Particion, uint8) {
	var aux [16]byte
	copy(aux[:], nombreBuscar)
	for x := 0; x < len(m.Particiones); x++ {
		if m.Particiones[x].Status == 'y' && (m.Particiones[x].Tipo == 'p' || m.Particiones[x].Tipo == 'P') { // PRIMARIAS
			if string(m.Particiones[x].Nombre[:]) == string(aux[:]) {
				return m.Particiones[x], uint8(x)
			}
		} else if m.Particiones[x].Status == 'y' && (m.Particiones[x].Tipo == 'e' || m.Particiones[x].Tipo == 'E') {

			if string(m.Particiones[x].Nombre[:]) == string(aux[:]) {
				return m.Particiones[x], uint8(x)
			}
		}
	}
	nada := Particion{}
	return nada, uint8(0)
}

func (m *TipoMbr) eliminarFullLogica(nombreBuscar string, archivoDisco *os.File) bool { // retorna si fue posible la ELIMINACION
	var aux [16]byte
	copy(aux[:], nombreBuscar)
	for x := 0; x < len(m.Particiones); x++ {
		if m.Particiones[x].Status == 'y' && (m.Particiones[x].Tipo == 'E' || m.Particiones[x].Tipo == 'e') {
			archivoDisco.Seek(m.Particiones[x].Inicio, 0)
			ebrAux := Ebr{}
			tamanioEBR := binary.Size(ebrAux) //tamanio de lo que ire a traer
			ebr_en_bytes := leerBytePorByte(archivoDisco, tamanioEBR)
			buff := bytes.NewBuffer(ebr_en_bytes)               // lo convierto a buffer porque eso pedia la funcion
			err := binary.Read(buff, binary.BigEndian, &ebrAux) //ya tengo el original
			if err != nil {
				fmt.Println("error en lectura ebr ")
			}
			if string(ebrAux.Nombre[:]) == string(aux[:]) { // el primer EBR NO PUEDE SER ELIMINADO
				println(color.Red + "EL PRIMER EBR NO SE PUEDE ELIMINAR , por tanto se marcara como que no la encontro" + color.Reset)
			} else if ebrAux.Next == -1 {
				continue
			} else {
				// toca recorrer en busca de la logica
				auxAnterior := Ebr{}
				for string(ebrAux.Nombre[:]) != string(aux[:]) && ebrAux.Next != -1 {
					auxAnterior = ebrAux
					// LEER EBR POR EBR
					archivoDisco.Seek(ebrAux.Next, 0)
					tamanioEBR := binary.Size(ebrAux) //tamanio de lo que ire a traer
					ebr_en_bytes := leerBytePorByte(archivoDisco, tamanioEBR)
					buff := bytes.NewBuffer(ebr_en_bytes)              // lo convierto a buffer porque eso pedia la funcion
					err = binary.Read(buff, binary.BigEndian, &ebrAux) //ya tengo el original
					fmt.Printf(color.Cyan+"EBR actual: %s\n", ebrAux.Nombre)
				}
				// 															LA PREGUNTA ES ¿EL QUE SALIO ES EL QUE ?
				println("SALIO EL EBR" + color.Reset)
				if ebrAux.Next == -1 && string(ebrAux.Nombre[:]) == string(aux[:]) { // 										si es el ultimo solo hago esto :v
					auxAnterior.Next = -1
					escribirUnEBR(archivoDisco, (auxAnterior.Inicio - int64(binary.Size(auxAnterior))), auxAnterior) // 	REFRESCTO SU NEXT , CON SOLO ESTO BASTA
					inicio := ebrAux.Inicio
					fin := ebrAux.Size
					ebrAux.deleteFastMenosElNext()
					escribirUnEBR(archivoDisco, (ebrAux.Inicio - int64(binary.Size(ebrAux))), ebrAux) // 					LA PONGO EN INHABILITADA

					archivoDisco.Seek(inicio, 0)
					var ceros []byte
					for r := 0; r < int(fin); r++ {
						ceros = append(ceros, 0)
					}
					var nuevoEscritor bytes.Buffer
					binary.Write(&nuevoEscritor, binary.BigEndian, &ceros)
					escribirBinariamente(archivoDisco, nuevoEscritor.Bytes())
					println(color.Blue + "Particion eliminada" + color.Reset)
					return true
				} else if string(ebrAux.Nombre[:]) == string(aux[:]) {
					auxAnterior.Next = ebrAux.Next
					escribirUnEBR(archivoDisco, (auxAnterior.Inicio - int64(binary.Size(auxAnterior))), auxAnterior) // 	REFRESCTO SU NEXT
					inicio := ebrAux.Inicio
					fin := ebrAux.Size
					ebrAux.deleteFastMenosElNext()
					escribirUnEBR(archivoDisco, (ebrAux.Inicio - int64(binary.Size(ebrAux))), ebrAux) // LA PONGO EN INHABILITADA
					archivoDisco.Seek(inicio, 0)
					var ceros []byte
					for r := 0; r < int(fin); r++ {
						ceros = append(ceros, 0)
					}
					var nuevoEscritor bytes.Buffer
					binary.Write(&nuevoEscritor, binary.BigEndian, &ceros)
					escribirBinariamente(archivoDisco, nuevoEscritor.Bytes())
					println(color.Blue + "Particion eliminada" + color.Reset)
					return true
				}

			}

		}
	}
	return false
}

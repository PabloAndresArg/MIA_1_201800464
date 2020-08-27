package An

import (
	"encoding/binary"
	"fmt"
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

// Montura me sirve para el comando mount
type Montura struct {
	// path , nombre , id , y arreglo dinamico de particiones
	PathDisco string
	Nombre    [16]byte
	id        string
}

// Particion primaria, extendida o logica
type Particion struct {
	Status byte
	Fit    byte // son char de GOLANG
	Inicio int64
	Size   int64
	Nombre [16]byte
	Tipo   byte
	// PUEDO TENER UN VECTOR DINAMICO DE EXTENDED_B_R limitado solo por el tamaño de mi archivo
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

func (m TipoMbr) crearParticion(fit string, size int64, nombre string, tipo byte) TipoMbr { // retornar si si pudo agregar la particion o si no
	for x := 0; x < len(m.Particiones); x++ { // NECESITO inicioParticion int64

		if m.Particiones[x].Status == 'n' { // ingresa en la libre y le cambia el status , PRIMER AJUSTE
			m.Particiones[x].Status = 'y'
			m.Particiones[x].Tipo = tipo
			m.Particiones[x].Fit = getFit(fit)
			m.Particiones[x].Size = size
			copy(m.Particiones[x].Nombre[:], nombre)
			m.Particiones[x].Inicio = m.getInicio(uint8(x))
			m.Particiones[x].imprimirDatosParticion()
			println(color.Yellow + "PARTICION PRIMARIA CREADA CON EXITO" + color.Reset)
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
			m.Particiones[x].imprimirDatosParticion()
			println(color.Yellow + "PARTICION EXTENDIDA CREADA CON EXITO" + color.Reset)
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

func (mont Montura) imprimirMontura() {
	print(color.Green + "ID: " + mont.id + " Path: " + mont.PathDisco + color.Reset)
	fmt.Printf(" Nombre: %s\n", mont.Nombre)
}

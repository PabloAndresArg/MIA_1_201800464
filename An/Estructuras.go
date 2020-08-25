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
	Next   int64
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
		fmt.Println("\n" + fmt.Sprint("tamanio del disco: ", m.Tamanio) + fmt.Sprint(" tamanio antes:", tamanoOcupado) + fmt.Sprint(" espacio Adicionado: ", nuevoEspacio))
		return true // si cabe una particion de ese tamanio
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

func (m TipoMbr) crearParticion(fit string, size int64, nombre string, tipo byte) { // retornar si si pudo agregar la particion o si no
	for x := 0; x < len(m.Particiones); x++ { // NECESITO inicioParticion int64

		if m.Particiones[x].Status == 'n' { // ingresa en la libre y le cambia el status , PRIMER AJUSTE
			m.Particiones[x].Status = 'y'
			m.Particiones[x].Tipo = tipo
			m.Particiones[x].Fit = getFit(fit)
			m.Particiones[x].Size = size
			copy(m.Particiones[x].Nombre[:], nombre)
			// ahora lo mas dificil
			m.getDatosPrint(uint8(x))
			println(color.Green + "PARTICION Creada Con Exito" + color.Reset)
			break
		}
	}
}
func (m TipoMbr) getDatosPrint(indice uint8) {
	fmt.Printf("Status: %c\n", m.Particiones[indice].Status)
	fmt.Printf("Tipo: %c\n", m.Particiones[indice].Tipo)
	fmt.Printf("Fit: %c\n", m.Particiones[indice].Fit)
	fmt.Printf("inicio: %d\n", m.Particiones[indice].Inicio)
	fmt.Printf("size: %d\n", m.Particiones[indice].Size)
	fmt.Printf("Nombre: %s\n", m.Particiones[indice].Nombre)
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

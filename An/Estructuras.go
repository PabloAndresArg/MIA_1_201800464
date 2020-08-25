//An   ESTE ARCHIVO TIENE LOS STRUCTS
package An

import (
	"encoding/binary"
	"fmt"

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
		fmt.Println(fmt.Sprint("tamanio del disco: ", m.Tamanio) + fmt.Sprint(" tamanio antes:", tamanoOcupado) + fmt.Sprint(" espacio Adicionado: ", nuevoEspacio))
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

//fit string, inicioParticion int64, size int64, nombre string, tipo byte
func (m TipoMbr) crearParticion() { // retornar si si pudo agregar la particion o si no
	for x := 0; x < len(m.Particiones); x++ {

		if m.Particiones[x].Status == 'n' { // ingresa en la libre y le cambia el status

			m.Particiones[x].Status = 'y'
			println(color.Green + "PARTICION Creada Con Exito" + color.Reset)
			break
		}
	}
}

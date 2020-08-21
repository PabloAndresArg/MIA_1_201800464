//An   ESTE ARCHIVO TIENE LOS STRUCTS
package An

// TipoMbr es un mbr: tabla de particiones y que tiene info del archivo , es lo primero que se guarda dentro de un disco(en este caso archivo por ser simulacion )
type TipoMbr struct {
	Tamanio       int64
	Fecha         [11]byte
	DiskSignature int64
	Particiones   [4]Particion // DE TIPO PRIMARIA
}

// Particion primaria, extendida o logica
type Particion struct {
	Status byte
	Fit    byte // son char de GOLANG
	Inicio int64
	Size   byte
	Nombre [16]byte
	Tipo   byte
	// PUEDO TENER UN VECTOR DINAMICO DE EXTENDED_B_R limitado solo por el tamaño de mi archivo
}

//Ebr es un  EBR solo existen adentro de las particiones extendidas
type Ebr struct {
	Status byte
	Fit    byte // son char de GOLANG
	Inicio int64
	Size   byte
	Nombre [16]byte
	Next   int64
}

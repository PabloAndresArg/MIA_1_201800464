package An

//Inodo es para los i-nodos :v
type Inodo struct {
	NumeroInodo            int64
	SizeArchivo            int64
	NumeroBloquesAsignados int64
	ArrayAptBloques        [4]int64
	AptIndirecto           int64
	IDProper               [16]byte
}

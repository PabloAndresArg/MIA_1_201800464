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

func (i *Inodo) crearPrimerInodo() {
	i.NumeroInodo = 1
	i.SizeArchivo = 34
	i.NumeroBloquesAsignados = 2
	i.AptIndirecto = 0
	copy(i.IDProper[:], "root")
}

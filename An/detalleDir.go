package An

import "time"

//DetalleDir es el detalle de directorio
type DetalleDir struct {
	AptDetalleDir int64 // APUNTA HACIA OTRA
	ArrayFiles    [5]infoDetalleDir
}

type infoDetalleDir struct {
	NombreArchivo     [20]byte
	FechaCreacion     [19]byte
	FechaModificacion [19]byte
	AptInodo          int64
}

func (dDir *DetalleDir) llenarDatosUsertxt(apuntadorInodo int64) {
	fechaModificacin := time.Now()
	copy(dDir.ArrayFiles[0].FechaModificacion[:], fechaModificacin.String())
	copy(dDir.ArrayFiles[0].NombreArchivo[:], "users.txt")
	dDir.ArrayFiles[0].AptInodo = apuntadorInodo
}

package An

//DetalleDir es el detalle de directorio
type DetalleDir struct {
	AptDetalleDir int64 // APUNTA HACIA OTRA
	ArrayFiles    [5]infoDetalleDir
}

type infoDetalleDir struct {
	NombreArchivo     [16]byte
	FechaCreacion     [19]byte
	FechaModificacion [19]byte
	AptInodo          int64
}

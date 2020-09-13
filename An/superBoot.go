package An

//Super BOOT
type SuperB struct {
	SbNombre          [16]byte
	SbAVDcount        int64
	SbDetalleDirCount int64
	SbInodosCount     int64
	SbBloquesCount    int64
	// LOS FREE
	SbAVDfree        int64
	SbDetalleDirFree int64
	SbInodosFree     int64
	SbBloquesFree    int64
	// datos
	SbFechaCreacion      [19]byte
	SbFechaUltimoMontaje [19]byte
	SbMontajesCount      int64 // me da duda
	//ARBOL DE DIRECTORIO
	AptBitMapAVD int64
	AptAVD       int64
	//DETALLE DIRECTORIO
	AptBitMapDetalleDir int64
	AptDetalleDir       int64
	//tabla i-nodos
	AptBitMapInodos      int64
	AptTablaInicioInodos int64
	//bloques
	AptBitMapBloques int64
	AptInicioBloques int64
	//LOG o BITACORA
	AptLog int64
	// SIZES
	SizeAVD        int64
	SizeDetalleDir int64
	SizeInodo      int64
	SizeBloque     int64
	// PRIMEROS LIBRES
	FirstLibreAVD        int64
	FirstLibreDetalleDir int64
	FirstLibreTablaInodo int64
	FirstLibreBloque     int64
	SbMagicNum           int64 // MI NUMERO DE CARNET 201800464
}

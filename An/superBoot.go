package An

//Super BOOT
type Super struct {
	SbNombre            [16]byte
	SbAVDcount          int64
	SbInodosCount       int64
	SbBloquesCount      int64
	SbAVDfree           int64
	SbInodosFree        int64
	SbBloquesFree       int64
	SbDate              [19]byte
	SbDateUltimoMontaje [19]byte
	SbMontajesCount     int64
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

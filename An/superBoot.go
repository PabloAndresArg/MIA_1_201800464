package An

import (
	"fmt"

	"github.com/TwinProduction/go-color"
)

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

func (sb SuperB) imprimirDatosBoot() {
	print(color.Yellow)
	fmt.Println("\n*********************************** SUPER BOOT **************************************")
	fmt.Printf("Nombre: %s\n", sb.SbNombre)
	fmt.Printf("numero AVD: %d\n", sb.SbAVDcount)
	fmt.Printf("numero Detalles: %d\n", sb.SbDetalleDirCount)
	fmt.Printf("numero I-nodos: %d\n", sb.SbInodosCount)
	fmt.Printf("numero bloques: %d\n", sb.SbBloquesCount)
	fmt.Printf("AVD Disponibles: %d\n", sb.SbAVDfree)
	fmt.Printf("Detalles Disponibles: %d\n", sb.SbDetalleDirFree)
	fmt.Printf("I-nodos Disponibles: %d\n", sb.SbInodosFree)
	fmt.Printf("Bloques Disponibles: %d\n", sb.SbBloquesFree)
	fmt.Printf("Fecha Creacion: %s\n", sb.SbFechaCreacion)
	fmt.Printf("Fecha UltimoMonjate: %s\n", sb.SbFechaUltimoMontaje)
	fmt.Printf("Cantidad de Montajes: %d\n", sb.SbMontajesCount)
	fmt.Printf("Puntero Bitmap AVD: %d\n", sb.AptBitMapAVD)
	fmt.Printf("Puntero AVD: %d\n", sb.AptAVD)
	fmt.Printf("Puntero Bitmap DetalleDir: %d\n", sb.AptBitMapDetalleDir)
	fmt.Printf("Puntero DetalleDir: %d\n", sb.AptDetalleDir)
	fmt.Printf("Puntero Bitmap I-nodos: %d\n", sb.AptBitMapInodos)
	fmt.Printf("Puntero Inicio I-nodos: %d\n", sb.AptTablaInicioInodos)
	fmt.Printf("Puntero Bitmap Bloques: %d\n", sb.AptBitMapInodos)
	fmt.Printf("Puntero Inicio Bloques: %d\n", sb.AptTablaInicioInodos)
	fmt.Printf("Puntero Inicio Bitacora: %d\n", sb.AptLog)
	fmt.Printf("Size AVD: %d\n", sb.SizeAVD)
	fmt.Printf("Size DetalleDir: %d\n", sb.SizeDetalleDir)
	fmt.Printf("Size I-nodo: %d\n", sb.SizeInodo)
	fmt.Printf("Size Bloque: %d\n", sb.SizeBloque)
	/*
		// PRIMEROS LIBRES
		FirstLibreAVD        int64
		FirstLibreDetalleDir int64
		FirstLibreTablaInodo int64
		FirstLibreBloque     int64
	*/
	fmt.Printf("SbMagic: %d\n", sb.SbMagicNum)
	fmt.Println("*************************************************************************************")
	print(color.Reset)
}

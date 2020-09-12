package An

// ES EL SUPER BOOT
type Super struct {
	SbNombre                           [16]byte
	SbAVDcount                         int64
	SbInodosCount                      int64
	SbBloquesCount                     int64
	SbAVDfree                          int64
	SbInodosFree                       int64
	SbBloquesFree                      int64
	SbDate                             [19]byte
	SbDateUltimoMontaje                [19]byte
	SbMontajesCount                    int64
	SbApuntadorBitMapAVD               int64
	SbApuntadorAVD                     int64
	SbApuntadorBitMapDetalleDirectorio int64
	SbApuntadorDetalleDirectorio       int64
}

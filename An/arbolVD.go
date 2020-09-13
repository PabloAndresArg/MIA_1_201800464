package An

import "time"

type AVD struct {
	FechaCreacion       [19]byte
	NombreDirectorio    [16]byte
	SubDirectorios      [6]int64
	ApuntadorDetalleDir int64
	ApuntadorAVDextra   int64
	Proper              [16]byte
}

func (avd *AVD) crearRoot() {
	fechaActual := time.Now()
	copy(avd.FechaCreacion[:], fechaActual.String())
	copy(avd.NombreDirectorio[:], "/")
	copy(avd.Proper[:], "root")
}

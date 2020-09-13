package An

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/TwinProduction/go-color"
)

func metodoMKFS(id string, tipo string, add string, unit string) {
	if len(id) != 0 {
		if validacionQueEsteMontada(string(id[2]), id) {
			dis, part := getDiscoYparticionDelMount(string(id[2]), id)
			archivoDisco, err := os.OpenFile(QuitarComillas(dis.Path), os.O_RDWR, 0644) // TIEENE PERMISOS DE ESCRITURA Y DE LECTURA PERRO :V
			defer archivoDisco.Close()
			if err != nil {
				log.Fatal(err)
			}

			superBoot := SuperB{}
			fmt.Println("NOMBRE: " + dis.getOnlyName())
			fmt.Println(fmt.Sprint(getNumedds(part.Parti.Size)))
			if part.Tipo == 'l' { // LOGICA
				numeroEdds := getNumedds(part.PartiLogica.Size)
				superBoot.SbMagicNum = 201800464
				copy(superBoot.SbNombre[:], dis.getOnlyName())
				superBoot.SbAVDcount = numeroEdds
				superBoot.SbDetalleDirCount = numeroEdds
				superBoot.SbInodosCount = 4 * numeroEdds
				superBoot.SbBloquesCount = 5 * (4 * numeroEdds) // por cada avd hay 4 i-nodos y por cada i-nodo hay 5 bloques
				// los free
				superBoot.SbAVDfree = numeroEdds
				superBoot.SbDetalleDirFree = numeroEdds
				superBoot.SbInodosFree = 4 * numeroEdds
				superBoot.SbBloquesFree = 5 * (4 * numeroEdds)
				// ahora los size
				superBoot.SizeAVD = int64(binary.Size(AVD{}))
				superBoot.SizeDetalleDir = int64(binary.Size(DetalleDir{}))
				superBoot.SizeInodo = int64(binary.Size(Inodo{}))
				superBoot.SizeBloque = int64(binary.Size(Bloque{}))
				// ahora apuntadores

				// fechas
				FechaFormatoTime := time.Now()
				copy(superBoot.SbFechaUltimoMontaje[:], FechaFormatoTime.String())
				copy(superBoot.SbFechaCreacion[:], FechaFormatoTime.String())
				// APUNTADORES AVD
				superBoot.AptBitMapAVD = part.PartiLogica.Inicio + int64(binary.Size(superBoot)) + 1
				superBoot.AptAVD = superBoot.AptBitMapAVD + superBoot.SbAVDcount + 1
				// APUNTADORES DETALLE
				superBoot.AptBitMapDetalleDir = superBoot.AptAVD + int64(superBoot.SbAVDcount*superBoot.SizeAVD) + 1
				superBoot.AptDetalleDir = superBoot.AptBitMapDetalleDir + superBoot.SbDetalleDirCount + 1
				// apuntadores i-nodos
				superBoot.AptBitMapInodos = superBoot.AptDetalleDir + int64(superBoot.SbDetalleDirCount*superBoot.SizeDetalleDir) + 1
				superBoot.AptTablaInicioInodos = superBoot.AptBitMapInodos + superBoot.SbInodosCount + 1
				// apuntadores bloques
				superBoot.AptBitMapBloques = superBoot.AptTablaInicioInodos + int64(superBoot.SbInodosCount*superBoot.SizeInodo) + 1
				superBoot.AptInicioBloques = superBoot.AptBitMapBloques + superBoot.SbBloquesCount + 1
				// bitacora
				superBoot.AptLog = superBoot.AptInicioBloques + int64(superBoot.SbBloquesCount*superBoot.SizeBloque) + 1
				// AHORA LOS FREE PRIMER BYTE
				superBoot.FirstLibreAVD = superBoot.AptBitMapAVD
				superBoot.FirstLibreDetalleDir = superBoot.AptBitMapDetalleDir
				superBoot.FirstLibreTablaInodo = superBoot.AptBitMapInodos
				superBoot.FirstLibreBloque = superBoot.AptBitMapBloques
				// FIN
				archivoDisco.Seek(part.PartiLogica.Inicio, 0)
				var escritor bytes.Buffer
				binary.Write(&escritor, binary.BigEndian, &superBoot)
				escribirBinariamente(archivoDisco, escritor.Bytes())
			} else { // PRIMARIA
				numeroEdds := getNumedds(part.Parti.Size)
				superBoot.SbMagicNum = 201800464
				copy(superBoot.SbNombre[:], dis.getOnlyName())
				superBoot.SbAVDcount = numeroEdds
				superBoot.SbDetalleDirCount = numeroEdds
				superBoot.SbInodosCount = 4 * numeroEdds
				superBoot.SbBloquesCount = 5 * (4 * numeroEdds) // por cada avd hay 4 i-nodos y por cada i-nodo hay 5 bloques
				// los free
				superBoot.SbAVDfree = numeroEdds
				superBoot.SbDetalleDirFree = numeroEdds
				superBoot.SbInodosFree = 4 * numeroEdds
				superBoot.SbBloquesFree = 5 * (4 * numeroEdds)
				// ahora los size
				superBoot.SizeAVD = int64(binary.Size(AVD{}))
				superBoot.SizeDetalleDir = int64(binary.Size(DetalleDir{}))
				superBoot.SizeInodo = int64(binary.Size(Inodo{}))
				superBoot.SizeBloque = int64(binary.Size(Bloque{}))
				// ahora apuntadores

				// fechas
				FechaFormatoTime := time.Now()
				copy(superBoot.SbFechaUltimoMontaje[:], FechaFormatoTime.String())
				copy(superBoot.SbFechaCreacion[:], FechaFormatoTime.String())
				archivoDisco.Seek(part.Parti.Inicio, 0)
				var escritor bytes.Buffer
				binary.Write(&escritor, binary.BigEndian, &superBoot)
				escribirBinariamente(archivoDisco, escritor.Bytes())
			}

		}
	} else {
		println(color.Red + "-----------------------------------" + color.Reset)
		println(color.Red + "NO PUSO UN PARAMETRO OBLIGATORIO :)" + color.Reset)
		println(color.Red + "-----------------------------------" + color.Reset)
	}
	limpiarVariablesMKFS()
}

func limpiarVariablesMKFS() {
	Id_vdlentraNumero_ = ""
	type_ = "full"
	add_ = ""
	Unit_k_ = "k"
}

func getNumedds(tamanioParticion int64) int64 {
	cantidadEdds := int64(0)
	superbloque := SuperB{}
	arbolVD := AVD{}
	detalle := DetalleDir{}
	bloque := Bloque{}
	bita := Bitacora{}
	inodo := Inodo{}
	tamanioSuperBloque := int64(binary.Size(superbloque))
	tamanioArbolVD := int64(binary.Size(arbolVD))
	tamanioDetalleDir := int64(binary.Size(detalle))
	tamanioBloque := int64(binary.Size(bloque))
	tamanioInodo := int64(binary.Size(inodo))
	tamanioBita := int64(binary.Size(bita))
	cantidadEdds = (tamanioParticion - (2 * tamanioSuperBloque)) / (27 + tamanioArbolVD + tamanioDetalleDir + (5*tamanioInodo + (20 * tamanioBloque) + tamanioBita))
	return cantidadEdds
}

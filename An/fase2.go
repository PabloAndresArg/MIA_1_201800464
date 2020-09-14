package An

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
				fmt.Println("EL DISCO YA NO EXISTE , O EL PATH FUE INCORRECTO")
				return
			}
			superBoot := SuperB{}

			fmt.Println("NOMBRE: " + dis.getOnlyName())
			if part.Tipo == 'l' { // LOGICA

				numeroEdds := getNumedds(part.PartiLogica.Size)
				superBoot.SbMagicNum = 201800464
				copy(superBoot.SbNombre[:], dis.getOnlyName())
				superBoot.SbAVDcount = numeroEdds
				superBoot.SbDetalleDirCount = numeroEdds
				superBoot.SbInodosCount = 4 * numeroEdds
				superBoot.SbBloquesCount = 5 * (4 * numeroEdds) // por cada avd hay 4 i-nodos y por cada i-nodo hay 5 bloques
				// los free
				superBoot.SbAVDfree = numeroEdds - 1           // POR EL ROOT
				superBoot.SbDetalleDirFree = numeroEdds - 1    // por el dir de primero
				superBoot.SbInodosFree = 4*numeroEdds - 1      // por el primer i-nodo que se crea
				superBoot.SbBloquesFree = 5*(4*numeroEdds) - 2 // porque necesito dos para el toor
				// ahora los size
				superBoot.SizeAVD = int64(binary.Size(AVD{}))
				superBoot.SizeDetalleDir = int64(binary.Size(DetalleDir{}))
				superBoot.SizeInodo = int64(binary.Size(Inodo{}))
				superBoot.SizeBloque = int64(binary.Size(Bloque{}))
				// ahora apuntadores

				// fechas
				FechaFormatoTime := time.Now()
				copy(superBoot.SbFechaUltimoMontaje[:], FechaFormatoTime.String())
				if superBoot.SbMontajesCount == 0 {
					copy(superBoot.SbFechaCreacion[:], FechaFormatoTime.String())
				}
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
				superBoot.FirstLibreAVD = superBoot.AptBitMapAVD + int64(binary.Size(AVD{})) + 1
				superBoot.FirstLibreDetalleDir = superBoot.AptBitMapDetalleDir + int64(binary.Size(DetalleDir{})) + 1
				superBoot.FirstLibreTablaInodo = superBoot.AptBitMapInodos + int64(binary.Size(Inodo{})) + 1
				superBoot.FirstLibreBloque = superBoot.AptBitMapBloques + 2*int64(binary.Size(Inodo{})) + 2
				// FIN
				superBoot.SbMontajesCount++
				// ESCRIBIENDO EL SUPER BOOT :v
				archivoDisco.Seek(part.PartiLogica.Inicio, 0)
				var escritor bytes.Buffer
				binary.Write(&escritor, binary.BigEndian, &superBoot)
				escribirBinariamente(archivoDisco, escritor.Bytes())
				//***************************************************************************************************************************************************************************************
				//************************************************************************ ESCRIBIENDO EL BITMAP DEL ARBOL DE DIRECTORIO
				var bitMapAVD []byte
				for r := 0; r < int(superBoot.SbAVDcount); r++ {
					if r == 0 {
						bitMapAVD = append(bitMapAVD, 1)
					} else {
						bitMapAVD = append(bitMapAVD, 0)
					}
				}
				escribirBitMap(archivoDisco, superBoot.AptBitMapAVD, bitMapAVD)
				//************************************************************************ ESCRIBIENDO EL ROOT - ARBOL DE DIRECTORIO
				avd := AVD{}
				avd.crearRoot()
				avd.ApuntadorDetalleDir = superBoot.AptDetalleDir
				escribirUnAVD(archivoDisco, superBoot.AptAVD, avd)

				//*********************************************************************** ESCRIBIENDO  BIT MAP DETALLE DIR
				var bitMapDetalleDir []byte
				for r := 0; r < int(superBoot.SbDetalleDirCount); r++ {
					if r == 0 {
						bitMapDetalleDir = append(bitMapDetalleDir, 1)
					} else {
						bitMapDetalleDir = append(bitMapDetalleDir, 0)
					}
				}
				escribirBitMap(archivoDisco, superBoot.AptBitMapDetalleDir, bitMapDetalleDir)
				//*********************************************************************** ESCRIBIENDO UN DETALLE DIR
				detalle := DetalleDir{}

				fcreacion := time.Now()
				copy(detalle.ArrayFiles[0].FechaCreacion[:], fcreacion.String())

				detalle.llenarDatosUsertxt(superBoot.AptTablaInicioInodos)
				escribirDetalleDir(archivoDisco, superBoot.AptDetalleDir, detalle)
				//*********************************************************************** BIT MAP TABLA I-NODO
				var bitMapTablaInodo []byte
				for r := 0; r < int(superBoot.SbInodosCount); r++ {
					if r == 0 {
						bitMapTablaInodo = append(bitMapTablaInodo, 1)
					} else {
						bitMapTablaInodo = append(bitMapTablaInodo, 0)
					}
				}
				escribirBitMap(archivoDisco, superBoot.AptBitMapInodos, bitMapTablaInodo)
				//********************************************************************** ESCRIBIENOD EL PRIMER I-NODO

				iNode := Inodo{}
				iNode.crearPrimerInodo()
				escribirUnInodo(archivoDisco, superBoot.AptTablaInicioInodos, iNode)
				//*********************************************************************** BIT MAP BLOQUES
				var bitMapBloques []byte
				for r := 0; r < int(superBoot.SbBloquesCount); r++ {
					if r == 0 || r == 1 {
						bitMapBloques = append(bitMapBloques, 1)
					} else {
						bitMapBloques = append(bitMapBloques, 0)
					}
				}
				escribirBitMap(archivoDisco, superBoot.AptBitMapBloques, bitMapBloques)
				//********************************************************************** ESCRIBIENDO 2 BLOQUES :'V
				/*

					1, G, root\n
					1, U, root, root , 201800464\n

				*/
				bloque1 := Bloque{}
				copy(bloque1.DBdata[:], "1,G,root\\n1,U,root,root,2")
				bloque2 := Bloque{}
				copy(bloque2.DBdata[:], "01800464")
				escribirUnBloque(archivoDisco, superBoot.AptInicioBloques, bloque1)
				escribirUnBloque(archivoDisco, superBoot.AptInicioBloques+1+int64(binary.Size(bloque2)), bloque2)
				iNode.ArrayAptBloques[0] = superBoot.AptInicioBloques
				iNode.ArrayAptBloques[1] = superBoot.AptInicioBloques + 1 + int64(binary.Size(bloque2))
				escribirUnInodo(archivoDisco, superBoot.AptTablaInicioInodos, iNode) // PODRIA DEJAR SOLO ESTE PARA "NO" ESCRIBIR DOS VECES EL I-NODO
				//********************************************************************** ESCRIBIENDO LA BITACORA O LOG
				log1 := Bitacora{}
				copy(log1.TipoOpe[:], "mkdir")
				copy(log1.Nombre[:], "root")
				copy(log1.Tipo[:], "Directorio")
				copy(log1.fecha[:], time.Now().String())
				log1.Contenido = bloque1.DBdata
				escribirBitacora(archivoDisco, superBoot.AptLog, log1)
				log2 := Bitacora{}
				copy(log2.TipoOpe[:], "mkfile")
				copy(log2.Nombre[:], "users.txt")
				copy(log2.Tipo[:], "Archivo")
				copy(log2.fecha[:], time.Now().String())
				log2.Contenido = bloque2.DBdata
				escribirBitacora(archivoDisco, superBoot.AptLog+1+int64(binary.Size(log1)), log2)
				escribirSB(archivoDisco, part.PartiLogica.Inicio, superBoot)
				superBoot.imprimirDatosBoot()
				println(color.Green + "-----------------------------------" + color.Reset)
				println(color.Green + "FORMETO DE UNA PARTICION REALIZADO " + color.Reset)
				println(color.Green + "-----------------------------------" + color.Reset)
			} else { // PRIMARIA
				numeroEdds := getNumedds(part.Parti.Size)
				superBoot.SbMagicNum = 201800464
				copy(superBoot.SbNombre[:], dis.getOnlyName())
				superBoot.SbAVDcount = numeroEdds
				superBoot.SbDetalleDirCount = numeroEdds
				superBoot.SbInodosCount = 4 * numeroEdds
				superBoot.SbBloquesCount = 5 * (4 * numeroEdds) // por cada avd hay 4 i-nodos y por cada i-nodo hay 5 bloques
				// los free
				superBoot.SbAVDfree = numeroEdds - 1           // POR EL ROOT
				superBoot.SbDetalleDirFree = numeroEdds - 1    // por el dir de primero
				superBoot.SbInodosFree = 4*numeroEdds - 1      // por el primer i-nodo que se crea
				superBoot.SbBloquesFree = 5*(4*numeroEdds) - 2 // porque necesito dos para el toor
				// ahora los size
				superBoot.SizeAVD = int64(binary.Size(AVD{}))
				superBoot.SizeDetalleDir = int64(binary.Size(DetalleDir{}))
				superBoot.SizeInodo = int64(binary.Size(Inodo{}))
				superBoot.SizeBloque = int64(binary.Size(Bloque{}))
				// ahora apuntadores
				// fechas
				FechaFormatoTime := time.Now()
				copy(superBoot.SbFechaUltimoMontaje[:], FechaFormatoTime.String())
				if superBoot.SbMontajesCount == 0 {
					copy(superBoot.SbFechaCreacion[:], FechaFormatoTime.String())
				}
				// APUNTADORES AVD
				superBoot.AptBitMapAVD = part.Parti.Inicio + int64(binary.Size(superBoot)) + 1
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
				superBoot.FirstLibreAVD = superBoot.AptBitMapAVD + int64(binary.Size(AVD{})) + 1
				superBoot.FirstLibreDetalleDir = superBoot.AptBitMapDetalleDir + int64(binary.Size(DetalleDir{})) + 1
				superBoot.FirstLibreTablaInodo = superBoot.AptBitMapInodos + int64(binary.Size(Inodo{})) + 1
				superBoot.FirstLibreBloque = superBoot.AptBitMapBloques + 2*int64(binary.Size(Inodo{})) + 2
				// FIN
				superBoot.SbMontajesCount++

				// fechas
				fech := time.Now()
				copy(superBoot.SbFechaUltimoMontaje[:], fech.String())
				copy(superBoot.SbFechaCreacion[:], fech.String())
				// ESCRIBIENDO EL SUPER BOOT HASTA ABAJO
				//************************************************************************ ESCRIBIENDO EL BITMAP DEL ARBOL DE DIRECTORIO
				var bitMapAVD []byte
				for r := 0; r < int(superBoot.SbAVDcount); r++ {
					if r == 0 {
						bitMapAVD = append(bitMapAVD, 1)
					} else {
						bitMapAVD = append(bitMapAVD, 0)
					}
				}
				escribirBitMap(archivoDisco, superBoot.AptBitMapAVD, bitMapAVD)
				//************************************************************************ ESCRIBIENDO EL ROOT - ARBOL DE DIRECTORIO
				avd := AVD{}
				avd.crearRoot()
				avd.ApuntadorDetalleDir = superBoot.AptDetalleDir
				escribirUnAVD(archivoDisco, superBoot.AptAVD, avd)

				//*********************************************************************** ESCRIBIENDO  BIT MAP DETALLE DIR
				var bitMapDetalleDir []byte
				for r := 0; r < int(superBoot.SbDetalleDirCount); r++ {
					if r == 0 {
						bitMapDetalleDir = append(bitMapDetalleDir, 1)
					} else {
						bitMapDetalleDir = append(bitMapDetalleDir, 0)
					}
				}
				escribirBitMap(archivoDisco, superBoot.AptBitMapDetalleDir, bitMapDetalleDir)
				//*********************************************************************** ESCRIBIENDO UN DETALLE DIR
				detalle := DetalleDir{}

				fcreacion := time.Now()
				copy(detalle.ArrayFiles[0].FechaCreacion[:], fcreacion.String())

				detalle.llenarDatosUsertxt(superBoot.AptTablaInicioInodos)
				escribirDetalleDir(archivoDisco, superBoot.AptDetalleDir, detalle)
				//*********************************************************************** BIT MAP TABLA I-NODO
				var bitMapTablaInodo []byte
				for r := 0; r < int(superBoot.SbInodosCount); r++ {
					if r == 0 {
						bitMapTablaInodo = append(bitMapTablaInodo, 1)
					} else {
						bitMapTablaInodo = append(bitMapTablaInodo, 0)
					}
				}
				escribirBitMap(archivoDisco, superBoot.AptBitMapInodos, bitMapTablaInodo)
				//********************************************************************** ESCRIBIENOD EL PRIMER I-NODO

				iNode := Inodo{}
				iNode.crearPrimerInodo()
				escribirUnInodo(archivoDisco, superBoot.AptTablaInicioInodos, iNode)
				//*********************************************************************** BIT MAP BLOQUES
				var bitMapBloques []byte
				for r := 0; r < int(superBoot.SbBloquesCount); r++ {
					if r == 0 || r == 1 {
						bitMapBloques = append(bitMapBloques, 1)
					} else {
						bitMapBloques = append(bitMapBloques, 0)
					}
				}
				escribirBitMap(archivoDisco, superBoot.AptBitMapBloques, bitMapBloques)
				//********************************************************************** ESCRIBIENDO 2 BLOQUES :'V
				/*

					1, G, root\n
					1, U, root, root , 201800464\n

				*/
				bloque1 := Bloque{}
				copy(bloque1.DBdata[:], "1,G,root\\n1,U,root,root,2")
				bloque2 := Bloque{}
				copy(bloque2.DBdata[:], "01800464")
				escribirUnBloque(archivoDisco, superBoot.AptInicioBloques, bloque1)
				escribirUnBloque(archivoDisco, superBoot.AptInicioBloques+1+int64(binary.Size(bloque2)), bloque2)
				iNode.ArrayAptBloques[0] = superBoot.AptInicioBloques
				iNode.ArrayAptBloques[1] = superBoot.AptInicioBloques + 1 + int64(binary.Size(bloque2))
				escribirUnInodo(archivoDisco, superBoot.AptTablaInicioInodos, iNode) // PODRIA DEJAR SOLO ESTE PARA "NO" ESCRIBIR DOS VECES EL I-NODO
				//********************************************************************** ESCRIBIENDO LA BITACORA O LOG
				log1 := Bitacora{}
				copy(log1.TipoOpe[:], "mkdir")
				copy(log1.Nombre[:], "root")
				copy(log1.Tipo[:], "Directorio")
				copy(log1.fecha[:], time.Now().String())
				log1.Contenido = bloque1.DBdata
				escribirBitacora(archivoDisco, superBoot.AptLog, log1)
				log2 := Bitacora{}
				copy(log2.TipoOpe[:], "mkfile")
				copy(log2.Nombre[:], "users.txt")
				copy(log2.Tipo[:], "Archivo")
				copy(log2.fecha[:], time.Now().String())
				log2.Contenido = bloque2.DBdata
				escribirBitacora(archivoDisco, superBoot.AptLog+1+int64(binary.Size(log1)), log2)
				escribirSB(archivoDisco, part.Parti.Inicio, superBoot)
				superBoot.imprimirDatosBoot()
				println(color.Green + "-----------------------------------" + color.Reset)
				println(color.Green + "FORMETO DE UNA PARTICION REALIZADO " + color.Reset)
				println(color.Green + "-----------------------------------" + color.Reset)

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

func escribirUnAVD(archivoDisco *os.File, desde int64, objeto AVD) {
	archivoDisco.Seek(desde, 0)
	var escritor bytes.Buffer
	binary.Write(&escritor, binary.BigEndian, &objeto)
	escribirBinariamente(archivoDisco, escritor.Bytes())
}

func escribirBitMap(archivoDisco *os.File, desde int64, bitMap []byte) {
	archivoDisco.Seek(desde, 0)
	var escritor bytes.Buffer
	binary.Write(&escritor, binary.BigEndian, &bitMap)
	escribirBinariamente(archivoDisco, escritor.Bytes())
}

func escribirDetalleDir(archivoDisco *os.File, desde int64, d DetalleDir) {
	archivoDisco.Seek(desde, 0)
	var escritor bytes.Buffer
	binary.Write(&escritor, binary.BigEndian, &d)
	escribirBinariamente(archivoDisco, escritor.Bytes())
}

func escribirUnInodo(archivoDisco *os.File, desde int64, i Inodo) {
	archivoDisco.Seek(desde, 0)
	var escritor bytes.Buffer
	binary.Write(&escritor, binary.BigEndian, &i)
	escribirBinariamente(archivoDisco, escritor.Bytes())
}

func escribirUnBloque(archivoDisco *os.File, desde int64, bloque Bloque) {
	archivoDisco.Seek(desde, 0)
	var escritor bytes.Buffer
	binary.Write(&escritor, binary.BigEndian, &bloque)
	escribirBinariamente(archivoDisco, escritor.Bytes())
}
func escribirBitacora(archivoDisco *os.File, desde int64, log Bitacora) {
	archivoDisco.Seek(desde, 0)
	var escritor bytes.Buffer
	binary.Write(&escritor, binary.BigEndian, &log)
	escribirBinariamente(archivoDisco, escritor.Bytes())
}

func escribirSB(archivoDisco *os.File, desde int64, sb SuperB) {
	archivoDisco.Seek(desde, 0)
	var escritor bytes.Buffer
	binary.Write(&escritor, binary.BigEndian, &sb)
	escribirBinariamente(archivoDisco, escritor.Bytes())
}

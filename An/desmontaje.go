package An

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"

	"github.com/TwinProduction/go-color"
)

func desmontar(id string) {
	print(color.Yellow + "Desmontando: " + color.Reset)
	println(id)
	letra := string(id[2])
	if len(DiscosMontados_) == 0 {
		println(color.Red + "---------------------------")
		println(color.Red + "no hay particiones montadas ")
		println(color.Red + "---------------------------" + color.Reset)
		return
	}
	if !(validacionQueEsteMontada(letra, id)) {
		return
	}

	for x := 0; x < len(DiscosMontados_); x++ { // primero ver si esta en la lista , luego ver si existe el disco  ,luego ver si es primaria o logica, y por ultimo escribir easy :'v , quitar de la lista 2
		if DiscosMontados_[x].Letra == letra { // encuentro el disco
			if _, err := os.Stat(DiscosMontados_[x].Path); !(os.IsNotExist(err)) {

				for u_u := 0; u_u < len(DiscosMontados_[x].ParticionesMontadas); u_u++ {
					if DiscosMontados_[x].ParticionesMontadas[u_u].Id == id { // tengo la particion manipulada :v osea que si existe
						if DiscosMontados_[x].ParticionesMontadas[u_u].Tipo == 'p' {
							archivoDisco, err := os.OpenFile(DiscosMontados_[x].Path, os.O_RDWR, 0644) // TIEENE PERMISOS DE ESCRITURA Y DE LECTURA PERRO :V
							defer archivoDisco.Close()
							if err != nil {
								log.Fatal(err)
							}
							mrbAuxiliar := TipoMbr{}
							tamanioMbr := binary.Size(mrbAuxiliar) // este mbrAuxiliar aunque este vacio ya tiene el tamanio por defecto por eso aca se usa
							datosEnBytes := leerBytePorByte(archivoDisco, tamanioMbr)
							buff := bytes.NewBuffer(datosEnBytes)                   // lo convierto a buffer porque eso pedia la funcion
							err = binary.Read(buff, binary.BigEndian, &mrbAuxiliar) //se decodifica y se guarda en el mbrAuxiliar , asi que despues de aca ya tengo el original
							if err != nil {
								log.Fatal("error de lectura", err)
							}
							if mrbAuxiliar.buscarExistenciaEnParticiones(string(DiscosMontados_[x].ParticionesMontadas[u_u].Nombre[:])) {
								mrbAuxiliar.Particiones[DiscosMontados_[x].ParticionesMontadas[u_u].PosArray] = DiscosMontados_[x].ParticionesMontadas[u_u].Parti
								escribirMBR(archivoDisco, mrbAuxiliar)
								fmt.Println("---------------------------------")
								println(color.Blue + "Particion Primaria desmontada con Exito" + color.Reset)
								fmt.Println("---------------------------------")
								fmt.Println()
							} else {
								println(color.Red + "Lo siento al parecerer la particion fue eliminada y ya no se encuentra en el disco , de igual mandera se desmontara(Quitara de la lista)" + color.Reset)
							}
							// DE UNA O DE OTRA MANERA TENGO QUE QUITAR ESTE ELEMENTO DE  MI ARRAY :'V
							DiscosMontados_[x].ParticionesMontadas = QuitarMontaje(DiscosMontados_[x].ParticionesMontadas, u_u)
						} else {
							archivoDisco, err := os.OpenFile(DiscosMontados_[x].Path, os.O_RDWR, 0644) // TIEENE PERMISOS DE ESCRITURA Y DE LECTURA PERRO :V
							defer archivoDisco.Close()
							if err != nil {
								log.Fatal(err)
							}
							mrbAuxiliar := TipoMbr{}
							tamanioMbr := binary.Size(mrbAuxiliar) // este mbrAuxiliar aunque este vacio ya tiene el tamanio por defecto por eso aca se usa
							datosEnBytes := leerBytePorByte(archivoDisco, tamanioMbr)
							buff := bytes.NewBuffer(datosEnBytes)                   // lo convierto a buffer porque eso pedia la funcion
							err = binary.Read(buff, binary.BigEndian, &mrbAuxiliar) //se decodifica y se guarda en el mbrAuxiliar , asi que despues de aca ya tengo el original
							if err != nil {
								log.Fatal("error de lectura", err)
							}

							_, existeAun := mrbAuxiliar.getLOGICA(string(DiscosMontados_[x].ParticionesMontadas[u_u].Nombre[:]), archivoDisco)
							if existeAun {
								ebrAux := DiscosMontados_[x].ParticionesMontadas[u_u].PartiLogica
								escribirUnEBR(archivoDisco, (ebrAux.Inicio - int64(binary.Size(ebrAux))), ebrAux)
								fmt.Println("---------------------------------")
								println(color.Blue + "Particion Logica desmontada con Exito" + color.Reset)
								fmt.Println("---------------------------------")
								fmt.Println()
							} else {
								println(color.Red + "Lo siento al parecerer la particion fue eliminada y ya no se encuentra en el disco  , de igual mandera se desmontara(Quitara de la lista)" + color.Reset)
							}
							// igual se desmonta :v
							DiscosMontados_[x].ParticionesMontadas = QuitarMontaje(DiscosMontados_[x].ParticionesMontadas, u_u)
						}
						break
					}
				}

			} else {
				println(color.Red + "----------------------------")
				println(color.Red + "-- Ya no existe ese disco --")
				println(color.Red + "----------------------------" + color.Reset)
			}

		}
	}

}

// QuitarMontaje  quita la particion montada :v
func QuitarMontaje(mon []Montura, indiceQuit int) []Montura {
	return append(mon[:indiceQuit], mon[indiceQuit+1:]...)
}

func validacionQueEsteMontada(letra string, id string) bool {
	for x := 0; x < len(DiscosMontados_); x++ { // primero ver si esta en la lista , luego ver si existe el disco  ,luego ver si es primaria o logica, y por ultimo escribir easy :'v , quitar de la lista 2
		if DiscosMontados_[x].Letra == letra { // encuentro el disco
			if _, err := os.Stat(DiscosMontados_[x].Path); !(os.IsNotExist(err)) {
				for u_u := 0; u_u < len(DiscosMontados_[x].ParticionesMontadas); u_u++ {
					if DiscosMontados_[x].ParticionesMontadas[u_u].Id == id { // tengo la particion manipulada :v osea que si existe
						return true
					}
				}

			}
		}
	}
	println(color.Red + "-------------------------------")
	println(color.Red + "LA PARTICION NO ESTABA MONTADA")
	println(color.Red + "-------------------------------" + color.Reset)
	return false
}

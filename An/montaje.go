package An

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"

	"github.com/TwinProduction/go-color"
)

type Montura struct {
	// path , nombre , id
	PathDisco     string // me va servir para ir a escribir al mero disco cuando desmonte
	Nombre        [16]byte
	Id            string
	LetraDelDisco string
	Parti         Particion
	PosArray      uint8
	//cuando la desmonte la voy a escribir al disco  , debo GUARDAR LA POSICION PARA LA HORA QUE HAGA EL DESMONTAJE
	PartiLogica Ebr
	Tipo        byte
}

type disco struct {
	Path                string // EL PATH ES PARA IR A TRAER EL ARCHIVO Y LEERLO , aca creo que no es necesario  porque tengo la letra
	Letra               string // POR LA MANERA EN LA QUE LA ESTOY TRABAJANDO FIJO NECESITO SABER LA LETRA
	ParticionesMontadas []Montura
	CantidadPartciones  int64 // para el id
}

func (d *disco) agregarParticionMontada(path string, nombre string, id string, posArray uint8, partition Particion) { // ACA DEBO ENVIAR UN EBR TAMBIEN

	mon := Montura{PathDisco: path, Id: id, PosArray: posArray, Parti: partition, Tipo: 'p'}
	copy(mon.Nombre[:], nombre)
	d.ParticionesMontadas = append(d.ParticionesMontadas, mon)
}
func (d *disco) agregarParticionMontadaLOGICA(path string, nombre string, id string, ebrMontura Ebr) { // ACA DEBO ENVIAR UN EBR TAMBIEN
	mon := Montura{PathDisco: path, Id: id, PartiLogica: ebrMontura, Tipo: 'l'}
	copy(mon.Nombre[:], nombre)
	d.ParticionesMontadas = append(d.ParticionesMontadas, mon)
}

func (d disco) imprimirMontura() {
	for x := 0; x < len(d.ParticionesMontadas); x++ {
		d.ParticionesMontadas[x].getImpresionMontura()
	}
}

func crearMontaje(path string, nombre string) { // montaje
	if len(path) != 0 && len(nombre) != 0 {
		if _, err := os.Stat(path); !(os.IsNotExist(err)) {
			verificarSiExisteParticion(path, nombre)
		} else {
			println(color.Red + "--No existe ese disco--" + color.Reset)
		}
	} else {
		fmt.Println("falto especificar un parametro obligatorio")
	}

	Path_ = ""
	Name_ = ""
}

func verificarSiExisteParticion(direccion_archivo_binario string, nombreBuscar string) {
	archivoDisco, err := os.OpenFile(QuitarComillas(direccion_archivo_binario), os.O_RDWR, 0644)
	defer archivoDisco.Close()
	if err != nil {
		log.Fatal(err)
		return // ya no hay nada por hacer si no se pudo abrir el archivo
	}
	mrbAuxiliar := TipoMbr{}
	tamanioMbr := binary.Size(mrbAuxiliar) // este mbrAuxiliar aunque este vacio ya tiene el tamanio por defecto por eso aca se usa
	datosEnBytes := leerBytePorByte(archivoDisco, tamanioMbr)
	buff := bytes.NewBuffer(datosEnBytes)                   // lo convierto a buffer porque eso pedia la funcion
	err = binary.Read(buff, binary.BigEndian, &mrbAuxiliar) //se decodifica y se guarda en el mbrAuxiliar , asi que despues de aca ya tengo el original
	if err != nil {
		log.Fatal("error---", err)
	}
	archivoDisco.Seek(0, 0)
	respuesta := mrbAuxiliar.buscarExistenciaEnParticiones(nombreBuscar) // obtengo el mbr y reviso las particiones
	//traigo la particion y el indice
	//mrbAuxiliar.imprimirDatosMBR()                                       // solo para verificar si lo que retorno esta bien
	if respuesta {
		parti, posArray := mrbAuxiliar.GetParticionYposicion(nombreBuscar)
		if yaRegistreElPathEnElMount(direccion_archivo_binario) {
			// si ya lo registre solo retorno ese disco y le asigno su nuevos atributos
			discoYaMontado := getDiscoMontadoPorPath(direccion_archivo_binario) // ANTES DE METER UNA NUEVA PARTICION TENGO QUE IR A VER SI EXISTE EN MI LISTA
			validacionYaMonteEsaParticion := false
			var aux [16]byte
			copy(aux[:], nombreBuscar)
			for i := 0; i < len(discoYaMontado.ParticionesMontadas); i++ {
				if string(discoYaMontado.ParticionesMontadas[i].Nombre[:]) == string(aux[:]) {
					validacionYaMonteEsaParticion = true
				}

			}
			if validacionYaMonteEsaParticion == false {
				discoYaMontado.CantidadPartciones++
				var idPartition string = "vd" + discoYaMontado.Letra
				idPartition = fmt.Sprint(idPartition, discoYaMontado.CantidadPartciones)
				discoYaMontado.agregarParticionMontada(direccion_archivo_binario, nombreBuscar, idPartition, posArray, parti)
				println(color.Yellow + "Particion Montada" + color.Reset)
			} else {
				println(color.Red + "\n---------------------------------------------------------------")
				println(color.Red + "ESTA PARTICION YA ESTA MONTADA , NO LA PUEDES VOLVERLA A MONTAR")
				println(color.Red + "---------------------------------------------------------------" + color.Reset)
			}
		} else {
			discoNuevo := disco{Path: direccion_archivo_binario, Letra: getLetra(), CantidadPartciones: 0}
			discoNuevo.CantidadPartciones++
			var idPartition string = "vd" + discoNuevo.Letra
			idPartition = fmt.Sprint(idPartition, discoNuevo.CantidadPartciones)
			discoNuevo.agregarParticionMontada(direccion_archivo_binario, nombreBuscar, idPartition, posArray, parti)
			addMonturaDisco(discoNuevo)
			println(color.Yellow + "Particion Montada" + color.Reset)
		}
	} else {
		// PUEDO HACER ALGO PARA LEER LAS LOGICAS , SI NO ESTA EN LAS LOGICAS SI F no esta :'v
		ebrParticionLogica, res := mrbAuxiliar.getLOGICA(nombreBuscar, archivoDisco)
		if res {
			// AHORA TENGO QUE VER SI YA REGISTRE EL DISCO
			if yaRegistreElPathEnElMount(direccion_archivo_binario) {
				// si ya lo registre solo retorno ese disco y le asigno su nuevos atributos
				discoYaMontado := getDiscoMontadoPorPath(direccion_archivo_binario)
				validacionYaMonteEsaParticion := false
				var aux [16]byte
				copy(aux[:], nombreBuscar)
				for i := 0; i < len(discoYaMontado.ParticionesMontadas); i++ {
					if string(discoYaMontado.ParticionesMontadas[i].Nombre[:]) == string(aux[:]) {
						validacionYaMonteEsaParticion = true
					}

				}
				if validacionYaMonteEsaParticion == false {
					discoYaMontado.CantidadPartciones++
					var idPartition string = "vd" + discoYaMontado.Letra
					idPartition = fmt.Sprint(idPartition, discoYaMontado.CantidadPartciones)
					discoYaMontado.agregarParticionMontadaLOGICA(direccion_archivo_binario, nombreBuscar, idPartition, ebrParticionLogica)
					println(color.Yellow + "Particion Montada" + color.Reset)
				} else {
					println(color.Red + "\n---------------------------------------------------------------")
					println(color.Red + "ESTA PARTICION YA ESTA MONTADA , NO LA PUEDES VOLVERLA A MONTAR")
					println(color.Red + "---------------------------------------------------------------" + color.Reset)
				}
			} else {
				discoNuevo := disco{Path: direccion_archivo_binario, Letra: getLetra(), CantidadPartciones: 0}
				discoNuevo.CantidadPartciones++
				var idPartition string = "vd" + discoNuevo.Letra
				idPartition = fmt.Sprint(idPartition, discoNuevo.CantidadPartciones)
				discoNuevo.agregarParticionMontadaLOGICA(direccion_archivo_binario, nombreBuscar, idPartition, ebrParticionLogica)
				addMonturaDisco(discoNuevo)
				println(color.Yellow + "Particion Montada" + color.Reset)
			}
		} else {
			println(color.Red + "Error ese nombre no se encontro en ninguna particion de este disco" + color.Reset)
		}
	}
}

func mostrarMounts() {
	fmt.Println("***************PARTICIONES MONTADOS****************")
	b := false
	for u_u := 0; u_u < len(DiscosMontados_); u_u++ {
		b = true
		DiscosMontados_[u_u].imprimirMontura() // ACA HACER UN FOR CON LAS MONTURAS DEL DISCO MONTADAS E MOSTRARLAS
	}
	if !(b) {
		println(color.Yellow + "POR EL MOMENTO NO HAY PARTICIONES MONTADAS" + color.Reset)
	}
	fmt.Println("**********************************************")
}

func yaRegistreElPathEnElMount(path string) bool {
	for x := 0; x < len(DiscosMontados_); x++ {
		if DiscosMontados_[x].Path == path {
			return true
		}
	}
	return false
}

func getDiscoMontadoPorPath(path string) *disco {
	for x := 0; x < len(DiscosMontados_); x++ {
		if DiscosMontados_[x].Path == path {
			return &DiscosMontados_[x]
		}
	}
	discoVacio := disco{Letra: "NOENCONTRADO"}
	return &discoVacio // en teoria nunca va pasar
}
func getDiscoMontadoPorLetraID(letraID string) *disco {
	for x := 0; x < len(DiscosMontados_); x++ {
		if DiscosMontados_[x].Letra == letraID {
			return &DiscosMontados_[x]
		}
	}
	discoVacio := disco{Letra: "NOENCONTRADO"}
	return &discoVacio // en teoria nunca va pasar
}

func addMonturaDisco(disk disco) {
	DiscosMontados_ = append(DiscosMontados_, disk)
}

/*
	ALGORITMO PARA LAS PARTICIONES

*/

// SI EXISTE LA PARTICION , ENTONCES BUSCO EL PATH SI YA TENGO REGISTRADO EL PATH SOLO LO GUARDO EN MI disco particionesMontadas
// si es un nuevo disco mi contador de letrass incrementa y pasa a a ser el b ,
// si el path no esta registrado pero si existe entoneces LO AGREGO AL ARREGLO DE DISCOS DE DONDE SACO MIS PARTICIONES

func getLetra() string {
	aux := " " // entra al switch y cambia su valor siempre
	switch CONT_lETRA {
	case 0:
		aux = "a"
	case 1:
		aux = "b"
	case 2:
		aux = "c"
	case 3:
		aux = "d"
	case 4:
		aux = "e"
	case 5:
		aux = "f"
	case 6:
		aux = "g"
	case 7:
		aux = "h"
	case 8:
		aux = "i"
	case 9:
		aux = "j"
	case 10:
		aux = "k"
	case 11:
		aux = "l"
	case 12:
		aux = "m"
	case 13:
		aux = "n"
	case 14:
		aux = "o"
	case 15:
		aux = "p"
	case 16:
		aux = "q"
	case 17:
		aux = "r"
	case 18:
		aux = "s"
	case 19:
		aux = "t"
	case 20:
		aux = "u"
	case 21:
		aux = "v"
	case 22:
		aux = "w"
	case 23:
		aux = "x"
	case 24:
		aux = "y"
	case 25:
		aux = "z"
	default:
		aux = " "
		println(color.Red + " YA NO HAY MAS LETRAS PARA NOMBRAR PARTICIONES " + color.Reset)
	} // tengo que llegar al 25
	CONT_lETRA++ // asi siempre crece
	return aux
}

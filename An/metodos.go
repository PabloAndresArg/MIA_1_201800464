// Code generated by PABLO. DO NOT EDIT.

package An

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/TwinProduction/go-color"
)

// LAS FUNCIONES DE ESTE ARCHIVO ESTAN COMPARTIDAS PORQUE PERTENECEN AL MISMO PACKAGE , siempre deben de iniciar con MAYUSCULA el nombre del metodo para ser exportado

// QuitarComillas: igual para la declaracion de variables MAYUSCULA AL INICIO para usar ne otras partes

func init() {
	// es una funcion especial que permite iniciazar variables o estructuras
}

//quita \* de continuacion de linea
func QuitarSimboloNextLine(cadena string) string {
	salida := ""
	for i := 0; i < len(cadena); i++ {
		if (i != len(cadena)-2) && (i != len(cadena)-1) {
			salida += string(cadena[i])
		}
	}
	return salida
}

// QuitarComillas lo que hace es quitar comillas xd
func QuitarComillas(ruta string) string {
	ruta = strings.TrimSpace(ruta)
	salida := ""
	if ruta[0] == '"' {
		for i := 0; i < len(ruta); i++ {
			if i != 0 && (i != len(ruta)-1) {
				salida += string(ruta[i])
			}

		}
	} else {
		salida = ruta
	}
	return salida
}

func verificarRuta(ruta string) {
	rutas := strings.Split(ruta, "/")

	var temporal string = ""
	if ruta[len(ruta)-1] != '/' {
		i := 0
		for i < len(rutas) {
			temporal += rutas[i] + "/"
			CrearDirectorio_si_no_exist(temporal)
			//mkdisk-SIze->1-path->/home/pablo/Escritorio/carpetae/hola/-name->prub.dsk-unit->K
			i++
		}
	} else {
		i := 0
		for i < len(rutas) {
			if i != (len(rutas) - 1) {
				temporal += rutas[i] + "/"
				CrearDirectorio_si_no_exist(temporal)
			}
			i++
		}
	}
}
func CrearDisco(numero string, ruta string, nombre string, K_o_M string) {
	ruta = QuitarComillas(ruta)
	tamanio, _ := strconv.ParseInt(numero, 10, 64)
	verificarRuta(ruta)
	size := int64(0)
	if K_o_M == "K" || K_o_M == "k" {
		size = int64(tamanio * 1024)
	} else { // SINO SON MEGABYTES
		size = int64(tamanio * 1024 * 1024)
	}
	rutaCompleta := ruta + nombre
	fichero, err := os.Create(rutaCompleta)
	defer fichero.Close()
	if err != nil {
		log.Fatal("fallo creando el archivo de salida")
	} else {
		fmt.Println("---------------------------------")
		println(color.Yellow + " Disco creado Correctamente: " + color.Reset)
	}
	var cero int8 = 0 // asignando el cero
	direccion_cero := &cero
	var binario_ bytes.Buffer
	binary.Write(&binario_, binary.BigEndian, direccion_cero) // SE ESCRIBE UN CERO AL INICIO DEL ARCHIVO
	escribirBinariamente(fichero, binario_.Bytes())
	fichero.Seek(size-1, 0) // posicionarse en la pos 0
	var bin2_ bytes.Buffer  // se escribe un cero al final del archivo
	binary.Write(&bin2_, binary.BigEndian, direccion_cero)
	escribirBinariamente(fichero, bin2_.Bytes())

	/*



		METIENDO EL STRUCT AL DISCO , 	// SEREALIZACION DEL STRUCT , escribir al inicio del archivo el struct




	*/
	fichero.Seek(0, 0) // POS AL INICIO DEL ARCHIVO

	FechaFormatoTime := time.Now()
	mbr := TipoMbr{Tamanio: size, DiskSignature: dameUnNumeroRandom()}
	copy(mbr.Fecha[:], FechaFormatoTime.String())

	for i := 0; i < 4; i++ {
		mbr.Particiones[i] = Particion{Status: 'n', Size: 0} // PARA MI N ES QUE NO HAY , Y es de yes que si hay xd
	}

	var bin3_ bytes.Buffer
	binary.Write(&bin3_, binary.BigEndian, &mbr)
	escribirBinariamente(fichero, bin3_.Bytes())
	fmt.Println("Nombre: " + nombre)
	fmt.Printf("\nFECHA: %s\nTamanio: %v\n", mbr.Fecha, mbr.Tamanio)
	fmt.Printf("Signature: %d\n", mbr.DiskSignature)
	fmt.Printf("hay pariciones(n/y): %c\n", mbr.Particiones[0].Status)
	fmt.Println("---------------------------------")
	// limpiar variables
	Name_ = ""
	Path_ = ""
	Size_ = ""
	Unit_m_ = "M"

}

func escribirBinariamente(fichero *os.File, bytes []byte) {
	_, erro := fichero.Write(bytes)
	if erro != nil {
		log.Fatal(erro)
	}
}

// elimina un disco duro o archivo
func EliminarDisco(ruta_absoluta string) {
	ruta_absoluta = QuitarComillas(ruta_absoluta)
	if _, err := os.Stat(ruta_absoluta); !(os.IsNotExist(err)) {
		fmt.Println("¿ESTA SEGURO DE QUERER ELIMINAR ESTE DISCO? ")
		fmt.Print("Presione 1 para confirmar, dsino presione 0")
		var decision int
		fmt.Scanln(&decision)
		if decision == 1 {
			erro := os.Remove(ruta_absoluta)
			if erro == nil {
				fmt.Println("Disco eliminado...")
			} else {
				fmt.Printf("ERROR , NO SE PUEDO ELIMINAR EL DISCO: %v\n", erro)
			}
		}
	} else {
		println(color.Red + "-- No existe ese disco --" + color.Reset)
	}
}

// crea un direcctorio si no encuentra la ruta
func CrearDirectorio_si_no_exist(dir__ string) {

	if _, err := os.Stat(dir__); os.IsNotExist(err) {

		err = os.Mkdir(dir__, 0755)
		fmt.Println("Crea la carpeta:  " + dir__)
		if err != nil {
			panic(err)
		}
	}

}

// LeerBinariamente es para leer archivos binarios
func LeerBinariamenteMimbr(direccion_archivo_binario string) {
	archivoDisco, err := os.OpenFile(QuitarComillas(direccion_archivo_binario), os.O_RDWR, 0644) // TIEENE PERMISOS DE ESCRITURA Y DE LECTURA PERRO :V
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
		log.Fatal("error al leer", err)
	}
	mrbAuxiliar.imprimirDatosMBR()
	//  v – formats the value in a default format
	//  d – formats decimal integers
	//  g – formats the floating-point numbers
	//  b – formats base 222 numbers
	//  o – formats base 888 numbers
	//  t – formats true or false values
	//  s – formats string values
	//  c- caracteres
	/*												ESCRIBIR LA PARTICION 						*/
}

func leerBytePorByte(archivoDisco *os.File, tamanio int) []byte {
	bytes := make([]byte, tamanio) // hago un arreglo dinamico de bytes
	_, err := archivoDisco.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}
func dameUnNumeroRandom() int64 { // para el signature del mbr , tener una lista para asegurarme que no se  repitan los ids
	num := int64(rand.Intn((100)-1) + (10 / 10))
	for esRepetido(num) {
		num = int64(rand.Intn((100)-1) + (10 / 10))
	}
	return num
}

// MetodosParticiones es para crear o eliminar particiones
func MetodosParticiones(rutaPath string, nombreName string, sizeTamanio string, fit string, delete string, add string, tipo__ string, unit string) {
	// atributos obligatorios : NAME , PATH  SIZE
	// OPCIONALES  unit , type , fit , delete , add
	if len(nombreName) != 0 {
		nombreName = QuitarComillas(nombreName)
	}

	if len(rutaPath) != 0 && len(nombreName) != 0 && len(sizeTamanio) != 0 { // si vienen los obligatorios
		archivoDisco, err := os.OpenFile(QuitarComillas(rutaPath), os.O_RDWR, 0644) // TIEENE PERMISOS DE ESCRITURA Y DE LECTURA PERRO :V
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
		/*
			CREANDO LA PARTICION Y VOLVIENDO A ESCRIBIR EN EL ARCHIVO
		*/
		if len(delete) == 0 && len(add) == 0 {
			if mrbAuxiliar.hayUnaParticionDisponible() {
				size, _ := strconv.ParseInt(sizeTamanio, 10, 64)
				size = getSizeConUnidad(size, unit)
				if mrbAuxiliar.hayEspacioSuficiente(size) {
					////fit string , size int64, nombre string, tipo byte
					tipoParticionByte := getTipoEnBytes(tipo__) // TENGO QUE VOLVER A MI TIPO UN BIYE  pero primero ver si viene algo o es el default
					switch tipoParticionByte {
					case 'p':
						mrbAuxiliar = mrbAuxiliar.crearParticion(fit, size, nombreName, tipoParticionByte) // ES POSIBLE CREAR LA PARTICION

					case 'e':
						if !(mrbAuxiliar.yaExisteUnaExtendida()) { // si no existe una extendida pues la puede crear
							pos := mrbAuxiliar.crearParticionExtendida(fit, size, nombreName, tipoParticionByte) // le mande directo e

							/* TENGO QUE CREAR EL EBR DE INICIO */
							desde := int64(mrbAuxiliar.Particiones[pos].Inicio)
							ebr := Ebr{Inicio: desde, Status: 'n', Size: 0, Next: -1, Fit: 'w'} // el name esta vacio
							escribirUnEBR(archivoDisco, desde, ebr)
						} else {
							println(color.Red + "Lo siento solo se puede tener una particion extendida por disco" + color.Reset)
						}

					case 'l':
						if mrbAuxiliar.yaExisteUnaExtendida() { // si EXISTE es posible crear una logica
							extendida := mrbAuxiliar.getExtendida() // NECESITO EL INICIO DE LA EXTENDIDA PARA POSICIONARME EN EL PRIMER EBR
							if size < extendida.Size {
								archivoDisco.Seek(extendida.Inicio, 0)
								ebrAux := Ebr{}
								tamanioEBR := binary.Size(ebrAux) //tamanio de lo que ire a traer
								ebr_en_bytes := leerBytePorByte(archivoDisco, tamanioEBR)
								buff := bytes.NewBuffer(ebr_en_bytes)              // lo convierto a buffer porque eso pedia la funcion
								err = binary.Read(buff, binary.BigEndian, &ebrAux) //ya tengo el original
								ocupado := int64(binary.Size(ebrAux))
								/*fmt.Println("--------------- PRIMER EBR ----------------")
								ebrAux.imprimirDatosEbr()
								fmt.Println("-------------------------------------------")*/
								if ebrAux.Status == 'n' {
									if (ocupado + size) < extendida.Size {
										ebrAux.Status = 'y'
										ebrAux.Fit = getFit(fit)
										ebrAux.Inicio = (extendida.Inicio + ocupado) // el ocupado por el momento tendria solo el tamanio del ebr
										ebrAux.Size = size
										copy(ebrAux.Nombre[:], nombreName)
										ebrAux.Next = -1
										escribirUnEBR(archivoDisco, extendida.Inicio, ebrAux)
										fmt.Println("---------------------------------")
										println(color.Green + "PARTICION LOGICA CREADA CON EXITO" + color.Reset)
										ebrAux.imprimirDatosEbr()
										fmt.Println("---------------------------------")
									} else {
										println(color.Red + "NO CABE UNA LOGICA DE ESE SIZE" + color.Reset)
									}

								} else {
									// tengo que recorrer los EBR
									if ebrAux.Next == -1 {
										ocupado += ebrAux.Size + int64(binary.Size(ebrAux))
									} else {
										ocupado = 0
									}
									//fmt.Printf(color.Yellow+"EBR actual: %s\n", ebrAux.Nombre)
									for ebrAux.Next != -1 {
										ocupado += ebrAux.Size + int64(binary.Size(ebrAux)) // por cada iteracion se va acumulando el espacio ocupad
										if (ocupado + size + 1) > extendida.Size {
											println(color.Red + "NO CABE UNA LOGICA DE ESE SIZE" + color.Reset)
											return
										}
										// LEER EBR POR EBR
										posicionFinalEbr := int64(ebrAux.Inicio + ebrAux.Size)
										archivoDisco.Seek(posicionFinalEbr+1, 0)
										tamanioEBR := binary.Size(ebrAux) //tamanio de lo que ire a traer
										ebr_en_bytes := leerBytePorByte(archivoDisco, tamanioEBR)
										buff := bytes.NewBuffer(ebr_en_bytes)              // lo convierto a buffer porque eso pedia la funcion
										err = binary.Read(buff, binary.BigEndian, &ebrAux) //ya tengo el original
										//fmt.Printf(color.Yellow+"EBR actual: %s\n", ebrAux.Nombre)
									}
									// AL SALIR DEL FOR EN TEORIA TENDRIA EL EBR QUE TIENE COMO SIGUIENTE A -1
									//	println("SALIO EL EBR" + color.Reset)
									//ebrAux.imprimirDatosEbr()
									ebrAux.Next = ebrAux.Inicio + ebrAux.Size + 1
									if (ocupado + size + 1) < extendida.Size { // ahora este es el ultimo
										ebrNuevo := Ebr{}
										ebrNuevo.Status = 'y'
										ebrNuevo.Fit = getFit(fit)
										ebrNuevo.Inicio = ebrAux.Next + int64(binary.Size(ebrNuevo)) // el ocupado por el momento tendria solo el tamanio del ebr
										ebrNuevo.Size = size
										copy(ebrNuevo.Nombre[:], nombreName)
										ebrNuevo.Next = -1
										escribirUnEBR(archivoDisco, (ebrAux.Inicio - int64(binary.Size(ebrAux))), ebrAux) // REFRESCTO SU NEXT
										escribirUnEBR(archivoDisco, ebrAux.Next, ebrNuevo)
										fmt.Println("---------------------------------")
										println(color.Green + "PARTICION LOGICA CREADA CON EXITO" + color.Reset)
										ebrNuevo.imprimirDatosEbr()
										fmt.Println("---------------------------------")
									} else {
										println(color.Red + "NO CABE UNA LOGICA DE ESE SIZE" + color.Reset)
									}

									// FIN CREACION DE PARTICION LOGICA
								}

							} else {
								println(color.Red + "EL SIZE DE ESTA PARTICION SUPERA AL SIZE DE LA EXTENDIDA" + color.Reset)
							}
						} else {
							println(color.Red + "No puedes Crear una Particion Logica sin antes tener una extendida" + color.Reset)
						}
					default:
						println("Error en el tipo de particion")
					}

					/*escribiendo el MBR */
					archivoDisco.Seek(0, 0) // al inicio del archivo para sobreescribir mi disco
					var escritor bytes.Buffer
					binary.Write(&escritor, binary.BigEndian, &mrbAuxiliar)
					escribirBinariamente(archivoDisco, escritor.Bytes())

				} else {
					println(color.Red + "Espacio insuficiente" + color.Reset)
				}
			} else {
				println(color.Yellow + "Lo siento no es posible porque ya tiene 4 particiones en este disco" + color.Reset)
			} /*





			 */
		} else if len(delete) != 0 && len(add) == 0 { // vengo a borrar una particion
			delete = strings.ToLower(delete)
			switch delete {
			case "fast":
				println(color.Yellow + "¿ESTAS SEGURO DE QUERER ELIMINAR ESTA PARTICION EN MODO FAST?" + color.Reset)
				fmt.Println("Presiona 1 para confirmar")
				fmt.Println("Presiona 0 para no hacerlo")
				lector := bufio.NewReader(os.Stdin)
				entradaOP, _ := lector.ReadString('\n')
				menu := strings.TrimRight(entradaOP, "\r\n") // quita el salto de linea de la parte derecha
				switch menu {
				case "0":
					fmt.Println("ya no se elimino la particion :D")
				case "1":
					res := mrbAuxiliar.eliminarFast(nombreName, archivoDisco)
					if res {
						archivoDisco.Seek(0, 0) // al inicio del archivo para sobreescribir mi disco
						var escritor bytes.Buffer
						binary.Write(&escritor, binary.BigEndian, &mrbAuxiliar)
						escribirBinariamente(archivoDisco, escritor.Bytes())
						println(color.Blue + "Particion eliminada" + color.Reset)

					} else {
						println(color.Red + "no se encontro la particion" + color.Reset)
					}

				default:
					fmt.Println("opcion no encontrada..")
				}

			case "full":
				println(color.Yellow + "¿ESTAS SEGURO DE QUERER ELIMINAR ESTA PARTICION EN MODO FULL?" + color.Reset)
				fmt.Println("Presiona 1 para confirmar")
				fmt.Println("Presiona 0 para no hacerlo")
				lector := bufio.NewReader(os.Stdin)
				entradaOP, _ := lector.ReadString('\n')
				menu := strings.TrimRight(entradaOP, "\r\n") // quita el salto de linea de la parte derecha
				switch menu {
				case "0":
					fmt.Println("ya no se elimino la particion :D")
				case "1":
					if mrbAuxiliar.buscarExistenciaEnParticiones(nombreName) {
						particionDelete, _ := mrbAuxiliar.GetParticionYposicion(nombreName) // TENGO SU PARTICION ASI QUE PUEDO ACTUAR CON LO DEL TIPO S
						inicio := particionDelete.Inicio
						fin := particionDelete.Size
						res := mrbAuxiliar.eliminarFast(nombreName, archivoDisco)
						if res {
							archivoDisco.Seek(0, 0) // al inicio del archivo para sobreescribir mi disco
							var escritor bytes.Buffer
							binary.Write(&escritor, binary.BigEndian, &mrbAuxiliar)
							escribirBinariamente(archivoDisco, escritor.Bytes())
							/*
								LLENANDO LA PARTICION DE 0 para dejarla limpia
							*/
							archivoDisco.Seek(inicio, 0)
							var ceros []byte
							for r := 0; r < int(fin); r++ {
								ceros = append(ceros, 0)
							}
							var nuevoEscritor bytes.Buffer
							binary.Write(&nuevoEscritor, binary.BigEndian, &ceros)
							escribirBinariamente(archivoDisco, nuevoEscritor.Bytes())
							println(color.Blue + "Particion eliminada" + color.Reset)
						} else {
							println(color.Red + "no se encontro la particion" + color.Reset)
						}
					} else {
						// TONS PUEDE QUE SEA LOGICA , YA SI NO ES LOGICA SI TIRO EL ERROR A LO MASHO :v
						res := mrbAuxiliar.eliminarFast(nombreName, archivoDisco)
						if res {

						} else {
							println(color.Red + "no se encontro la particion con ese nombre , no se podra hacer la eliminacion FULL " + color.Reset)
						}
					}

				default:
					fmt.Println("opcion no encontrada..")
				}

			default:
				println(color.Red + "parametro indefinido :v" + color.Reset)
			}

		} else if len(add) != 0 && len(delete) == 0 { // vengo a dar mas espacio a una particion  o disminuir
			add = strings.TrimSpace(add)
			if add[0] == '-' { // resta
				add, _ := strconv.ParseInt(add, 10, 64) // lo devuelvo como un entero
				add = getSizeConUnidad(add, unit)
				// ya tengo mi tamaño en bytes

				// PARA REDUCIR SOLO TENGO QUE VER SI LO QUE LE QUITARE NO ES MENOR A 0 , PRIMERO IR A BUSCAR SI EXISTE LA PARTICION
				if mrbAuxiliar.buscarExistenciaEnParticiones(nombreName) { // es das PRINCIPALES
					// ENTONCES NECESITO ESA PRIMARIA O EXT y LE RESTO Y SI ES MAYOR A CERO ESTA PERMITIDO
					parti, pos := mrbAuxiliar.GetParticionYposicion(nombreName)
					if parti.Tipo == 'E' || parti.Tipo == 'e' {
						// ACA FUNCIONARIA UN POCO DIFERENTE EL REDUCIR PORQUE NO SE SI HAY ALGUN EBR OCUPANDO ESE ESPACIO , IGUAL SERIA DE HACERLO CON UN RANGO
						ebrEncont, band := mrbAuxiliar.getUltimoEbrDeLasLogicas(archivoDisco)
						if band {
							limite := ebrEncont.Inicio + ebrEncont.Size // EL VALOR NUEVO TIENE QUE SER MAYOR AL DEL limite
							if parti.Inicio+parti.Size+add >= limite {  // tons si , size realtivo 												 TOMAR EN CUENTA STRUCST PARA ESPACIO DISPONIBLE EN EL DISCO Y EN LA EXTENDIDA
								mrbAuxiliar.Particiones[pos].Size = parti.Size + add // EN REALIDAD ES UNA RESTA
								fmt.Println("---------------------------------")
								fmt.Println("la particion actualmente es de: " + fmt.Sprint(parti.Size) + " bytes")
								fmt.Println("Disponible: " + fmt.Sprint(parti.Size-mrbAuxiliar.getTamanioOcupadoDeLosEbrs(archivoDisco)))
								mrbAuxiliar.Particiones[pos].Size = parti.Size + add // EN REALIDAD ES UNA RESTA
								escribirMBR(archivoDisco, mrbAuxiliar)
								println(color.Yellow + "REDUCIENDO LA PARTICION EXTENDIDA: " + nombreName + " en " + fmt.Sprint(add) + " bytes" + color.Reset)
								fmt.Println("ahora esta paricion Es de: " + fmt.Sprint(mrbAuxiliar.Particiones[pos].Size) + " bytes")
								fmt.Println("---------------------------------")
							} else {
								println(color.Red + "No puedo reducir la extendida porque estaria perdiendo LOGICAS")
								fmt.Println("La ultima particion logica llega a ocupar hasta el byte " + fmt.Sprint(limite) + " y si dejo hacer la reduccion mi extendida llega hasta el byte " + fmt.Sprint(parti.Inicio+parti.Size+add) + color.Reset)
							}

						} else {
							// pues si no hay ultimo ni modo le permito reducir xd
							if parti.Size+add <= (int64(binary.Size(tamanioMbr)) + 1) { // TENGO QUE TENER ESPACIO PARA AUNQUE SEA UN EBR SINO NO TENDRIA SENTIDO Y AL MENOS UNA PARTICION DE 1 BYTE
								println(color.Red + "ERROR ESPACIO NEGATIVO o INSUFICIENTE PARA ALMACENAR UN EBR con una particion minima " + color.Reset)
							} else {
								mrbAuxiliar.Particiones[pos].Size = parti.Size + add // EN REALIDAD ES UNA RESTA
								fmt.Println("---------------------------------")
								fmt.Println("la particion actualmente es de: " + fmt.Sprint(parti.Size) + " bytes")
								fmt.Println("Sobran: " + fmt.Sprint(parti.Size-mrbAuxiliar.getTamanioOcupadoDeLosEbrs(archivoDisco)) + " bytes disponibles")
								mrbAuxiliar.Particiones[pos].Size = parti.Size + add // EN REALIDAD ES UNA RESTA
								escribirMBR(archivoDisco, mrbAuxiliar)
								println(color.Yellow + "REDUCIENDO LA PARTICION EXTENDIDA: " + nombreName + " en " + fmt.Sprint(add) + " bytes" + color.Reset)
								fmt.Println("ahora esta paricion tiene: " + fmt.Sprint(mrbAuxiliar.Particiones[pos].Size) + " bytes")
								fmt.Println("---------------------------------")
							}
						}
					} else {
						if parti.Size+add <= 0 { // LA RESTA ES UNA SUMA CON SIGNO DE PRECEDENCIA :V
							println(color.Red + "ERROR LA PARTICION SE QUEDA SIN ESPACIO O CON ESPACIO NEGATIVO" + color.Reset)
						} else {

							fmt.Println("---------------------------------")
							fmt.Println("la particion actualmente es de: " + fmt.Sprint(parti.Size) + " bytes")
							mrbAuxiliar.Particiones[pos].Size = parti.Size + add // EN REALIDAD ES UNA RESTA
							escribirMBR(archivoDisco, mrbAuxiliar)
							println(color.Yellow + "REDUCIENDO LA PARTICION PRIMARIA: " + nombreName + " en " + fmt.Sprint(add) + " bytes" + color.Reset)
							fmt.Println("ahora esta paricion tiene: " + fmt.Sprint(mrbAuxiliar.Particiones[pos].Size) + " bytes")
							fmt.Println("---------------------------------")
						}

					}

				} else { // DEBO BUSCAR EN LAS LOGICAS  SINO ESTAN ES QUE NO EXISTE ESA PARTICION EN EL DISCO
					if mrbAuxiliar.yaExisteUnaExtendida() {
						ebrEncontrado, bandera := mrbAuxiliar.getLOGICA(nombreName, archivoDisco)
						if bandera {
							if ebrEncontrado.Size+add > 0 { // LA RESTA ES PERMITIDA
								fmt.Println("---------------------------------")
								fmt.Println("la particion actualmente es de: " + fmt.Sprint(ebrEncontrado.Size) + " bytes")
								ebrEncontrado.Size = ebrEncontrado.Size + add
								escribirUnEBR(archivoDisco, (ebrEncontrado.Inicio - int64(binary.Size(ebrEncontrado))), ebrEncontrado) // SIEMPRE ES RELATIVO A SU INICIO
								println(color.Yellow + "REDUCIENDO LA PARTICION LOGICA: " + nombreName + " en " + fmt.Sprint(add) + " bytes" + color.Reset)
								fmt.Println("ahora esta paricion tiene: " + fmt.Sprint(ebrEncontrado.Size) + " bytes")
								fmt.Println("---------------------------------")
							} else {
								println(color.Red + "ERROR LA PARTICION SE QUEDA SIN ESPACIO O CON ESPACIO NEGATIVO" + color.Reset)
							}
						} else {
							println(color.Red + "Esa particion no esta en este disco.. (l) " + color.Reset)
						}
					} else {
						println(color.Red + "Esa particion no esta en este disco" + color.Reset)
					}
				}

			} else { // agranda
				add, _ := strconv.ParseInt(add, 10, 64) // lo devuelvo como un entero
				add = getSizeConUnidad(add, unit)
				if mrbAuxiliar.buscarExistenciaEnParticiones(nombreName) { // es das PRINCIPALES
					// ya se que es TIPO E o TIPO P
					rangos := mrbAuxiliar.getRangosParticiones(nombreName)
					if len(rangos) != 0 {
						_, pos := mrbAuxiliar.GetParticionYposicion(nombreName) // DISQUE SOLO EXISTE UNA PARTICION
						ban := false
						nuevoLimiteSuperior := mrbAuxiliar.Particiones[pos].Inicio + mrbAuxiliar.Particiones[pos].Size + add
						for x := 0; x < len(rangos); x++ { // pregunto posiciones relativas :v
							//fmt.Println(fmt.Sprint(nuevoLimiteSuperior) + ">=" + fmt.Sprint(rangos[x].LimiteInferior) + " && " + fmt.Sprint(nuevoLimiteSuperior) + "<=" + fmt.Sprint(rangos[x].LimiteSuperior))
							if nuevoLimiteSuperior >= rangos[x].LimiteInferior && nuevoLimiteSuperior <= rangos[x].LimiteSuperior { // QUIERE DECIR QUE ESTOY OCUPANDO ESPACIO QUE NO ES MIO
								ban = true
								break
							}
						}
						if ban {
							println(color.Red + "No hay espacio suficiente a la derecha, estarias ocupando espacio de otra particion" + color.Reset)
						} else {
							// SE PUEDE AGRANDAR LA PARTICION
							if mrbAuxiliar.hayEspacioSuficienteAdd(add) && nuevoLimiteSuperior < mrbAuxiliar.Tamanio {
								fmt.Println("---------------------------------")
								fmt.Println("la particion actualmente es de: " + fmt.Sprint(mrbAuxiliar.Particiones[pos].Size) + " bytes")
								mrbAuxiliar.Particiones[pos].Size = mrbAuxiliar.Particiones[pos].Size + add // corro a la derecha
								escribirMBR(archivoDisco, mrbAuxiliar)
								println(color.Blue + "INCREMENTANDO LA PARTICION: " + nombreName + " en " + fmt.Sprint(add) + " bytes" + color.Reset)
								fmt.Println("ahora esta paricion tiene: " + fmt.Sprint(mrbAuxiliar.Particiones[pos].Size) + " bytes")
								fmt.Println("---------------------------------")
							} else {
								println(color.Red + "Espacio insuficiente en el DISCO" + color.Reset)
							}

						}

					} else { // DE UNA CRECE
						_, pos := mrbAuxiliar.GetParticionYposicion(nombreName) // DISQUE SOLO EXISTE UNA PARTICION
						// validando que no sobre pase el tamaño del disco-mbr size
						if mrbAuxiliar.Particiones[pos].Size+add <= mrbAuxiliar.Tamanio-int64(binary.Size(mrbAuxiliar)) {
							fmt.Println("---------------------------------")
							fmt.Println("la particion actualmente es de: " + fmt.Sprint(mrbAuxiliar.Particiones[pos].Size) + " bytes")
							mrbAuxiliar.Particiones[pos].Size = mrbAuxiliar.Particiones[pos].Size + add // corro a la derecha
							escribirMBR(archivoDisco, mrbAuxiliar)
							println(color.Blue + "INCREMENTANDO LA PARTICION: " + nombreName + " en " + fmt.Sprint(add) + " bytes" + color.Reset)
							fmt.Println("ahora esta paricion tiene: " + fmt.Sprint(mrbAuxiliar.Particiones[pos].Size) + " bytes")
							fmt.Println("---------------------------------")
						} else {
							println(color.Red + "Espacio insuficiente en el Disco" + color.Reset)
						}
					}

				} else { // DEBO BUSCAR EN LAS LOGICAS  SINO ESTAN ES QUE NO EXISTE ESA PARTICION EN EL DISCO
					//..
					if mrbAuxiliar.yaExisteUnaExtendida() {
						ebrEncontrado, bandera := mrbAuxiliar.getLOGICA(nombreName, archivoDisco)
						if bandera {
							//-------------------- inicia la adicion a las logicas
							rangos := mrbAuxiliar.getRangosParticionesLogicas(archivoDisco, nombreName)
							if len(rangos) != 0 {
								ban := false
								nuevoLimite := ebrEncontrado.Inicio + ebrEncontrado.Size + add
								for x := 0; x < len(rangos); x++ { // pregunto posiciones relativas :v
									//fmt.Println(fmt.Sprint(nuevoLimite) + ">=" + fmt.Sprint((rangos[x].LimiteInferior - int64(binary.Size(ebrEncontrado)))) + " && " + fmt.Sprint(nuevoLimite) + "<=" + fmt.Sprint(rangos[x].LimiteSuperior))
									if nuevoLimite >= (rangos[x].LimiteInferior-int64(binary.Size(ebrEncontrado))) && nuevoLimite <= rangos[x].LimiteSuperior { // QUIERE DECIR QUE ESTOY OCUPANDO ESPACIO QUE NO ES MIO
										ban = true
										break
									}
								}
								if ban {
									println(color.Red + "No hay espacio suficiente a la derecha, estarias ocupando espacio de otra particion o de otro ebr" + color.Reset)
								} else {
									nuevoLimite := ebrEncontrado.Inicio + ebrEncontrado.Size + add
									limiteExtendida := mrbAuxiliar.getExtendida().Inicio + mrbAuxiliar.getExtendida().Size
									if nuevoLimite <= limiteExtendida { // ES PERMITIDO PORQUE PASO LO ANTERIOR
										fmt.Println("---------------------------------")
										fmt.Println("la particion actualmente es de: " + fmt.Sprint(ebrEncontrado.Size) + " bytes")
										ebrEncontrado.Size = ebrEncontrado.Size + add
										escribirUnEBR(archivoDisco, (ebrEncontrado.Inicio - int64(binary.Size(ebrEncontrado))), ebrEncontrado) // SIEMPRE ES RELATIVO A SU INICIO
										println(color.Blue + "INCREMENTANDO LA PARTICION LOGICA: " + nombreName + " en " + fmt.Sprint(add) + " bytes" + color.Reset)
										fmt.Println("ahora esta paricion tiene: " + fmt.Sprint(ebrEncontrado.Size) + " bytes")
										fmt.Println("---------------------------------")
									} else {
										println(color.Red + "Espacio insuficiente en LA EXTENDIDA" + color.Reset)
									}

								}

							} else { // QUIERE DECIR QUE SOLO ESTA UN EBR , solo tengo que tener cuidado que no supere al contenedor en este caso LA EXTENDIDA
								nuevoLimite := ebrEncontrado.Inicio + ebrEncontrado.Size + add
								limiteExtendida := mrbAuxiliar.getExtendida().Inicio + mrbAuxiliar.getExtendida().Size
								if nuevoLimite <= limiteExtendida { // ES PERMITIDO
									fmt.Println("---------------------------------")
									fmt.Println("la particion actualmente es de: " + fmt.Sprint(ebrEncontrado.Size) + " bytes")
									ebrEncontrado.Size = ebrEncontrado.Size + add
									escribirUnEBR(archivoDisco, (ebrEncontrado.Inicio - int64(binary.Size(ebrEncontrado))), ebrEncontrado) // SIEMPRE ES RELATIVO A SU INICIO
									println(color.Blue + "INCREMENTANDO LA PARTICION LOGICA: " + nombreName + " en " + fmt.Sprint(add) + " bytes" + color.Reset)
									fmt.Println("ahora esta paricion tiene: " + fmt.Sprint(ebrEncontrado.Size) + " bytes")
									fmt.Println("---------------------------------")
								} else {
									println(color.Red + "Lo siento incrementaste demasiado a la logica , escribiria afuera de la extendida por tanto no lo permite " + color.Reset)
								}

							}

							//----------------------- fin de la adicion
						} else {
							println(color.Red + "Esa particion no esta en este disco.. (l) " + color.Reset)
						}
					} else {
						println(color.Red + "Esa particion no esta en este disco Ni si quiera hay una Extendida" + color.Reset)
					}
				}

			}

		}

	} else {
		println(color.Red + "ERROR FALTO UN PARAMETRO OBLIGATORIO..!" + color.Reset)
	}
	//LeerBinariamenteMimbr(rutaPath)
	limpiarVariableFdisk()
}

func escribirUnEBR(archivoDisco *os.File, desde int64, objeto Ebr) {
	archivoDisco.Seek(desde, 0)
	var escritor bytes.Buffer
	binary.Write(&escritor, binary.BigEndian, &objeto)
	escribirBinariamente(archivoDisco, escritor.Bytes())
}
func escribirMBR(archivoDisco *os.File, objeto TipoMbr) {
	archivoDisco.Seek(0, 0)
	var escritor bytes.Buffer
	binary.Write(&escritor, binary.BigEndian, &objeto)
	escribirBinariamente(archivoDisco, escritor.Bytes())
}

func limpiarVariableFdisk() {
	Name_ = ""
	Path_ = ""
	Unit_k_ = "k"
	tipo_particion_ = "p"
	FIT_ = "wf"
	OPCION_DELETE_ = ""
	add_ = ""
}

func getSizeConUnidad(size int64, unit string) int64 {
	switch strings.ToLower(unit) {
	case "b":
		return size
	case "k":
		return (size * 1024)
	case "m":
		return (size * 1024 * 1024)
	default:
		return (size * 1024)
	}
}

func pausar_() {
	fmt.Println("---------------------------------")
	println(color.Blue + "--Presiona Enter para continuar--" + color.Reset)
	fmt.Println("---------------------------------")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

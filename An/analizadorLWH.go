package An

/*
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// analizador sin herramienta
func Analizar_comando(entrada string) {
	/*fmt.Println("::" + entrada)
	comandos := strings.Split(entrada, "-")
	fmt.Println(comandos)
	switch strings.ToLower(comandos[0]) {
	case "exect":
		fmt.Println("CREAR DISCO")
		for i := 0 ;
	}
}

//crack
func IniciarArchivos() {
	puntero_lector := bufio.NewReader(os.NewFile(0, "stdin"))
	for { // ciclo infinito
		var entrada string
		var bandera_todo_bien bool
		fmt.Printf(">> ")

		if entrada, bandera_todo_bien = leerLineComando(puntero_lector); bandera_todo_bien {
			entrada = strings.TrimSpace(entrada)

			if len(entrada) >= 2 && entrada[len(entrada)-1] == '*' && entrada[len(entrada)-2] == '\\' {
				entrada = QuitarSimboloNextLine(entrada)
				for {
					fmt.Print("continua esa linea>>")
					temporal := ""
					temporal, bandera_todo_bien = leerLineComando(puntero_lector)
					entrada += temporal
					entrada = strings.TrimSpace(entrada)
					if entrada[len(entrada)-1] != '*' && entrada[len(entrada)-2] != '\\' {
						break
					} else {
						entrada = QuitarSimboloNextLine(entrada)
					}
				}
			} // ACA YA TENGO LA LINEA COMPLETA QUE QUIERO MANDAR A ANALIZAR
			Analizar_comando(entrada)

		} else {
			break
		}
	}
}


func leerLineComando(puntero_lector *bufio.Reader) (string, bool) {
	salida, err := puntero_lector.ReadString('\n')
	if err != nil {
		return "", false
	}
	return salida, true
}

*/

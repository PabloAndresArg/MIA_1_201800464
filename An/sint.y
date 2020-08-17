%{
package An

import (
  "bytes"
  "fmt"
  "bufio" // para esperar una entrada 
  "os"
  "strings" // PARA HACER EL STRIM() EN LAS CADENAS 
)
/*
un archivo .y esta compuesto por 4 secciones 
- importes , uniones o declaraciones de tokenes , declaracion de gramatica , Segmento de codigo  para las funciones 
*/
%}



%union{
    NoTerminal string
    str string // DEFINO EL TIPO  DE MIS TERMINALES , EN ESTE CASO TODOS LOS QUE ESTEN EN %type<token> lo que va devolver es un tipo string 
    tokenEntero  int64
}

// tokens o terminales , doble declaracion..
%token NUMERO EXTENSION_DSK RF BF FF WF K M CHGRP GRP PAUSE COMANDO_ID R CHOWN CP DEST FIND CAT MV RM REN P MKFILE MKDIR LOGOUT ID  FILE_N EDIT MKGRP RMGRP USR  MOUNT RMDISK FLECHA PATH ADD   EXEC RUTA MKDISK SIZE NAME UNIT FDISK TYPE FIT DELETE fast full UNMOUNT MKFS  PWD RMUSR MKURS CHMOD UGO CONT
%type <str>NUMERO  EXTENSION_DSK RF BF K M FF WF CHGRP GRP PAUSE COMANDO_ID R CHOWN CP DEST FIND CAT MV RM REN P MKFILE MKDIR LOGOUT ID FILE_N EDIT MKGRP RMGRP USR MOUNT RMDISK FLECHA PATH ADD    EXEC RUTA MKDISK SIZE  NAME UNIT FDISK TYPE FIT DELETE fast full UNMOUNT MKFS  PWD RMUSR MKURS CHMOD UGO CONT
// producciones o no terminales 
%type <NoTerminal> INICIO MENU_COMANDOS CREAR_DISCO TAM
/* % = es lo mismo que %prec  , y este significa que no tienen precedencia ni asociatividad :v  */



%start INICIO

%%


INICIO: /* epsilon , gramatica decendente :D */ { }
      | EXEC '-' PATH FLECHA RUTA { leerArchivoDeEntrada($5)}
      | MENU_COMANDOS 
	  ;
	  
MENU_COMANDOS:  ID '}' {fmt.Print("JEJE")}
    |  RMDISK ':' '{' '}' {fmt.Println("produccion de una funcion... creando archivo ntt ")}
    |  MOUNT KI{ fmt.Println("MONTANDO EL YIP YIP ")}
	|  FILE_N  R {  fmt.Println(" ----OK--- ")}
	|  PAUSE { pausar_() }
	|  CREAR_DISCO
    ;
KI: RMDISK{ prob() }; 
CREAR_DISCO: MKDISK '-'SIZE FLECHA NUMERO '-' PATH FLECHA RUTA '-' NAME  FLECHA EXTENSION_DSK  { CrearDisco($5 , $9 , $13 , "M" )}
           | MKDISK '-'SIZE FLECHA NUMERO '-' PATH FLECHA RUTA '-' NAME  FLECHA EXTENSION_DSK '-' UNIT FLECHA TAM { CrearDisco($5 , $9 , $13 , $17 ) }
		   ;

 TAM: K {$$ = $1}
    | M {$$ = $1}
	;



/* TERMINA LA SECCION DE LA  GRAMATICA Y COMIENZA LA DE LAS FUNCIONES */
%%

func pausar_(){
	fmt.Println("--Presiona enter para continuar--")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func prob(){
  fmt.Print(" desde una funcion :D ")
}

func leerArchivoDeEntrada(ruta string){
	fmt.Println("							.... Analizando un archivo ...")
	fmt.Println("")
    ARCHIVO, error := os.Open(ruta)
	algo_salio_mal:= false
	if error != nil {
		fmt.Println("ERROR REPORTADO")
		algo_salio_mal = true 
	}
	if  !(algo_salio_mal) {
	yyDebug = 0
	yyErrorVerbose = true
	scanner := bufio.NewScanner(ARCHIVO)
	entrada := ""
		for scanner.Scan() {
			linea_entrada := scanner.Text()
			linea_entrada = strings.TrimSpace(linea_entrada) // 	QUITO LOS ESPACIOS A LOS LADOS 
			entrada = entrada + linea_entrada
			var listo_para_analizar bool = true
			// PREGUNTA SI TIENE UN CARACTER PARA CONTINUAR CON LA SIGUIENTE LINEA 
			if len(entrada) >= 2 && entrada[len(entrada)-1] == '*' && entrada[len(entrada)-2] == '\\' {
			listo_para_analizar = false // TENGO QUE CONCATENAR LA ENTRADA ANTERIOR CON LA LINEA ACTUAL 
			entrada = QuitarSimboloNextLine(entrada)
			}	

			if listo_para_analizar{
			fmt.Println("EJECUTANDO>> " + entrada)
			l := nuevo_lexico__(bytes.NewBufferString(entrada), os.Stdout, "file.name") // ESTA FUNCION VIENE DEL ANALIZADOR LEXICO 
			yyParse(l)
			entrada = "" //limpio la entrada 
			}
		}
	}
	fmt.Println("")
	fmt.Println("							...Archivo terminado de analizar...")
	fmt.Println("")
}


func AnalizarComando() {
	puntero_lector := bufio.NewReader(os.NewFile(0, "stdin"))
	yyDebug = 0
	yyErrorVerbose = true
	for { // ciclo infinito 
		var entrada string
		var bandera_todo_bien bool

		fmt.Printf(">> ")
		if entrada, bandera_todo_bien = leerLineComando(puntero_lector); bandera_todo_bien {
			entrada= strings.TrimSpace(entrada)

			if len(entrada) >= 2 && entrada[len(entrada)-1] == '*' && entrada[len(entrada)-2] == '\\' {
			entrada = QuitarSimboloNextLine(entrada)
			for {
				fmt.Print("continua esa linea>>")
				temporal:= ""
				temporal, bandera_todo_bien = leerLineComando(puntero_lector)
				entrada += temporal
				entrada = strings.TrimSpace(entrada)
				if entrada[len(entrada)-1] != '*' && entrada[len(entrada)-2] != '\\' {
					break
				}else{
					entrada = QuitarSimboloNextLine(entrada)	
				}
			}	

				
			}
			l := nuevo_lexico__(bytes.NewBufferString(entrada), os.Stdout, "file.name") // ESTA FUNCION VIENE DEL ANALIZADOR LEXICO 
			yyParse(l)
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


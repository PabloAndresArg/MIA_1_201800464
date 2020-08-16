%{
package An

import (
  "bytes"
  "fmt"
  "bufio" // para esperar una entrada 
  "os"
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
%token RF CHGRP GRP PAUSE COMANDO_ID R CHOWN CP DEST FIND CAT MV RM REN P MKFILE MKDIR LOGOUT ID  FILE_N EDIT MKGRP RMGRP USR  MOUNT RMDISK FLECHA PATH ADD  NUMERO EXEC RUTA MKDISK SIZE NAME UNIT FDISK TYPE FIT DELETE fast full UNMOUNT MKFS  PWD RMUSR MKURS CHMOD UGO CONT
%type <str> RF CHGRP GRP PAUSE COMANDO_ID R CHOWN CP DEST FIND CAT MV RM REN P MKFILE MKDIR LOGOUT ID FILE_N EDIT MKGRP RMGRP USR MOUNT RMDISK FLECHA PATH ADD   NUMERO EXEC RUTA MKDISK SIZE  NAME UNIT FDISK TYPE FIT DELETE fast full UNMOUNT MKFS  PWD RMUSR MKURS CHMOD UGO CONT
// producciones o no terminales 
%type <NoTerminal> INICIO MENU_COMANDOS 
/* % = es lo mismo que %prec  , y este significa que no tienen precedencia ni asociatividad :v  */

%start INICIO

%%


INICIO: /* epsilon , gramatica decendente :D */ { }
      | EXEC '-' PATH FLECHA RUTA { leerArchivoDeEntrada($5)}
      | MENU_COMANDOS  { fmt.Println("menu") }
      ;

MENU_COMANDOS:  ID '}' {fmt.Print("JEJE")}
    |  RMDISK ':' '{' '}' {fmt.Println("produccion de una funcion... creando archivo ntt ")}
    |  MOUNT KI{ fmt.Println("MONTANDO EL YIP YIP ")}
	|  FILE_N  R {  fmt.Println(" ----OK--- ")}
	|  PAUSE { pausar_() }
    ;
KI: RMDISK{ prob() }
  ;





/* TERMINA LA SECCION DE LA  GRAMATICA Y COMIENZA LA DE LAS FUNCIONES */
%%

func pausar_(){
	fmt.Println("Presiona enter para continuar")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func prob(){
  fmt.Print(" desde una funcion :D ")
}

func leerArchivoDeEntrada(ruta string){

	fmt.Println(" EJECUTO LA FUNCION PARA LEER UN ARCHIVO DE UNA :D ")

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
		for scanner.Scan() {
			linea_entrada := scanner.Text()
			l := nuevo_lexico__(bytes.NewBufferString(linea_entrada), os.Stdout, "file.name") // ESTA FUNCION VIENE DEL ANALIZADOR LEXICO 
			yyParse(l)
			
		}
	}
	fmt.Println("...Archivo terminado de analizar...")
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


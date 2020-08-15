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
%token COMANDO_ID R CAT RM P MKFILE LOGOUT ID  FILE_N MKGRP RMGRP USR  MOUNT RMDISK FLECHA PATH ADD  NUMERO EXEC RUTA MKDISK SIZE NAME UNIT FDISK TYPE FIT DELETE fast full UNMOUNT MKFS  PWD RMUSR MKURS CHMOD UGO CONT
%type <str> COMANDO_ID R CAT RM P MKFILE LOGOUT ID FILE_N MKGRP RMGRP USR MOUNT RMDISK FLECHA PATH ADD   NUMERO EXEC RUTA MKDISK SIZE  NAME UNIT FDISK TYPE FIT DELETE fast full UNMOUNT MKFS  PWD RMUSR MKURS CHMOD UGO CONT
// producciones o no terminales 
%type <NoTerminal> INICIO MENU_COMANDOS 
/* % = es lo mismo que %prec  , y este significa que no tienen precedencia ni asociatividad :v  */

%start INICIO

%%


INICIO: /* epsilon , gramatica decendente :D */ { }
     | EXEC '-' PATH FLECHA RUTA { leerArchivoDeEntrada($5)}
     | MENU_COMANDOS  {fmt.Println("menu")}
     ;

MENU_COMANDOS:  ID '}' {fmt.Print("JEJE")}
    |  RMDISK ':' '{' '}' {fmt.Println("produccion de una funcion... creando archivo ntt ")}
    |  MOUNT KI{ fmt.Println("MONTANDO EL YIP YIP ")}
	|  FILE_N  ';' {  fmt.Println("----OK---")}
    ;
KI: RMDISK{ prob() }




/* TERMINA LA SECCION DE LA  GRAMATICA Y COMIENZA LA DE LAS FUNCIONES */
%%



func prob(){
  fmt.Print(" desde una funcion :D ")
}

func leerArchivoDeEntrada(entrada string){
fmt.Println(" EJECUTO LA FUNCION PARA LEER UN ARCHIVO DE UNA :D ")
 fmt.Println("A LEER: "+ entrada)
}

func AnalizarComando() {
	fi := bufio.NewReader(os.NewFile(0, "stdin"))
	yyDebug = 0
	yyErrorVerbose = true
	for { // ciclo infinito 
		var entrada string
		var bandera_todo_bien bool

		fmt.Printf("Ingrese el comando: ")
		if entrada, bandera_todo_bien = leerLineComando(fi); bandera_todo_bien {
			l := nuevo_lexico__(bytes.NewBufferString(entrada), os.Stdout, "file.name") // ESTA FUNCION VIENE DEL ANALIZADOR LEXICO 
			yyParse(l)
		} else {
			break
		}
	}

}


func leerLineComando(fi *bufio.Reader) (string, bool) {
	s, err := fi.ReadString('\n')
	if err != nil {
		return "", false
	}
	return s, true
}


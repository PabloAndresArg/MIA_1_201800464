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
%token NUMERO EXTENSION_DSK B E L  RF BF FF WF K M CHGRP GRP PAUSE COMANDO_ID R CHOWN CP DEST FIND CAT MV RM REN P MKFILE MKDIR LOGOUT ID  FILE_N EDIT MKGRP RMGRP USR  MOUNT RMDISK FLECHA PATH ADD   EXEC RUTA MKDISK SIZE NAME UNIT FDISK TYPE FIT DELETE fast full UNMOUNT MKFS  PWD RMUSR MKURS CHMOD UGO CONT
%type <str>NUMERO  EXTENSION_DSK B E L  RF BF K M FF WF CHGRP GRP PAUSE COMANDO_ID R CHOWN CP DEST FIND CAT MV RM REN P MKFILE MKDIR LOGOUT ID FILE_N EDIT MKGRP RMGRP USR MOUNT RMDISK FLECHA PATH ADD    EXEC RUTA MKDISK SIZE  NAME UNIT FDISK TYPE FIT DELETE fast full UNMOUNT MKFS  PWD RMUSR MKURS CHMOD UGO CONT
// producciones o no terminales 
%type <NoTerminal> INICIO OPCIONES_FIT MONTAR  MENU_COMANDOS CREAR_DISCO TAM ELIMINAR_DISCO PARAMETROS_MKDISK P_MKDISK ADMINISTRAR_PARTICIONES TAM2 TYPE_PARTICION OPCIONES_DELETE
/* % = es lo mismo que %prec  , y este significa que no tienen precedencia ni asociatividad :v  */



%start INICIO

%%


INICIO: /* epsilon , gramatica decendente :D */ { }
      | EXEC '-' PATH FLECHA RUTA { leerArchivoDeEntrada($5)}
      | MENU_COMANDOS 
	  ;
	  
MENU_COMANDOS:  CREAR_DISCO
   		     |  ELIMINAR_DISCO
			 |  ADMINISTRAR_PARTICIONES
			 |  MONTAR
   	         ;

MONTAR: MOUNT PARAMETROS_MONTAR PARAMETROS_MONTAR { fmt.Println("INSTRUCCION"); fmt.Println("MONTAR")}
      ;
PARAMETROS_MONTAR:'-' PATH FLECHA RUTA  { Path_ = $4 }
		         |'-' NAME  FLECHA ID { Name_ = $4}
				 ;

ELIMINAR_DISCO: RMDISK '-' PATH FLECHA RUTA { EliminarDisco($5) } ;

CREAR_DISCO: MKDISK PARAMETROS_MKDISK { CrearDisco(Size_ , Path_ , Name_, Unit_m_ )}; 
PARAMETROS_MKDISK: PARAMETROS_MKDISK  P_MKDISK 
				 | P_MKDISK
          		 ; 

P_MKDISK:'-' SIZE FLECHA NUMERO { Size_ = $4 }
		|'-' PATH FLECHA RUTA  { Path_ = $4 }
		|'-' NAME  FLECHA EXTENSION_DSK { Name_ = $4}
		|'-' UNIT FLECHA TAM { Unit_m_ = $4 }
		;



TAM: K {$$ = $1}
   | M {$$ = $1}
   ;

TAM2: K {$$ = $1}
    | M {$$ = $1}
	| B {$$ = $1}
	;

ADMINISTRAR_PARTICIONES: FDISK PARAMETROS_FDISK { /*LeerBinariamente_el_disco("/home/pablo/Escritorio/disco.dsk")*/ };
PARAMETROS_FDISK: PARAMETROS_FDISK  P_FDISK 
				 | P_FDISK
				 ; 
P_FDISK: '-' SIZE FLECHA NUMERO { Size_ = $4 }
		|'-' PATH FLECHA RUTA  { Path_ = $4 }
		|'-' NAME  FLECHA ID { Name_ = $4}
		|'-' UNIT FLECHA TAM2 { Unit_k_ = $4 }
		|'-' TYPE FLECHA TYPE_PARTICION {  tipo_particion_ = $4 }
		|'-' FIT FLECHA OPCIONES_FIT { FIT_ = $4}
		|'-' DELETE FLECHA OPCIONES_DELETE { OPCION_DELETE_ = $4 }
		|'-' ADD FLECHA NUMERO { add_ = $4  }
		; 

OPCIONES_DELETE: fast  {$$=$1}
			   | full  {$$=$1}
			   ;

OPCIONES_FIT: BF {$$ = $1}
		    | FF {$$ = $1}
			| WF {$$ = $1}
			;

TYPE_PARTICION: P {$$ = $1}
			  | E {$$ = $1}
			  | L {$$ = $1}
			  ;






/* TERMINA LA SECCION DE LA  GRAMATICA Y COMIENZA LA DE LAS FUNCIONES */
%%

func pausar_(){
	fmt.Println("--Presiona enter para continuar--")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
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


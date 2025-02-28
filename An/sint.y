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
%token UNMOUNT ID_N NUMERO REP CM_RUTA NUMERO_NEGATIVO VD_ID_NUMERO EXTENSION_DSK B E L  RF BF FF WF K M CHGRP GRP PAUSE COMANDO_ID R CHOWN CP DEST FIND CAT MV RM REN P MKFILE MKDIR LOGOUT ID  FILE_N EDIT MKGRP RMGRP USR  MOUNT RMDISK FLECHA PATH ADD   EXEC RUTA MKDISK SIZE NAME UNIT FDISK TYPE FIT DELETE fast full  MKFS  PWD RMUSR MKURS CHMOD UGO CONT
%type <str> UNMOUNT  ID_N NUMERO CM_RUTA NUMERO_NEGATIVO REP VD_ID_NUMERO EXTENSION_DSK B E L  RF BF K M FF WF CHGRP GRP PAUSE COMANDO_ID R CHOWN CP DEST FIND CAT MV RM REN P MKFILE MKDIR LOGOUT ID FILE_N EDIT MKGRP RMGRP USR MOUNT RMDISK FLECHA PATH ADD    EXEC RUTA MKDISK SIZE  NAME UNIT FDISK TYPE FIT DELETE fast full  MKFS  PWD RMUSR MKURS CHMOD UGO CONT
// producciones o no terminales 
%type <NoTerminal> FORMATEAR_DISCO P_MKFS DESMONTAR  PARAMETROS_REPORTES  P_FDISK REPORTES INICIO OPCIONES_FIT MONTAR CADENA_O_ID  MENU_COMANDOS CREAR_DISCO TAM ELIMINAR_DISCO PARAMETROS_MKDISK P_MKDISK ADMINISTRAR_PARTICIONES TAM2 TYPE_PARTICION OPCIONES_DELETE P_REPORTE
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
			 |  PAUSE {pausar_()}
			 |  REPORTES
			 |  DESMONTAR
			 |  FORMATEAR_DISCO
   	         ;


DESMONTAR: UNMOUNT PARAMETROS_UNMOUNT {/*NO HACE NADA PORQUE LO HAGO CADA QUE RECIBE UNA PRODUCCION ABAJO :v */};
PARAMETROS_UNMOUNT: PARAMETROS_UNMOUNT P_UNMOUNT
				  | P_UNMOUNT
				  ; 
P_UNMOUNT:  ID_N FLECHA VD_ID_NUMERO {desmontar($3)};


MONTAR: MOUNT PARAMETROS_MONTAR PARAMETROS_MONTAR {crearMontaje(QuitarComillas(Path_) ,QuitarComillas(Name_))}
      | MOUNT { mostrarMounts() }
	  ;
PARAMETROS_MONTAR:'-' PATH FLECHA RUTA  { Path_ = $4 }
		         |'-' NAME  FLECHA CADENA_O_ID { Name_ = $4}
				 ;
FORMATEAR_DISCO: MKFS PARAMETROS_MKFS { metodoMKFS( Id_vdlentraNumero_ , type_ , add_ , Unit_k_ ) }
				;
PARAMETROS_MKFS: PARAMETROS_MKFS P_MKFS
				| P_MKFS
				;
P_MKFS:'-' TYPE FLECHA OPCIONES_DELETE { type_ = $4 }
	  |'-' ADD FLECHA NUMERO { add_ = $4  }
	  |'-' ADD FLECHA NUMERO_NEGATIVO { add_ = $4}
	  |'-' UNIT FLECHA TAM2 { Unit_k_ = $4 }
	  |'-' COMANDO_ID FLECHA VD_ID_NUMERO { Id_vdlentraNumero_ = $4}
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


ADMINISTRAR_PARTICIONES: FDISK PARAMETROS_FDISK { MetodosParticiones(Path_ , Name_ , Size_ , FIT_ , OPCION_DELETE_ , add_ , tipo_particion_ , Unit_k_)}	;


PARAMETROS_FDISK: PARAMETROS_FDISK  P_FDISK 
				 | P_FDISK
				 ; 
P_FDISK: '-' SIZE FLECHA NUMERO { Size_ = $4 }
		|'-' PATH FLECHA RUTA  { Path_ = $4 }
		|'-' NAME  FLECHA CADENA_O_ID { Name_ = $4}
		|'-' UNIT FLECHA TAM2 { Unit_k_ = $4 }
		|'-' TYPE FLECHA TYPE_PARTICION {  tipo_particion_ = $4 }
		|'-' FIT FLECHA OPCIONES_FIT { FIT_ = $4}
		|'-' DELETE FLECHA OPCIONES_DELETE { OPCION_DELETE_ = $4 }
		|'-' ADD FLECHA NUMERO { add_ = $4  }
		|'-' ADD FLECHA NUMERO_NEGATIVO { add_ = $4}
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


CADENA_O_ID: RUTA { $$ = QuitarComillas($1) }
		   | ID   { $$ = QuitarComillas($1) }
		   ;
REPORTES: REP PARAMETROS_REPORTES {generarReporte()}; 

PARAMETROS_REPORTES : PARAMETROS_REPORTES P_REPORTE
                    | P_REPORTE
					;
P_REPORTE: '-' NAME FLECHA  ID { Name_ = $4}
		 | '-' PATH FLECHA RUTA     { Path_= QuitarComillas($4)}
		 | '-' COMANDO_ID FLECHA VD_ID_NUMERO { Id_vdlentraNumero_ = $4}
		 | '-' CM_RUTA FLECHA    ID { Commando_Ruta_ = $4 /*aun no se usara*/}
		 ;


/* TERMINA LA SECCION DE LA  GRAMATICA Y COMIENZA LA DE LAS FUNCIONES */
%%







func leerArchivoDeEntrada(ruta string){
	fmt.Println("");fmt.Println("")
	fmt.Println("							||||||||||||||||||||||||||||||||||")
	fmt.Println("							|||||| Analizando un archivo |||||")
    fmt.Println("							||||||||||||||||||||||||||||||||||")
	fmt.Println("")
    ARCHIVO, error := os.Open(QuitarComillas(ruta))
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
			l := nuevo_lexico__(bytes.NewBufferString(entrada), os.Stdout, "archivo.PabloAndres") // ESTA FUNCION VIENE DEL ANALIZADOR LEXICO 
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


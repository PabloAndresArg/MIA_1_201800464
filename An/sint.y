%{
package An

import (
  "fmt"
  "bytes"
  "io"
  "bufio"
	"os"
)

type node struct {
  name string
  children []node
}

func (n node) String() string {
  buf := new(bytes.Buffer)
  n.print(buf, " ")
  return buf.String()
}

func (n node) print(out io.Writer, indent string) {
  fmt.Fprintf(out, "\n%v%v", indent, n.name)
  for _, nn := range n.children { nn.print(out, indent + "  ") }
}

func Node(name string) node { return node{name: name} }
func (n node) append(nn...node) node { n.children = append(n.children, nn...); return n }




%}

%union{
    node node
    token string
}

// tokens o terminales , doble declaracion..
%token  ID   MOUNT SKDIR FLECHA PATH AND INORMAL NUMERO EXEC RUTA MKDISK SIZE
%type <token>  ID  MOUNT SKDIR FLECHA PATH AND  INORMAL NUMERO EXEC RUTA MKDISK SIZE 
// producciones o no terminales 
%type <node> INICIO MENU_COMANDOS 

%%
INICIO: /* epsilon , gramatica decendente :D */ { }
     | EXEC '-' PATH FLECHA RUTA { leerArchivoDeEntrada(string($5))}
     | MENU_COMANDOS  {fmt.Println("menu")}
     ;
//DIGAMOS AQUI LO QUE HACEMOS ES QUE TIENE QUE RECONOCER int InT, FlOat, CHAR,Char, no importa porque en el .l le agrege opcion de case insentive 
MENU_COMANDOS:  ID '}' {$$ = Node("identifacador")}
    |  SKDIR ':' '{' '}' {fmt.Println("produccion de una funcion... creando archivo ntt ")}
    |  MOUNT KI{ fmt.Println("MONTANDO EL YIP YIP ")}
    ;
KI: SKDIR{ skdir_fun() }

%% 
func skdir_fun(){
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
	for {
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


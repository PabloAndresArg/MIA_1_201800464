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
%token  ID   MOUNT SKDIR FLECHA PATH AND INORMAL NUMERO
%type <token>  ID  MOUNT SKDIR FLECHA PATH AND  INORMAL NUMERO 
// producciones o no terminales 
%type <node> INICIO PRODUC 

%%
INICIO: /* epsilon , gramatica decendente :D */ { }
     | PRODUC FLECHA {fmt.Println($2)}
     | PATH FLECHA {fmt.Println("OK PATH CON FLECHITA")}
     | AND {fmt.Println("UN AND LOGICO EN PRODUCCION")}
     | '&' NUMERO {fmt.Println(" LA Y   NORMAL  ademas tiene un numero :D")}
     ;
//DIGAMOS AQUI LO QUE HACEMOS ES QUE TIENE QUE RECONOCER int InT, FlOat, CHAR,Char, no importa porque en el .l le agrege opcion de case insentive 
PRODUC:  ID '}' {$$ = Node("identifacador")}
    |  SKDIR ':' '{' '}' {fmt.Println("produccion de una funcion... creando archivo ntt ")}
    |  MOUNT KI{ fmt.Println("MONTANDO EL YIP YIP ")}
    ;
KI: SKDIR{ skdir_fun() }

%% 
func skdir_fun(){
  fmt.Print(" desde una funcion :D ")
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


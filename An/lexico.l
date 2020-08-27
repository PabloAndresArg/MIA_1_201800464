%{

package An

import (
  "modernc.org/golex/lex"
  "bufio"
  "go/token"
  "unicode"
  "io"
)

func nuevo_lexico__(Fuente_ io.Reader, destino_ io.Writer, fName string) *lexer { 
  file := token.NewFileSet().AddFile(fName, -1, 1<<31-1)
  lexic_, err := lex.New(file, bufio.NewReader(Fuente_), lex.RuneClass(miniLexico))
  if (err != nil) { panic(err) }
  return &lexer{lexic_}
  // necesaria para el sintactico :D 
}


type lexer struct {
	*lex.Lexer
}
// propio de lex 
const (
	classUnicodeLeter = iota + 0x80
	classUnicodeDigit
	classOther
)
func miniLexico(caracter rune) int { // necesario para hacer el nuevo_lex
	if caracter >= 0 && caracter < 0x80 { 
		return int(caracter)
	}
	if unicode.IsLetter(caracter) {
		return classUnicodeLeter
	}
	if unicode.IsDigit(caracter) {
		return classUnicodeDigit
	}
	return classOther
}



func (l *lexer) Lex( yylval *yySymType) int {
  c := l.Enter()
%}


%option case-insensitive
%yyc c
%yyn c = l.Next()
%yym l.Mark()

utfDigito  \x81 
utfLetra \x80
digit         [0-9]+|{utfDigito}+
id    {letter}({letter}|{digit})*
letter        [a-zA-Z_]|{utfLetra}
flecha 		   (\-\>) | (\-\>)
ruta			(((\/)({letter})({letter}|{digit}|\.|(\/))+)|{cadenaLiteral})
cadenaLiteral     (\"[^"]*\")
comentario		  \#(.)*
file_n 				[f][i][l][e]([0-9])+
extension_dsk   	{id}(\.)dsk
continuacio_de_linea  \\\*
vd_ID_NUMERO       [v][d][a-zA-Z_][0-9]+
%%
	c = l.Rule0()

[ \t\r\n]+
{continuacio_de_linea} 
{comentario} yylval.str = string(l.TokenBytes(nil)); 
{flecha}  	 yylval.str = string(l.TokenBytes(nil)); return FLECHA //  yylval ES MI OBJETO %union que esta en el sint.y
{ruta}		 yylval.str = string(l.TokenBytes(nil)); return RUTA  // NO ES LO MISMO literal pero  los atributos del union los tiene el yyval , el yyval puede tener el nombre que queramos solo lo deje asi por lo que decia le manual xd 
mount 		 yylval.str = string(l.TokenBytes(nil)); return MOUNT
path 		 yylval.str = string(l.TokenBytes(nil)); return PATH 
rmdisk 		 yylval.str = string(l.TokenBytes(nil)); return RMDISK
exec		 yylval.str = string(l.TokenBytes(nil)); return EXEC
mkdisk		 yylval.str = string(l.TokenBytes(nil)); return MKDISK
size		 yylval.str = string(l.TokenBytes(nil)); return SIZE
name		 yylval.str = string(l.TokenBytes(nil)); return NAME
unit		 yylval.str = string(l.TokenBytes(nil)); return UNIT
fdisk		 yylval.str = string(l.TokenBytes(nil)); return FDISK
type		 yylval.str = string(l.TokenBytes(nil)); return TYPE
fit          yylval.str = string(l.TokenBytes(nil)); return FIT
delete 		 yylval.str = string(l.TokenBytes(nil)); return DELETE 
add          yylval.str = string(l.TokenBytes(nil)); return ADD
fast		 yylval.str = string(l.TokenBytes(nil)); return fast // TAMBIEN SE PUEDE EN MINUSCULAS	
full		 yylval.str = string(l.TokenBytes(nil)); return full
unmount		 yylval.str = string(l.TokenBytes(nil)); return UNMOUNT
mkfs		 yylval.str = string(l.TokenBytes(nil)); return MKFS
id 		     yylval.str = string(l.TokenBytes(nil)); return COMANDO_ID
usr 		 yylval.str = string(l.TokenBytes(nil)); return USR
pwd		     yylval.str = string(l.TokenBytes(nil)); return PWD 
logout		 yylval.str = string(l.TokenBytes(nil)); return LOGOUT
mkgrp		 yylval.str = string(l.TokenBytes(nil)); return MKGRP
mkurs		 yylval.str = string(l.TokenBytes(nil)); return MKURS
rmusr		 yylval.str = string(l.TokenBytes(nil)); return RMUSR
chmod		 yylval.str = string(l.TokenBytes(nil)); return CHMOD
ugo			 yylval.str = string(l.TokenBytes(nil)); return UGO 
rmgrp		 yylval.str = string(l.TokenBytes(nil)); return RMGRP // PAGINA 14 DEL ENCUNCIADO :'V '
cont		 yylval.str = string(l.TokenBytes(nil)); return CONT
mkfile       yylval.str = string(l.TokenBytes(nil)); return MKFILE
cat 		 yylval.str = string(l.TokenBytes(nil)); return CAT
rm           yylval.str = string(l.TokenBytes(nil)); return RM 
rf           yylval.str = string(l.TokenBytes(nil)); return RF // PARAMETRO 
edit	     yylval.str = string(l.TokenBytes(nil)); return EDIT 
ren 		 yylval.str = string(l.TokenBytes(nil)); return REN // pag 20 
mkdir		 yylval.str = string(l.TokenBytes(nil)); return MKDIR
{file_n} 	 yylval.str = string(l.TokenBytes(nil)); return FILE_N
R 			 yylval.str = string(l.TokenBytes(nil)); return R // OPCIONALES DENTRO DE LA GRAMATICA 
P			 yylval.str = string(l.TokenBytes(nil)); return P 
e 			 yylval.str = string(l.TokenBytes(nil)); return E
l			 yylval.str = string(l.TokenBytes(nil)); return L 
k			 yylval.str = string(l.TokenBytes(nil)); return K 
m	 		 yylval.str = string(l.TokenBytes(nil)); return M   
b			 yylval.str = string(l.TokenBytes(nil)); return B
cp 		     yylval.str = string(l.TokenBytes(nil)); return CP 
dest         yylval.str = string(l.TokenBytes(nil)); return DEST
find         yylval.str = string(l.TokenBytes(nil)); return FIND
chown        yylval.str = string(l.TokenBytes(nil)); return CHOWN
chgrp		 yylval.str = string(l.TokenBytes(nil)); return CHGRP 	
mv			 yylval.str = string(l.TokenBytes(nil)); return MV
grp			 yylval.str = string(l.TokenBytes(nil)); return GRP
pause        yylval.str = string(l.TokenBytes(nil)); return PAUSE
bf			 yylval.str = string(l.TokenBytes(nil)); return BF
ff			 yylval.str = string(l.TokenBytes(nil)); return FF
wf 		 	 yylval.str = string(l.TokenBytes(nil)); return WF // opcionales para el fit pero mmm quedaria mejor en la gramatica con comillas simples 
nombre  	 yylval.str = string(l.TokenBytes(nil)); return CM_NOMBRE
ruta yylval.str = string(l.TokenBytes(nil));return CM_RUTA 
rep			 yylval.str = string(l.TokenBytes(nil)); return REP 
{vd_ID_NUMERO} yylval.str = string(l.TokenBytes(nil)); return VD_ID_NUMERO
{extension_dsk} yylval.str = string(l.TokenBytes(nil));return EXTENSION_DSK

{id}         yylval.str = string(l.TokenBytes(nil)); return ID 
{digit}      yylval.str = string(l.TokenBytes(nil)); return NUMERO // strconv.Atoi("1")

%%
	if u_u, nitido := l.Abort(); nitido { return int(u_u) }
	goto yyAction
}
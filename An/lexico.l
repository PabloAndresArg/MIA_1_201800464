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
letter        [_a-zA-Z]|{utfLetra}
flecha 		   (\-\>) | (\-\>)
ruta			(((\/)({letter})({letter}|{digit}|\.|(\/))+)|{cadenaLiteral})
cadenaLiteral     (\"[^"]*\")

%%
	c = l.Rule0()

[ \t\r\n]+
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
{id}         yylval.str = string(l.TokenBytes(nil)); return ID 
{digit}      yylval.str = string(l.TokenBytes(nil)); return NUMERO

%%
	if u_u, nitido := l.Abort(); nitido { return int(u_u) }
	goto yyAction
}
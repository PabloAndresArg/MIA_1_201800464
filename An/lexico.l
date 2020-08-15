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
func miniLexico(caracter rune) int {
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



func (l *lexer) Lex(lval *yySymType) int {
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
{flecha}  	lval.token = string(l.TokenBytes(nil)); return FLECHA
{ruta}		lval.token = string(l.TokenBytes(nil)); return RUTA
mount 		lval.token = string(l.TokenBytes(nil)); return MOUNT
path 		lval.token = string(l.TokenBytes(nil)); return PATH 
skdir 		lval.token= string(l.TokenBytes(nil)); return SKDIR
exec		lval.token = string(l.TokenBytes(nil)); return EXEC
mkdisk		lval.token = string(l.TokenBytes(nil)); return MKDISK
size		lval.token = string(l.TokenBytes(nil)); return SIZE
{id} lval.token = string(l.TokenBytes(nil)); return ID 
{digit}  lval.token = string(l.TokenBytes(nil)); return NUMERO

%%
	if u_u, nitido := l.Abort(); nitido { return int(u_u) }
	goto yyAction
}
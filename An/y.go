// Code generated by goyacc sint.y. DO NOT EDIT.

//line sint.y:2
package An

import __yyfmt__ "fmt"

//line sint.y:2

import (
	"bufio" // para esperar una entrada
	"bytes"
	"fmt"
	"os"
)

/*
un archivo .y esta compuesto por 4 secciones
- importes , uniones o declaraciones de tokenes , declaracion de gramatica , Segmento de codigo  para las funciones
*/

//line sint.y:18
type yySymType struct {
	yys         int
	NoTerminal  string
	str         string // DEFINO EL TIPO  DE MIS TERMINALES , EN ESTE CASO TODOS LOS QUE ESTEN EN %type<token> lo que va devolver es un tipo string
	tokenEntero int64
}

const COMANDO_ID = 57346
const R = 57347
const CAT = 57348
const P = 57349
const MKFILE = 57350
const LOGOUT = 57351
const ID = 57352
const FILE_N = 57353
const MKGRP = 57354
const RMGRP = 57355
const USR = 57356
const MOUNT = 57357
const RMDISK = 57358
const FLECHA = 57359
const PATH = 57360
const ADD = 57361
const NUMERO = 57362
const EXEC = 57363
const RUTA = 57364
const MKDISK = 57365
const SIZE = 57366
const NAME = 57367
const UNIT = 57368
const FDISK = 57369
const TYPE = 57370
const FIT = 57371
const DELETE = 57372
const fast = 57373
const full = 57374
const UNMOUNT = 57375
const MKFS = 57376
const PWD = 57377
const RMUSR = 57378
const MKURS = 57379
const CHMOD = 57380
const UGO = 57381
const CONT = 57382

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"COMANDO_ID",
	"R",
	"CAT",
	"P",
	"MKFILE",
	"LOGOUT",
	"ID",
	"FILE_N",
	"MKGRP",
	"RMGRP",
	"USR",
	"MOUNT",
	"RMDISK",
	"FLECHA",
	"PATH",
	"ADD",
	"NUMERO",
	"EXEC",
	"RUTA",
	"MKDISK",
	"SIZE",
	"NAME",
	"UNIT",
	"FDISK",
	"TYPE",
	"FIT",
	"DELETE",
	"fast",
	"full",
	"UNMOUNT",
	"MKFS",
	"PWD",
	"RMUSR",
	"MKURS",
	"CHMOD",
	"UGO",
	"CONT",
	"'-'",
	"'}'",
	"':'",
	"'{'",
	"';'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line sint.y:52

func prob() {
	fmt.Print(" desde una funcion :D ")
}

func leerArchivoDeEntrada(entrada string) {
	fmt.Println(" EJECUTO LA FUNCION PARA LEER UN ARCHIVO DE UNA :D ")
	fmt.Println("A LEER: " + entrada)
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

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 18

var yyAct = [...]int{
	13, 15, 10, 17, 9, 8, 4, 7, 18, 14,
	16, 6, 5, 12, 11, 3, 1, 2,
}

var yyPact = [...]int{
	-4, -1000, -36, -1000, -38, -41, -3, -45, -9, -1000,
	-43, -1000, -1000, -1000, -7, -39, -14, -1000, -1000,
}

var yyPgo = [...]int{
	0, 16, 15, 14,
}

var yyR1 = [...]int{
	0, 1, 1, 1, 2, 2, 2, 2, 3,
}

var yyR2 = [...]int{
	0, 0, 5, 1, 2, 4, 2, 2, 1,
}

var yyChk = [...]int{
	-1000, -1, 21, -2, 10, 16, 15, 11, 41, 42,
	43, -3, 16, 45, 18, 44, 17, 42, 22,
}

var yyDef = [...]int{
	1, -2, 0, 3, 0, 0, 0, 0, 0, 4,
	0, 6, 8, 7, 0, 0, 0, 5, 2,
}

var yyTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 41, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 43, 45,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 44, 3, 42,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40,
}

var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-0 : yypt+1]
//line sint.y:36
		{
		}
	case 2:
		yyDollar = yyS[yypt-5 : yypt+1]
//line sint.y:37
		{
			leerArchivoDeEntrada(yyDollar[5].str)
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line sint.y:38
		{
			fmt.Println("menu")
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
//line sint.y:41
		{
			fmt.Print("JEJE")
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
//line sint.y:42
		{
			fmt.Println("produccion de una funcion... creando archivo ntt ")
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line sint.y:43
		{
			fmt.Println("MONTANDO EL YIP YIP ")
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
//line sint.y:44
		{
			fmt.Println("----OK---")
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line sint.y:46
		{
			prob()
		}
	}
	goto yystack /* stack new state and value */
}

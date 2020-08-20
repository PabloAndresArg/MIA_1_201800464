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
	"strings" // PARA HACER EL STRIM() EN LAS CADENAS
)

/*
un archivo .y esta compuesto por 4 secciones
- importes , uniones o declaraciones de tokenes , declaracion de gramatica , Segmento de codigo  para las funciones
*/

//line sint.y:19
type yySymType struct {
	yys         int
	NoTerminal  string
	str         string // DEFINO EL TIPO  DE MIS TERMINALES , EN ESTE CASO TODOS LOS QUE ESTEN EN %type<token> lo que va devolver es un tipo string
	tokenEntero int64
}

const NUMERO = 57346
const EXTENSION_DSK = 57347
const RF = 57348
const BF = 57349
const FF = 57350
const WF = 57351
const K = 57352
const M = 57353
const CHGRP = 57354
const GRP = 57355
const PAUSE = 57356
const COMANDO_ID = 57357
const R = 57358
const CHOWN = 57359
const CP = 57360
const DEST = 57361
const FIND = 57362
const CAT = 57363
const MV = 57364
const RM = 57365
const REN = 57366
const P = 57367
const MKFILE = 57368
const MKDIR = 57369
const LOGOUT = 57370
const ID = 57371
const FILE_N = 57372
const EDIT = 57373
const MKGRP = 57374
const RMGRP = 57375
const USR = 57376
const MOUNT = 57377
const RMDISK = 57378
const FLECHA = 57379
const PATH = 57380
const ADD = 57381
const EXEC = 57382
const RUTA = 57383
const MKDISK = 57384
const SIZE = 57385
const NAME = 57386
const UNIT = 57387
const FDISK = 57388
const TYPE = 57389
const FIT = 57390
const DELETE = 57391
const fast = 57392
const full = 57393
const UNMOUNT = 57394
const MKFS = 57395
const PWD = 57396
const RMUSR = 57397
const MKURS = 57398
const CHMOD = 57399
const UGO = 57400
const CONT = 57401

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMERO",
	"EXTENSION_DSK",
	"RF",
	"BF",
	"FF",
	"WF",
	"K",
	"M",
	"CHGRP",
	"GRP",
	"PAUSE",
	"COMANDO_ID",
	"R",
	"CHOWN",
	"CP",
	"DEST",
	"FIND",
	"CAT",
	"MV",
	"RM",
	"REN",
	"P",
	"MKFILE",
	"MKDIR",
	"LOGOUT",
	"ID",
	"FILE_N",
	"EDIT",
	"MKGRP",
	"RMGRP",
	"USR",
	"MOUNT",
	"RMDISK",
	"FLECHA",
	"PATH",
	"ADD",
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
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line sint.y:76

func pausar_() {
	fmt.Println("--Presiona enter para continuar--")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func leerArchivoDeEntrada(ruta string) {
	fmt.Println("							.... Analizando un archivo ...")
	fmt.Println("")
	ARCHIVO, error := os.Open(ruta)
	algo_salio_mal := false
	if error != nil {
		fmt.Println("ERROR REPORTADO")
		algo_salio_mal = true
	}
	if !(algo_salio_mal) {
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

			if listo_para_analizar {
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
			entrada = strings.TrimSpace(entrada)

			if len(entrada) >= 2 && entrada[len(entrada)-1] == '*' && entrada[len(entrada)-2] == '\\' {
				entrada = QuitarSimboloNextLine(entrada)
				for {
					fmt.Print("continua esa linea>>")
					temporal := ""
					temporal, bandera_todo_bien = leerLineComando(puntero_lector)
					entrada += temporal
					entrada = strings.TrimSpace(entrada)
					if entrada[len(entrada)-1] != '*' && entrada[len(entrada)-2] != '\\' {
						break
					} else {
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

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 42

var yyAct = [...]int{
	13, 19, 20, 12, 24, 7, 41, 27, 36, 23,
	25, 26, 34, 21, 33, 32, 31, 30, 29, 28,
	4, 6, 15, 18, 16, 37, 5, 11, 39, 40,
	35, 2, 14, 10, 17, 9, 38, 8, 3, 1,
	0, 22,
}

var yyPact = [...]int{
	-9, -1000, -57, -1000, -61, -14, 8, -1000, -1000, -1000,
	-59, -58, -25, -1000, -1000, -1000, -1000, -59, -1000, -34,
	-31, -18, -1000, -19, -20, -21, -22, -23, -29, 26,
	-33, 20, 18, -35, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000,
}

var yyPgo = [...]int{
	0, 39, 38, 37, 36, 35, 34, 23, 32,
}

var yyR1 = [...]int{
	0, 1, 1, 1, 2, 2, 2, 2, 2, 2,
	3, 6, 6, 7, 7, 7, 7, 4, 4, 5,
	8,
}

var yyR2 = [...]int{
	0, 0, 5, 1, 2, 2, 2, 1, 1, 1,
	2, 2, 1, 4, 4, 4, 4, 1, 1, 5,
	1,
}

var yyChk = [...]int{
	-1000, -1, 40, -2, 29, 35, 30, 14, -3, -5,
	42, 36, 60, 61, -8, 36, 16, -6, -7, 60,
	60, 38, -7, 43, 38, 44, 45, 38, 37, 37,
	37, 37, 37, 37, 41, 4, 41, 5, -4, 10,
	11, 41,
}

var yyDef = [...]int{
	1, -2, 0, 3, 0, 0, 0, 7, 8, 9,
	0, 0, 0, 4, 5, 20, 6, 10, 12, 0,
	0, 0, 11, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 2, 13, 14, 15, 16, 17,
	18, 19,
}

var yyTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 60, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 61,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59,
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
//line sint.y:39
		{
		}
	case 2:
		yyDollar = yyS[yypt-5 : yypt+1]
//line sint.y:40
		{
			leerArchivoDeEntrada(yyDollar[5].str)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
//line sint.y:44
		{
			fmt.Print("JEJE")
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line sint.y:45
		{
			fmt.Println("MONTANDO EL YIP YIP ")
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line sint.y:46
		{
			fmt.Println(" ----OK--- ")
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line sint.y:47
		{
			pausar_()
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
//line sint.y:52
		{
			CrearDisco(Size_, Path_, Name_, Unit_)
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
//line sint.y:57
		{
			Size_ = yyDollar[4].str
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
//line sint.y:58
		{
			Path_ = yyDollar[4].str
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
//line sint.y:59
		{
			Name_ = yyDollar[4].str
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
//line sint.y:60
		{
			Unit_ = yyDollar[4].NoTerminal
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line sint.y:64
		{
			yyVAL.NoTerminal = yyDollar[1].str
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line sint.y:65
		{
			yyVAL.NoTerminal = yyDollar[1].str
		}
	case 19:
		yyDollar = yyS[yypt-5 : yypt+1]
//line sint.y:68
		{
			EliminarDisco(yyDollar[5].str)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line sint.y:72
		{
		}
	}
	goto yystack /* stack new state and value */
}

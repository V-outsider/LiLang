// Code generated from LiLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type LiLangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var LiLangLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lilanglexerLexerInit() {
	staticData := &LiLangLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'.'", "'var'", "'function'", "'return'", "'if'", "'else'", "'switch'",
		"'case'", "'default'", "'int'", "'list'", "'add'", "'remove'", "'get'",
		"'size'", "'print'", "", "", "", "", "", "'+'", "'-'", "'*'", "'/'",
		"'='", "'('", "')'", "'{'", "'}'", "':'", "';'", "','", "'&'", "'['",
		"']'", "'<'", "'>'",
	}
	staticData.SymbolicNames = []string{
		"", "", "VAR", "FUNCTION", "RETURN", "IF", "ELSE", "SWITCH", "CASE",
		"DEFAULT", "INT", "LIST", "ADD", "REMOVE", "GET", "SIZE", "PRINT", "WS",
		"COMMENT", "NUMBER", "IDENTIFIER", "STRING", "PLUS", "MINUS", "MULT",
		"DIV", "ASSIGN", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "COLON", "SEMICOLON",
		"COMMA", "AMPERSAND", "SLBRACE", "LLBRACE", "LT", "GT",
	}
	staticData.RuleNames = []string{
		"T__0", "VAR", "FUNCTION", "RETURN", "IF", "ELSE", "SWITCH", "CASE",
		"DEFAULT", "INT", "LIST", "ADD", "REMOVE", "GET", "SIZE", "PRINT", "WS",
		"COMMENT", "NUMBER", "IDENTIFIER", "STRING", "PLUS", "MINUS", "MULT",
		"DIV", "ASSIGN", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "COLON", "SEMICOLON",
		"COMMA", "AMPERSAND", "SLBRACE", "LLBRACE", "LT", "GT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 38, 235, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 4, 16, 164,
		8, 16, 11, 16, 12, 16, 165, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5,
		17, 174, 8, 17, 10, 17, 12, 17, 177, 9, 17, 1, 17, 1, 17, 1, 18, 4, 18,
		182, 8, 18, 11, 18, 12, 18, 183, 1, 19, 1, 19, 5, 19, 188, 8, 19, 10, 19,
		12, 19, 191, 9, 19, 1, 20, 1, 20, 5, 20, 195, 8, 20, 10, 20, 12, 20, 198,
		9, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 196, 0, 38, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13,
		27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22,
		45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31,
		63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 1, 0, 5, 3, 0,
		9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 1, 0, 48, 57, 3, 0, 65, 90,
		95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 239, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0,
		0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1,
		0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71,
		1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 1, 77, 1, 0, 0, 0, 3,
		79, 1, 0, 0, 0, 5, 83, 1, 0, 0, 0, 7, 92, 1, 0, 0, 0, 9, 99, 1, 0, 0, 0,
		11, 102, 1, 0, 0, 0, 13, 107, 1, 0, 0, 0, 15, 114, 1, 0, 0, 0, 17, 119,
		1, 0, 0, 0, 19, 127, 1, 0, 0, 0, 21, 131, 1, 0, 0, 0, 23, 136, 1, 0, 0,
		0, 25, 140, 1, 0, 0, 0, 27, 147, 1, 0, 0, 0, 29, 151, 1, 0, 0, 0, 31, 156,
		1, 0, 0, 0, 33, 163, 1, 0, 0, 0, 35, 169, 1, 0, 0, 0, 37, 181, 1, 0, 0,
		0, 39, 185, 1, 0, 0, 0, 41, 192, 1, 0, 0, 0, 43, 201, 1, 0, 0, 0, 45, 203,
		1, 0, 0, 0, 47, 205, 1, 0, 0, 0, 49, 207, 1, 0, 0, 0, 51, 209, 1, 0, 0,
		0, 53, 211, 1, 0, 0, 0, 55, 213, 1, 0, 0, 0, 57, 215, 1, 0, 0, 0, 59, 217,
		1, 0, 0, 0, 61, 219, 1, 0, 0, 0, 63, 221, 1, 0, 0, 0, 65, 223, 1, 0, 0,
		0, 67, 225, 1, 0, 0, 0, 69, 227, 1, 0, 0, 0, 71, 229, 1, 0, 0, 0, 73, 231,
		1, 0, 0, 0, 75, 233, 1, 0, 0, 0, 77, 78, 5, 46, 0, 0, 78, 2, 1, 0, 0, 0,
		79, 80, 5, 118, 0, 0, 80, 81, 5, 97, 0, 0, 81, 82, 5, 114, 0, 0, 82, 4,
		1, 0, 0, 0, 83, 84, 5, 102, 0, 0, 84, 85, 5, 117, 0, 0, 85, 86, 5, 110,
		0, 0, 86, 87, 5, 99, 0, 0, 87, 88, 5, 116, 0, 0, 88, 89, 5, 105, 0, 0,
		89, 90, 5, 111, 0, 0, 90, 91, 5, 110, 0, 0, 91, 6, 1, 0, 0, 0, 92, 93,
		5, 114, 0, 0, 93, 94, 5, 101, 0, 0, 94, 95, 5, 116, 0, 0, 95, 96, 5, 117,
		0, 0, 96, 97, 5, 114, 0, 0, 97, 98, 5, 110, 0, 0, 98, 8, 1, 0, 0, 0, 99,
		100, 5, 105, 0, 0, 100, 101, 5, 102, 0, 0, 101, 10, 1, 0, 0, 0, 102, 103,
		5, 101, 0, 0, 103, 104, 5, 108, 0, 0, 104, 105, 5, 115, 0, 0, 105, 106,
		5, 101, 0, 0, 106, 12, 1, 0, 0, 0, 107, 108, 5, 115, 0, 0, 108, 109, 5,
		119, 0, 0, 109, 110, 5, 105, 0, 0, 110, 111, 5, 116, 0, 0, 111, 112, 5,
		99, 0, 0, 112, 113, 5, 104, 0, 0, 113, 14, 1, 0, 0, 0, 114, 115, 5, 99,
		0, 0, 115, 116, 5, 97, 0, 0, 116, 117, 5, 115, 0, 0, 117, 118, 5, 101,
		0, 0, 118, 16, 1, 0, 0, 0, 119, 120, 5, 100, 0, 0, 120, 121, 5, 101, 0,
		0, 121, 122, 5, 102, 0, 0, 122, 123, 5, 97, 0, 0, 123, 124, 5, 117, 0,
		0, 124, 125, 5, 108, 0, 0, 125, 126, 5, 116, 0, 0, 126, 18, 1, 0, 0, 0,
		127, 128, 5, 105, 0, 0, 128, 129, 5, 110, 0, 0, 129, 130, 5, 116, 0, 0,
		130, 20, 1, 0, 0, 0, 131, 132, 5, 108, 0, 0, 132, 133, 5, 105, 0, 0, 133,
		134, 5, 115, 0, 0, 134, 135, 5, 116, 0, 0, 135, 22, 1, 0, 0, 0, 136, 137,
		5, 97, 0, 0, 137, 138, 5, 100, 0, 0, 138, 139, 5, 100, 0, 0, 139, 24, 1,
		0, 0, 0, 140, 141, 5, 114, 0, 0, 141, 142, 5, 101, 0, 0, 142, 143, 5, 109,
		0, 0, 143, 144, 5, 111, 0, 0, 144, 145, 5, 118, 0, 0, 145, 146, 5, 101,
		0, 0, 146, 26, 1, 0, 0, 0, 147, 148, 5, 103, 0, 0, 148, 149, 5, 101, 0,
		0, 149, 150, 5, 116, 0, 0, 150, 28, 1, 0, 0, 0, 151, 152, 5, 115, 0, 0,
		152, 153, 5, 105, 0, 0, 153, 154, 5, 122, 0, 0, 154, 155, 5, 101, 0, 0,
		155, 30, 1, 0, 0, 0, 156, 157, 5, 112, 0, 0, 157, 158, 5, 114, 0, 0, 158,
		159, 5, 105, 0, 0, 159, 160, 5, 110, 0, 0, 160, 161, 5, 116, 0, 0, 161,
		32, 1, 0, 0, 0, 162, 164, 7, 0, 0, 0, 163, 162, 1, 0, 0, 0, 164, 165, 1,
		0, 0, 0, 165, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 1, 0, 0,
		0, 167, 168, 6, 16, 0, 0, 168, 34, 1, 0, 0, 0, 169, 170, 5, 47, 0, 0, 170,
		171, 5, 47, 0, 0, 171, 175, 1, 0, 0, 0, 172, 174, 8, 1, 0, 0, 173, 172,
		1, 0, 0, 0, 174, 177, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 175, 176, 1, 0,
		0, 0, 176, 178, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 178, 179, 6, 17, 0, 0,
		179, 36, 1, 0, 0, 0, 180, 182, 7, 2, 0, 0, 181, 180, 1, 0, 0, 0, 182, 183,
		1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 38, 1, 0,
		0, 0, 185, 189, 7, 3, 0, 0, 186, 188, 7, 4, 0, 0, 187, 186, 1, 0, 0, 0,
		188, 191, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190,
		40, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 192, 196, 5, 34, 0, 0, 193, 195,
		9, 0, 0, 0, 194, 193, 1, 0, 0, 0, 195, 198, 1, 0, 0, 0, 196, 197, 1, 0,
		0, 0, 196, 194, 1, 0, 0, 0, 197, 199, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0,
		199, 200, 5, 34, 0, 0, 200, 42, 1, 0, 0, 0, 201, 202, 5, 43, 0, 0, 202,
		44, 1, 0, 0, 0, 203, 204, 5, 45, 0, 0, 204, 46, 1, 0, 0, 0, 205, 206, 5,
		42, 0, 0, 206, 48, 1, 0, 0, 0, 207, 208, 5, 47, 0, 0, 208, 50, 1, 0, 0,
		0, 209, 210, 5, 61, 0, 0, 210, 52, 1, 0, 0, 0, 211, 212, 5, 40, 0, 0, 212,
		54, 1, 0, 0, 0, 213, 214, 5, 41, 0, 0, 214, 56, 1, 0, 0, 0, 215, 216, 5,
		123, 0, 0, 216, 58, 1, 0, 0, 0, 217, 218, 5, 125, 0, 0, 218, 60, 1, 0,
		0, 0, 219, 220, 5, 58, 0, 0, 220, 62, 1, 0, 0, 0, 221, 222, 5, 59, 0, 0,
		222, 64, 1, 0, 0, 0, 223, 224, 5, 44, 0, 0, 224, 66, 1, 0, 0, 0, 225, 226,
		5, 38, 0, 0, 226, 68, 1, 0, 0, 0, 227, 228, 5, 91, 0, 0, 228, 70, 1, 0,
		0, 0, 229, 230, 5, 93, 0, 0, 230, 72, 1, 0, 0, 0, 231, 232, 5, 60, 0, 0,
		232, 74, 1, 0, 0, 0, 233, 234, 5, 62, 0, 0, 234, 76, 1, 0, 0, 0, 6, 0,
		165, 175, 183, 189, 196, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// LiLangLexerInit initializes any static state used to implement LiLangLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLiLangLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LiLangLexerInit() {
	staticData := &LiLangLexerLexerStaticData
	staticData.once.Do(lilanglexerLexerInit)
}

// NewLiLangLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLiLangLexer(input antlr.CharStream) *LiLangLexer {
	LiLangLexerInit()
	l := new(LiLangLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &LiLangLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "LiLang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LiLangLexer tokens.
const (
	LiLangLexerT__0       = 1
	LiLangLexerVAR        = 2
	LiLangLexerFUNCTION   = 3
	LiLangLexerRETURN     = 4
	LiLangLexerIF         = 5
	LiLangLexerELSE       = 6
	LiLangLexerSWITCH     = 7
	LiLangLexerCASE       = 8
	LiLangLexerDEFAULT    = 9
	LiLangLexerINT        = 10
	LiLangLexerLIST       = 11
	LiLangLexerADD        = 12
	LiLangLexerREMOVE     = 13
	LiLangLexerGET        = 14
	LiLangLexerSIZE       = 15
	LiLangLexerPRINT      = 16
	LiLangLexerWS         = 17
	LiLangLexerCOMMENT    = 18
	LiLangLexerNUMBER     = 19
	LiLangLexerIDENTIFIER = 20
	LiLangLexerSTRING     = 21
	LiLangLexerPLUS       = 22
	LiLangLexerMINUS      = 23
	LiLangLexerMULT       = 24
	LiLangLexerDIV        = 25
	LiLangLexerASSIGN     = 26
	LiLangLexerLPAREN     = 27
	LiLangLexerRPAREN     = 28
	LiLangLexerLBRACE     = 29
	LiLangLexerRBRACE     = 30
	LiLangLexerCOLON      = 31
	LiLangLexerSEMICOLON  = 32
	LiLangLexerCOMMA      = 33
	LiLangLexerAMPERSAND  = 34
	LiLangLexerSLBRACE    = 35
	LiLangLexerLLBRACE    = 36
	LiLangLexerLT         = 37
	LiLangLexerGT         = 38
)

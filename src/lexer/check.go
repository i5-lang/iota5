// SPDX-License-Identifier: GPL-3.0-or-later
package lexer

import "github.com/i5/i5/src/types"

const (
	EQ       = 61
	NOT      = 33
	PLUS     = 43
	MINUS    = 45
	MULTIPLY = 42
	DIVIDE   = 47
	MODULO   = 37
	AND      = 38
	OR       = 124
	XOR      = 94
	BNOT     = 126
	LT       = 60
	GT       = 62
	COLON    = 58
	DOT      = 46
	COMMA    = 44
	LPAREN   = 40
	RPAREN   = 41
	LBRACE   = 123
	RBRACE   = 125
	LBRACKET = 91
	RBRACKET = 93
	QM       = 63
)

func IsKeyword(char string) (string, bool) {
	switch char {
	case types.IF:
		return types.IF, true
	case types.ELIF:
		return types.ELIF, true
	case types.ELSE:
		return types.ELSE, true
	case types.SWITCH:
		return types.SWITCH, true
	case types.CASE:
		return types.CASE, true
	case types.LOOP:
		return types.LOOP, true
	case types.FN:
		return types.FN, true
	case types.RETURN:
		return types.RETURN, true
	case types.IMPORT:
		return types.IMPORT, true
	case types.THROW:
		return types.THROW, true
	case types.TRY:
		return types.TRY, true
	case types.CATCH:
		return types.CATCH, true
	case types.FINALLY:
		return types.FINALLY, true
	case types.TRUE:
		return types.TRUE, true
	case types.FALSE:
		return types.FALSE, true
	case types.ANDAND:
		return types.ANDAND, true
	case types.OROR:
		return types.OROR, true
	default:
		return "", false
	}
}

func escape(char byte) string {
	switch char {
	case 116:
		return string(9) // if char is 't' return '\t'
	case 110:
		return string(10) // if char is 'n' return '\n'
	case 114:
		return string(13) // if char is 'r' return '\r'
	default:
		return string(char) // else return string(char)
	}
}

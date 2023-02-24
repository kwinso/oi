// Code generated by "stringer -type=TokenType"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ILLEGAL-0]
	_ = x[EOF-1]
	_ = x[NEWLINE-2]
	_ = x[IDENT-3]
	_ = x[INT-4]
	_ = x[FLOAT-5]
	_ = x[TRUE-6]
	_ = x[FALSE-7]
	_ = x[STRING-8]
	_ = x[COMMA-9]
	_ = x[DOT-10]
	_ = x[LPAREN-11]
	_ = x[RPAREN-12]
	_ = x[LBRACE-13]
	_ = x[RBRACE-14]
	_ = x[ASSIGN-15]
	_ = x[PLUS-16]
	_ = x[MINUS-17]
	_ = x[MULTIPLY-18]
	_ = x[DIVIDE-19]
	_ = x[POWER-20]
	_ = x[AND-21]
	_ = x[OR-22]
	_ = x[NOT-23]
	_ = x[EQ-24]
	_ = x[NEQ-25]
	_ = x[LT-26]
	_ = x[GT-27]
	_ = x[LTE-28]
	_ = x[GTE-29]
	_ = x[LET-30]
	_ = x[FN-31]
	_ = x[RETURN-32]
	_ = x[IF-33]
	_ = x[ELSE-34]
	_ = x[PIPE_CTX-35]
	_ = x[PIPE_FN-36]
	_ = x[PIPE_OP-37]
}

const _TokenType_name = "ILLEGALEOFNEWLINEIDENTINTFLOATTRUEFALSESTRINGCOMMADOTLPARENRPARENLBRACERBRACEASSIGNPLUSMINUSMULTIPLYDIVIDEPOWERANDORNOTEQNEQLTGTLTEGTELETFNRETURNIFELSEPIPE_CTXPIPE_FNPIPE_OP"

var _TokenType_index = [...]uint8{0, 7, 10, 17, 22, 25, 30, 34, 39, 45, 50, 53, 59, 65, 71, 77, 83, 87, 92, 100, 106, 111, 114, 116, 119, 121, 124, 126, 128, 131, 134, 137, 139, 145, 147, 151, 159, 166, 173}

func (i TokenType) String() string {
	if i < 0 || i >= TokenType(len(_TokenType_index)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}

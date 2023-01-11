package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
}

const (
	ILLEGAL = "HARAMU"
	EOF     = "MWISHO"

	// Identifiers + literals
	IDENT  = "KITAMBULISHI"
	INT    = "NAMBA"
	STRING = "NENO"
	FLOAT  = "DESIMALI"

	// Operators
	ASSIGN          = "="
	PLUS            = "+"
	MINUS           = "-"
	BANG            = "!"
	ASTERISK        = "*"
	POW             = "**"
	SLASH           = "/"
	MODULUS         = "%"
	LT              = "<"
	LTE             = "<="
	GT              = ">"
	GTE             = ">="
	EQ              = "=="
	NOT_EQ          = "!="
	AND             = "&&"
	OR              = "||"
	PLUS_ASSIGN     = "+="
	PLUS_PLUS       = "++"
	MINUS_ASSIGN    = "-="
	MINUS_MINUS     = "--"
	ASTERISK_ASSIGN = "*="
	SLASH_ASSIGN    = "/="
	MODULUS_ASSIGN  = "%="

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	COLON     = ":"
	DOT       = "."

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "FANYA"
	TRUE     = "KWELI"
	FALSE    = "SIKWELI"
	IF       = "KAMA"
	ELSE     = "SIVYO"
	RETURN   = "RUDISHA"
	WHILE    = "WAKATI"
	NULL     = "TUPU"
	BREAK    = "VUNJA"
	CONTINUE = "ENDELEA"
	IN       = "KTK"
	FOR      = "KWA"
	SWITCH   = "BADILI"
	CASE     = "IKIWA"
	DEFAULT  = "KAWAIDA"
	TIME     = "MUDA"
	IMPORT   = "TUMIA"
)

var keywords = map[string]TokenType{
	"unda":    FUNCTION,
	"fanya":   LET,
	"kweli":   TRUE,
	"sikweli": FALSE,
	"kama":    IF,
	"au":      ELSE,
	"sivyo":   ELSE,
	"wakati":  WHILE,
	"rudisha": RETURN,
	"vunja":   BREAK,
	"endelea": CONTINUE,
	"tupu":    NULL,
	"ktk":     IN,
	"kwa":     FOR,
	"badili":  SWITCH,
	"ikiwa":   CASE,
	"kawaida": DEFAULT,
	"muda":    TIME,
	"tumia":   IMPORT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

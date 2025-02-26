package tokens

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL     = "ILLEGAL"
	EOF         = ""
	DASH        = "DASH"
	DOUBLE_DASH = "DOUBLE_DASH"
	IDENT       = "IDENT" // command and values
)

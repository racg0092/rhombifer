// Abstract Syntax Tree
package ast

import "github.com/racg0092/rhombifer/tokens"

type Node interface {
	TokenLiteral() string
}

type Cmd interface {
	Node
	commandNode()
}

type Flag interface {
	Node
	flagNode()
}

type Program struct {
	Commands []Cmd
}

func (p Program) TokenLiteral() string {
	//NOTE: it can probable be handly differently
	if len(p.Commands) > 0 {
		return p.Commands[0].TokenLiteral()
	} else {
		return ""
	}
}

type Command struct {
	Token tokens.Token
	Name  string
	// Sub   bool //NOTE: may not be needed
}

type Identifier struct {
	Token tokens.Token
	Value string
}

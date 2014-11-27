package vmparse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const formula = `
f:= a*b 
g := sin(1+2)^2/sqrt(9*x)
`

//go:generate go tool yacc -o vmparser.y.go vmparser.y
//go:generate nex vmlexer.nex
func Run(args []string) {
	lex := NewContext(NewLexer(strings.NewReader(formula)))
	yyParse(lex)
	if len(lex.errors) > 0 {
		log.Fatal(lex.errors)
	}
	for i, s := range lex.statements {
		fmt.Printf("statement %d. %s\n", i, s.Formula())
	}
}

func NewContext(y yyLexer) *context {
	return &context{yyLexer: y}
}
func (c *context) Error(e string) {
	c.Error2(fmt.Errorf("%s", e))
}

func (c *context) Error2(e error) {
	log.Printf("oops: %v\n", e)
	c.errors = append(c.errors, e)
}

type context struct {
	statements []*Node
	yyLexer
	errors []error
}

type NodeType string

const (
	numberNT            NodeType = "NUMBER"
	identifierNT        NodeType = "IDENTIFIER"
	indexedIdentifierNT NodeType = "INDEXED IDENTIFIER"
	functionNT          NodeType = "FUNCTION"
	statementNT         NodeType = "STATEMENT"
)

type Node struct {
	Type     NodeType `json:"T,omitempty"`
	S        string   `json:",omitempty"`
	F        float64  `json:",omitempty"`
	I        int      `json:",omitempty"`
	Children []*Node  `json:"C,omitempty"`
	Name     string   `json:"N,omitempty"` // name of variable assigned by parser

	// for vm-assembly version:
	register, output int
}

func (n Node) String() string {
	buf, _ := json.Marshal(n)
	return string(buf)
}

func (n Node) Formula() string {
	buf := new(bytes.Buffer)
	switch n.Type {
	case numberNT:
		fmt.Fprintf(buf, "%f", n.F)
	case identifierNT:
		fmt.Fprintf(buf, "%s", n.S)
	case indexedIdentifierNT:
		fmt.Fprintf(buf, "%s[%d]", n.S, n.I)
	case functionNT:
		op := func(x string) {
			fmt.Fprintf(buf, "(%s %s %s)", n.Children[0].Formula(), x, n.Children[1].Formula())
		}
		switch n.S {
		case "multiply":
			op("*")
		case "divide":
			op("/")
		case "subtract":
			op("-")
		case "add":
			op("+")
		case "pow":
			op("^")
		default:
			var args []string
			for _, c := range n.Children {
				args = append(args, c.Formula())
			}
			fmt.Fprintf(buf, "%s(%s)", n.S, strings.Join(args, ","))
		}
	case statementNT:
		fmt.Fprintf(buf, "%s := %s", n.Children[0].Formula(), n.Children[1].Formula())
	default:
		panic("illegal type: " + n.Type)
	}
	return buf.String()

}

func LexNumber(s string) *Node {
	n, _ := strconv.ParseFloat(s, 64)
	return Number(n)
}

func LexIdentifier(s string) *Node {
	return &Node{
		Type: identifierNT,
		S:    s,
	}
}

func Number(n float64) *Node {
	return &Node{
		Type: numberNT,
		F:    n,
	}
}

func NewStatement(lhs, rhs *Node) *Node {
	return &Node{
		Type:     statementNT,
		Children: []*Node{lhs, rhs},
	}
}

func IndexedIdentifier(ident, index *Node) *Node {
	return &Node{
		Type: indexedIdentifierNT,
		S:    ident.S,
		I:    int(index.F),
	}
}

func Function(ident string, args ...*Node) *Node {
	return &Node{
		Type:     functionNT,
		S:        ident,
		Children: args,
	}
}

func Negate(a *Node) *Node {
	return Function("multiply", Number(-1), a)
}

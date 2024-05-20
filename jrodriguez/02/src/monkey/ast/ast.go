package ast 


// Ast conists of nodes connected with each other
// nodes are of the following two types : expressions and statements 
// Every node in the ast will have to implement the Node interface 


// The base node interface 

type Node interface {

	// TokenLiteral() will be used only for testing and debugging 
	TokenLiteral () string 
}

// dummy methods that are not strictly necessary but will guide the Go compiler ( producing errors )

type Statement interface {
	Node 
	statementNode()
}
type Expression interface {
	Node 
	expressionNode ()
}

type Program struct {

	// Every valid monkey program is a series of statements
	// These statements are conained in program.statements 
	Statements [] Statement
}



func (p * Program ) TokenLiteral() string {
	if len(p.Statemenst)> 0 
	{
		return p.Statements[0].TokenLiteral()
	}
	else 
	{
		return ""
	}
		
}

type LetStatement struct {

	Token token.Token 
	Name *Identifier 
	Value Expression
}
func (ls *LetStatement) statementNode()
package ast 


// Ast conists of nodes connected with each other
// nodes are of the follwoing two types : expressions and statements 
// Every node in the ast will have to implement the node interface 


// The base node interface 

type Node interface {

	// TokenLiteral() will be used only for testing and debugging 
	TokenLiteral () string 
}

type Statement interface {
	Node 
	statementNode()
}
type Expression interface {
	Node 
	expressionNode ()
}
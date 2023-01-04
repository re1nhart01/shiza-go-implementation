package ast_tree

type ExpressionNode struct {
}

type IToken interface {
	Print()
}

type AssignNode struct {
	ExpressionNode
	Operator  IToken
	LeftNode  ExpressionNode
	RightNode ExpressionNode
}

type BinaryOperation struct {
	ExpressionNode
	Operator  IToken
	LeftNode  ExpressionNode
	RightNode ExpressionNode
}

type VariableNode struct {
	ExpressionNode
	Keyword  IToken
	Variable IToken
}
type NumberNode struct {
	ExpressionNode
	Number IToken
}
type StringNode struct {
	ExpressionNode
	String IToken
}
type BooleanNode struct {
	ExpressionNode
	Boolean IToken
}
type StatementsNode struct {
	ExpressionNode
	CodeStrings []ExpressionNode
}

func (stNode *StatementsNode) AddNode(node ExpressionNode) {
	stNode.CodeStrings = append(stNode.CodeStrings, node)
}

type UnaryOperationNode struct {
	ExpressionNode
	Operator IToken
	Operand  ExpressionNode
}

type ConstructionNode struct {
	ExpressionNode
	Operator IToken
	Query    []ExpressionNode
}

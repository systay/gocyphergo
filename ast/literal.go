package ast

type Expression interface {
	Value() interface{}
	expressionNode()
}

package parser

type Position struct {
	Line   int
	Column int
}

type Node interface {
	Kind() string
}

type Value interface {
	Node
	valueNode()
}

type ObjectField struct {
	Key   *StringNode
	Value Value
}

type ObjectNode struct {
	Fields   []ObjectField
	Position Position
}

func (*ObjectNode) Kind() string { return "Object" }
func (*ObjectNode) valueNode()   {}

type ArrayNode struct {
	Elements []Value
	Position Position
}

func (*ArrayNode) Kind() string { return "Array" }
func (*ArrayNode) valueNode()   {}

type StringNode struct {
	Value    string
	Position Position
}

func (*StringNode) Kind() string { return "String" }
func (*StringNode) valueNode()   {}

type NumberNode struct {
	Raw      string
	Position Position
}

func (*NumberNode) Kind() string { return "Number" }
func (*NumberNode) valueNode()   {}

type BooleanNode struct {
	Value    bool
	Position Position
}

func (*BooleanNode) Kind() string { return "Boolean" }
func (*BooleanNode) valueNode()   {}

type NullNode struct {
	Position Position
}

func (*NullNode) Kind() string { return "Null" }
func (*NullNode) valueNode()   {}
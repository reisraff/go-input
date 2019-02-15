package node

import "github.com/reisraff/goinput/input/interfaces"
import "github.com/reisraff/goinput/input/constraints"

func CreateFloatNode() interfaces.NodeInterface {
    node := FloatNode{}
    node.AddConstraint(constraints.ConstraintType("float"))

    return &node
}

type FloatNode struct {
    BaseNode
}
package tree

import "fmt"

type NodeState int

const (
	SUCCESS NodeState = iota
	FAILURE
	CONTINUE
)

type NodeExecution interface {
	jumpToParent()
	executeNode()
}

type Behaviour interface {
	executionBehaviour(n *Node) NodeState
}

type Node struct {
	parent    *Node
	Children  []*Node
	behaviour Behaviour
	nodeState NodeState
}

func (n *Node) jumpToParent() {
	if n.parent != nil {
		n.parent.executeNode()
	}
	fmt.Println("You are at the beginning, forward is your only option!")
	n.executeNode()
}

func (n *Node) executeNode() {
	n.behaviour.executionBehaviour(n)
}

type SelectorBehaviour struct{}

type SequenceBehaviour struct{}

func (sb *SelectorBehaviour) executionBehaviour(n *Node) NodeState {
	for _, childNode := range n.Children {
		childNode.executeNode()
		if childNode.nodeState == SUCCESS {
			return SUCCESS
		} else if childNode.nodeState == CONTINUE {
			return CONTINUE
		}
	}
	return FAILURE
}

func (sb *SequenceBehaviour) executionBehaviour(n *Node) NodeState {
	for _, childNode := range n.Children {
		childNode.executeNode()
		if childNode.nodeState == CONTINUE {
			return CONTINUE
		} else if childNode.nodeState == FAILURE {
			return FAILURE
		}
	}
	return SUCCESS
}

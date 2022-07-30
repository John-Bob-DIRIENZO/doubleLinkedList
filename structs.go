package doubleLinkedList

import "fmt"

type Node struct {
	Value    int
	prevNode *Node
	nextNode *Node
}

type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

func (n *Node) GetValue(index int) (int, error) {
	if n.GetNode(index) == nil {
		return 0, &MyError{Message: "Je n'ai rien trouvé ici"}
	}

	return n.GetNode(index).Value, nil
}

func (n *Node) GetNode(index int) *Node {
	if index == 0 {
		return n
	}

	temp := n

	for i := 1; i <= index; i++ {
		temp = temp.nextNode
		if temp == nil {
			return nil
		}
	}

	return temp
}

func (n *Node) AddNode(newNode *Node) {
	lastNode := n.GetLastNode()
	newNode.prevNode = lastNode
	lastNode.nextNode = newNode
}

func (n *Node) DeleteNode(index int) error {
	if n.GetNode(index) == nil {
		return &MyError{Message: "Je n'ai rien trouvé ici"}
	}

	nodeToDelete := n.GetNode(index)
	nodeToDelete.prevNode.nextNode = nodeToDelete.nextNode
	nodeToDelete.nextNode.prevNode = nodeToDelete.prevNode
	return nil
}

func (n *Node) GetLastNode() *Node {
	parsedNode := n

	for parsedNode.nextNode != nil {
		parsedNode = parsedNode.nextNode
	}

	return parsedNode
}

func (n *Node) TraverseList() {
	parsedNode := n

	fmt.Println(parsedNode.Value)

	for parsedNode.nextNode != nil {
		parsedNode = parsedNode.nextNode
		fmt.Println(parsedNode.Value)
	}
}

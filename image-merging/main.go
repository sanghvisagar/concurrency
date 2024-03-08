// node 
// parent
// child 
// value




package main


type Node struct {
	Value int
	//Parent Node
	Child []Node
}

func main() {

	tree1 := Node{}
	tree2 := Node{}

	resultTree := Node{}

	treeTraversal(tree1, tree2, &resultTree)





}

Tree1
             x

	  x      1    0     1
     1 0 1 1


Tree2 
              1



resultTree
	



              





func treeTraversal(tree1 Node, tree2 Node, resultTree *Node)  {
	n := len(tree1.Child)
	m := len(tree2.Child)

	if n==0 and m==0 {
		resultTree := Node{
			Value: operate(tree1.Value, tree2.Value)
			Child: []Node{}
		}
		// resultTree.Child = append(resultTree.Child, newNode)
		return 
	}

	for idx, childNode := range max(tree1,tree2) {
		
		if idx >= m {
			treeTraversal(childNode, Node{Value: tree2.Value}, resultTree.Child[idx])
		} else {
			treeTraversal(childNode, tree2.Child[idx],  resultTree.Child[idx])
		}
	}

	//if tree1.value == tree2.value {
	// create a new node - get its value via operate function 
	//    update the child of existing node 
	//   update resultTree to new node
	newNode := Node{
		Value: operate(tree1.Value, tree2.Value)
		Child: []Node{}
	}
	resultTree.Child = append(resultTree.Child, newNode)
	// resultTree = newNode


	


	// if both are nil , just return 

	// if one of tree is nil and other is non nil 

	// if both are non nil
	// call the operate function
	// return the value after applying 

	

	// newNode := Node{}
	// resultTree.value = operate(tree1.val, tree2.val)
	





	// if n>m {
	
	// }




	

}


func operate(val1 int, val2 int) {

}
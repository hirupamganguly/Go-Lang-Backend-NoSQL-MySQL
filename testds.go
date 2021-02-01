package main

// Node ...
type Node struct {
	key  int
	data string
	l    *Node
	r    *Node
}

// Bst ...
type Bst struct {
	root *Node
}

// derefferncing will done by the compiler automatically
func (n *Node) insertRecursively(node Node) {
	// function to find the correct place for a node in a tree

	// when new node is lesser than n -> go to left subtree of n and
	//again
	//start the function by making - left child of n ->as n
	// restore Point(Previous Execution State) stored at callstack.. it restored again when
	// newly function call finshed..
	// when new node is greater than n -> go to right subtree of n and
	//again
	//start the function by making - right child of n ->as n
	// restore Point(Previous Execution State) stored at callstack.. it restored again when
	// newly function call finshed..
	if n == nil {
		return //  base condition of Recursive Problem->    if(n==null) return ;
	} else if node.key <= n.key {
		if n.l == nil {
			n.l = &node
		} else {
			n.l.insertRecursively(node)
		}
	} else {
		if n.r == nil {
			n.r = &node
		} else {
			n.r.insertRecursively(node)
		}
	}
}

func (bst *Bst) insert(keyy int, value string) *Bst {
	// insert function create Node then add the node to its correct position by calling insertRecursively()

	node := &Node{key: keyy, data: value}
	if bst.root == nil {
		bst.root = &Node{key: keyy, data: value}
	} else {
		bst.root.insertRecursively(*node)
	}
	return bst
}
func searchRecursively(keyy int, n *Node) bool {

	if n == nil {
		return false //// base case of recursive problem.
	}
	// recursively call for left subtree and right subtree, by checking the key is smaller or greater than n
	//(as it is recursive function- here n variable contains actually curent node each time..)
	// when key is lesser than n -> go to left subtree of n and
	//again
	//start the function by making - left child of n ->as n
	// when key is greater than n -> go to right subtree of n and
	// again
	//start the function by making - right child of n ->as n
	if keyy < n.key {
		return searchRecursively(keyy, n.l)
	}
	if keyy > n.key {
		return searchRecursively(keyy, n.r)
	}
	// we find the match so return true
	return true
}

func (bst *Bst) search(keyy int) bool {
	return searchRecursively(keyy, bst.root)
}
func removeRecursively(keyy int, n *Node) *Node {
	if n == nil {
		// base case of recursive Problem.
		return nil
	}
	if keyy < n.key {
		// find the element at left subtree
		n.l = removeRecursively(keyy, n.l)
		return n
	}
	if keyy > n.key {
		// find the element at right subtree
		n.r = removeRecursively(keyy, n.r)
		return n
	}
	//when found then:-
	if keyy == n.key {
		// no child
		if n.l == nil && n.r == nil {
			n = nil
			return nil
		}
		// single child
		if n.l == nil {
			// if it has only right child then replaced it with right child.
			n = n.r
			return n
		}
		if n.r == nil {
			// if it has only left child then replaced it with left child.
			n = n.l
			return n
		}
		// two children

		// try to find the next small node after the node which we want to delete
		// here we create and initialize lmostOfRight variable as -> n.right, means we just have to find
		// the smallest of n.right, because that will be the in-Order Successor
		lmostOfRight := n.r
		for {
			if lmostOfRight != nil && lmostOfRight.l != nil {
				lmostOfRight = lmostOfRight.l // just incrementing towards left...
			} else {
				break
			}
		}
		// Found it,
		//then... replace with it
		//as it has both child then replaced it with in order successor. then delete that inorder successor as
		// it is still prsent at its position.
		//Now the twist is as it is inorder successor so it can have right child or no chils, as it is left most of n.right.
		// Now again we call delete function by passing the inorder successor's key.
		n.key, n.data = lmostOfRight.key, lmostOfRight.data
		//then... delete it
		n.r = removeRecursively(n.key, n.r)
		return n
	}
	return n
}
func (bst *Bst) remove(keyy int) {
	removeRecursively(keyy, bst.root)
}

package main

import "fmt"

// countBST is a function that returns
// a function that returns the number of nodes in a BST.
//currently only for balanced BST
func countBST() func() int {

}

func main() {
	f := countBST()

}

type BinarySearchTree BinarySearchTree;
func (bst *BinarySearchTree) Integer() int {
    return 
}


/* JAVA EQUIVALENT */
class BinarySearchTree {
  private TreeNode root;

  public BinarySearchTree(){
    this.root = new TreeNode();
  }
  public BinarySearchTree(TreeNode root){
    this.root = root;
  }
  public BinarySearchTree(BinarySearchTree bst){
    //TODO: deep copy of existing BST
  }

  public int countNodes(){
    if
  }

}

class TreeNode implements Node {
  private int value;
  private TreeNode left;
  private TreeNode right;

  public TreeNode(){}
  public TreeNode(int val){
    this.value = val;
  }
  public TreeNode(TreeNode node){
    this.value = node.getValue();
    this.left = node.getLeft();
    this.right = node.getRight();
  }

  public int getValue(){
    return this.value;
  }
  public TreeNode getLeft(){
    return this.left;
  }
  public TreeNode getRight(){
    return this.right;
  }

  public void setValue(int val){
    this.value = val;
  }
  public void setLeft(TreeNode node){
    this.left = node;
  }
  public void setRight(TreeNode node){
    this.right = node;
  }

  //hasLeft
  //hasRight

}

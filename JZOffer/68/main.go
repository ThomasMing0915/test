package main

// 题目需要使用c语言实现
/*
struct TreeNode* lowestCommonAncestor(struct TreeNode* root, struct TreeNode* p, struct TreeNode* q){
if ((root == NULL) || (root == p) || (root == q)) {
return root;
}
struct TreeNode *Left;
struct TreeNode *Right;

Left = lowestCommonAncestor(root->left,p,q);
Right = lowestCommonAncestor(root->right,p,q);

if (Left == NULL) {
return Right;
}

if (Right == NULL) {
return Left;
}

return root;
}
*/

//test pass

//struct TreeNode* lowestCommonAncestor(struct TreeNode* root, struct TreeNode* p, struct TreeNode* q) {
//if (root->val > p->val && root->val > q->val){
//return  lowestCommonAncestor(root->left,p,q);
//}else if (root->val < p->val && root->val < q->val) {
//return  lowestCommonAncestor(root->right,p,q);
//}else{
//return root;
//}
//}

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode{
	if root==nil || p==nil || q==nil{
		return nil
	}

	if root.Val<p.Val && root.Val<q.Val{
		return lowestCommonAncestor(root.Right,p,q)
	}else if root.Val>p.Val && root.Val>q.Val{
		return lowestCommonAncestor(root.Left,p,q)
	}else{
		return root
	}
}

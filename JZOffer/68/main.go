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

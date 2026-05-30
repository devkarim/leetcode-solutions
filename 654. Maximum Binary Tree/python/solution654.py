import pytest


# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution(object):
    def constructMaximumBinaryTree(self, nums):
        """
        :type nums: List[int]
        :rtype: Optional[TreeNode]
        """
        st = []
        for num in nums:
            left = None
            while st and st[-1].val < num:
                left = st.pop()
            tree = TreeNode(val=num)
            tree.left = left

            if st:
                st[-1].right = tree

            st.append(tree)

        return st[0]


def tree_to_list(root):
    if not root:
        return []
    result, queue = [], [root]
    while queue:
        node = queue.pop(0)
        if node:
            result.append(node.val)
            queue.append(node.left)
            queue.append(node.right)
        else:
            result.append(None)

    while result and result[-1] is None:
        result.pop()

    return result


@pytest.fixture
def sol():
    return Solution()


def test_example1(sol):
    result = sol.constructMaximumBinaryTree([3, 2, 1, 6, 0, 5])
    assert tree_to_list(result) == [6, 3, 5, None, 2, 0, None, None, 1]

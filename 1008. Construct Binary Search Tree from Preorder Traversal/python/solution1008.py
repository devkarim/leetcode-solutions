import pytest


# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution(object):
    def bstFromPreorder(self, preorder):
        """
        :type preorder: List[int]
        :rtype: Optional[TreeNode]
        """
        root = TreeNode(val=preorder[0])
        st = [root]
        for num in preorder[1:]:
            if num < st[-1].val:
                st[-1].left = TreeNode(val=num)
                st.append(st[-1].left)
            else:
                last = None
                while st and st[-1].val < num:
                    last = st.pop()
                if last:
                    last.right = TreeNode(val=num)
                    st.append(last.right)
        return root


def tree_to_list(root):
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


def test_example1(sol: Solution):
    root = sol.bstFromPreorder([8, 5, 1, 7, 10, 12])
    expected = [8, 5, 10, 1, 7, None, 12]
    assert tree_to_list(root) == expected
